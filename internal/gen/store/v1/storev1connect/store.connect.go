//***************************************************************************************************************
//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc/proto/store/v1
//	@license	Copyright © 2021-2024 observerly
//***************************************************************************************************************

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: store/v1/store.proto

//***************************************************************************************************************

package storev1connect

import (
	v1 "birpc/internal/gen/store/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// StorageServiceName is the fully-qualified name of the StorageService service.
	StorageServiceName = "store.v1.StorageService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StorageServiceGetFITSAsTIFFHandlerProcedure is the fully-qualified name of the StorageService's
	// GetFITSAsTIFFHandler RPC.
	StorageServiceGetFITSAsTIFFHandlerProcedure = "/store.v1.StorageService/GetFITSAsTIFFHandler"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	storageServiceServiceDescriptor                    = v1.File_store_v1_store_proto.Services().ByName("StorageService")
	storageServiceGetFITSAsTIFFHandlerMethodDescriptor = storageServiceServiceDescriptor.Methods().ByName("GetFITSAsTIFFHandler")
)

// StorageServiceClient is a client for the store.v1.StorageService service.
type StorageServiceClient interface {
	GetFITSAsTIFFHandler(context.Context, *connect.Request[v1.GetFITSAsGenericHandlerRequest]) (*connect.Response[v1.GetFITSAsGenericHandlerResponse], error)
}

// NewStorageServiceClient constructs a client for the store.v1.StorageService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStorageServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) StorageServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &storageServiceClient{
		getFITSAsTIFFHandler: connect.NewClient[v1.GetFITSAsGenericHandlerRequest, v1.GetFITSAsGenericHandlerResponse](
			httpClient,
			baseURL+StorageServiceGetFITSAsTIFFHandlerProcedure,
			connect.WithSchema(storageServiceGetFITSAsTIFFHandlerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// storageServiceClient implements StorageServiceClient.
type storageServiceClient struct {
	getFITSAsTIFFHandler *connect.Client[v1.GetFITSAsGenericHandlerRequest, v1.GetFITSAsGenericHandlerResponse]
}

// GetFITSAsTIFFHandler calls store.v1.StorageService.GetFITSAsTIFFHandler.
func (c *storageServiceClient) GetFITSAsTIFFHandler(ctx context.Context, req *connect.Request[v1.GetFITSAsGenericHandlerRequest]) (*connect.Response[v1.GetFITSAsGenericHandlerResponse], error) {
	return c.getFITSAsTIFFHandler.CallUnary(ctx, req)
}

// StorageServiceHandler is an implementation of the store.v1.StorageService service.
type StorageServiceHandler interface {
	GetFITSAsTIFFHandler(context.Context, *connect.Request[v1.GetFITSAsGenericHandlerRequest]) (*connect.Response[v1.GetFITSAsGenericHandlerResponse], error)
}

// NewStorageServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStorageServiceHandler(svc StorageServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	storageServiceGetFITSAsTIFFHandlerHandler := connect.NewUnaryHandler(
		StorageServiceGetFITSAsTIFFHandlerProcedure,
		svc.GetFITSAsTIFFHandler,
		connect.WithSchema(storageServiceGetFITSAsTIFFHandlerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/store.v1.StorageService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StorageServiceGetFITSAsTIFFHandlerProcedure:
			storageServiceGetFITSAsTIFFHandlerHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStorageServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStorageServiceHandler struct{}

func (UnimplementedStorageServiceHandler) GetFITSAsTIFFHandler(context.Context, *connect.Request[v1.GetFITSAsGenericHandlerRequest]) (*connect.Response[v1.GetFITSAsGenericHandlerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("store.v1.StorageService.GetFITSAsTIFFHandler is not implemented"))
}
