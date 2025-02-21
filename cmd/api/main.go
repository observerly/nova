/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/nova/cmd/api/main
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

	"nova/internal/adapters"
	"nova/internal/gen/solve/v1/solvev1connect"
	"nova/internal/gen/store/v1/storev1connect"
	"nova/service/model"
	"nova/service/solve"
	"nova/service/storage"
)

/*****************************************************************************************************************/

const (
	serviceName = "nova"
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
	storePath, storeHandler := storev1connect.NewStorageServiceHandler(
		storage.NewStorageServer(app, client),
	)

	// Register our Solve service:
	solvePath, solveHandler := solvev1connect.NewSolveServiceHandler(
		solve.NewSolveServer(app, client),
	)

	reflector := grpcreflect.NewStaticReflector(
		storev1connect.StorageServiceName,
	)

	checker := grpchealth.NewStaticChecker(
		storev1connect.StorageServiceName,
	)

	mux := http.NewServeMux()

	// Register our store service with the gRPC server:
	mux.Handle(storePath, storeHandler)

	// Register our solve service with the gRPC server:
	mux.Handle(solvePath, solveHandler)

	// Register reflection service on gRPC server:
	mux.Handle(grpcreflect.NewHandlerV1(reflector))

	// Register health check service on gRPC server:
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// Register health check service on gRPC server:
	mux.Handle(grpchealth.NewHandler(checker))

	// Start gRPC http server h2c handler - run in a go function so we can better handle SIGTERM:
	go func() {
		log.Info().Msgf("Server running on %v", fmt.Sprintf("http://%s:%d", config.Host, config.Port))

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
