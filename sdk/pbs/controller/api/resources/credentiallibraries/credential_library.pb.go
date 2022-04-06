// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: controller/api/resources/credentiallibraries/v1/credential_library.proto

package credentiallibraries

import (
	scopes "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/scopes"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CredentialLibrary contains all fields related to an Credential Library resource
type CredentialLibrary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The ID of the Credential Library.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty" class:"public"` // @gotags: `class:"public"`
	// The ID of the Credential Store of which this Credential Library is a part.
	CredentialStoreId string `protobuf:"bytes,20,opt,name=credential_store_id,proto3" json:"credential_store_id,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. Scope information for this Credential Library.
	Scope *scopes.ScopeInfo `protobuf:"bytes,30,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional name for identification purposes.
	Name *wrapperspb.StringValue `protobuf:"bytes,40,opt,name=name,proto3" json:"name,omitempty" class:"public"` // @gotags: `class:"public"`
	// Optional user-set description for identification purposes.
	Description *wrapperspb.StringValue `protobuf:"bytes,50,opt,name=description,proto3" json:"description,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The time this resource was created.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,60,opt,name=created_time,proto3" json:"created_time,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The time this resource was last updated.
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,70,opt,name=updated_time,proto3" json:"updated_time,omitempty" class:"public"` // @gotags: `class:"public"`
	// Version is used in mutation requests, after the initial creation, to ensure this resource has not changed.
	// The mutation will fail if the version does not match the latest known good version.
	Version uint32 `protobuf:"varint,80,opt,name=version,proto3" json:"version,omitempty" class:"public"` // @gotags: `class:"public"`
	// The Credential Library type.
	Type string `protobuf:"bytes,90,opt,name=type,proto3" json:"type,omitempty" class:"public"` // @gotags: `class:"public"`
	// Types that are assignable to Attrs:
	//	*CredentialLibrary_Attributes
	//	*CredentialLibrary_VaultCredentialLibraryAttributes
	Attrs isCredentialLibrary_Attrs `protobuf_oneof:"attrs"`
	// Output only. The available actions on this resource for this user.
	AuthorizedActions []string `protobuf:"bytes,300,rep,name=authorized_actions,proto3" json:"authorized_actions,omitempty" class:"public"` // @gotags: `class:"public"`
	// The type of credential this library will issue, defaults to Unspecified
	CredentialType string `protobuf:"bytes,310,opt,name=credential_type,proto3" json:"credential_type,omitempty" class:"public"` // @gotags: `class:"public"`
	// The credential mapping overrides
	CredentialMappingOverrides *structpb.Struct `protobuf:"bytes,320,opt,name=credential_mapping_overrides,proto3" json:"credential_mapping_overrides,omitempty"`
}

func (x *CredentialLibrary) Reset() {
	*x = CredentialLibrary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_credentiallibraries_v1_credential_library_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialLibrary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialLibrary) ProtoMessage() {}

func (x *CredentialLibrary) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_credentiallibraries_v1_credential_library_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialLibrary.ProtoReflect.Descriptor instead.
func (*CredentialLibrary) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialLibrary) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CredentialLibrary) GetCredentialStoreId() string {
	if x != nil {
		return x.CredentialStoreId
	}
	return ""
}

func (x *CredentialLibrary) GetScope() *scopes.ScopeInfo {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *CredentialLibrary) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *CredentialLibrary) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *CredentialLibrary) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *CredentialLibrary) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *CredentialLibrary) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CredentialLibrary) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (m *CredentialLibrary) GetAttrs() isCredentialLibrary_Attrs {
	if m != nil {
		return m.Attrs
	}
	return nil
}

func (x *CredentialLibrary) GetAttributes() *structpb.Struct {
	if x, ok := x.GetAttrs().(*CredentialLibrary_Attributes); ok {
		return x.Attributes
	}
	return nil
}

func (x *CredentialLibrary) GetVaultCredentialLibraryAttributes() *VaultCredentialLibraryAttributes {
	if x, ok := x.GetAttrs().(*CredentialLibrary_VaultCredentialLibraryAttributes); ok {
		return x.VaultCredentialLibraryAttributes
	}
	return nil
}

func (x *CredentialLibrary) GetAuthorizedActions() []string {
	if x != nil {
		return x.AuthorizedActions
	}
	return nil
}

func (x *CredentialLibrary) GetCredentialType() string {
	if x != nil {
		return x.CredentialType
	}
	return ""
}

func (x *CredentialLibrary) GetCredentialMappingOverrides() *structpb.Struct {
	if x != nil {
		return x.CredentialMappingOverrides
	}
	return nil
}

type isCredentialLibrary_Attrs interface {
	isCredentialLibrary_Attrs()
}

type CredentialLibrary_Attributes struct {
	// The attributes that are applicable for the specific Credential Library type.
	Attributes *structpb.Struct `protobuf:"bytes,100,opt,name=attributes,proto3,oneof"`
}

type CredentialLibrary_VaultCredentialLibraryAttributes struct {
	VaultCredentialLibraryAttributes *VaultCredentialLibraryAttributes `protobuf:"bytes,101,opt,name=vault_credential_library_attributes,json=vaultCredentialLibraryAttributes,proto3,oneof"`
}

func (*CredentialLibrary_Attributes) isCredentialLibrary_Attrs() {}

func (*CredentialLibrary_VaultCredentialLibraryAttributes) isCredentialLibrary_Attrs() {}

// The attributes of a vault typed Credential Library.
type VaultCredentialLibraryAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path in Vault to request credentials from.
	Path *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=path,proto3" json:"path,omitempty" class:"public"` // @gotags: `class:"public"`
	// The HTTP method the library uses to communicate with Vault.
	HttpMethod *wrapperspb.StringValue `protobuf:"bytes,20,opt,name=http_method,proto3" json:"http_method,omitempty" class:"public"` // @gotags: `class:"public"`
	// The body of the HTTP request the library sends to vault. When set http_method must be "POST"
	HttpRequestBody *wrapperspb.StringValue `protobuf:"bytes,30,opt,name=http_request_body,proto3" json:"http_request_body,omitempty" class:"secret"` // @gotags: `class:"secret"`
}

func (x *VaultCredentialLibraryAttributes) Reset() {
	*x = VaultCredentialLibraryAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_credentiallibraries_v1_credential_library_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultCredentialLibraryAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultCredentialLibraryAttributes) ProtoMessage() {}

func (x *VaultCredentialLibraryAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_credentiallibraries_v1_credential_library_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultCredentialLibraryAttributes.ProtoReflect.Descriptor instead.
func (*VaultCredentialLibraryAttributes) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDescGZIP(), []int{1}
}

func (x *VaultCredentialLibraryAttributes) GetPath() *wrapperspb.StringValue {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *VaultCredentialLibraryAttributes) GetHttpMethod() *wrapperspb.StringValue {
	if x != nil {
		return x.HttpMethod
	}
	return nil
}

func (x *VaultCredentialLibraryAttributes) GetHttpRequestBody() *wrapperspb.StringValue {
	if x != nil {
		return x.HttpRequestBody
	}
	return nil
}

var File_controller_api_resources_credentiallibraries_v1_credential_library_proto protoreflect.FileDescriptor

var file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDesc = []byte{
	0x0a, 0x48, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x07, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x13, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xa0, 0xe3, 0x29, 0x01, 0x52, 0x13, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69,
	0x64, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x14, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x0c, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x62,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x22, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x1a, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x50, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x4a, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x0f, 0xa0,
	0xda, 0x29, 0x01, 0x9a, 0xe3, 0x29, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0xc1, 0x01, 0x0a,
	0x23, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x1d, 0xa0,
	0xda, 0x29, 0x01, 0x9a, 0xe3, 0x29, 0x05, 0x76, 0x61, 0x75, 0x6c, 0x74, 0xfa, 0xd2, 0xe4, 0x93,
	0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x48, 0x00, 0x52, 0x20,
	0x76, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x2f, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xac, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x2f, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0xb6, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xa0, 0xda, 0x29,
	0x01, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x62, 0x0a, 0x1c, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x73, 0x18, 0xc0, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x42, 0x04, 0xa0, 0xda, 0x29, 0x01, 0x52, 0x1c, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x22,
	0xee, 0x02, 0x0a, 0x20, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x24, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x1c, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x12, 0x09, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x6c, 0x0a, 0x0b,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x2c, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x24, 0x0a, 0x16, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x0a, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0b, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x83, 0x01, 0x0a, 0x11, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x37, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x2f, 0x0a, 0x1c,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0f, 0x48, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x11, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79,
	0x42, 0x68, 0x5a, 0x66, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x3b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDescOnce sync.Once
	file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDescData = file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDesc
)

func file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDescGZIP() []byte {
	file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDescData)
	})
	return file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDescData
}

var file_controller_api_resources_credentiallibraries_v1_credential_library_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_api_resources_credentiallibraries_v1_credential_library_proto_goTypes = []interface{}{
	(*CredentialLibrary)(nil),                // 0: controller.api.resources.credentiallibraries.v1.CredentialLibrary
	(*VaultCredentialLibraryAttributes)(nil), // 1: controller.api.resources.credentiallibraries.v1.VaultCredentialLibraryAttributes
	(*scopes.ScopeInfo)(nil),                 // 2: controller.api.resources.scopes.v1.ScopeInfo
	(*wrapperspb.StringValue)(nil),           // 3: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),            // 4: google.protobuf.Timestamp
	(*structpb.Struct)(nil),                  // 5: google.protobuf.Struct
}
var file_controller_api_resources_credentiallibraries_v1_credential_library_proto_depIdxs = []int32{
	2,  // 0: controller.api.resources.credentiallibraries.v1.CredentialLibrary.scope:type_name -> controller.api.resources.scopes.v1.ScopeInfo
	3,  // 1: controller.api.resources.credentiallibraries.v1.CredentialLibrary.name:type_name -> google.protobuf.StringValue
	3,  // 2: controller.api.resources.credentiallibraries.v1.CredentialLibrary.description:type_name -> google.protobuf.StringValue
	4,  // 3: controller.api.resources.credentiallibraries.v1.CredentialLibrary.created_time:type_name -> google.protobuf.Timestamp
	4,  // 4: controller.api.resources.credentiallibraries.v1.CredentialLibrary.updated_time:type_name -> google.protobuf.Timestamp
	5,  // 5: controller.api.resources.credentiallibraries.v1.CredentialLibrary.attributes:type_name -> google.protobuf.Struct
	1,  // 6: controller.api.resources.credentiallibraries.v1.CredentialLibrary.vault_credential_library_attributes:type_name -> controller.api.resources.credentiallibraries.v1.VaultCredentialLibraryAttributes
	5,  // 7: controller.api.resources.credentiallibraries.v1.CredentialLibrary.credential_mapping_overrides:type_name -> google.protobuf.Struct
	3,  // 8: controller.api.resources.credentiallibraries.v1.VaultCredentialLibraryAttributes.path:type_name -> google.protobuf.StringValue
	3,  // 9: controller.api.resources.credentiallibraries.v1.VaultCredentialLibraryAttributes.http_method:type_name -> google.protobuf.StringValue
	3,  // 10: controller.api.resources.credentiallibraries.v1.VaultCredentialLibraryAttributes.http_request_body:type_name -> google.protobuf.StringValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_controller_api_resources_credentiallibraries_v1_credential_library_proto_init() }
func file_controller_api_resources_credentiallibraries_v1_credential_library_proto_init() {
	if File_controller_api_resources_credentiallibraries_v1_credential_library_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_resources_credentiallibraries_v1_credential_library_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialLibrary); i {
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
		file_controller_api_resources_credentiallibraries_v1_credential_library_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultCredentialLibraryAttributes); i {
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
	file_controller_api_resources_credentiallibraries_v1_credential_library_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CredentialLibrary_Attributes)(nil),
		(*CredentialLibrary_VaultCredentialLibraryAttributes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_credentiallibraries_v1_credential_library_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_credentiallibraries_v1_credential_library_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_credentiallibraries_v1_credential_library_proto_msgTypes,
	}.Build()
	File_controller_api_resources_credentiallibraries_v1_credential_library_proto = out.File
	file_controller_api_resources_credentiallibraries_v1_credential_library_proto_rawDesc = nil
	file_controller_api_resources_credentiallibraries_v1_credential_library_proto_goTypes = nil
	file_controller_api_resources_credentiallibraries_v1_credential_library_proto_depIdxs = nil
}
