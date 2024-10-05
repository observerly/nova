/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package storage

/*****************************************************************************************************************/

import (
	pb "birpc/internal/gen/store/v1/storev1connect"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
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
