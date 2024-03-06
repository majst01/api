// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/project.proto

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
	// ProjectServiceName is the fully-qualified name of the ProjectService service.
	ProjectServiceName = "api.v1.ProjectService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProjectServiceListProcedure is the fully-qualified name of the ProjectService's List RPC.
	ProjectServiceListProcedure = "/api.v1.ProjectService/List"
	// ProjectServiceGetProcedure is the fully-qualified name of the ProjectService's Get RPC.
	ProjectServiceGetProcedure = "/api.v1.ProjectService/Get"
	// ProjectServiceCreateProcedure is the fully-qualified name of the ProjectService's Create RPC.
	ProjectServiceCreateProcedure = "/api.v1.ProjectService/Create"
	// ProjectServiceDeleteProcedure is the fully-qualified name of the ProjectService's Delete RPC.
	ProjectServiceDeleteProcedure = "/api.v1.ProjectService/Delete"
	// ProjectServiceUpdateProcedure is the fully-qualified name of the ProjectService's Update RPC.
	ProjectServiceUpdateProcedure = "/api.v1.ProjectService/Update"
	// ProjectServiceRemoveMemberProcedure is the fully-qualified name of the ProjectService's
	// RemoveMember RPC.
	ProjectServiceRemoveMemberProcedure = "/api.v1.ProjectService/RemoveMember"
	// ProjectServiceUpdateMemberProcedure is the fully-qualified name of the ProjectService's
	// UpdateMember RPC.
	ProjectServiceUpdateMemberProcedure = "/api.v1.ProjectService/UpdateMember"
	// ProjectServiceInviteProcedure is the fully-qualified name of the ProjectService's Invite RPC.
	ProjectServiceInviteProcedure = "/api.v1.ProjectService/Invite"
	// ProjectServiceInviteAcceptProcedure is the fully-qualified name of the ProjectService's
	// InviteAccept RPC.
	ProjectServiceInviteAcceptProcedure = "/api.v1.ProjectService/InviteAccept"
	// ProjectServiceInviteDeleteProcedure is the fully-qualified name of the ProjectService's
	// InviteDelete RPC.
	ProjectServiceInviteDeleteProcedure = "/api.v1.ProjectService/InviteDelete"
	// ProjectServiceInvitesListProcedure is the fully-qualified name of the ProjectService's
	// InvitesList RPC.
	ProjectServiceInvitesListProcedure = "/api.v1.ProjectService/InvitesList"
	// ProjectServiceInviteGetProcedure is the fully-qualified name of the ProjectService's InviteGet
	// RPC.
	ProjectServiceInviteGetProcedure = "/api.v1.ProjectService/InviteGet"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	projectServiceServiceDescriptor            = v1.File_api_v1_project_proto.Services().ByName("ProjectService")
	projectServiceListMethodDescriptor         = projectServiceServiceDescriptor.Methods().ByName("List")
	projectServiceGetMethodDescriptor          = projectServiceServiceDescriptor.Methods().ByName("Get")
	projectServiceCreateMethodDescriptor       = projectServiceServiceDescriptor.Methods().ByName("Create")
	projectServiceDeleteMethodDescriptor       = projectServiceServiceDescriptor.Methods().ByName("Delete")
	projectServiceUpdateMethodDescriptor       = projectServiceServiceDescriptor.Methods().ByName("Update")
	projectServiceRemoveMemberMethodDescriptor = projectServiceServiceDescriptor.Methods().ByName("RemoveMember")
	projectServiceUpdateMemberMethodDescriptor = projectServiceServiceDescriptor.Methods().ByName("UpdateMember")
	projectServiceInviteMethodDescriptor       = projectServiceServiceDescriptor.Methods().ByName("Invite")
	projectServiceInviteAcceptMethodDescriptor = projectServiceServiceDescriptor.Methods().ByName("InviteAccept")
	projectServiceInviteDeleteMethodDescriptor = projectServiceServiceDescriptor.Methods().ByName("InviteDelete")
	projectServiceInvitesListMethodDescriptor  = projectServiceServiceDescriptor.Methods().ByName("InvitesList")
	projectServiceInviteGetMethodDescriptor    = projectServiceServiceDescriptor.Methods().ByName("InviteGet")
)

// ProjectServiceClient is a client for the api.v1.ProjectService service.
type ProjectServiceClient interface {
	// List all accessible projects
	List(context.Context, *connect.Request[v1.ProjectServiceListRequest]) (*connect.Response[v1.ProjectServiceListResponse], error)
	// Get a project
	Get(context.Context, *connect.Request[v1.ProjectServiceGetRequest]) (*connect.Response[v1.ProjectServiceGetResponse], error)
	// Create a project
	Create(context.Context, *connect.Request[v1.ProjectServiceCreateRequest]) (*connect.Response[v1.ProjectServiceCreateResponse], error)
	// Delete a project
	Delete(context.Context, *connect.Request[v1.ProjectServiceDeleteRequest]) (*connect.Response[v1.ProjectServiceDeleteResponse], error)
	// Update a project
	Update(context.Context, *connect.Request[v1.ProjectServiceUpdateRequest]) (*connect.Response[v1.ProjectServiceUpdateResponse], error)
	// RemoveMember remove a user from a project
	RemoveMember(context.Context, *connect.Request[v1.ProjectServiceRemoveMemberRequest]) (*connect.Response[v1.ProjectServiceRemoveMemberResponse], error)
	// UpdateMember update a user for a project
	UpdateMember(context.Context, *connect.Request[v1.ProjectServiceUpdateMemberRequest]) (*connect.Response[v1.ProjectServiceUpdateMemberResponse], error)
	// Invite a user to a project
	Invite(context.Context, *connect.Request[v1.ProjectServiceInviteRequest]) (*connect.Response[v1.ProjectServiceInviteResponse], error)
	// InviteAccept is called from a user to accept a invitation
	InviteAccept(context.Context, *connect.Request[v1.ProjectServiceInviteAcceptRequest]) (*connect.Response[v1.ProjectServiceInviteAcceptResponse], error)
	// InviteDelete deletes a pending invitation
	InviteDelete(context.Context, *connect.Request[v1.ProjectServiceInviteDeleteRequest]) (*connect.Response[v1.ProjectServiceInviteDeleteResponse], error)
	// InvitesList list all invites to a project
	InvitesList(context.Context, *connect.Request[v1.ProjectServiceInvitesListRequest]) (*connect.Response[v1.ProjectServiceInvitesListResponse], error)
	// InviteGet get an invite
	InviteGet(context.Context, *connect.Request[v1.ProjectServiceInviteGetRequest]) (*connect.Response[v1.ProjectServiceInviteGetResponse], error)
}

// NewProjectServiceClient constructs a client for the api.v1.ProjectService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProjectServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProjectServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &projectServiceClient{
		list: connect.NewClient[v1.ProjectServiceListRequest, v1.ProjectServiceListResponse](
			httpClient,
			baseURL+ProjectServiceListProcedure,
			connect.WithSchema(projectServiceListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[v1.ProjectServiceGetRequest, v1.ProjectServiceGetResponse](
			httpClient,
			baseURL+ProjectServiceGetProcedure,
			connect.WithSchema(projectServiceGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		create: connect.NewClient[v1.ProjectServiceCreateRequest, v1.ProjectServiceCreateResponse](
			httpClient,
			baseURL+ProjectServiceCreateProcedure,
			connect.WithSchema(projectServiceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[v1.ProjectServiceDeleteRequest, v1.ProjectServiceDeleteResponse](
			httpClient,
			baseURL+ProjectServiceDeleteProcedure,
			connect.WithSchema(projectServiceDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[v1.ProjectServiceUpdateRequest, v1.ProjectServiceUpdateResponse](
			httpClient,
			baseURL+ProjectServiceUpdateProcedure,
			connect.WithSchema(projectServiceUpdateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		removeMember: connect.NewClient[v1.ProjectServiceRemoveMemberRequest, v1.ProjectServiceRemoveMemberResponse](
			httpClient,
			baseURL+ProjectServiceRemoveMemberProcedure,
			connect.WithSchema(projectServiceRemoveMemberMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateMember: connect.NewClient[v1.ProjectServiceUpdateMemberRequest, v1.ProjectServiceUpdateMemberResponse](
			httpClient,
			baseURL+ProjectServiceUpdateMemberProcedure,
			connect.WithSchema(projectServiceUpdateMemberMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		invite: connect.NewClient[v1.ProjectServiceInviteRequest, v1.ProjectServiceInviteResponse](
			httpClient,
			baseURL+ProjectServiceInviteProcedure,
			connect.WithSchema(projectServiceInviteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		inviteAccept: connect.NewClient[v1.ProjectServiceInviteAcceptRequest, v1.ProjectServiceInviteAcceptResponse](
			httpClient,
			baseURL+ProjectServiceInviteAcceptProcedure,
			connect.WithSchema(projectServiceInviteAcceptMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		inviteDelete: connect.NewClient[v1.ProjectServiceInviteDeleteRequest, v1.ProjectServiceInviteDeleteResponse](
			httpClient,
			baseURL+ProjectServiceInviteDeleteProcedure,
			connect.WithSchema(projectServiceInviteDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		invitesList: connect.NewClient[v1.ProjectServiceInvitesListRequest, v1.ProjectServiceInvitesListResponse](
			httpClient,
			baseURL+ProjectServiceInvitesListProcedure,
			connect.WithSchema(projectServiceInvitesListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		inviteGet: connect.NewClient[v1.ProjectServiceInviteGetRequest, v1.ProjectServiceInviteGetResponse](
			httpClient,
			baseURL+ProjectServiceInviteGetProcedure,
			connect.WithSchema(projectServiceInviteGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// projectServiceClient implements ProjectServiceClient.
type projectServiceClient struct {
	list         *connect.Client[v1.ProjectServiceListRequest, v1.ProjectServiceListResponse]
	get          *connect.Client[v1.ProjectServiceGetRequest, v1.ProjectServiceGetResponse]
	create       *connect.Client[v1.ProjectServiceCreateRequest, v1.ProjectServiceCreateResponse]
	delete       *connect.Client[v1.ProjectServiceDeleteRequest, v1.ProjectServiceDeleteResponse]
	update       *connect.Client[v1.ProjectServiceUpdateRequest, v1.ProjectServiceUpdateResponse]
	removeMember *connect.Client[v1.ProjectServiceRemoveMemberRequest, v1.ProjectServiceRemoveMemberResponse]
	updateMember *connect.Client[v1.ProjectServiceUpdateMemberRequest, v1.ProjectServiceUpdateMemberResponse]
	invite       *connect.Client[v1.ProjectServiceInviteRequest, v1.ProjectServiceInviteResponse]
	inviteAccept *connect.Client[v1.ProjectServiceInviteAcceptRequest, v1.ProjectServiceInviteAcceptResponse]
	inviteDelete *connect.Client[v1.ProjectServiceInviteDeleteRequest, v1.ProjectServiceInviteDeleteResponse]
	invitesList  *connect.Client[v1.ProjectServiceInvitesListRequest, v1.ProjectServiceInvitesListResponse]
	inviteGet    *connect.Client[v1.ProjectServiceInviteGetRequest, v1.ProjectServiceInviteGetResponse]
}

// List calls api.v1.ProjectService.List.
func (c *projectServiceClient) List(ctx context.Context, req *connect.Request[v1.ProjectServiceListRequest]) (*connect.Response[v1.ProjectServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Get calls api.v1.ProjectService.Get.
func (c *projectServiceClient) Get(ctx context.Context, req *connect.Request[v1.ProjectServiceGetRequest]) (*connect.Response[v1.ProjectServiceGetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// Create calls api.v1.ProjectService.Create.
func (c *projectServiceClient) Create(ctx context.Context, req *connect.Request[v1.ProjectServiceCreateRequest]) (*connect.Response[v1.ProjectServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Delete calls api.v1.ProjectService.Delete.
func (c *projectServiceClient) Delete(ctx context.Context, req *connect.Request[v1.ProjectServiceDeleteRequest]) (*connect.Response[v1.ProjectServiceDeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// Update calls api.v1.ProjectService.Update.
func (c *projectServiceClient) Update(ctx context.Context, req *connect.Request[v1.ProjectServiceUpdateRequest]) (*connect.Response[v1.ProjectServiceUpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// RemoveMember calls api.v1.ProjectService.RemoveMember.
func (c *projectServiceClient) RemoveMember(ctx context.Context, req *connect.Request[v1.ProjectServiceRemoveMemberRequest]) (*connect.Response[v1.ProjectServiceRemoveMemberResponse], error) {
	return c.removeMember.CallUnary(ctx, req)
}

// UpdateMember calls api.v1.ProjectService.UpdateMember.
func (c *projectServiceClient) UpdateMember(ctx context.Context, req *connect.Request[v1.ProjectServiceUpdateMemberRequest]) (*connect.Response[v1.ProjectServiceUpdateMemberResponse], error) {
	return c.updateMember.CallUnary(ctx, req)
}

// Invite calls api.v1.ProjectService.Invite.
func (c *projectServiceClient) Invite(ctx context.Context, req *connect.Request[v1.ProjectServiceInviteRequest]) (*connect.Response[v1.ProjectServiceInviteResponse], error) {
	return c.invite.CallUnary(ctx, req)
}

// InviteAccept calls api.v1.ProjectService.InviteAccept.
func (c *projectServiceClient) InviteAccept(ctx context.Context, req *connect.Request[v1.ProjectServiceInviteAcceptRequest]) (*connect.Response[v1.ProjectServiceInviteAcceptResponse], error) {
	return c.inviteAccept.CallUnary(ctx, req)
}

// InviteDelete calls api.v1.ProjectService.InviteDelete.
func (c *projectServiceClient) InviteDelete(ctx context.Context, req *connect.Request[v1.ProjectServiceInviteDeleteRequest]) (*connect.Response[v1.ProjectServiceInviteDeleteResponse], error) {
	return c.inviteDelete.CallUnary(ctx, req)
}

// InvitesList calls api.v1.ProjectService.InvitesList.
func (c *projectServiceClient) InvitesList(ctx context.Context, req *connect.Request[v1.ProjectServiceInvitesListRequest]) (*connect.Response[v1.ProjectServiceInvitesListResponse], error) {
	return c.invitesList.CallUnary(ctx, req)
}

// InviteGet calls api.v1.ProjectService.InviteGet.
func (c *projectServiceClient) InviteGet(ctx context.Context, req *connect.Request[v1.ProjectServiceInviteGetRequest]) (*connect.Response[v1.ProjectServiceInviteGetResponse], error) {
	return c.inviteGet.CallUnary(ctx, req)
}

// ProjectServiceHandler is an implementation of the api.v1.ProjectService service.
type ProjectServiceHandler interface {
	// List all accessible projects
	List(context.Context, *connect.Request[v1.ProjectServiceListRequest]) (*connect.Response[v1.ProjectServiceListResponse], error)
	// Get a project
	Get(context.Context, *connect.Request[v1.ProjectServiceGetRequest]) (*connect.Response[v1.ProjectServiceGetResponse], error)
	// Create a project
	Create(context.Context, *connect.Request[v1.ProjectServiceCreateRequest]) (*connect.Response[v1.ProjectServiceCreateResponse], error)
	// Delete a project
	Delete(context.Context, *connect.Request[v1.ProjectServiceDeleteRequest]) (*connect.Response[v1.ProjectServiceDeleteResponse], error)
	// Update a project
	Update(context.Context, *connect.Request[v1.ProjectServiceUpdateRequest]) (*connect.Response[v1.ProjectServiceUpdateResponse], error)
	// RemoveMember remove a user from a project
	RemoveMember(context.Context, *connect.Request[v1.ProjectServiceRemoveMemberRequest]) (*connect.Response[v1.ProjectServiceRemoveMemberResponse], error)
	// UpdateMember update a user for a project
	UpdateMember(context.Context, *connect.Request[v1.ProjectServiceUpdateMemberRequest]) (*connect.Response[v1.ProjectServiceUpdateMemberResponse], error)
	// Invite a user to a project
	Invite(context.Context, *connect.Request[v1.ProjectServiceInviteRequest]) (*connect.Response[v1.ProjectServiceInviteResponse], error)
	// InviteAccept is called from a user to accept a invitation
	InviteAccept(context.Context, *connect.Request[v1.ProjectServiceInviteAcceptRequest]) (*connect.Response[v1.ProjectServiceInviteAcceptResponse], error)
	// InviteDelete deletes a pending invitation
	InviteDelete(context.Context, *connect.Request[v1.ProjectServiceInviteDeleteRequest]) (*connect.Response[v1.ProjectServiceInviteDeleteResponse], error)
	// InvitesList list all invites to a project
	InvitesList(context.Context, *connect.Request[v1.ProjectServiceInvitesListRequest]) (*connect.Response[v1.ProjectServiceInvitesListResponse], error)
	// InviteGet get an invite
	InviteGet(context.Context, *connect.Request[v1.ProjectServiceInviteGetRequest]) (*connect.Response[v1.ProjectServiceInviteGetResponse], error)
}

// NewProjectServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProjectServiceHandler(svc ProjectServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	projectServiceListHandler := connect.NewUnaryHandler(
		ProjectServiceListProcedure,
		svc.List,
		connect.WithSchema(projectServiceListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceGetHandler := connect.NewUnaryHandler(
		ProjectServiceGetProcedure,
		svc.Get,
		connect.WithSchema(projectServiceGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceCreateHandler := connect.NewUnaryHandler(
		ProjectServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(projectServiceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceDeleteHandler := connect.NewUnaryHandler(
		ProjectServiceDeleteProcedure,
		svc.Delete,
		connect.WithSchema(projectServiceDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceUpdateHandler := connect.NewUnaryHandler(
		ProjectServiceUpdateProcedure,
		svc.Update,
		connect.WithSchema(projectServiceUpdateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceRemoveMemberHandler := connect.NewUnaryHandler(
		ProjectServiceRemoveMemberProcedure,
		svc.RemoveMember,
		connect.WithSchema(projectServiceRemoveMemberMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceUpdateMemberHandler := connect.NewUnaryHandler(
		ProjectServiceUpdateMemberProcedure,
		svc.UpdateMember,
		connect.WithSchema(projectServiceUpdateMemberMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInviteHandler := connect.NewUnaryHandler(
		ProjectServiceInviteProcedure,
		svc.Invite,
		connect.WithSchema(projectServiceInviteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInviteAcceptHandler := connect.NewUnaryHandler(
		ProjectServiceInviteAcceptProcedure,
		svc.InviteAccept,
		connect.WithSchema(projectServiceInviteAcceptMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInviteDeleteHandler := connect.NewUnaryHandler(
		ProjectServiceInviteDeleteProcedure,
		svc.InviteDelete,
		connect.WithSchema(projectServiceInviteDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInvitesListHandler := connect.NewUnaryHandler(
		ProjectServiceInvitesListProcedure,
		svc.InvitesList,
		connect.WithSchema(projectServiceInvitesListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectServiceInviteGetHandler := connect.NewUnaryHandler(
		ProjectServiceInviteGetProcedure,
		svc.InviteGet,
		connect.WithSchema(projectServiceInviteGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.ProjectService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProjectServiceListProcedure:
			projectServiceListHandler.ServeHTTP(w, r)
		case ProjectServiceGetProcedure:
			projectServiceGetHandler.ServeHTTP(w, r)
		case ProjectServiceCreateProcedure:
			projectServiceCreateHandler.ServeHTTP(w, r)
		case ProjectServiceDeleteProcedure:
			projectServiceDeleteHandler.ServeHTTP(w, r)
		case ProjectServiceUpdateProcedure:
			projectServiceUpdateHandler.ServeHTTP(w, r)
		case ProjectServiceRemoveMemberProcedure:
			projectServiceRemoveMemberHandler.ServeHTTP(w, r)
		case ProjectServiceUpdateMemberProcedure:
			projectServiceUpdateMemberHandler.ServeHTTP(w, r)
		case ProjectServiceInviteProcedure:
			projectServiceInviteHandler.ServeHTTP(w, r)
		case ProjectServiceInviteAcceptProcedure:
			projectServiceInviteAcceptHandler.ServeHTTP(w, r)
		case ProjectServiceInviteDeleteProcedure:
			projectServiceInviteDeleteHandler.ServeHTTP(w, r)
		case ProjectServiceInvitesListProcedure:
			projectServiceInvitesListHandler.ServeHTTP(w, r)
		case ProjectServiceInviteGetProcedure:
			projectServiceInviteGetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProjectServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProjectServiceHandler struct{}

func (UnimplementedProjectServiceHandler) List(context.Context, *connect.Request[v1.ProjectServiceListRequest]) (*connect.Response[v1.ProjectServiceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.List is not implemented"))
}

func (UnimplementedProjectServiceHandler) Get(context.Context, *connect.Request[v1.ProjectServiceGetRequest]) (*connect.Response[v1.ProjectServiceGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.Get is not implemented"))
}

func (UnimplementedProjectServiceHandler) Create(context.Context, *connect.Request[v1.ProjectServiceCreateRequest]) (*connect.Response[v1.ProjectServiceCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.Create is not implemented"))
}

func (UnimplementedProjectServiceHandler) Delete(context.Context, *connect.Request[v1.ProjectServiceDeleteRequest]) (*connect.Response[v1.ProjectServiceDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.Delete is not implemented"))
}

func (UnimplementedProjectServiceHandler) Update(context.Context, *connect.Request[v1.ProjectServiceUpdateRequest]) (*connect.Response[v1.ProjectServiceUpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.Update is not implemented"))
}

func (UnimplementedProjectServiceHandler) RemoveMember(context.Context, *connect.Request[v1.ProjectServiceRemoveMemberRequest]) (*connect.Response[v1.ProjectServiceRemoveMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.RemoveMember is not implemented"))
}

func (UnimplementedProjectServiceHandler) UpdateMember(context.Context, *connect.Request[v1.ProjectServiceUpdateMemberRequest]) (*connect.Response[v1.ProjectServiceUpdateMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.UpdateMember is not implemented"))
}

func (UnimplementedProjectServiceHandler) Invite(context.Context, *connect.Request[v1.ProjectServiceInviteRequest]) (*connect.Response[v1.ProjectServiceInviteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.Invite is not implemented"))
}

func (UnimplementedProjectServiceHandler) InviteAccept(context.Context, *connect.Request[v1.ProjectServiceInviteAcceptRequest]) (*connect.Response[v1.ProjectServiceInviteAcceptResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.InviteAccept is not implemented"))
}

func (UnimplementedProjectServiceHandler) InviteDelete(context.Context, *connect.Request[v1.ProjectServiceInviteDeleteRequest]) (*connect.Response[v1.ProjectServiceInviteDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.InviteDelete is not implemented"))
}

func (UnimplementedProjectServiceHandler) InvitesList(context.Context, *connect.Request[v1.ProjectServiceInvitesListRequest]) (*connect.Response[v1.ProjectServiceInvitesListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.InvitesList is not implemented"))
}

func (UnimplementedProjectServiceHandler) InviteGet(context.Context, *connect.Request[v1.ProjectServiceInviteGetRequest]) (*connect.Response[v1.ProjectServiceInviteGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.ProjectService.InviteGet is not implemented"))
}