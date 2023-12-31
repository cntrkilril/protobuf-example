// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: role_service_v1.proto

package role_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// roles
type RoleEntityV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Alias       string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RoleEntityV1) Reset() {
	*x = RoleEntityV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleEntityV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleEntityV1) ProtoMessage() {}

func (x *RoleEntityV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleEntityV1.ProtoReflect.Descriptor instead.
func (*RoleEntityV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{0}
}

func (x *RoleEntityV1) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RoleEntityV1) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RoleEntityV1) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *RoleEntityV1) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AllRoleEntityV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []*RoleEntityV1 `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *AllRoleEntityV1) Reset() {
	*x = AllRoleEntityV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRoleEntityV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRoleEntityV1) ProtoMessage() {}

func (x *AllRoleEntityV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRoleEntityV1.ProtoReflect.Descriptor instead.
func (*AllRoleEntityV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{1}
}

func (x *AllRoleEntityV1) GetRoles() []*RoleEntityV1 {
	if x != nil {
		return x.Roles
	}
	return nil
}

type CreateRoleRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Alias       string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateRoleRequestV1) Reset() {
	*x = CreateRoleRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequestV1) ProtoMessage() {}

func (x *CreateRoleRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequestV1.ProtoReflect.Descriptor instead.
func (*CreateRoleRequestV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoleRequestV1) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRoleRequestV1) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *CreateRoleRequestV1) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetRoleByIDRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetRoleByIDRequestV1) Reset() {
	*x = GetRoleByIDRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleByIDRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleByIDRequestV1) ProtoMessage() {}

func (x *GetRoleByIDRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleByIDRequestV1.ProtoReflect.Descriptor instead.
func (*GetRoleByIDRequestV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoleByIDRequestV1) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteRoleRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteRoleRequestV1) Reset() {
	*x = DeleteRoleRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequestV1) ProtoMessage() {}

func (x *DeleteRoleRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequestV1.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequestV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRoleRequestV1) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

// check
type CheckAdminRoleRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminID string `protobuf:"bytes,1,opt,name=adminID,proto3" json:"adminID,omitempty"`
	Alias   string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *CheckAdminRoleRequestV1) Reset() {
	*x = CheckAdminRoleRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAdminRoleRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAdminRoleRequestV1) ProtoMessage() {}

func (x *CheckAdminRoleRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAdminRoleRequestV1.ProtoReflect.Descriptor instead.
func (*CheckAdminRoleRequestV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CheckAdminRoleRequestV1) GetAdminID() string {
	if x != nil {
		return x.AdminID
	}
	return ""
}

func (x *CheckAdminRoleRequestV1) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type CheckAdminRoleResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CheckAdminRoleResponseV1) Reset() {
	*x = CheckAdminRoleResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAdminRoleResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAdminRoleResponseV1) ProtoMessage() {}

func (x *CheckAdminRoleResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAdminRoleResponseV1.ProtoReflect.Descriptor instead.
func (*CheckAdminRoleResponseV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{6}
}

func (x *CheckAdminRoleResponseV1) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

// admins with alias
type CreateAdminRoleRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminID string `protobuf:"bytes,1,opt,name=adminID,proto3" json:"adminID,omitempty"`
	RoleID  int32  `protobuf:"varint,2,opt,name=roleID,proto3" json:"roleID,omitempty"`
}

func (x *CreateAdminRoleRequestV1) Reset() {
	*x = CreateAdminRoleRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminRoleRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminRoleRequestV1) ProtoMessage() {}

func (x *CreateAdminRoleRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminRoleRequestV1.ProtoReflect.Descriptor instead.
func (*CreateAdminRoleRequestV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{7}
}

func (x *CreateAdminRoleRequestV1) GetAdminID() string {
	if x != nil {
		return x.AdminID
	}
	return ""
}

func (x *CreateAdminRoleRequestV1) GetRoleID() int32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

type AdminRoleEntityV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	AdminID string `protobuf:"bytes,2,opt,name=adminID,proto3" json:"adminID,omitempty"`
	RoleID  int32  `protobuf:"varint,3,opt,name=roleID,proto3" json:"roleID,omitempty"`
}

func (x *AdminRoleEntityV1) Reset() {
	*x = AdminRoleEntityV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRoleEntityV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRoleEntityV1) ProtoMessage() {}

func (x *AdminRoleEntityV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRoleEntityV1.ProtoReflect.Descriptor instead.
func (*AdminRoleEntityV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{8}
}

func (x *AdminRoleEntityV1) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AdminRoleEntityV1) GetAdminID() string {
	if x != nil {
		return x.AdminID
	}
	return ""
}

func (x *AdminRoleEntityV1) GetRoleID() int32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

type AllAdminRoleEntityV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminRoles []*AdminRoleEntityV1 `protobuf:"bytes,1,rep,name=admin_roles,json=adminRoles,proto3" json:"admin_roles,omitempty"`
}

func (x *AllAdminRoleEntityV1) Reset() {
	*x = AllAdminRoleEntityV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllAdminRoleEntityV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllAdminRoleEntityV1) ProtoMessage() {}

func (x *AllAdminRoleEntityV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllAdminRoleEntityV1.ProtoReflect.Descriptor instead.
func (*AllAdminRoleEntityV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{9}
}

func (x *AllAdminRoleEntityV1) GetAdminRoles() []*AdminRoleEntityV1 {
	if x != nil {
		return x.AdminRoles
	}
	return nil
}

type DeleteAdminRoleRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminID string `protobuf:"bytes,1,opt,name=adminID,proto3" json:"adminID,omitempty"`
}

func (x *DeleteAdminRoleRequestV1) Reset() {
	*x = DeleteAdminRoleRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_service_v1_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminRoleRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminRoleRequestV1) ProtoMessage() {}

func (x *DeleteAdminRoleRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_role_service_v1_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminRoleRequestV1.ProtoReflect.Descriptor instead.
func (*DeleteAdminRoleRequestV1) Descriptor() ([]byte, []int) {
	return file_role_service_v1_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAdminRoleRequestV1) GetAdminID() string {
	if x != nil {
		return x.AdminID
	}
	return ""
}

var File_role_service_v1_proto protoreflect.FileDescriptor

var file_role_service_v1_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0c, 0x52, 0x6f,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x52,
	0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x12, 0x28, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x22, 0x49, 0x0a, 0x17, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x22, 0x2a, 0x0a, 0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x22, 0x4c, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x55,
	0x0a, 0x11, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x50, 0x0a, 0x14, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x12, 0x38, 0x0a,
	0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x52, 0x0a, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x56, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44, 0x32, 0xc7, 0x05,
	0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12,
	0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x19, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x41, 0x6c, 0x6c,
	0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31,
	0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x56, 0x31, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x1e, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x1e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a,
	0x17, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x1a, 0x17, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x31, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x31, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x56, 0x31, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_role_service_v1_proto_rawDescOnce sync.Once
	file_role_service_v1_proto_rawDescData = file_role_service_v1_proto_rawDesc
)

func file_role_service_v1_proto_rawDescGZIP() []byte {
	file_role_service_v1_proto_rawDescOnce.Do(func() {
		file_role_service_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_service_v1_proto_rawDescData)
	})
	return file_role_service_v1_proto_rawDescData
}

var file_role_service_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_role_service_v1_proto_goTypes = []interface{}{
	(*RoleEntityV1)(nil),             // 0: role.RoleEntityV1
	(*AllRoleEntityV1)(nil),          // 1: role.AllRoleEntityV1
	(*CreateRoleRequestV1)(nil),      // 2: role.CreateRoleRequestV1
	(*GetRoleByIDRequestV1)(nil),     // 3: role.GetRoleByIDRequestV1
	(*DeleteRoleRequestV1)(nil),      // 4: role.DeleteRoleRequestV1
	(*CheckAdminRoleRequestV1)(nil),  // 5: role.CheckAdminRoleRequestV1
	(*CheckAdminRoleResponseV1)(nil), // 6: role.CheckAdminRoleResponseV1
	(*CreateAdminRoleRequestV1)(nil), // 7: role.CreateAdminRoleRequestV1
	(*AdminRoleEntityV1)(nil),        // 8: role.AdminRoleEntityV1
	(*AllAdminRoleEntityV1)(nil),     // 9: role.AllAdminRoleEntityV1
	(*DeleteAdminRoleRequestV1)(nil), // 10: role.DeleteAdminRoleRequestV1
	(*emptypb.Empty)(nil),            // 11: google.protobuf.Empty
}
var file_role_service_v1_proto_depIdxs = []int32{
	0,  // 0: role.AllRoleEntityV1.roles:type_name -> role.RoleEntityV1
	8,  // 1: role.AllAdminRoleEntityV1.admin_roles:type_name -> role.AdminRoleEntityV1
	2,  // 2: role.RoleServiceV1.CreateRole:input_type -> role.CreateRoleRequestV1
	3,  // 3: role.RoleServiceV1.GetRoleByID:input_type -> role.GetRoleByIDRequestV1
	11, // 4: role.RoleServiceV1.GetAllRole:input_type -> google.protobuf.Empty
	0,  // 5: role.RoleServiceV1.UpdateRole:input_type -> role.RoleEntityV1
	4,  // 6: role.RoleServiceV1.DeleteRole:input_type -> role.DeleteRoleRequestV1
	5,  // 7: role.RoleServiceV1.CheckAdminRole:input_type -> role.CheckAdminRoleRequestV1
	7,  // 8: role.RoleServiceV1.CreateAdminRole:input_type -> role.CreateAdminRoleRequestV1
	8,  // 9: role.RoleServiceV1.UpdateAdminRole:input_type -> role.AdminRoleEntityV1
	10, // 10: role.RoleServiceV1.DeleteAdminRole:input_type -> role.DeleteAdminRoleRequestV1
	11, // 11: role.RoleServiceV1.GetAllAdminRole:input_type -> google.protobuf.Empty
	0,  // 12: role.RoleServiceV1.CreateRole:output_type -> role.RoleEntityV1
	0,  // 13: role.RoleServiceV1.GetRoleByID:output_type -> role.RoleEntityV1
	1,  // 14: role.RoleServiceV1.GetAllRole:output_type -> role.AllRoleEntityV1
	0,  // 15: role.RoleServiceV1.UpdateRole:output_type -> role.RoleEntityV1
	11, // 16: role.RoleServiceV1.DeleteRole:output_type -> google.protobuf.Empty
	6,  // 17: role.RoleServiceV1.CheckAdminRole:output_type -> role.CheckAdminRoleResponseV1
	8,  // 18: role.RoleServiceV1.CreateAdminRole:output_type -> role.AdminRoleEntityV1
	8,  // 19: role.RoleServiceV1.UpdateAdminRole:output_type -> role.AdminRoleEntityV1
	11, // 20: role.RoleServiceV1.DeleteAdminRole:output_type -> google.protobuf.Empty
	9,  // 21: role.RoleServiceV1.GetAllAdminRole:output_type -> role.AllAdminRoleEntityV1
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_role_service_v1_proto_init() }
func file_role_service_v1_proto_init() {
	if File_role_service_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_service_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleEntityV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllRoleEntityV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleByIDRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAdminRoleRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAdminRoleResponseV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdminRoleRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRoleEntityV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllAdminRoleEntityV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_role_service_v1_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdminRoleRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_role_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_service_v1_proto_goTypes,
		DependencyIndexes: file_role_service_v1_proto_depIdxs,
		MessageInfos:      file_role_service_v1_proto_msgTypes,
	}.Build()
	File_role_service_v1_proto = out.File
	file_role_service_v1_proto_rawDesc = nil
	file_role_service_v1_proto_goTypes = nil
	file_role_service_v1_proto_depIdxs = nil
}
