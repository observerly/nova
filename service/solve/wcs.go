/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/nova
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package solve

import (
	"context"
	"fmt"
	"math"
	"time"

	"connectrpc.com/connect"

	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/skysolve/pkg/astrometry"
	"github.com/observerly/skysolve/pkg/catalog"
	"github.com/observerly/skysolve/pkg/fov"
	"github.com/observerly/skysolve/pkg/solve"
	"github.com/rs/zerolog/log"

	pb "nova/internal/gen/solve/v1"
	pbWCS "nova/internal/gen/wcs/v1"
	"nova/internal/middleware"
	"nova/internal/stores"
)

/*****************************************************************************************************************/

func (s *server) solveForWCSFITSHandler(ctx context.Context, req *connect.Request[pb.SolveForWCSFITSHandlerRequest]) (*connect.Response[pb.SolveForWCSFITSHandlerResponse], error) {
	now := time.Now()

	s.Logger = log.With().Str("owner", req.Msg.Owner).Str("bucket", req.Msg.BucketName).Str("location", req.Msg.Location).Str("rfc3339", now.Format(time.RFC3339)).Logger()

	// Get the FITS file from the storage client:
	fit, err := s.RetrieveFITSFromStorage(ctx, req.Msg.BucketName, req.Msg.Location)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to retrieve FITS from storage")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// We know the image is 2D, so we can extract the width from the fits image:
	width := fit.Header.Naxis1

	// We know the image is 2D, so we can extract the height from the fits image:
	height := fit.Header.Naxis2

	if fit.Pixels != width*height {
		s.Logger.Error().Msg("Failed to read exposure data")
		return nil, fmt.Errorf("failed to read exposure data as the number of pixels does not match the width and height")
	}

	ra := math.Inf(1)

	// Attempt to get the RA header from the FITS file:
	if hdr, ok := fit.Header.Floats["RA"]; ok {
		ra = float64(hdr.Value)
	}

	if req.Msg.Ra != nil {
		ra = req.Msg.GetRa()
	}

	if ra == math.MaxFloat32 {
		s.Logger.Warn().Msg("No RA in proto or FITS header; defaulting to 0°")
	}

	// Ensure that the Right Ascension is between 0 and 360 degrees:
	if ra < 0 || ra > 360 {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("ra must be between 0 and 360 degrees"))
	}

	dec := math.Inf(1)

	// Attempt to get the Dec header from the FITS file:
	if hdr, ok := fit.Header.Floats["DEC"]; ok {
		dec = float64(hdr.Value)
	}

	if req.Msg.Dec != nil {
		dec = req.Msg.GetDec()
	}

	if dec == math.MaxFloat32 {
		s.Logger.Warn().Msg("No Dec in proto or FITS header; defaulting to 0°")
	}

	// Ensure that the Declination is between -90 and 90 degrees:
	if dec < -90 || dec > 90 {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("dec must be between -90 and 90 degrees"))
	}

	pixelScaleX := 0.000540 // 0.000540 degrees per pixel in the x-axis is a sensible default

	if req.Msg.PixelScaleX != nil {
		pixelScaleX = req.Msg.GetPixelScaleX()
	}

	pixelScaleY := 0.000540 // 0.000540 degrees per pixel in the y-axis is a sensible default

	if req.Msg.PixelScaleY != nil {
		pixelScaleY = req.Msg.GetPixelScaleY()
	}

	// Create a new GAIA service client:
	service := catalog.NewCatalogService(catalog.GAIA, catalog.Params{
		Limit:     60, // Limit the number of records to ~100
		Threshold: 16, // Limiting Magntiude, filter out any stars that are magnitude 16 or above (fainter)
	})

	sigma := 2.5

	if req.Msg.Sigma != nil {
		sigma = req.Msg.GetSigma()
	}

	radius := 16.0

	if req.Msg.Radius != nil {
		radius = req.Msg.GetRadius()
	}

	// Attempt to create a new PlateSolver:
	solver, err := solve.NewPlateSolver(solve.Params{
		Data:                fit.Data,               // The exposure data from the fits image
		Width:               int(fit.Header.Naxis1), // The width of the image
		Height:              int(fit.Header.Naxis2), // The height of the image
		PixelScaleX:         pixelScaleX,            // The pixel scale in the x-axis
		PixelScaleY:         pixelScaleY,            // The pixel scale in the y-axis
		ADU:                 fit.ADU,                // The analog-to-digital unit of the image
		ExtractionThreshold: 16,                     // Extract a minimum of 20 of the brightest stars
		Radius:              radius,                 // 16 pixels radius for the star extraction
		Sigma:               sigma,                  // 8 pixels sigma for the Gaussian kernel
	})

	if err != nil {
		fmt.Printf("there was an error while creating the plate solver: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Define the tolerances parameters for the solver:
	tolerance := solve.ToleranceParams{
		QuadTolerance:           0.02,
		EuclidianPixelTolerance: 10,
	}

	// Whilst we have no matches, and whilst we are within 1 degree of the initial { ra, dec } guess, keep solving:
	eq := astrometry.ICRSEquatorialCoordinate{
		RA:  ra,
		Dec: dec,
	}

	searchRadius := fov.GetRadialExtent(
		float64(fit.Header.Naxis1),
		float64(fit.Header.Naxis2),
		fov.PixelScale{X: pixelScaleX, Y: pixelScaleY},
	)

	// Perform a radial search with the given center and radius, for all sources with a magnitude less than 10:
	sources, err := service.PerformRadialSearch(eq, searchRadius)

	if err != nil {
		fmt.Printf("there was an error while performing the radial search: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	solver.Sources = append(solver.Sources, sources...)

	wcs, _, err := solver.Solve(tolerance, 3)

	if err != nil {
		fmt.Printf("there was an error while solving for the WCS: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if wcs == nil {
		fmt.Printf("wcs is nil")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("wcs is nil"))
	}

	// At this point, we have successfully solved for the WCS, so we can update the FITS header with the WCS information:
	fit.Header.Floats["WCSAXES"] = fits.FITSHeaderFloat{
		Value:   2,
		Comment: "Number of World Coordinate System axes",
	}

	fit.Header.Floats["CRPIX1"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CRPIX1),
		Comment: "X-coordinate of the reference pixel",
	}

	fit.Header.Floats["CRPIX2"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CRPIX2),
		Comment: "Y-coordinate of the reference pixel",
	}

	fit.Header.Floats["CRVAL1"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CRVAL1),
		Comment: "Right Ascension of the reference pixel",
	}

	fit.Header.Floats["CRVAL2"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CRVAL2),
		Comment: "Declination of the reference pixel",
	}

	fit.Header.Strings["CTYPE1"] = fits.FITSHeaderString{
		Value:   wcs.CTYPE1,
		Comment: "Coordinate type for axis 1, typically RA with TAN projection",
	}

	fit.Header.Strings["CTYPE2"] = fits.FITSHeaderString{
		Value:   wcs.CTYPE2,
		Comment: "Coordinate type for axis 2, typically RA with TAN projection",
	}

	fit.Header.Floats["CDELT1"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CDELT1),
		Comment: "Coordinate increment for axis 1",
	}

	fit.Header.Floats["CDELT2"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CDELT2),
		Comment: "Coordinate increment for axis 2",
	}

	fit.Header.Strings["CUNIT1"] = fits.FITSHeaderString{
		Value:   wcs.CUNIT1,
		Comment: "Coordinate unit for axis 1",
	}

	fit.Header.Strings["CUNIT2"] = fits.FITSHeaderString{
		Value:   wcs.CUNIT2,
		Comment: "Coordinate unit for axis 2",
	}

	fit.Header.Floats["CD1_1"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CD1_1),
		Comment: "Coordinate transformation matrix element",
	}

	fit.Header.Floats["CD1_2"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CD1_2),
		Comment: "Coordinate transformation matrix element",
	}

	fit.Header.Floats["CD2_1"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CD2_1),
		Comment: "Coordinate transformation matrix element",
	}

	fit.Header.Floats["CD2_2"] = fits.FITSHeaderFloat{
		Value:   float32(wcs.CD2_2),
		Comment: "Coordinate transformation matrix element",
	}

	fit.Header.Floats["RADESYS"] = fits.FITSHeaderFloat{
		Value:   2000.0,
		Comment: "Equatorial coordinate system",
	}

	// fit.Header.Strings["A_ORDER"] = fits.FITSHeaderString{
	// 	Value:   fmt.Sprintf("%d", wcs.FSIP.AOrder),
	// 	Comment: "SIP polynomial order for A",
	// }

	// for key, sips := range wcs.FSIP.APower {
	// 	fit.Header.Floats[fmt.Sprintf("A_%s", key)] = fits.FITSHeaderFloat{
	// 		Value:   float32(sips),
	// 		Comment: fmt.Sprintf("SIP coefficient for A order %s", key),
	// 	}
	// }

	// fit.Header.Strings["B_ORDER"] = fits.FITSHeaderString{
	// 	Value:   fmt.Sprintf("%d", wcs.FSIP.BOrder),
	// 	Comment: "SIP polynomial order for A",
	// }

	// for key, sips := range wcs.FSIP.BPower {
	// 	fit.Header.Floats[fmt.Sprintf("B_%s", key)] = fits.FITSHeaderFloat{
	// 		Value:   float32(sips),
	// 		Comment: fmt.Sprintf("SIP coefficient for B order %s", key),
	// 	}
	// }

	// fit.Header.Strings["AP_ORDER"] = fits.FITSHeaderString{
	// 	Value:   fmt.Sprintf("%d", wcs.ISIP.APOrder),
	// 	Comment: "SIP polynomial order for A",
	// }

	// for key, sips := range wcs.ISIP.APPower {
	// 	fit.Header.Floats[fmt.Sprintf("AP_%s", key)] = fits.FITSHeaderFloat{
	// 		Value:   float32(sips),
	// 		Comment: fmt.Sprintf("SIP coefficient for AP order %s", key),
	// 	}
	// }

	// fit.Header.Strings["BP_ORDER"] = fits.FITSHeaderString{
	// 	Value:   fmt.Sprintf("%d", wcs.ISIP.BPOrder),
	// 	Comment: "SIP polynomial order for A",
	// }

	// for key, sips := range wcs.ISIP.BPPower {
	// 	fit.Header.Floats[fmt.Sprintf("BP_%s", key)] = fits.FITSHeaderFloat{
	// 		Value:   float32(sips),
	// 		Comment: fmt.Sprintf("SIP coefficient for BP order %s", key),
	// 	}
	// }

	// Save the FITS file back to the storage client with the updated WCS information:
	go func(fit *fits.FITSImage) {
		buff, err := fit.WriteToBuffer()

		if err != nil {
			s.Logger.Error().Err(err).Msg("Failed to write FITS to buffer")
		}

		// Store the exposure image in Firebase Storage:
		err = s.Client.StoreBuffer(ctx, buff, req.Msg.BucketName, req.Msg.Location, stores.StoreBufferParams{
			ContentType: "application/fits",
			Owner:       req.Msg.Owner,
		})

		if err != nil {
			s.Logger.Error().Err(err).Msg("Failed to store buffer")
		}
	}(fit)

	fsip := &pbWCS.SIP2DForwardParameters{
		AOrder: int32(wcs.FSIP.AOrder),
		APower: wcs.FSIP.APower,
		BOrder: int32(wcs.FSIP.BOrder),
		BPower: wcs.FSIP.BPower,
	}

	isip := &pbWCS.SIP2DInverseParameters{
		APOrder: int32(wcs.ISIP.APOrder),
		APPower: wcs.ISIP.APPower,
		BPOrder: int32(wcs.ISIP.BPOrder),
		BPPower: wcs.ISIP.BPPower,
	}

	return connect.NewResponse(&pb.SolveForWCSFITSHandlerResponse{
		Wcs: &pbWCS.WCS{
			WCAXES: int32(wcs.WCAXES),
			CRPIX1: wcs.CRPIX1,
			CRPIX2: wcs.CRPIX2,
			CRVAL1: wcs.CRVAL1,
			CRVAL2: wcs.CRVAL2,
			CTYPE1: wcs.CTYPE1,
			CTYPE2: wcs.CTYPE2,
			CDELT1: wcs.CDELT1,
			CDELT2: wcs.CDELT2,
			CUNIT1: wcs.CUNIT1,
			CUNIT2: wcs.CUNIT2,
			CD1_1:  wcs.CD1_1,
			CD1_2:  wcs.CD1_2,
			CD2_1:  wcs.CD2_1,
			CD2_2:  wcs.CD2_2,
			E:      wcs.E,
			F:      wcs.F,
			FSIP:   fsip,
			ISIP:   isip,
		},
	}), nil
}

/*****************************************************************************************************************/

func (s *server) SolveForWCSFITSHandler(ctx context.Context, req *connect.Request[pb.SolveForWCSFITSHandlerRequest]) (*connect.Response[pb.SolveForWCSFITSHandlerResponse], error) {
	auth, err := s.App.Auth(ctx)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to authenticate user")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to authenticate user: %w", err))
	}

	return middleware.MustHaveAuthentication(ctx, req, auth, func() (*connect.Response[pb.SolveForWCSFITSHandlerResponse], error) {
		return s.solveForWCSFITSHandler(ctx, req)
	})
}

/*****************************************************************************************************************/
