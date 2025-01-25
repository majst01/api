// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/health.proto

package apiv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/metal-stack/api/go/api/v1"
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
	// HealthServiceName is the fully-qualified name of the HealthService service.
	HealthServiceName = "api.v1.HealthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HealthServiceGetProcedure is the fully-qualified name of the HealthService's Get RPC.
	HealthServiceGetProcedure = "/api.v1.HealthService/Get"
)

// HealthServiceClient is a client for the api.v1.HealthService service.
type HealthServiceClient interface {
	// Get the health of the platform
	Get(context.Context, *connect.Request[v1.HealthServiceGetRequest]) (*connect.Response[v1.HealthServiceGetResponse], error)
}

// NewHealthServiceClient constructs a client for the api.v1.HealthService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHealthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) HealthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	healthServiceMethods := v1.File_api_v1_health_proto.Services().ByName("HealthService").Methods()
	return &healthServiceClient{
		get: connect.NewClient[v1.HealthServiceGetRequest, v1.HealthServiceGetResponse](
			httpClient,
			baseURL+HealthServiceGetProcedure,
			connect.WithSchema(healthServiceMethods.ByName("Get")),
			connect.WithClientOptions(opts...),
		),
	}
}

// healthServiceClient implements HealthServiceClient.
type healthServiceClient struct {
	get *connect.Client[v1.HealthServiceGetRequest, v1.HealthServiceGetResponse]
}

// Get calls api.v1.HealthService.Get.
func (c *healthServiceClient) Get(ctx context.Context, req *connect.Request[v1.HealthServiceGetRequest]) (*connect.Response[v1.HealthServiceGetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// HealthServiceHandler is an implementation of the api.v1.HealthService service.
type HealthServiceHandler interface {
	// Get the health of the platform
	Get(context.Context, *connect.Request[v1.HealthServiceGetRequest]) (*connect.Response[v1.HealthServiceGetResponse], error)
}

// NewHealthServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHealthServiceHandler(svc HealthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	healthServiceMethods := v1.File_api_v1_health_proto.Services().ByName("HealthService").Methods()
	healthServiceGetHandler := connect.NewUnaryHandler(
		HealthServiceGetProcedure,
		svc.Get,
		connect.WithSchema(healthServiceMethods.ByName("Get")),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.HealthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HealthServiceGetProcedure:
			healthServiceGetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHealthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHealthServiceHandler struct{}

func (UnimplementedHealthServiceHandler) Get(context.Context, *connect.Request[v1.HealthServiceGetRequest]) (*connect.Response[v1.HealthServiceGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.HealthService.Get is not implemented"))
}
