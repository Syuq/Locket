// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v2/activity_service.proto

package apiv2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId  int32                  `protobuf:"varint,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Type       string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Level      string                 `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Payload    *ActivityPayload       `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_activity_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_activity_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_api_v2_activity_service_proto_rawDescGZIP(), []int{0}
}

func (x *Activity) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Activity) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Activity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Activity) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Activity) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Activity) GetPayload() *ActivityPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ActivityLocketCommentPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocketId        int32 `protobuf:"varint,1,opt,name=locket_id,json=locketId,proto3" json:"locket_id,omitempty"`
	RelatedLocketId int32 `protobuf:"varint,2,opt,name=related_locket_id,json=relatedLocketId,proto3" json:"related_locket_id,omitempty"`
}

func (x *ActivityLocketCommentPayload) Reset() {
	*x = ActivityLocketCommentPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_activity_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityLocketCommentPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityLocketCommentPayload) ProtoMessage() {}

func (x *ActivityLocketCommentPayload) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_activity_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityLocketCommentPayload.ProtoReflect.Descriptor instead.
func (*ActivityLocketCommentPayload) Descriptor() ([]byte, []int) {
	return file_api_v2_activity_service_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityLocketCommentPayload) GetLocketId() int32 {
	if x != nil {
		return x.LocketId
	}
	return 0
}

func (x *ActivityLocketCommentPayload) GetRelatedLocketId() int32 {
	if x != nil {
		return x.RelatedLocketId
	}
	return 0
}

type ActivityVersionUpdatePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ActivityVersionUpdatePayload) Reset() {
	*x = ActivityVersionUpdatePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_activity_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityVersionUpdatePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityVersionUpdatePayload) ProtoMessage() {}

func (x *ActivityVersionUpdatePayload) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_activity_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityVersionUpdatePayload.ProtoReflect.Descriptor instead.
func (*ActivityVersionUpdatePayload) Descriptor() ([]byte, []int) {
	return file_api_v2_activity_service_proto_rawDescGZIP(), []int{2}
}

func (x *ActivityVersionUpdatePayload) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ActivityPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocketComment *ActivityLocketCommentPayload `protobuf:"bytes,1,opt,name=locket_comment,json=locketComment,proto3" json:"locket_comment,omitempty"`
	VersionUpdate *ActivityVersionUpdatePayload `protobuf:"bytes,2,opt,name=version_update,json=versionUpdate,proto3" json:"version_update,omitempty"`
}

func (x *ActivityPayload) Reset() {
	*x = ActivityPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_activity_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityPayload) ProtoMessage() {}

func (x *ActivityPayload) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_activity_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityPayload.ProtoReflect.Descriptor instead.
func (*ActivityPayload) Descriptor() ([]byte, []int) {
	return file_api_v2_activity_service_proto_rawDescGZIP(), []int{3}
}

func (x *ActivityPayload) GetLocketComment() *ActivityLocketCommentPayload {
	if x != nil {
		return x.LocketComment
	}
	return nil
}

func (x *ActivityPayload) GetVersionUpdate() *ActivityVersionUpdatePayload {
	if x != nil {
		return x.VersionUpdate
	}
	return nil
}

type GetActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetActivityRequest) Reset() {
	*x = GetActivityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_activity_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActivityRequest) ProtoMessage() {}

func (x *GetActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_activity_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActivityRequest.ProtoReflect.Descriptor instead.
func (*GetActivityRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_activity_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetActivityRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activity *Activity `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *GetActivityResponse) Reset() {
	*x = GetActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_activity_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActivityResponse) ProtoMessage() {}

func (x *GetActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_activity_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActivityResponse.ProtoReflect.Descriptor instead.
func (*GetActivityResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_activity_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetActivityResponse) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

var File_api_v2_activity_service_proto protoreflect.FileDescriptor

var file_api_v2_activity_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x67, 0x0a, 0x1c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x38,
	0x0a, 0x1c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xbb, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x53, 0x0a, 0x0e,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x53, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x32, 0x8b, 0x01, 0x0a, 0x0f, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x22, 0x2e, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0xda, 0x41, 0x02, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xb3, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x14,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x75, 0x71, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b,
	0x61, 0x70, 0x69, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x4c, 0x41, 0x58, 0xaa, 0x02, 0x0e, 0x4c, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0e, 0x4c,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1a,
	0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x4c, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v2_activity_service_proto_rawDescOnce sync.Once
	file_api_v2_activity_service_proto_rawDescData = file_api_v2_activity_service_proto_rawDesc
)

func file_api_v2_activity_service_proto_rawDescGZIP() []byte {
	file_api_v2_activity_service_proto_rawDescOnce.Do(func() {
		file_api_v2_activity_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_activity_service_proto_rawDescData)
	})
	return file_api_v2_activity_service_proto_rawDescData
}

var file_api_v2_activity_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v2_activity_service_proto_goTypes = []interface{}{
	(*Activity)(nil),                     // 0: lockets.api.v2.Activity
	(*ActivityLocketCommentPayload)(nil), // 1: lockets.api.v2.ActivityLocketCommentPayload
	(*ActivityVersionUpdatePayload)(nil), // 2: lockets.api.v2.ActivityVersionUpdatePayload
	(*ActivityPayload)(nil),              // 3: lockets.api.v2.ActivityPayload
	(*GetActivityRequest)(nil),           // 4: lockets.api.v2.GetActivityRequest
	(*GetActivityResponse)(nil),          // 5: lockets.api.v2.GetActivityResponse
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
}
var file_api_v2_activity_service_proto_depIdxs = []int32{
	6, // 0: lockets.api.v2.Activity.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: lockets.api.v2.Activity.payload:type_name -> lockets.api.v2.ActivityPayload
	1, // 2: lockets.api.v2.ActivityPayload.locket_comment:type_name -> lockets.api.v2.ActivityLocketCommentPayload
	2, // 3: lockets.api.v2.ActivityPayload.version_update:type_name -> lockets.api.v2.ActivityVersionUpdatePayload
	0, // 4: lockets.api.v2.GetActivityResponse.activity:type_name -> lockets.api.v2.Activity
	4, // 5: lockets.api.v2.ActivityService.GetActivity:input_type -> lockets.api.v2.GetActivityRequest
	5, // 6: lockets.api.v2.ActivityService.GetActivity:output_type -> lockets.api.v2.GetActivityResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_v2_activity_service_proto_init() }
func file_api_v2_activity_service_proto_init() {
	if File_api_v2_activity_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v2_activity_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_api_v2_activity_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityLocketCommentPayload); i {
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
		file_api_v2_activity_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityVersionUpdatePayload); i {
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
		file_api_v2_activity_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityPayload); i {
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
		file_api_v2_activity_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActivityRequest); i {
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
		file_api_v2_activity_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActivityResponse); i {
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
			RawDescriptor: file_api_v2_activity_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_activity_service_proto_goTypes,
		DependencyIndexes: file_api_v2_activity_service_proto_depIdxs,
		MessageInfos:      file_api_v2_activity_service_proto_msgTypes,
	}.Build()
	File_api_v2_activity_service_proto = out.File
	file_api_v2_activity_service_proto_rawDesc = nil
	file_api_v2_activity_service_proto_goTypes = nil
	file_api_v2_activity_service_proto_depIdxs = nil
}
