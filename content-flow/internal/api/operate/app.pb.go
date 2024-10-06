// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: operate/app.proto

package operate

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

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 索引ID
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// 内容ID
	ContentID string `protobuf:"bytes,2,opt,name=contentID,proto3" json:"contentID,omitempty"`
	// 内容标题
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// 视频播放URL
	VideoURL string `protobuf:"bytes,4,opt,name=VideoURL,proto3" json:"VideoURL,omitempty"`
	// 作者
	Author string `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	// 内容描述
	Description string `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	// 封面图URL
	Thumbnail string `protobuf:"bytes,7,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// 内容分类
	Category string `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	// 内容时长
	Duration int64 `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	// 分辨率 如720p、1080p
	Resolution string `protobuf:"bytes,10,opt,name=resolution,proto3" json:"resolution,omitempty"`
	// 文件大小
	FileSize int64 `protobuf:"varint,11,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	// 文件格式 如MP4、AVI
	Format string `protobuf:"bytes,12,opt,name=format,proto3" json:"format,omitempty"`
	// // 视频质量 1-高清 2-标清
	Quality int32 `protobuf:"varint,13,opt,name=quality,proto3" json:"quality,omitempty"`
	// 审核状态 1-审核中 2-审核通过 3-审核不通过
	ApprovalStatus int32 `protobuf:"varint,14,opt,name=approvalStatus,proto3" json:"approvalStatus,omitempty"`
	// 创建时间
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// 审核状态 1-审核中 2-审核通过 3-审核不通过
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Content) GetContentID() string {
	if x != nil {
		return x.ContentID
	}
	return ""
}

func (x *Content) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Content) GetVideoURL() string {
	if x != nil {
		return x.VideoURL
	}
	return ""
}

func (x *Content) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Content) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Content) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *Content) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Content) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Content) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

func (x *Content) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *Content) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Content) GetQuality() int32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

func (x *Content) GetApprovalStatus() int32 {
	if x != nil {
		return x.ApprovalStatus
	}
	return 0
}

func (x *Content) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Content) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// 创建
type CreateContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content *Content `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateContentReq) Reset() {
	*x = CreateContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentReq) ProtoMessage() {}

func (x *CreateContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentReq.ProtoReflect.Descriptor instead.
func (*CreateContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContentReq) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type CreateContentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdxID int64 `protobuf:"varint,1,opt,name=idxID,proto3" json:"idxID,omitempty"`
}

func (x *CreateContentRsp) Reset() {
	*x = CreateContentRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentRsp) ProtoMessage() {}

func (x *CreateContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentRsp.ProtoReflect.Descriptor instead.
func (*CreateContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{2}
}

func (x *CreateContentRsp) GetIdxID() int64 {
	if x != nil {
		return x.IdxID
	}
	return 0
}

// 更新
type UpdateContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdxID   int64    `protobuf:"varint,1,opt,name=idxID,proto3" json:"idxID,omitempty"`
	Content *Content `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateContentReq) Reset() {
	*x = UpdateContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContentReq) ProtoMessage() {}

func (x *UpdateContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContentReq.ProtoReflect.Descriptor instead.
func (*UpdateContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateContentReq) GetIdxID() int64 {
	if x != nil {
		return x.IdxID
	}
	return 0
}

func (x *UpdateContentReq) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type UpdateContentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateContentRsp) Reset() {
	*x = UpdateContentRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContentRsp) ProtoMessage() {}

func (x *UpdateContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContentRsp.ProtoReflect.Descriptor instead.
func (*UpdateContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{4}
}

// 删除
type DeleteContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdxID int64 `protobuf:"varint,1,opt,name=IdxID,proto3" json:"IdxID,omitempty"`
}

func (x *DeleteContentReq) Reset() {
	*x = DeleteContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContentReq) ProtoMessage() {}

func (x *DeleteContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContentReq.ProtoReflect.Descriptor instead.
func (*DeleteContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteContentReq) GetIdxID() int64 {
	if x != nil {
		return x.IdxID
	}
	return 0
}

type DeleteContentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteContentRsp) Reset() {
	*x = DeleteContentRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContentRsp) ProtoMessage() {}

func (x *DeleteContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContentRsp.ProtoReflect.Descriptor instead.
func (*DeleteContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{6}
}

// 查找
type FindContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdxID    int64  `protobuf:"varint,1,opt,name=IdxID,proto3" json:"IdxID,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author   string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Page     int64  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64  `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *FindContentReq) Reset() {
	*x = FindContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindContentReq) ProtoMessage() {}

func (x *FindContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindContentReq.ProtoReflect.Descriptor instead.
func (*FindContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{7}
}

func (x *FindContentReq) GetIdxID() int64 {
	if x != nil {
		return x.IdxID
	}
	return 0
}

func (x *FindContentReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FindContentReq) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *FindContentReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindContentReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type FindContentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []*Content `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *FindContentRsp) Reset() {
	*x = FindContentRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindContentRsp) ProtoMessage() {}

func (x *FindContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindContentRsp.ProtoReflect.Descriptor instead.
func (*FindContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{8}
}

func (x *FindContentRsp) GetContent() []*Content {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_operate_app_proto protoreflect.FileDescriptor

var file_operate_app_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x04, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x55, 0x52, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x4a, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x64, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x69, 0x64, 0x78, 0x49, 0x44, 0x22, 0x60, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x64, 0x78,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x64, 0x78, 0x49, 0x44, 0x12,
	0x36, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x28, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x49, 0x64, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x49, 0x64, 0x78, 0x49, 0x44, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x46, 0x69,
	0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x49, 0x64, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x49, 0x64, 0x78,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x48, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x73, 0x70, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xfb, 0x02, 0x0a, 0x03, 0x41,
	0x70, 0x70, 0x12, 0x5d, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x5d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70,
	0x12, 0x5d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12,
	0x57, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x42, 0x43, 0x0a, 0x13, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50,
	0x01, 0x5a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x3b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operate_app_proto_rawDescOnce sync.Once
	file_operate_app_proto_rawDescData = file_operate_app_proto_rawDesc
)

func file_operate_app_proto_rawDescGZIP() []byte {
	file_operate_app_proto_rawDescOnce.Do(func() {
		file_operate_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_operate_app_proto_rawDescData)
	})
	return file_operate_app_proto_rawDescData
}

var file_operate_app_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_operate_app_proto_goTypes = []any{
	(*Content)(nil),               // 0: api.content.operate.Content
	(*CreateContentReq)(nil),      // 1: api.content.operate.CreateContentReq
	(*CreateContentRsp)(nil),      // 2: api.content.operate.CreateContentRsp
	(*UpdateContentReq)(nil),      // 3: api.content.operate.UpdateContentReq
	(*UpdateContentRsp)(nil),      // 4: api.content.operate.UpdateContentRsp
	(*DeleteContentReq)(nil),      // 5: api.content.operate.DeleteContentReq
	(*DeleteContentRsp)(nil),      // 6: api.content.operate.DeleteContentRsp
	(*FindContentReq)(nil),        // 7: api.content.operate.FindContentReq
	(*FindContentRsp)(nil),        // 8: api.content.operate.FindContentRsp
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_operate_app_proto_depIdxs = []int32{
	9, // 0: api.content.operate.Content.createdAt:type_name -> google.protobuf.Timestamp
	9, // 1: api.content.operate.Content.updatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: api.content.operate.CreateContentReq.content:type_name -> api.content.operate.Content
	0, // 3: api.content.operate.UpdateContentReq.content:type_name -> api.content.operate.Content
	0, // 4: api.content.operate.FindContentRsp.content:type_name -> api.content.operate.Content
	1, // 5: api.content.operate.App.CreateContent:input_type -> api.content.operate.CreateContentReq
	3, // 6: api.content.operate.App.UpdateContent:input_type -> api.content.operate.UpdateContentReq
	5, // 7: api.content.operate.App.DeleteContent:input_type -> api.content.operate.DeleteContentReq
	7, // 8: api.content.operate.App.FindContent:input_type -> api.content.operate.FindContentReq
	2, // 9: api.content.operate.App.CreateContent:output_type -> api.content.operate.CreateContentRsp
	4, // 10: api.content.operate.App.UpdateContent:output_type -> api.content.operate.UpdateContentRsp
	6, // 11: api.content.operate.App.DeleteContent:output_type -> api.content.operate.DeleteContentRsp
	8, // 12: api.content.operate.App.FindContent:output_type -> api.content.operate.FindContentRsp
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_operate_app_proto_init() }
func file_operate_app_proto_init() {
	if File_operate_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operate_app_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Content); i {
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
		file_operate_app_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateContentReq); i {
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
		file_operate_app_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateContentRsp); i {
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
		file_operate_app_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateContentReq); i {
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
		file_operate_app_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateContentRsp); i {
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
		file_operate_app_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteContentReq); i {
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
		file_operate_app_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteContentRsp); i {
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
		file_operate_app_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*FindContentReq); i {
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
		file_operate_app_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*FindContentRsp); i {
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
			RawDescriptor: file_operate_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operate_app_proto_goTypes,
		DependencyIndexes: file_operate_app_proto_depIdxs,
		MessageInfos:      file_operate_app_proto_msgTypes,
	}.Build()
	File_operate_app_proto = out.File
	file_operate_app_proto_rawDesc = nil
	file_operate_app_proto_goTypes = nil
	file_operate_app_proto_depIdxs = nil
}
