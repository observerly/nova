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
	"strings"
	"time"

	pb "birpc/internal/gen/store/v1"
	"birpc/internal/stores"

	cloud "cloud.google.com/go/storage"
	"connectrpc.com/connect"
	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/iris/pkg/image"
	"github.com/rs/zerolog/log"
	"golang.org/x/image/tiff"
)

/*****************************************************************************************************************/

func (s *server) GetFITSAsTIFFHandler(ctx context.Context, req *connect.Request[pb.GetFITSAsGenericHandlerRequest]) (*connect.Response[pb.GetFITSAsGenericHandlerResponse], error) {
	now := time.Now()

	// Assume an image of 2x2 pixels with 16-bit depth, and no offset:
	fit := fits.NewFITSImage(2, 0, 0, 65535)

	logger := log.With().Str("owner", req.Msg.Owner).Str("bucket", req.Msg.BucketName).Str("location", req.Msg.Location).Str("rfc3339", now.Format(time.RFC3339)).Logger()

	url := fmt.Sprintf("https://%s/%s", req.Msg.BucketName, req.Msg.Location)

	width := int32(0)

	height := int32(0)

	storage, err := s.App.Storage(ctx)

	if err != nil {
		logger.Error().Err(err).Msg("Failed to get storage client")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Get the bucket from the storage client:
	bucket, err := storage.Bucket(req.Msg.BucketName)

	if err != nil {
		logger.Error().Err(err).Msg("Failed to get bucket")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Create a new Firebase Storage client:
	client := stores.NewFirebaseStorageClient(storage)

	// Get the buffer from the storage client:
	buff, err := client.RetriveBuffer(ctx, req.Msg.BucketName, req.Msg.Location)

	if err != nil {
		logger.Error().Err(err).Msg("Failed to get buffer")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Read in our exposure data into the image:
	err = fit.Read(buff)

	if err != nil {
		logger.Error().Err(err).Msg("Failed to read exposure data")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// We know the image is 2D, so we can extract the width from the fits image:
	width = fit.Header.Naxis1

	// We know the image is 2D, so we can extract the height from the fits image:
	height = fit.Header.Naxis2

	if fit.Pixels != width*height {
		logger.Error().Msg("Failed to read exposure data")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to read exposure data as the number of pixels does not match the width and height"))
	}

	// [TBI]: Extract the fits.Date to a simple Gray16 image:
	// img, err := s.float32ArrayToGray16Image(fit.Data, int(width), int(height))
	img, err := image.NewGray16FromRawFloat32Pixels(fit.Data, int(width))

	if err != nil {
		logger.Error().Err(err).Msg("Failed to convert exposure data to image")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Create a new buffer to store the object data:
	tiffb := new(bytes.Buffer)

	// Encode the image as a TIFF:
	err = tiff.Encode(tiffb, img, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to encode image to TIFF: %w", err)
	}

	if err != nil {
		logger.Error().Err(err).Msg("Failed to encode image as TIFF")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Remove the .fits extension from the location:
	location := strings.Replace(req.Msg.Location, ".fits", ".tiff", 1)

	// Store the exposure image in Firebase Storage:
	err = client.StoreBuffer(ctx, tiffb, req.Msg.BucketName, location, stores.StoreBufferParams{
		ContentType: "image/tiff",
		Owner:       req.Msg.Owner,
	})

	if err != nil {
		logger.Error().Err(err).Msg("Failed to store buffer")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Generate a signed URL for the newly created image in Firebase Storage, valid for 1 minutes:
	url, err = bucket.SignedURL(location, &cloud.SignedURLOptions{
		Expires: now.Add(1 * time.Minute),
		Method:  "GET",
	})

	if err != nil {
		logger.Error().Err(err).Msg("Failed to generate signed URL")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Delete the newly created image from Firebase Storage after 1 minutes (until the signed URL expires):
	go func() {
		ctx := context.Background()

		// Sleep for the duration of the signed URL:
		time.Sleep(1 * time.Minute)

		// Delete the newly created image from Firebase Storage:
		err = bucket.Object(location).Delete(ctx)

		if err != nil {
			logger.Error().Err(err).Msg("Failed to delete object")
		}
	}()

	logger.Info().Msg("Returning TIFF Download URL")

	return connect.NewResponse(&pb.GetFITSAsGenericHandlerResponse{
		DownloadUrl: url,
		Height:      height,
		Width:       width,
	}), nil
}

/*****************************************************************************************************************/
