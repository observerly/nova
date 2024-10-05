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
	"github.com/rs/zerolog/log"
	"golang.org/x/image/tiff"
)

/*****************************************************************************************************************/

func (s *server) GetFITSAsTIFFHandler(ctx context.Context, req *connect.Request[pb.GetFITSAsGenericHandlerRequest]) (*connect.Response[pb.GetFITSAsGenericHandlerResponse], error) {
	now := time.Now()

	s.Logger = log.With().Str("owner", req.Msg.Owner).Str("bucket", req.Msg.BucketName).Str("location", req.Msg.Location).Str("rfc3339", now.Format(time.RFC3339)).Logger()

	// Get the bucket from the storage client:
	bucket, err := s.Storage.Bucket(req.Msg.BucketName)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to get bucket")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Get the image from the storage client as a 16-bit grayscale image:
	img, err := s.getFITSAsGray16Image(ctx, req.Msg.BucketName, req.Msg.Location)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to get image as 16-bit grayscale from FITS")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get image as 16-bit grayscale from FITS: %w", err))
	}

	// Create a new buffer to store the object data:
	tiffb := new(bytes.Buffer)

	// Encode the image as a TIFF:
	err = tiff.Encode(tiffb, img, nil)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to encode image as TIFF")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to encode image to TIFF: %w", err))
	}

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to encode image as TIFF")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Remove the .fits extension from the location:
	location := strings.Replace(req.Msg.Location, ".fits", ".tiff", 1)

	// Store the exposure image in Firebase Storage:
	err = s.Client.StoreBuffer(ctx, tiffb, req.Msg.BucketName, location, stores.StoreBufferParams{
		ContentType: "image/tiff",
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

	s.Logger.Info().Msg("Returning TIFF Download URL")

	return connect.NewResponse(&pb.GetFITSAsGenericHandlerResponse{
		DownloadUrl: url,
		Height:      int32(bounds.Dy()),
		Width:       int32(bounds.Dx()),
	}), nil
}

/*****************************************************************************************************************/
