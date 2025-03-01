// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v2/webhook_service.proto

package apiv2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

type Webhook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId   int32                  `protobuf:"varint,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	RowStatus   RowStatus              `protobuf:"varint,5,opt,name=row_status,json=rowStatus,proto3,enum=lockets.api.v2.RowStatus" json:"row_status,omitempty"`
	Name        string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Url         string                 `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Webhook) Reset() {
	*x = Webhook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Webhook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Webhook) ProtoMessage() {}

func (x *Webhook) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Webhook.ProtoReflect.Descriptor instead.
func (*Webhook) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{0}
}

func (x *Webhook) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Webhook) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Webhook) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Webhook) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Webhook) GetRowStatus() RowStatus {
	if x != nil {
		return x.RowStatus
	}
	return RowStatus_ROW_STATUS_UNSPECIFIED
}

func (x *Webhook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Webhook) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateWebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateWebhookRequest) Reset() {
	*x = CreateWebhookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWebhookRequest) ProtoMessage() {}

func (x *CreateWebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWebhookRequest.ProtoReflect.Descriptor instead.
func (*CreateWebhookRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWebhookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateWebhookRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateWebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhook *Webhook `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *CreateWebhookResponse) Reset() {
	*x = CreateWebhookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWebhookResponse) ProtoMessage() {}

func (x *CreateWebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWebhookResponse.ProtoReflect.Descriptor instead.
func (*CreateWebhookResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWebhookResponse) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type GetWebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWebhookRequest) Reset() {
	*x = GetWebhookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebhookRequest) ProtoMessage() {}

func (x *GetWebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebhookRequest.ProtoReflect.Descriptor instead.
func (*GetWebhookRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetWebhookRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetWebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhook *Webhook `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *GetWebhookResponse) Reset() {
	*x = GetWebhookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebhookResponse) ProtoMessage() {}

func (x *GetWebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebhookResponse.ProtoReflect.Descriptor instead.
func (*GetWebhookResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetWebhookResponse) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type ListWebhooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatorId int32 `protobuf:"varint,1,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *ListWebhooksRequest) Reset() {
	*x = ListWebhooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWebhooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWebhooksRequest) ProtoMessage() {}

func (x *ListWebhooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWebhooksRequest.ProtoReflect.Descriptor instead.
func (*ListWebhooksRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListWebhooksRequest) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

type ListWebhooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhooks []*Webhook `protobuf:"bytes,1,rep,name=webhooks,proto3" json:"webhooks,omitempty"`
}

func (x *ListWebhooksResponse) Reset() {
	*x = ListWebhooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWebhooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWebhooksResponse) ProtoMessage() {}

func (x *ListWebhooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWebhooksResponse.ProtoReflect.Descriptor instead.
func (*ListWebhooksResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListWebhooksResponse) GetWebhooks() []*Webhook {
	if x != nil {
		return x.Webhooks
	}
	return nil
}

type UpdateWebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhook    *Webhook               `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateWebhookRequest) Reset() {
	*x = UpdateWebhookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWebhookRequest) ProtoMessage() {}

func (x *UpdateWebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWebhookRequest.ProtoReflect.Descriptor instead.
func (*UpdateWebhookRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateWebhookRequest) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

func (x *UpdateWebhookRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateWebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhook *Webhook `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *UpdateWebhookResponse) Reset() {
	*x = UpdateWebhookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWebhookResponse) ProtoMessage() {}

func (x *UpdateWebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWebhookResponse.ProtoReflect.Descriptor instead.
func (*UpdateWebhookResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateWebhookResponse) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type DeleteWebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWebhookRequest) Reset() {
	*x = DeleteWebhookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWebhookRequest) ProtoMessage() {}

func (x *DeleteWebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWebhookRequest.ProtoReflect.Descriptor instead.
func (*DeleteWebhookRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteWebhookRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteWebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteWebhookResponse) Reset() {
	*x = DeleteWebhookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_webhook_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWebhookResponse) ProtoMessage() {}

func (x *DeleteWebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_webhook_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWebhookResponse.ProtoReflect.Descriptor instead.
func (*DeleteWebhookResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_webhook_service_proto_rawDescGZIP(), []int{10}
}

var File_api_v2_webhook_service_proto protoreflect.FileDescriptor

var file_api_v2_webhook_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x13,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02,
	0x0a, 0x07, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x3c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x4a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x22, 0x34,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x08, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x4a, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x07, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa1, 0x05, 0x0a, 0x0e, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x24, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x77, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x12, 0x21, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0xda, 0x41, 0x02, 0x69,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x73,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x23,
	0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x73, 0x12, 0xa2, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x24, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x44, 0xda, 0x41, 0x13, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28,
	0x3a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x32, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x7b, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x69, 0x64, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x24, 0x2e, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0xda, 0x41, 0x02, 0x69, 0x64, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xb2, 0x01, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x42, 0x13, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x75, 0x71, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x4c, 0x41, 0x58, 0xaa, 0x02,
	0x0e, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0xca,
	0x02, 0x0e, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32,
	0xe2, 0x02, 0x1a, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10,
	0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v2_webhook_service_proto_rawDescOnce sync.Once
	file_api_v2_webhook_service_proto_rawDescData = file_api_v2_webhook_service_proto_rawDesc
)

func file_api_v2_webhook_service_proto_rawDescGZIP() []byte {
	file_api_v2_webhook_service_proto_rawDescOnce.Do(func() {
		file_api_v2_webhook_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_webhook_service_proto_rawDescData)
	})
	return file_api_v2_webhook_service_proto_rawDescData
}

var file_api_v2_webhook_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v2_webhook_service_proto_goTypes = []interface{}{
	(*Webhook)(nil),               // 0: lockets.api.v2.Webhook
	(*CreateWebhookRequest)(nil),  // 1: lockets.api.v2.CreateWebhookRequest
	(*CreateWebhookResponse)(nil), // 2: lockets.api.v2.CreateWebhookResponse
	(*GetWebhookRequest)(nil),     // 3: lockets.api.v2.GetWebhookRequest
	(*GetWebhookResponse)(nil),    // 4: lockets.api.v2.GetWebhookResponse
	(*ListWebhooksRequest)(nil),   // 5: lockets.api.v2.ListWebhooksRequest
	(*ListWebhooksResponse)(nil),  // 6: lockets.api.v2.ListWebhooksResponse
	(*UpdateWebhookRequest)(nil),  // 7: lockets.api.v2.UpdateWebhookRequest
	(*UpdateWebhookResponse)(nil), // 8: lockets.api.v2.UpdateWebhookResponse
	(*DeleteWebhookRequest)(nil),  // 9: lockets.api.v2.DeleteWebhookRequest
	(*DeleteWebhookResponse)(nil), // 10: lockets.api.v2.DeleteWebhookResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(RowStatus)(0),                // 12: lockets.api.v2.RowStatus
	(*fieldmaskpb.FieldMask)(nil), // 13: google.protobuf.FieldMask
}
var file_api_v2_webhook_service_proto_depIdxs = []int32{
	11, // 0: lockets.api.v2.Webhook.created_time:type_name -> google.protobuf.Timestamp
	11, // 1: lockets.api.v2.Webhook.updated_time:type_name -> google.protobuf.Timestamp
	12, // 2: lockets.api.v2.Webhook.row_status:type_name -> lockets.api.v2.RowStatus
	0,  // 3: lockets.api.v2.CreateWebhookResponse.webhook:type_name -> lockets.api.v2.Webhook
	0,  // 4: lockets.api.v2.GetWebhookResponse.webhook:type_name -> lockets.api.v2.Webhook
	0,  // 5: lockets.api.v2.ListWebhooksResponse.webhooks:type_name -> lockets.api.v2.Webhook
	0,  // 6: lockets.api.v2.UpdateWebhookRequest.webhook:type_name -> lockets.api.v2.Webhook
	13, // 7: lockets.api.v2.UpdateWebhookRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 8: lockets.api.v2.UpdateWebhookResponse.webhook:type_name -> lockets.api.v2.Webhook
	1,  // 9: lockets.api.v2.WebhookService.CreateWebhook:input_type -> lockets.api.v2.CreateWebhookRequest
	3,  // 10: lockets.api.v2.WebhookService.GetWebhook:input_type -> lockets.api.v2.GetWebhookRequest
	5,  // 11: lockets.api.v2.WebhookService.ListWebhooks:input_type -> lockets.api.v2.ListWebhooksRequest
	7,  // 12: lockets.api.v2.WebhookService.UpdateWebhook:input_type -> lockets.api.v2.UpdateWebhookRequest
	9,  // 13: lockets.api.v2.WebhookService.DeleteWebhook:input_type -> lockets.api.v2.DeleteWebhookRequest
	2,  // 14: lockets.api.v2.WebhookService.CreateWebhook:output_type -> lockets.api.v2.CreateWebhookResponse
	4,  // 15: lockets.api.v2.WebhookService.GetWebhook:output_type -> lockets.api.v2.GetWebhookResponse
	6,  // 16: lockets.api.v2.WebhookService.ListWebhooks:output_type -> lockets.api.v2.ListWebhooksResponse
	8,  // 17: lockets.api.v2.WebhookService.UpdateWebhook:output_type -> lockets.api.v2.UpdateWebhookResponse
	10, // 18: lockets.api.v2.WebhookService.DeleteWebhook:output_type -> lockets.api.v2.DeleteWebhookResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_v2_webhook_service_proto_init() }
func file_api_v2_webhook_service_proto_init() {
	if File_api_v2_webhook_service_proto != nil {
		return
	}
	file_api_v2_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v2_webhook_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Webhook); i {
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
		file_api_v2_webhook_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWebhookRequest); i {
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
		file_api_v2_webhook_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWebhookResponse); i {
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
		file_api_v2_webhook_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebhookRequest); i {
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
		file_api_v2_webhook_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebhookResponse); i {
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
		file_api_v2_webhook_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWebhooksRequest); i {
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
		file_api_v2_webhook_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWebhooksResponse); i {
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
		file_api_v2_webhook_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWebhookRequest); i {
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
		file_api_v2_webhook_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWebhookResponse); i {
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
		file_api_v2_webhook_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWebhookRequest); i {
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
		file_api_v2_webhook_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWebhookResponse); i {
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
			RawDescriptor: file_api_v2_webhook_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_webhook_service_proto_goTypes,
		DependencyIndexes: file_api_v2_webhook_service_proto_depIdxs,
		MessageInfos:      file_api_v2_webhook_service_proto_msgTypes,
	}.Build()
	File_api_v2_webhook_service_proto = out.File
	file_api_v2_webhook_service_proto_rawDesc = nil
	file_api_v2_webhook_service_proto_goTypes = nil
	file_api_v2_webhook_service_proto_depIdxs = nil
}
