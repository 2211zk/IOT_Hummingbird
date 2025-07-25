// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.4
// source: user/v1/user.proto

package v1

import (
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

// 用户信息
type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName      string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserNickname  string                 `protobuf:"bytes,3,opt,name=user_nickname,json=userNickname,proto3" json:"user_nickname,omitempty"`
	Department    int32                  `protobuf:"varint,4,opt,name=department,proto3" json:"department,omitempty"`
	Mobile        string                 `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email         string                 `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Gender        string                 `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
	Role          int32                  `protobuf:"varint,9,opt,name=role,proto3" json:"role,omitempty"`
	UserStatus    string                 `protobuf:"bytes,10,opt,name=user_status,json=userStatus,proto3" json:"user_status,omitempty"`
	Comment       string                 `protobuf:"bytes,11,opt,name=comment,proto3" json:"comment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[0]
	if x != nil {
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
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetUserNickname() string {
	if x != nil {
		return x.UserNickname
	}
	return ""
}

func (x *User) GetDepartment() int32 {
	if x != nil {
		return x.Department
	}
	return 0
}

func (x *User) GetMobile() string {
	if x != nil {
		return x.Mobile
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
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *User) GetUserStatus() string {
	if x != nil {
		return x.UserStatus
	}
	return ""
}

func (x *User) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

// 用户注册请求（支持DTM分布式事务）
type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserNickname  string                 `protobuf:"bytes,2,opt,name=user_nickname,json=userNickname,proto3" json:"user_nickname,omitempty"`
	Department    int32                  `protobuf:"varint,3,opt,name=department,proto3" json:"department,omitempty"`
	Mobile        string                 `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Gender        string                 `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Role          int32                  `protobuf:"varint,8,opt,name=role,proto3" json:"role,omitempty"`
	Comment       string                 `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"`
	DtmGid        string                 `protobuf:"bytes,10,opt,name=dtm_gid,json=dtmGid,proto3" json:"dtm_gid,omitempty"` // DTM全局事务ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_user_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *RegisterRequest) GetUserNickname() string {
	if x != nil {
		return x.UserNickname
	}
	return ""
}

func (x *RegisterRequest) GetDepartment() int32 {
	if x != nil {
		return x.Department
	}
	return 0
}

func (x *RegisterRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *RegisterRequest) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *RegisterRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *RegisterRequest) GetDtmGid() string {
	if x != nil {
		return x.DtmGid
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	mi := &file_user_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 用户注册补偿请求（DTM分布式事务）
type RegisterCompensateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DtmGid        string                 `protobuf:"bytes,2,opt,name=dtm_gid,json=dtmGid,proto3" json:"dtm_gid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterCompensateRequest) Reset() {
	*x = RegisterCompensateRequest{}
	mi := &file_user_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterCompensateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCompensateRequest) ProtoMessage() {}

func (x *RegisterCompensateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCompensateRequest.ProtoReflect.Descriptor instead.
func (*RegisterCompensateRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterCompensateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegisterCompensateRequest) GetDtmGid() string {
	if x != nil {
		return x.DtmGid
	}
	return ""
}

type RegisterCompensateReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterCompensateReply) Reset() {
	*x = RegisterCompensateReply{}
	mi := &file_user_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterCompensateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCompensateReply) ProtoMessage() {}

func (x *RegisterCompensateReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCompensateReply.ProtoReflect.Descriptor instead.
func (*RegisterCompensateReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterCompensateReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 用户登录
type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_user_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *LoginRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName      string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Token         string                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	mi := &file_user_v1_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *LoginReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginReply) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 用户查询
type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_user_v1_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	mi := &file_user_v1_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserReply) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// 用户列表
type ListUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	mi := &file_user_v1_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{9}
}

func (x *ListUserRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListUserRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListUserReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*User                `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserReply) Reset() {
	*x = ListUserReply{}
	mi := &file_user_v1_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReply) ProtoMessage() {}

func (x *ListUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReply.ProtoReflect.Descriptor instead.
func (*ListUserReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{10}
}

func (x *ListUserReply) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ListUserReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 用户更新
type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserNickname  string                 `protobuf:"bytes,2,opt,name=user_nickname,json=userNickname,proto3" json:"user_nickname,omitempty"`
	Department    int32                  `protobuf:"varint,3,opt,name=department,proto3" json:"department,omitempty"`
	Mobile        string                 `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Gender        string                 `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Role          int32                  `protobuf:"varint,7,opt,name=role,proto3" json:"role,omitempty"`
	UserStatus    string                 `protobuf:"bytes,8,opt,name=user_status,json=userStatus,proto3" json:"user_status,omitempty"`
	Comment       string                 `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_user_v1_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserRequest) GetUserNickname() string {
	if x != nil {
		return x.UserNickname
	}
	return ""
}

func (x *UpdateUserRequest) GetDepartment() int32 {
	if x != nil {
		return x.Department
	}
	return 0
}

func (x *UpdateUserRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdateUserRequest) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *UpdateUserRequest) GetUserStatus() string {
	if x != nil {
		return x.UserStatus
	}
	return ""
}

func (x *UpdateUserRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type UpdateUserReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserReply) Reset() {
	*x = UpdateUserReply{}
	mi := &file_user_v1_user_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReply) ProtoMessage() {}

func (x *UpdateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReply.ProtoReflect.Descriptor instead.
func (*UpdateUserReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateUserReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 用户删除
type DeleteUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	mi := &file_user_v1_user_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserReply) Reset() {
	*x = DeleteUserReply{}
	mi := &file_user_v1_user_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReply) ProtoMessage() {}

func (x *DeleteUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReply.ProtoReflect.Descriptor instead.
func (*DeleteUserReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteUserReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_user_v1_user_proto protoreflect.FileDescriptor

const file_user_v1_user_proto_rawDesc = "" +
	"\n" +
	"\x12user/v1/user.proto\x12\auser.v1\"\xa9\x02\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\x12#\n" +
	"\ruser_nickname\x18\x03 \x01(\tR\fuserNickname\x12\x1e\n" +
	"\n" +
	"department\x18\x04 \x01(\x05R\n" +
	"department\x12\x16\n" +
	"\x06mobile\x18\x05 \x01(\tR\x06mobile\x12\x14\n" +
	"\x05email\x18\x06 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\a \x01(\tR\bpassword\x12\x16\n" +
	"\x06gender\x18\b \x01(\tR\x06gender\x12\x12\n" +
	"\x04role\x18\t \x01(\x05R\x04role\x12\x1f\n" +
	"\vuser_status\x18\n" +
	" \x01(\tR\n" +
	"userStatus\x12\x18\n" +
	"\acomment\x18\v \x01(\tR\acomment\"\x9c\x02\n" +
	"\x0fRegisterRequest\x12\x1b\n" +
	"\tuser_name\x18\x01 \x01(\tR\buserName\x12#\n" +
	"\ruser_nickname\x18\x02 \x01(\tR\fuserNickname\x12\x1e\n" +
	"\n" +
	"department\x18\x03 \x01(\x05R\n" +
	"department\x12\x16\n" +
	"\x06mobile\x18\x04 \x01(\tR\x06mobile\x12\x14\n" +
	"\x05email\x18\x05 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x06 \x01(\tR\bpassword\x12\x16\n" +
	"\x06gender\x18\a \x01(\tR\x06gender\x12\x12\n" +
	"\x04role\x18\b \x01(\x05R\x04role\x12\x18\n" +
	"\acomment\x18\t \x01(\tR\acomment\x12\x17\n" +
	"\adtm_gid\x18\n" +
	" \x01(\tR\x06dtmGid\"\x1f\n" +
	"\rRegisterReply\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"D\n" +
	"\x19RegisterCompensateRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x17\n" +
	"\adtm_gid\x18\x02 \x01(\tR\x06dtmGid\"3\n" +
	"\x17RegisterCompensateReply\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"G\n" +
	"\fLoginRequest\x12\x1b\n" +
	"\tuser_name\x18\x01 \x01(\tR\buserName\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"O\n" +
	"\n" +
	"LoginReply\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\x12\x14\n" +
	"\x05token\x18\x03 \x01(\tR\x05token\" \n" +
	"\x0eGetUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"1\n" +
	"\fGetUserReply\x12!\n" +
	"\x04user\x18\x01 \x01(\v2\r.user.v1.UserR\x04user\"B\n" +
	"\x0fListUserRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\"J\n" +
	"\rListUserReply\x12#\n" +
	"\x05users\x18\x01 \x03(\v2\r.user.v1.UserR\x05users\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total\"\xfd\x01\n" +
	"\x11UpdateUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12#\n" +
	"\ruser_nickname\x18\x02 \x01(\tR\fuserNickname\x12\x1e\n" +
	"\n" +
	"department\x18\x03 \x01(\x05R\n" +
	"department\x12\x16\n" +
	"\x06mobile\x18\x04 \x01(\tR\x06mobile\x12\x14\n" +
	"\x05email\x18\x05 \x01(\tR\x05email\x12\x16\n" +
	"\x06gender\x18\x06 \x01(\tR\x06gender\x12\x12\n" +
	"\x04role\x18\a \x01(\x05R\x04role\x12\x1f\n" +
	"\vuser_status\x18\b \x01(\tR\n" +
	"userStatus\x12\x18\n" +
	"\acomment\x18\t \x01(\tR\acomment\"+\n" +
	"\x0fUpdateUserReply\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"#\n" +
	"\x11DeleteUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"+\n" +
	"\x0fDeleteUserReply\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xdd\x03\n" +
	"\vUserService\x12<\n" +
	"\bRegister\x12\x18.user.v1.RegisterRequest\x1a\x16.user.v1.RegisterReply\x12Z\n" +
	"\x12RegisterCompensate\x12\".user.v1.RegisterCompensateRequest\x1a .user.v1.RegisterCompensateReply\x123\n" +
	"\x05Login\x12\x15.user.v1.LoginRequest\x1a\x13.user.v1.LoginReply\x129\n" +
	"\aGetUser\x12\x17.user.v1.GetUserRequest\x1a\x15.user.v1.GetUserReply\x12<\n" +
	"\bListUser\x12\x18.user.v1.ListUserRequest\x1a\x16.user.v1.ListUserReply\x12B\n" +
	"\n" +
	"UpdateUser\x12\x1a.user.v1.UpdateUserRequest\x1a\x18.user.v1.UpdateUserReply\x12B\n" +
	"\n" +
	"DeleteUser\x12\x1a.user.v1.DeleteUserRequest\x1a\x18.user.v1.DeleteUserReplyB\x17Z\x15kratos/api/user/v1;v1b\x06proto3"

var (
	file_user_v1_user_proto_rawDescOnce sync.Once
	file_user_v1_user_proto_rawDescData []byte
)

func file_user_v1_user_proto_rawDescGZIP() []byte {
	file_user_v1_user_proto_rawDescOnce.Do(func() {
		file_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_v1_user_proto_rawDesc), len(file_user_v1_user_proto_rawDesc)))
	})
	return file_user_v1_user_proto_rawDescData
}

var file_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_user_v1_user_proto_goTypes = []any{
	(*User)(nil),                      // 0: user.v1.User
	(*RegisterRequest)(nil),           // 1: user.v1.RegisterRequest
	(*RegisterReply)(nil),             // 2: user.v1.RegisterReply
	(*RegisterCompensateRequest)(nil), // 3: user.v1.RegisterCompensateRequest
	(*RegisterCompensateReply)(nil),   // 4: user.v1.RegisterCompensateReply
	(*LoginRequest)(nil),              // 5: user.v1.LoginRequest
	(*LoginReply)(nil),                // 6: user.v1.LoginReply
	(*GetUserRequest)(nil),            // 7: user.v1.GetUserRequest
	(*GetUserReply)(nil),              // 8: user.v1.GetUserReply
	(*ListUserRequest)(nil),           // 9: user.v1.ListUserRequest
	(*ListUserReply)(nil),             // 10: user.v1.ListUserReply
	(*UpdateUserRequest)(nil),         // 11: user.v1.UpdateUserRequest
	(*UpdateUserReply)(nil),           // 12: user.v1.UpdateUserReply
	(*DeleteUserRequest)(nil),         // 13: user.v1.DeleteUserRequest
	(*DeleteUserReply)(nil),           // 14: user.v1.DeleteUserReply
}
var file_user_v1_user_proto_depIdxs = []int32{
	0,  // 0: user.v1.GetUserReply.user:type_name -> user.v1.User
	0,  // 1: user.v1.ListUserReply.users:type_name -> user.v1.User
	1,  // 2: user.v1.UserService.Register:input_type -> user.v1.RegisterRequest
	3,  // 3: user.v1.UserService.RegisterCompensate:input_type -> user.v1.RegisterCompensateRequest
	5,  // 4: user.v1.UserService.Login:input_type -> user.v1.LoginRequest
	7,  // 5: user.v1.UserService.GetUser:input_type -> user.v1.GetUserRequest
	9,  // 6: user.v1.UserService.ListUser:input_type -> user.v1.ListUserRequest
	11, // 7: user.v1.UserService.UpdateUser:input_type -> user.v1.UpdateUserRequest
	13, // 8: user.v1.UserService.DeleteUser:input_type -> user.v1.DeleteUserRequest
	2,  // 9: user.v1.UserService.Register:output_type -> user.v1.RegisterReply
	4,  // 10: user.v1.UserService.RegisterCompensate:output_type -> user.v1.RegisterCompensateReply
	6,  // 11: user.v1.UserService.Login:output_type -> user.v1.LoginReply
	8,  // 12: user.v1.UserService.GetUser:output_type -> user.v1.GetUserReply
	10, // 13: user.v1.UserService.ListUser:output_type -> user.v1.ListUserReply
	12, // 14: user.v1.UserService.UpdateUser:output_type -> user.v1.UpdateUserReply
	14, // 15: user.v1.UserService.DeleteUser:output_type -> user.v1.DeleteUserReply
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_user_v1_user_proto_init() }
func file_user_v1_user_proto_init() {
	if File_user_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_v1_user_proto_rawDesc), len(file_user_v1_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_user_proto_goTypes,
		DependencyIndexes: file_user_v1_user_proto_depIdxs,
		MessageInfos:      file_user_v1_user_proto_msgTypes,
	}.Build()
	File_user_v1_user_proto = out.File
	file_user_v1_user_proto_goTypes = nil
	file_user_v1_user_proto_depIdxs = nil
}
