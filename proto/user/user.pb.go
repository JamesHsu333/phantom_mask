// @proto/ type following command to build pb.go
// protoc --proto_path=./ --go_out=./ --go-grpc_out=./ --grpc-gateway_out ./ --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --grpc-gateway_opt logtostderr=true --grpc-gateway_opt=paths=source_relative user/user.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0--rc2
// source: user/user.proto

package userProto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName  string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Name      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email     string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password  *string                `protobuf:"bytes,5,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Role      string                 `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	About     *string                `protobuf:"bytes,7,opt,name=about,proto3,oneof" json:"about,omitempty"`
	Avatar    *string                `protobuf:"bytes,8,opt,name=avatar,proto3,oneof" json:"avatar,omitempty"`
	Header    *string                `protobuf:"bytes,9,opt,name=header,proto3,oneof" json:"header,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	LoginDate *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=login_date,json=loginDate,proto3" json:"login_date,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetAbout() string {
	if x != nil && x.About != nil {
		return *x.About
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil && x.Avatar != nil {
		return *x.Avatar
	}
	return ""
}

func (x *User) GetHeader() string {
	if x != nil && x.Header != nil {
		return *x.Header
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetLoginDate() *timestamppb.Timestamp {
	if x != nil {
		return x.LoginDate
	}
	return nil
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName  string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Email     string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role      string                 `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	AuthToken string                 `protobuf:"bytes,5,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LoginDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=login_date,json=loginDate,proto3" json:"login_date,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *Auth) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Auth) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Auth) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Auth) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Auth) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *Auth) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Auth) GetLoginDate() *timestamppb.Timestamp {
	if x != nil {
		return x.LoginDate
	}
	return nil
}

type UsersList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int32   `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	TotalPages int32   `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Page       int32   `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Size       int32   `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	HasMore    bool    `protobuf:"varint,5,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Users      []*User `protobuf:"bytes,6,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersList) Reset() {
	*x = UsersList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersList) ProtoMessage() {}

func (x *UsersList) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersList.ProtoReflect.Descriptor instead.
func (*UsersList) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UsersList) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *UsersList) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *UsersList) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *UsersList) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UsersList) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *UsersList) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UploadInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File        []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size        string `protobuf:"bytes,3,opt,name=size,proto3" json:"size,omitempty"`
	ContentType string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *UploadInput) Reset() {
	*x = UploadInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInput) ProtoMessage() {}

func (x *UploadInput) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInput.ProtoReflect.Descriptor instead.
func (*UploadInput) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UploadInput) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *UploadInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadInput) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *UploadInput) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

var File_user_user_proto protoreflect.FileDescriptor

var file_user_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x62,
	0x6f, 0x75, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xfb, 0x01, 0x0a, 0x04, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x6c, 0x0a, 0x0b,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x6d, 0x65, 0x73, 0x48, 0x73,
	0x75, 0x33, 0x33, 0x33, 0x2f, 0x6b, 0x64, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData = file_user_user_proto_rawDesc
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_proto_rawDescData)
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_user_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: user.User
	(*Auth)(nil),                  // 1: user.Auth
	(*UsersList)(nil),             // 2: user.UsersList
	(*UploadInput)(nil),           // 3: user.UploadInput
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_user_user_proto_depIdxs = []int32{
	4, // 0: user.User.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: user.User.updated_at:type_name -> google.protobuf.Timestamp
	4, // 2: user.User.login_date:type_name -> google.protobuf.Timestamp
	4, // 3: user.Auth.created_at:type_name -> google.protobuf.Timestamp
	4, // 4: user.Auth.login_date:type_name -> google.protobuf.Timestamp
	0, // 5: user.UsersList.users:type_name -> user.User
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersList); i {
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
		file_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadInput); i {
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
	file_user_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_rawDesc = nil
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}