// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: api/v1/methods.proto

package apiv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MethodServiceListRequest is the request payload to list all public methods
type MethodServiceListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MethodServiceListRequest) Reset() {
	*x = MethodServiceListRequest{}
	mi := &file_api_v1_methods_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MethodServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodServiceListRequest) ProtoMessage() {}

func (x *MethodServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_methods_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodServiceListRequest.ProtoReflect.Descriptor instead.
func (*MethodServiceListRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_methods_proto_rawDescGZIP(), []int{0}
}

// MethodServiceListResponse is the response payload with all public visible methods
type MethodServiceListResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Methods is a list of methods public callable
	Methods       []string `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MethodServiceListResponse) Reset() {
	*x = MethodServiceListResponse{}
	mi := &file_api_v1_methods_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MethodServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodServiceListResponse) ProtoMessage() {}

func (x *MethodServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_methods_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodServiceListResponse.ProtoReflect.Descriptor instead.
func (*MethodServiceListResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_methods_proto_rawDescGZIP(), []int{1}
}

func (x *MethodServiceListResponse) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

// MethodServiceTokenScopedListRequest is the request payload to list all methods callable with the token present in the request
type MethodServiceTokenScopedListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MethodServiceTokenScopedListRequest) Reset() {
	*x = MethodServiceTokenScopedListRequest{}
	mi := &file_api_v1_methods_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MethodServiceTokenScopedListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodServiceTokenScopedListRequest) ProtoMessage() {}

func (x *MethodServiceTokenScopedListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_methods_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodServiceTokenScopedListRequest.ProtoReflect.Descriptor instead.
func (*MethodServiceTokenScopedListRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_methods_proto_rawDescGZIP(), []int{2}
}

// MethodServiceTokenScopedListResponse is the response payload which contains all methods which are callable with the given token
type MethodServiceTokenScopedListResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Permissions a list of methods which can be called
	Permissions []*MethodPermission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// ProjectRoles associates a project id with the corresponding role of the token owner
	ProjectRoles map[string]ProjectRole `protobuf:"bytes,3,rep,name=project_roles,json=projectRoles,proto3" json:"project_roles,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=api.v1.ProjectRole"`
	// TenantRoles associates a tenant id with the corresponding role of the token owner
	TenantRoles map[string]TenantRole `protobuf:"bytes,4,rep,name=tenant_roles,json=tenantRoles,proto3" json:"tenant_roles,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=api.v1.TenantRole"`
	// AdminRole defines the admin role of the token owner
	AdminRole     *AdminRole `protobuf:"varint,5,opt,name=admin_role,json=adminRole,proto3,enum=api.v1.AdminRole,oneof" json:"admin_role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MethodServiceTokenScopedListResponse) Reset() {
	*x = MethodServiceTokenScopedListResponse{}
	mi := &file_api_v1_methods_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MethodServiceTokenScopedListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodServiceTokenScopedListResponse) ProtoMessage() {}

func (x *MethodServiceTokenScopedListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_methods_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodServiceTokenScopedListResponse.ProtoReflect.Descriptor instead.
func (*MethodServiceTokenScopedListResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_methods_proto_rawDescGZIP(), []int{3}
}

func (x *MethodServiceTokenScopedListResponse) GetPermissions() []*MethodPermission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *MethodServiceTokenScopedListResponse) GetProjectRoles() map[string]ProjectRole {
	if x != nil {
		return x.ProjectRoles
	}
	return nil
}

func (x *MethodServiceTokenScopedListResponse) GetTenantRoles() map[string]TenantRole {
	if x != nil {
		return x.TenantRoles
	}
	return nil
}

func (x *MethodServiceTokenScopedListResponse) GetAdminRole() AdminRole {
	if x != nil && x.AdminRole != nil {
		return *x.AdminRole
	}
	return AdminRole_ADMIN_ROLE_UNSPECIFIED
}

var File_api_v1_methods_proto protoreflect.FileDescriptor

var file_api_v1_methods_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x13,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x35, 0x0a, 0x19, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22, 0x25, 0x0a, 0x23, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa3,
	0x04, 0x0a, 0x24, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x63, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x60, 0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c,
	0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x09, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x1a, 0x54, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x52, 0x0a, 0x10, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x32, 0xd6, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x04, 0xd8, 0xf3, 0x18, 0x01, 0x12, 0x72, 0x0a, 0x0f, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0xd8, 0xf3, 0x18, 0x03, 0x42, 0x7f, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2d, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x06,
	0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_v1_methods_proto_rawDescOnce sync.Once
	file_api_v1_methods_proto_rawDescData []byte
)

func file_api_v1_methods_proto_rawDescGZIP() []byte {
	file_api_v1_methods_proto_rawDescOnce.Do(func() {
		file_api_v1_methods_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_methods_proto_rawDesc), len(file_api_v1_methods_proto_rawDesc)))
	})
	return file_api_v1_methods_proto_rawDescData
}

var file_api_v1_methods_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1_methods_proto_goTypes = []any{
	(*MethodServiceListRequest)(nil),             // 0: api.v1.MethodServiceListRequest
	(*MethodServiceListResponse)(nil),            // 1: api.v1.MethodServiceListResponse
	(*MethodServiceTokenScopedListRequest)(nil),  // 2: api.v1.MethodServiceTokenScopedListRequest
	(*MethodServiceTokenScopedListResponse)(nil), // 3: api.v1.MethodServiceTokenScopedListResponse
	nil,                      // 4: api.v1.MethodServiceTokenScopedListResponse.ProjectRolesEntry
	nil,                      // 5: api.v1.MethodServiceTokenScopedListResponse.TenantRolesEntry
	(*MethodPermission)(nil), // 6: api.v1.MethodPermission
	(AdminRole)(0),           // 7: api.v1.AdminRole
	(ProjectRole)(0),         // 8: api.v1.ProjectRole
	(TenantRole)(0),          // 9: api.v1.TenantRole
}
var file_api_v1_methods_proto_depIdxs = []int32{
	6, // 0: api.v1.MethodServiceTokenScopedListResponse.permissions:type_name -> api.v1.MethodPermission
	4, // 1: api.v1.MethodServiceTokenScopedListResponse.project_roles:type_name -> api.v1.MethodServiceTokenScopedListResponse.ProjectRolesEntry
	5, // 2: api.v1.MethodServiceTokenScopedListResponse.tenant_roles:type_name -> api.v1.MethodServiceTokenScopedListResponse.TenantRolesEntry
	7, // 3: api.v1.MethodServiceTokenScopedListResponse.admin_role:type_name -> api.v1.AdminRole
	8, // 4: api.v1.MethodServiceTokenScopedListResponse.ProjectRolesEntry.value:type_name -> api.v1.ProjectRole
	9, // 5: api.v1.MethodServiceTokenScopedListResponse.TenantRolesEntry.value:type_name -> api.v1.TenantRole
	0, // 6: api.v1.MethodService.List:input_type -> api.v1.MethodServiceListRequest
	2, // 7: api.v1.MethodService.TokenScopedList:input_type -> api.v1.MethodServiceTokenScopedListRequest
	1, // 8: api.v1.MethodService.List:output_type -> api.v1.MethodServiceListResponse
	3, // 9: api.v1.MethodService.TokenScopedList:output_type -> api.v1.MethodServiceTokenScopedListResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1_methods_proto_init() }
func file_api_v1_methods_proto_init() {
	if File_api_v1_methods_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_token_proto_init()
	file_api_v1_methods_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_methods_proto_rawDesc), len(file_api_v1_methods_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_methods_proto_goTypes,
		DependencyIndexes: file_api_v1_methods_proto_depIdxs,
		MessageInfos:      file_api_v1_methods_proto_msgTypes,
	}.Build()
	File_api_v1_methods_proto = out.File
	file_api_v1_methods_proto_goTypes = nil
	file_api_v1_methods_proto_depIdxs = nil
}
