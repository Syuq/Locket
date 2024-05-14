// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: store/idp.proto

package store

import (
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

type IdentityProviderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Config:
	//
	//	*IdentityProviderConfig_Oauth2
	Config isIdentityProviderConfig_Config `protobuf_oneof:"config"`
}

func (x *IdentityProviderConfig) Reset() {
	*x = IdentityProviderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_idp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityProviderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProviderConfig) ProtoMessage() {}

func (x *IdentityProviderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_idp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProviderConfig.ProtoReflect.Descriptor instead.
func (*IdentityProviderConfig) Descriptor() ([]byte, []int) {
	return file_store_idp_proto_rawDescGZIP(), []int{0}
}

func (m *IdentityProviderConfig) GetConfig() isIdentityProviderConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *IdentityProviderConfig) GetOauth2() *IdentityProviderConfig_OAuth2 {
	if x, ok := x.GetConfig().(*IdentityProviderConfig_Oauth2); ok {
		return x.Oauth2
	}
	return nil
}

type isIdentityProviderConfig_Config interface {
	isIdentityProviderConfig_Config()
}

type IdentityProviderConfig_Oauth2 struct {
	Oauth2 *IdentityProviderConfig_OAuth2 `protobuf:"bytes,1,opt,name=oauth2,proto3,oneof"`
}

func (*IdentityProviderConfig_Oauth2) isIdentityProviderConfig_Config() {}

type IdentityProviderConfig_FieldMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier  string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *IdentityProviderConfig_FieldMapping) Reset() {
	*x = IdentityProviderConfig_FieldMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_idp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityProviderConfig_FieldMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProviderConfig_FieldMapping) ProtoMessage() {}

func (x *IdentityProviderConfig_FieldMapping) ProtoReflect() protoreflect.Message {
	mi := &file_store_idp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProviderConfig_FieldMapping.ProtoReflect.Descriptor instead.
func (*IdentityProviderConfig_FieldMapping) Descriptor() ([]byte, []int) {
	return file_store_idp_proto_rawDescGZIP(), []int{0, 0}
}

func (x *IdentityProviderConfig_FieldMapping) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *IdentityProviderConfig_FieldMapping) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *IdentityProviderConfig_FieldMapping) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type IdentityProviderConfig_OAuth2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string                               `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string                               `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	AuthUrl      string                               `protobuf:"bytes,3,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	TokenUrl     string                               `protobuf:"bytes,4,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	UserInfoUrl  string                               `protobuf:"bytes,5,opt,name=user_info_url,json=userInfoUrl,proto3" json:"user_info_url,omitempty"`
	Scopes       []string                             `protobuf:"bytes,6,rep,name=scopes,proto3" json:"scopes,omitempty"`
	FieldMapping *IdentityProviderConfig_FieldMapping `protobuf:"bytes,7,opt,name=field_mapping,json=fieldMapping,proto3" json:"field_mapping,omitempty"`
}

func (x *IdentityProviderConfig_OAuth2) Reset() {
	*x = IdentityProviderConfig_OAuth2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_idp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityProviderConfig_OAuth2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProviderConfig_OAuth2) ProtoMessage() {}

func (x *IdentityProviderConfig_OAuth2) ProtoReflect() protoreflect.Message {
	mi := &file_store_idp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProviderConfig_OAuth2.ProtoReflect.Descriptor instead.
func (*IdentityProviderConfig_OAuth2) Descriptor() ([]byte, []int) {
	return file_store_idp_proto_rawDescGZIP(), []int{0, 1}
}

func (x *IdentityProviderConfig_OAuth2) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *IdentityProviderConfig_OAuth2) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *IdentityProviderConfig_OAuth2) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *IdentityProviderConfig_OAuth2) GetTokenUrl() string {
	if x != nil {
		return x.TokenUrl
	}
	return ""
}

func (x *IdentityProviderConfig_OAuth2) GetUserInfoUrl() string {
	if x != nil {
		return x.UserInfoUrl
	}
	return ""
}

func (x *IdentityProviderConfig_OAuth2) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *IdentityProviderConfig_OAuth2) GetFieldMapping() *IdentityProviderConfig_FieldMapping {
	if x != nil {
		return x.FieldMapping
	}
	return nil
}

var File_store_idp_proto protoreflect.FileDescriptor

var file_store_idp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x22, 0xed, 0x03, 0x0a, 0x16, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x0a, 0x06, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x32, 0x1a, 0x67, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x97, 0x02, 0x0a,
	0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x72,
	0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x57, 0x0a,
	0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x42, 0x9a, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x08, 0x49, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x79, 0x75, 0x71, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x4c, 0x53, 0x58,
	0xaa, 0x02, 0x0d, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0xca, 0x02, 0x0d, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0xe2, 0x02, 0x19, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4c,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_idp_proto_rawDescOnce sync.Once
	file_store_idp_proto_rawDescData = file_store_idp_proto_rawDesc
)

func file_store_idp_proto_rawDescGZIP() []byte {
	file_store_idp_proto_rawDescOnce.Do(func() {
		file_store_idp_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_idp_proto_rawDescData)
	})
	return file_store_idp_proto_rawDescData
}

var file_store_idp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_store_idp_proto_goTypes = []interface{}{
	(*IdentityProviderConfig)(nil),              // 0: lockets.store.IdentityProviderConfig
	(*IdentityProviderConfig_FieldMapping)(nil), // 1: lockets.store.IdentityProviderConfig.FieldMapping
	(*IdentityProviderConfig_OAuth2)(nil),       // 2: lockets.store.IdentityProviderConfig.OAuth2
}
var file_store_idp_proto_depIdxs = []int32{
	2, // 0: lockets.store.IdentityProviderConfig.oauth2:type_name -> lockets.store.IdentityProviderConfig.OAuth2
	1, // 1: lockets.store.IdentityProviderConfig.OAuth2.field_mapping:type_name -> lockets.store.IdentityProviderConfig.FieldMapping
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_store_idp_proto_init() }
func file_store_idp_proto_init() {
	if File_store_idp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_idp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityProviderConfig); i {
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
		file_store_idp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityProviderConfig_FieldMapping); i {
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
		file_store_idp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityProviderConfig_OAuth2); i {
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
	file_store_idp_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*IdentityProviderConfig_Oauth2)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_idp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_idp_proto_goTypes,
		DependencyIndexes: file_store_idp_proto_depIdxs,
		MessageInfos:      file_store_idp_proto_msgTypes,
	}.Build()
	File_store_idp_proto = out.File
	file_store_idp_proto_rawDesc = nil
	file_store_idp_proto_goTypes = nil
	file_store_idp_proto_depIdxs = nil
}
