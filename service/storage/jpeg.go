/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package storage

/*****************************************************************************************************************/

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"strings"
	"time"

	pb "birpc/internal/gen/store/v1"
	"birpc/internal/stores"

	cloud "cloud.google.com/go/storage"
	"connectrpc.com/connect"
	"github.com/observerly/iris/pkg/image"
	"github.com/rs/zerolog/log"
)

/*****************************************************************************************************************/

func (s *server) GetFITSAsJPEGHandler(ctx context.Context, req *connect.Request[pb.GetFITSAsGenericHandlerRequest]) (*connect.Response[pb.GetFITSAsGenericHandlerResponse], error) {
	now := time.Now()

	s.Logger = log.With().Str("owner", req.Msg.Owner).Str("bucket", req.Msg.BucketName).Str("location", req.Msg.Location).Str("rfc3339", now.Format(time.RFC3339)).Logger()

	// Get the bucket from the storage client:
	bucket, err := s.Storage.Bucket(req.Msg.BucketName)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to get bucket")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

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

	// Convert the FITS exposure data to a 16-bit grayscale image:
	img, err := image.NewGray16FromRawFloat32Pixels(fit.Data, int(fit.Header.Naxis1))

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to convert exposure data to image")
		return nil, err
	}

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to get image as 16-bit grayscale from FITS")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get image as 16-bit grayscale from FITS: %w", err))
	}

	// Create a new buffer to store the object data:
	jpegb := new(bytes.Buffer)

	// Encode the image as a JPEG:
	err = jpeg.Encode(jpegb, img, nil)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to encode image as JPEG")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to encode image to JPEG: %w", err))
	}

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to encode image as JPEG")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Remove the .fits extension from the location:
	location := strings.Replace(req.Msg.Location, ".fits", ".jpeg", 1)

	// Store the exposure image in Firebase Storage:
	err = s.Client.StoreBuffer(ctx, jpegb, req.Msg.BucketName, location, stores.StoreBufferParams{
		ContentType: "image/jpeg",
		Owner:       req.Msg.Owner,
	})

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to store buffer")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Generate a signed URL for the newly created image in Firebase Storage, valid for 1 minutes:
	url, err := bucket.SignedURL(location, &cloud.SignedURLOptions{
		Expires: now.Add(1 * time.Minute),
		Method:  "GET",
	})

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to generate signed URL")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Delete the newly created image from Firebase Storage after 1 minutes (until the signed URL expires):
	go func() {
		ctx = context.Background()

		// Sleep for the duration of the signed URL:
		time.Sleep(1 * time.Minute)

		// Delete the newly created image from Firebase Storage:
		err = bucket.Object(location).Delete(ctx)

		if err != nil {
			s.Logger.Error().Err(err).Msg("Failed to delete object")
		}
	}()

	bounds := img.Bounds()

	s.Logger.Info().Msg("Returning JPEG Download URL")

	return connect.NewResponse(&pb.GetFITSAsGenericHandlerResponse{
		DownloadUrl: url,
		Height:      int32(bounds.Dy()),
		Width:       int32(bounds.Dx()),
	}), nil
}

/*****************************************************************************************************************/
