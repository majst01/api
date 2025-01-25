// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/methods.proto

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
	// MethodServiceName is the fully-qualified name of the MethodService service.
	MethodServiceName = "api.v1.MethodService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MethodServiceListProcedure is the fully-qualified name of the MethodService's List RPC.
	MethodServiceListProcedure = "/api.v1.MethodService/List"
	// MethodServiceTokenScopedListProcedure is the fully-qualified name of the MethodService's
	// TokenScopedList RPC.
	MethodServiceTokenScopedListProcedure = "/api.v1.MethodService/TokenScopedList"
)

// MethodServiceClient is a client for the api.v1.MethodService service.
type MethodServiceClient interface {
	// List all public visible methods
	List(context.Context, *connect.Request[v1.MethodServiceListRequest]) (*connect.Response[v1.MethodServiceListResponse], error)
	// TokenScopedList all methods callable with the token present in the request
	TokenScopedList(context.Context, *connect.Request[v1.MethodServiceTokenScopedListRequest]) (*connect.Response[v1.MethodServiceTokenScopedListResponse], error)
}

// NewMethodServiceClient constructs a client for the api.v1.MethodService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMethodServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MethodServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	methodServiceMethods := v1.File_api_v1_methods_proto.Services().ByName("MethodService").Methods()
	return &methodServiceClient{
		list: connect.NewClient[v1.MethodServiceListRequest, v1.MethodServiceListResponse](
			httpClient,
			baseURL+MethodServiceListProcedure,
			connect.WithSchema(methodServiceMethods.ByName("List")),
			connect.WithClientOptions(opts...),
		),
		tokenScopedList: connect.NewClient[v1.MethodServiceTokenScopedListRequest, v1.MethodServiceTokenScopedListResponse](
			httpClient,
			baseURL+MethodServiceTokenScopedListProcedure,
			connect.WithSchema(methodServiceMethods.ByName("TokenScopedList")),
			connect.WithClientOptions(opts...),
		),
	}
}

// methodServiceClient implements MethodServiceClient.
type methodServiceClient struct {
	list            *connect.Client[v1.MethodServiceListRequest, v1.MethodServiceListResponse]
	tokenScopedList *connect.Client[v1.MethodServiceTokenScopedListRequest, v1.MethodServiceTokenScopedListResponse]
}

// List calls api.v1.MethodService.List.
func (c *methodServiceClient) List(ctx context.Context, req *connect.Request[v1.MethodServiceListRequest]) (*connect.Response[v1.MethodServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// TokenScopedList calls api.v1.MethodService.TokenScopedList.
func (c *methodServiceClient) TokenScopedList(ctx context.Context, req *connect.Request[v1.MethodServiceTokenScopedListRequest]) (*connect.Response[v1.MethodServiceTokenScopedListResponse], error) {
	return c.tokenScopedList.CallUnary(ctx, req)
}

// MethodServiceHandler is an implementation of the api.v1.MethodService service.
type MethodServiceHandler interface {
	// List all public visible methods
	List(context.Context, *connect.Request[v1.MethodServiceListRequest]) (*connect.Response[v1.MethodServiceListResponse], error)
	// TokenScopedList all methods callable with the token present in the request
	TokenScopedList(context.Context, *connect.Request[v1.MethodServiceTokenScopedListRequest]) (*connect.Response[v1.MethodServiceTokenScopedListResponse], error)
}

// NewMethodServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMethodServiceHandler(svc MethodServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	methodServiceMethods := v1.File_api_v1_methods_proto.Services().ByName("MethodService").Methods()
	methodServiceListHandler := connect.NewUnaryHandler(
		MethodServiceListProcedure,
		svc.List,
		connect.WithSchema(methodServiceMethods.ByName("List")),
		connect.WithHandlerOptions(opts...),
	)
	methodServiceTokenScopedListHandler := connect.NewUnaryHandler(
		MethodServiceTokenScopedListProcedure,
		svc.TokenScopedList,
		connect.WithSchema(methodServiceMethods.ByName("TokenScopedList")),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.MethodService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MethodServiceListProcedure:
			methodServiceListHandler.ServeHTTP(w, r)
		case MethodServiceTokenScopedListProcedure:
			methodServiceTokenScopedListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMethodServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMethodServiceHandler struct{}

func (UnimplementedMethodServiceHandler) List(context.Context, *connect.Request[v1.MethodServiceListRequest]) (*connect.Response[v1.MethodServiceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.MethodService.List is not implemented"))
}

func (UnimplementedMethodServiceHandler) TokenScopedList(context.Context, *connect.Request[v1.MethodServiceTokenScopedListRequest]) (*connect.Response[v1.MethodServiceTokenScopedListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.MethodService.TokenScopedList is not implemented"))
}
