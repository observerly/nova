/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package storage

/*****************************************************************************************************************/

import (
	"context"
	"fmt"
	img "image"

	pb "birpc/internal/gen/store/v1/storev1connect"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/iris/pkg/image"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"birpc/internal/stores"
)

/*****************************************************************************************************************/

type server struct {
	pb.UnimplementedStorageServiceHandler
	App     *firebase.App
	Logger  zerolog.Logger
	Storage *storage.Client
	Client  *stores.FirebaseStore
}

/*****************************************************************************************************************/

func NewStorageServer(app *firebase.App, storage *storage.Client) *server {
	// Create a new logger:
	logger := log.With().Str("@observerly/stores", "rpc").Logger()

	// Create a new Firebase Storage client:
	client := stores.NewFirebaseStorageClient(storage)

	return &server{
		App:     app,
		Logger:  logger,
		Storage: storage,
		Client:  client,
	}
}

/*****************************************************************************************************************/

func (s *server) RetrieveFITSFromStorage(
	ctx context.Context,
	bucketName string,
	location string,
) (*fits.FITSImage, error) {
	// Assume an image of 2x2 pixels with 16-bit depth, and no offset:
	fit := fits.NewFITSImage(2, 0, 0, 65535)

	// Get the buffer from the storage client:
	buff, err := s.Client.RetriveBuffer(ctx, bucketName, location)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to get buffer")
		return nil, err
	}

	// Read in our exposure data into the image:
	err = fit.Read(buff)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to read exposure data")
		return nil, fmt.Errorf("failed to read exposure data: %w", err)
	}

	// We know the image is 2D, so we can extract the width from the fits image:
	width := fit.Header.Naxis1

	// We know the image is 2D, so we can extract the height from the fits image:
	height := fit.Header.Naxis2

	if fit.Pixels != width*height {
		s.Logger.Error().Msg("Failed to read exposure data")
		return nil, fmt.Errorf("failed to read exposure data as the number of pixels does not match the width and height")
	}

	return fit, nil
}

/*****************************************************************************************************************/

func (s *server) getFITSAsGray16Image(
	ctx context.Context,
	bucketName string,
	location string,
) (*img.Gray16, error) {
	// Assume an image of 2x2 pixels with 16-bit depth, and no offset:
	fit := fits.NewFITSImage(2, 0, 0, 65535)

	// Get the buffer from the storage client:
	buff, err := s.Client.RetriveBuffer(ctx, bucketName, location)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to get buffer")
		return nil, err
	}

	// Read in our exposure data into the image:
	err = fit.Read(buff)

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to read exposure data")
		return nil, fmt.Errorf("failed to read exposure data: %w", err)
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
	i, err := image.NewGray16FromRawFloat32Pixels(fit.Data, int(width))

	if err != nil {
		s.Logger.Error().Err(err).Msg("Failed to convert exposure data to image")
		return nil, err
	}

	return i, nil
}

/*****************************************************************************************************************/
