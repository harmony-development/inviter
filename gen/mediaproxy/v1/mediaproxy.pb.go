// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: mediaproxy/v1/mediaproxy.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/harmony-development/inviter/gen/harmonytypes/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SiteMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SiteTitle   string `protobuf:"bytes,1,opt,name=site_title,json=siteTitle,proto3" json:"site_title,omitempty"`
	PageTitle   string `protobuf:"bytes,2,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	Kind        string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Url         string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Image       string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *SiteMetadata) Reset() {
	*x = SiteMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiteMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiteMetadata) ProtoMessage() {}

func (x *SiteMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiteMetadata.ProtoReflect.Descriptor instead.
func (*SiteMetadata) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{0}
}

func (x *SiteMetadata) GetSiteTitle() string {
	if x != nil {
		return x.SiteTitle
	}
	return ""
}

func (x *SiteMetadata) GetPageTitle() string {
	if x != nil {
		return x.PageTitle
	}
	return ""
}

func (x *SiteMetadata) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *SiteMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SiteMetadata) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SiteMetadata) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type MediaMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mimetype string `protobuf:"bytes,1,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *MediaMetadata) Reset() {
	*x = MediaMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaMetadata) ProtoMessage() {}

func (x *MediaMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaMetadata.ProtoReflect.Descriptor instead.
func (*MediaMetadata) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{1}
}

func (x *MediaMetadata) GetMimetype() string {
	if x != nil {
		return x.Mimetype
	}
	return ""
}

func (x *MediaMetadata) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type FetchLinkMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchLinkMetadataRequest) Reset() {
	*x = FetchLinkMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchLinkMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLinkMetadataRequest) ProtoMessage() {}

func (x *FetchLinkMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLinkMetadataRequest.ProtoReflect.Descriptor instead.
func (*FetchLinkMetadataRequest) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{2}
}

func (x *FetchLinkMetadataRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type FetchLinkMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*FetchLinkMetadataResponse_IsSite
	//	*FetchLinkMetadataResponse_IsMedia
	Data isFetchLinkMetadataResponse_Data `protobuf_oneof:"data"`
}

func (x *FetchLinkMetadataResponse) Reset() {
	*x = FetchLinkMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchLinkMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLinkMetadataResponse) ProtoMessage() {}

func (x *FetchLinkMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLinkMetadataResponse.ProtoReflect.Descriptor instead.
func (*FetchLinkMetadataResponse) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{3}
}

func (m *FetchLinkMetadataResponse) GetData() isFetchLinkMetadataResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *FetchLinkMetadataResponse) GetIsSite() *SiteMetadata {
	if x, ok := x.GetData().(*FetchLinkMetadataResponse_IsSite); ok {
		return x.IsSite
	}
	return nil
}

func (x *FetchLinkMetadataResponse) GetIsMedia() *MediaMetadata {
	if x, ok := x.GetData().(*FetchLinkMetadataResponse_IsMedia); ok {
		return x.IsMedia
	}
	return nil
}

type isFetchLinkMetadataResponse_Data interface {
	isFetchLinkMetadataResponse_Data()
}

type FetchLinkMetadataResponse_IsSite struct {
	IsSite *SiteMetadata `protobuf:"bytes,1,opt,name=is_site,json=isSite,proto3,oneof"`
}

type FetchLinkMetadataResponse_IsMedia struct {
	IsMedia *MediaMetadata `protobuf:"bytes,2,opt,name=is_media,json=isMedia,proto3,oneof"`
}

func (*FetchLinkMetadataResponse_IsSite) isFetchLinkMetadataResponse_Data() {}

func (*FetchLinkMetadataResponse_IsMedia) isFetchLinkMetadataResponse_Data() {}

type InstantViewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *InstantViewRequest) Reset() {
	*x = InstantViewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantViewRequest) ProtoMessage() {}

func (x *InstantViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantViewRequest.ProtoReflect.Descriptor instead.
func (*InstantViewRequest) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{4}
}

func (x *InstantViewRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type InstantViewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *SiteMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Content  string        `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	IsValid  bool          `protobuf:"varint,3,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
}

func (x *InstantViewResponse) Reset() {
	*x = InstantViewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantViewResponse) ProtoMessage() {}

func (x *InstantViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantViewResponse.ProtoReflect.Descriptor instead.
func (*InstantViewResponse) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{5}
}

func (x *InstantViewResponse) GetMetadata() *SiteMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *InstantViewResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *InstantViewResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

type CanInstantViewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanInstantView bool `protobuf:"varint,1,opt,name=can_instant_view,json=canInstantView,proto3" json:"can_instant_view,omitempty"`
}

func (x *CanInstantViewResponse) Reset() {
	*x = CanInstantViewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanInstantViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanInstantViewResponse) ProtoMessage() {}

func (x *CanInstantViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanInstantViewResponse.ProtoReflect.Descriptor instead.
func (*CanInstantViewResponse) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{6}
}

func (x *CanInstantViewResponse) GetCanInstantView() bool {
	if x != nil {
		return x.CanInstantView
	}
	return false
}

var File_mediaproxy_v1_mediaproxy_proto protoreflect.FileDescriptor

var file_mediaproxy_v1_mediaproxy_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e,
	0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x53, 0x69, 0x74, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x74, 0x65,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x22, 0x47, 0x0a, 0x0d, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x18, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xa8, 0x01, 0x0a, 0x19, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x69,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x06, 0x69, 0x73, 0x53, 0x69, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x8c, 0x01, 0x0a,
	0x13, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x16, 0x43,
	0x61, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x63, 0x61, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x32,
	0xf8, 0x02, 0x0a, 0x11, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69,
	0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x05, 0x9a, 0x44, 0x02, 0x08, 0x01, 0x12, 0x6d, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x12, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x05,
	0x9a, 0x44, 0x02, 0x08, 0x01, 0x12, 0x73, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x12, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x05, 0x9a, 0x44, 0x02, 0x08, 0x01, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mediaproxy_v1_mediaproxy_proto_rawDescOnce sync.Once
	file_mediaproxy_v1_mediaproxy_proto_rawDescData = file_mediaproxy_v1_mediaproxy_proto_rawDesc
)

func file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP() []byte {
	file_mediaproxy_v1_mediaproxy_proto_rawDescOnce.Do(func() {
		file_mediaproxy_v1_mediaproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_mediaproxy_v1_mediaproxy_proto_rawDescData)
	})
	return file_mediaproxy_v1_mediaproxy_proto_rawDescData
}

var file_mediaproxy_v1_mediaproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mediaproxy_v1_mediaproxy_proto_goTypes = []interface{}{
	(*SiteMetadata)(nil),              // 0: protocol.mediaproxy.v1.SiteMetadata
	(*MediaMetadata)(nil),             // 1: protocol.mediaproxy.v1.MediaMetadata
	(*FetchLinkMetadataRequest)(nil),  // 2: protocol.mediaproxy.v1.FetchLinkMetadataRequest
	(*FetchLinkMetadataResponse)(nil), // 3: protocol.mediaproxy.v1.FetchLinkMetadataResponse
	(*InstantViewRequest)(nil),        // 4: protocol.mediaproxy.v1.InstantViewRequest
	(*InstantViewResponse)(nil),       // 5: protocol.mediaproxy.v1.InstantViewResponse
	(*CanInstantViewResponse)(nil),    // 6: protocol.mediaproxy.v1.CanInstantViewResponse
}
var file_mediaproxy_v1_mediaproxy_proto_depIdxs = []int32{
	0, // 0: protocol.mediaproxy.v1.FetchLinkMetadataResponse.is_site:type_name -> protocol.mediaproxy.v1.SiteMetadata
	1, // 1: protocol.mediaproxy.v1.FetchLinkMetadataResponse.is_media:type_name -> protocol.mediaproxy.v1.MediaMetadata
	0, // 2: protocol.mediaproxy.v1.InstantViewResponse.metadata:type_name -> protocol.mediaproxy.v1.SiteMetadata
	2, // 3: protocol.mediaproxy.v1.MediaProxyService.FetchLinkMetadata:input_type -> protocol.mediaproxy.v1.FetchLinkMetadataRequest
	4, // 4: protocol.mediaproxy.v1.MediaProxyService.InstantView:input_type -> protocol.mediaproxy.v1.InstantViewRequest
	4, // 5: protocol.mediaproxy.v1.MediaProxyService.CanInstantView:input_type -> protocol.mediaproxy.v1.InstantViewRequest
	3, // 6: protocol.mediaproxy.v1.MediaProxyService.FetchLinkMetadata:output_type -> protocol.mediaproxy.v1.FetchLinkMetadataResponse
	5, // 7: protocol.mediaproxy.v1.MediaProxyService.InstantView:output_type -> protocol.mediaproxy.v1.InstantViewResponse
	6, // 8: protocol.mediaproxy.v1.MediaProxyService.CanInstantView:output_type -> protocol.mediaproxy.v1.CanInstantViewResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mediaproxy_v1_mediaproxy_proto_init() }
func file_mediaproxy_v1_mediaproxy_proto_init() {
	if File_mediaproxy_v1_mediaproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SiteMetadata); i {
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
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaMetadata); i {
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
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchLinkMetadataRequest); i {
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
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchLinkMetadataResponse); i {
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
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantViewRequest); i {
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
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantViewResponse); i {
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
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanInstantViewResponse); i {
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
	file_mediaproxy_v1_mediaproxy_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*FetchLinkMetadataResponse_IsSite)(nil),
		(*FetchLinkMetadataResponse_IsMedia)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mediaproxy_v1_mediaproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mediaproxy_v1_mediaproxy_proto_goTypes,
		DependencyIndexes: file_mediaproxy_v1_mediaproxy_proto_depIdxs,
		MessageInfos:      file_mediaproxy_v1_mediaproxy_proto_msgTypes,
	}.Build()
	File_mediaproxy_v1_mediaproxy_proto = out.File
	file_mediaproxy_v1_mediaproxy_proto_rawDesc = nil
	file_mediaproxy_v1_mediaproxy_proto_goTypes = nil
	file_mediaproxy_v1_mediaproxy_proto_depIdxs = nil
}
