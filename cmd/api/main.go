/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc/cmd/api/main
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package main

/*****************************************************************************************************************/

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	"birpc/internal/adapters"
	"birpc/internal/gen/store/v1/storev1connect"
	"birpc/service/model"
	"birpc/service/storage"
)

/*****************************************************************************************************************/

const (
	serviceName = "storage"
)

/*****************************************************************************************************************/

func main() {
	// Set zerolog to use unix time
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var config model.Config

	// Parse the environment variables:
	err := envconfig.Process(serviceName, &config)

	if err != nil {
		// could not parse config
		log.Fatal().Err(err).Msg("Cannot read config")
	}

	// Setup the Firebase app:
	app, err := adapters.SetupFirebaseApp()

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot setup Firebase app")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)

	client, err := app.Storage(ctx)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot setup Firebase storage client")
	}

	defer stop()

	// Setup our base gRPC server:
	s := grpc.NewServer()

	// Register our Store service:
	path, handler := storev1connect.NewStorageServiceHandler(
		storage.NewStorageServer(app, client),
	)

	reflector := grpcreflect.NewStaticReflector(
		storev1connect.StorageServiceName,
	)

	checker := grpchealth.NewStaticChecker(
		storev1connect.StorageServiceName,
	)

	mux := http.NewServeMux()

	// Register our service with the gRPC server:
	mux.Handle(path, handler)

	// Register reflection service on gRPC server:
	mux.Handle(grpcreflect.NewHandlerV1(reflector))

	// Register health check service on gRPC server:
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// Register health check service on gRPC server:
	mux.Handle(grpchealth.NewHandler(checker))

	// Start gRPC http server h2c handler - run in a go function so we can better handle SIGTERM:
	go func() {
		log.Info().Msgf("Server running on %v%s", fmt.Sprintf("http://%s:%d", config.Host, config.Port), path)

		err := http.ListenAndServe(
			fmt.Sprintf("%s:%d", config.Host, config.Port),
			// Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(mux, &http2.Server{}),
		)
		if err != nil {
			log.Fatal().Err(err).Msgf("failed to start server")
		}

		ctx.Done()
	}()

	// Wait for Control C signal termination to exit:
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)

	// Block until a signal is received:
	<-ch
	log.Info().Msg("Stopping the server")
	// Wait for the server to stop gracefully:
	s.GracefulStop()
	log.Info().Msg("Server shutdown")
}

/*****************************************************************************************************************/
