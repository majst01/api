// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/ip.proto

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
	// IPServiceName is the fully-qualified name of the IPService service.
	IPServiceName = "api.v1.IPService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IPServiceGetProcedure is the fully-qualified name of the IPService's Get RPC.
	IPServiceGetProcedure = "/api.v1.IPService/Get"
	// IPServiceAllocateProcedure is the fully-qualified name of the IPService's Allocate RPC.
	IPServiceAllocateProcedure = "/api.v1.IPService/Allocate"
	// IPServiceUpdateProcedure is the fully-qualified name of the IPService's Update RPC.
	IPServiceUpdateProcedure = "/api.v1.IPService/Update"
	// IPServiceListProcedure is the fully-qualified name of the IPService's List RPC.
	IPServiceListProcedure = "/api.v1.IPService/List"
	// IPServiceDeleteProcedure is the fully-qualified name of the IPService's Delete RPC.
	IPServiceDeleteProcedure = "/api.v1.IPService/Delete"
)

// IPServiceClient is a client for the api.v1.IPService service.
type IPServiceClient interface {
	// Get a ip
	Get(context.Context, *connect.Request[v1.IPServiceGetRequest]) (*connect.Response[v1.IPServiceGetResponse], error)
	// Allocate a ip
	Allocate(context.Context, *connect.Request[v1.IPServiceAllocateRequest]) (*connect.Response[v1.IPServiceAllocateResponse], error)
	// Update a ip
	Update(context.Context, *connect.Request[v1.IPServiceUpdateRequest]) (*connect.Response[v1.IPServiceUpdateResponse], error)
	// List all ips
	List(context.Context, *connect.Request[v1.IPServiceListRequest]) (*connect.Response[v1.IPServiceListResponse], error)
	// Delete a ip
	Delete(context.Context, *connect.Request[v1.IPServiceDeleteRequest]) (*connect.Response[v1.IPServiceDeleteResponse], error)
}

// NewIPServiceClient constructs a client for the api.v1.IPService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIPServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) IPServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	iPServiceMethods := v1.File_api_v1_ip_proto.Services().ByName("IPService").Methods()
	return &iPServiceClient{
		get: connect.NewClient[v1.IPServiceGetRequest, v1.IPServiceGetResponse](
			httpClient,
			baseURL+IPServiceGetProcedure,
			connect.WithSchema(iPServiceMethods.ByName("Get")),
			connect.WithClientOptions(opts...),
		),
		allocate: connect.NewClient[v1.IPServiceAllocateRequest, v1.IPServiceAllocateResponse](
			httpClient,
			baseURL+IPServiceAllocateProcedure,
			connect.WithSchema(iPServiceMethods.ByName("Allocate")),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[v1.IPServiceUpdateRequest, v1.IPServiceUpdateResponse](
			httpClient,
			baseURL+IPServiceUpdateProcedure,
			connect.WithSchema(iPServiceMethods.ByName("Update")),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[v1.IPServiceListRequest, v1.IPServiceListResponse](
			httpClient,
			baseURL+IPServiceListProcedure,
			connect.WithSchema(iPServiceMethods.ByName("List")),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[v1.IPServiceDeleteRequest, v1.IPServiceDeleteResponse](
			httpClient,
			baseURL+IPServiceDeleteProcedure,
			connect.WithSchema(iPServiceMethods.ByName("Delete")),
			connect.WithClientOptions(opts...),
		),
	}
}

// iPServiceClient implements IPServiceClient.
type iPServiceClient struct {
	get      *connect.Client[v1.IPServiceGetRequest, v1.IPServiceGetResponse]
	allocate *connect.Client[v1.IPServiceAllocateRequest, v1.IPServiceAllocateResponse]
	update   *connect.Client[v1.IPServiceUpdateRequest, v1.IPServiceUpdateResponse]
	list     *connect.Client[v1.IPServiceListRequest, v1.IPServiceListResponse]
	delete   *connect.Client[v1.IPServiceDeleteRequest, v1.IPServiceDeleteResponse]
}

// Get calls api.v1.IPService.Get.
func (c *iPServiceClient) Get(ctx context.Context, req *connect.Request[v1.IPServiceGetRequest]) (*connect.Response[v1.IPServiceGetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// Allocate calls api.v1.IPService.Allocate.
func (c *iPServiceClient) Allocate(ctx context.Context, req *connect.Request[v1.IPServiceAllocateRequest]) (*connect.Response[v1.IPServiceAllocateResponse], error) {
	return c.allocate.CallUnary(ctx, req)
}

// Update calls api.v1.IPService.Update.
func (c *iPServiceClient) Update(ctx context.Context, req *connect.Request[v1.IPServiceUpdateRequest]) (*connect.Response[v1.IPServiceUpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// List calls api.v1.IPService.List.
func (c *iPServiceClient) List(ctx context.Context, req *connect.Request[v1.IPServiceListRequest]) (*connect.Response[v1.IPServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Delete calls api.v1.IPService.Delete.
func (c *iPServiceClient) Delete(ctx context.Context, req *connect.Request[v1.IPServiceDeleteRequest]) (*connect.Response[v1.IPServiceDeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// IPServiceHandler is an implementation of the api.v1.IPService service.
type IPServiceHandler interface {
	// Get a ip
	Get(context.Context, *connect.Request[v1.IPServiceGetRequest]) (*connect.Response[v1.IPServiceGetResponse], error)
	// Allocate a ip
	Allocate(context.Context, *connect.Request[v1.IPServiceAllocateRequest]) (*connect.Response[v1.IPServiceAllocateResponse], error)
	// Update a ip
	Update(context.Context, *connect.Request[v1.IPServiceUpdateRequest]) (*connect.Response[v1.IPServiceUpdateResponse], error)
	// List all ips
	List(context.Context, *connect.Request[v1.IPServiceListRequest]) (*connect.Response[v1.IPServiceListResponse], error)
	// Delete a ip
	Delete(context.Context, *connect.Request[v1.IPServiceDeleteRequest]) (*connect.Response[v1.IPServiceDeleteResponse], error)
}

// NewIPServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIPServiceHandler(svc IPServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	iPServiceMethods := v1.File_api_v1_ip_proto.Services().ByName("IPService").Methods()
	iPServiceGetHandler := connect.NewUnaryHandler(
		IPServiceGetProcedure,
		svc.Get,
		connect.WithSchema(iPServiceMethods.ByName("Get")),
		connect.WithHandlerOptions(opts...),
	)
	iPServiceAllocateHandler := connect.NewUnaryHandler(
		IPServiceAllocateProcedure,
		svc.Allocate,
		connect.WithSchema(iPServiceMethods.ByName("Allocate")),
		connect.WithHandlerOptions(opts...),
	)
	iPServiceUpdateHandler := connect.NewUnaryHandler(
		IPServiceUpdateProcedure,
		svc.Update,
		connect.WithSchema(iPServiceMethods.ByName("Update")),
		connect.WithHandlerOptions(opts...),
	)
	iPServiceListHandler := connect.NewUnaryHandler(
		IPServiceListProcedure,
		svc.List,
		connect.WithSchema(iPServiceMethods.ByName("List")),
		connect.WithHandlerOptions(opts...),
	)
	iPServiceDeleteHandler := connect.NewUnaryHandler(
		IPServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(iPServiceMethods.ByName("Delete")),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.IPService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IPServiceGetProcedure:
			iPServiceGetHandler.ServeHTTP(w, r)
		case IPServiceAllocateProcedure:
			iPServiceAllocateHandler.ServeHTTP(w, r)
		case IPServiceUpdateProcedure:
			iPServiceUpdateHandler.ServeHTTP(w, r)
		case IPServiceListProcedure:
			iPServiceListHandler.ServeHTTP(w, r)
		case IPServiceDeleteProcedure:
			iPServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIPServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIPServiceHandler struct{}

func (UnimplementedIPServiceHandler) Get(context.Context, *connect.Request[v1.IPServiceGetRequest]) (*connect.Response[v1.IPServiceGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.IPService.Get is not implemented"))
}

func (UnimplementedIPServiceHandler) Allocate(context.Context, *connect.Request[v1.IPServiceAllocateRequest]) (*connect.Response[v1.IPServiceAllocateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.IPService.Allocate is not implemented"))
}

func (UnimplementedIPServiceHandler) Update(context.Context, *connect.Request[v1.IPServiceUpdateRequest]) (*connect.Response[v1.IPServiceUpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.IPService.Update is not implemented"))
}

func (UnimplementedIPServiceHandler) List(context.Context, *connect.Request[v1.IPServiceListRequest]) (*connect.Response[v1.IPServiceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.IPService.List is not implemented"))
}

func (UnimplementedIPServiceHandler) Delete(context.Context, *connect.Request[v1.IPServiceDeleteRequest]) (*connect.Response[v1.IPServiceDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.IPService.Delete is not implemented"))
}
