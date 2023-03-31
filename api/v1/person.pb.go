//
//make sure you have the protoc compiler
//and install the go plugins with
//go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//then generate with:
//
//protoc -I./api/v1 --go_out=api/v1/. --go_opt=paths=source_relative --go-grpc_out=api/v1/. --go-grpc_opt=paths=source_relative api/v1/person.proto
//
//for the grpc-gateway install the following:
//
//go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
//go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
//go get google.golang.org/protobuf/cmd/protoc-gen-go
//go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
//
//and then run:
//
//protoc -I./api/v1 --grpc-gateway_out api/v1 --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true api/v1/person.proto
//
//which generates person.pb.gw.go in api/v1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: person.proto

package v1

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type IdRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *IdRef) Reset() {
	*x = IdRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRef) ProtoMessage() {}

func (x *IdRef) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRef.ProtoReflect.Descriptor instead.
func (*IdRef) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *IdRef) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IdRef) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// IMPORTANT: tag "json_name" only works for protojson
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Active             bool                   `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	DateCreated        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	DateUpdated        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
	FullName           string                 `protobuf:"bytes,5,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	FirstName          string                 `protobuf:"bytes,6,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName           string                 `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email              string                 `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Orcid              string                 `protobuf:"bytes,9,opt,name=orcid,proto3" json:"orcid,omitempty"`
	OrcidToken         string                 `protobuf:"bytes,10,opt,name=orcid_token,json=orcidToken,proto3" json:"orcid_token,omitempty"`
	PreferredFirstName string                 `protobuf:"bytes,11,opt,name=preferred_first_name,json=preferredFirstName,proto3" json:"preferred_first_name,omitempty"`
	PreferredLastName  string                 `protobuf:"bytes,12,opt,name=preferred_last_name,json=preferredLastName,proto3" json:"preferred_last_name,omitempty"`
	BirthDate          string                 `protobuf:"bytes,13,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Title              string                 `protobuf:"bytes,14,opt,name=title,proto3" json:"title,omitempty"`
	OtherId            []*IdRef               `protobuf:"bytes,15,rep,name=other_id,json=otherId,proto3" json:"other_id,omitempty"`
	OrganizationId     []string               `protobuf:"bytes,16,rep,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Person) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

func (x *Person) GetDateUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateUpdated
	}
	return nil
}

func (x *Person) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Person) GetOrcid() string {
	if x != nil {
		return x.Orcid
	}
	return ""
}

func (x *Person) GetOrcidToken() string {
	if x != nil {
		return x.OrcidToken
	}
	return ""
}

func (x *Person) GetPreferredFirstName() string {
	if x != nil {
		return x.PreferredFirstName
	}
	return ""
}

func (x *Person) GetPreferredLastName() string {
	if x != nil {
		return x.PreferredLastName
	}
	return ""
}

func (x *Person) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *Person) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Person) GetOtherId() []*IdRef {
	if x != nil {
		return x.OtherId
	}
	return nil
}

func (x *Person) GetOrganizationId() []string {
	if x != nil {
		return x.OrganizationId
	}
	return nil
}

type GetPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPersonRequest) Reset() {
	*x = GetPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRequest) ProtoMessage() {}

func (x *GetPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRequest.ProtoReflect.Descriptor instead.
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{2}
}

func (x *GetPersonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*GetPersonResponse_Person
	//	*GetPersonResponse_Error
	Response isGetPersonResponse_Response `protobuf_oneof:"response"`
}

func (x *GetPersonResponse) Reset() {
	*x = GetPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonResponse) ProtoMessage() {}

func (x *GetPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonResponse.ProtoReflect.Descriptor instead.
func (*GetPersonResponse) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{3}
}

func (m *GetPersonResponse) GetResponse() isGetPersonResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *GetPersonResponse) GetPerson() *Person {
	if x, ok := x.GetResponse().(*GetPersonResponse_Person); ok {
		return x.Person
	}
	return nil
}

func (x *GetPersonResponse) GetError() *status.Status {
	if x, ok := x.GetResponse().(*GetPersonResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isGetPersonResponse_Response interface {
	isGetPersonResponse_Response()
}

type GetPersonResponse_Person struct {
	Person *Person `protobuf:"bytes,1,opt,name=person,proto3,oneof"`
}

type GetPersonResponse_Error struct {
	Error *status.Status `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*GetPersonResponse_Person) isGetPersonResponse_Response() {}

func (*GetPersonResponse_Error) isGetPersonResponse_Response() {}

type GetAllPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllPersonRequest) Reset() {
	*x = GetAllPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPersonRequest) ProtoMessage() {}

func (x *GetAllPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPersonRequest.ProtoReflect.Descriptor instead.
func (*GetAllPersonRequest) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{4}
}

type GetAllPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *GetAllPersonResponse) Reset() {
	*x = GetAllPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPersonResponse) ProtoMessage() {}

func (x *GetAllPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPersonResponse.ProtoReflect.Descriptor instead.
func (*GetAllPersonResponse) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllPersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2b, 0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xbe, 0x04,
	0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x3d, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x72, 0x63, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72,
	0x63, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x63, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x63, 0x69, 0x64, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65,
	0x64, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x66, 0x52, 0x07, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x22,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x75, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x32, 0x97, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData = file_person_proto_rawDesc
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_proto_rawDescData)
	})
	return file_person_proto_rawDescData
}

var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_person_proto_goTypes = []interface{}{
	(*IdRef)(nil),                 // 0: api.v1.IdRef
	(*Person)(nil),                // 1: api.v1.Person
	(*GetPersonRequest)(nil),      // 2: api.v1.GetPersonRequest
	(*GetPersonResponse)(nil),     // 3: api.v1.GetPersonResponse
	(*GetAllPersonRequest)(nil),   // 4: api.v1.GetAllPersonRequest
	(*GetAllPersonResponse)(nil),  // 5: api.v1.GetAllPersonResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*status.Status)(nil),         // 7: google.rpc.Status
}
var file_person_proto_depIdxs = []int32{
	6, // 0: api.v1.Person.date_created:type_name -> google.protobuf.Timestamp
	6, // 1: api.v1.Person.date_updated:type_name -> google.protobuf.Timestamp
	0, // 2: api.v1.Person.other_id:type_name -> api.v1.IdRef
	1, // 3: api.v1.GetPersonResponse.person:type_name -> api.v1.Person
	7, // 4: api.v1.GetPersonResponse.error:type_name -> google.rpc.Status
	1, // 5: api.v1.GetAllPersonResponse.person:type_name -> api.v1.Person
	2, // 6: api.v1.People.GetPerson:input_type -> api.v1.GetPersonRequest
	4, // 7: api.v1.People.GetAllPerson:input_type -> api.v1.GetAllPersonRequest
	3, // 8: api.v1.People.GetPerson:output_type -> api.v1.GetPersonResponse
	5, // 9: api.v1.People.GetAllPerson:output_type -> api.v1.GetAllPersonResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRef); i {
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
		file_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonRequest); i {
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
		file_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonResponse); i {
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
		file_person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPersonRequest); i {
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
		file_person_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPersonResponse); i {
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
	file_person_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GetPersonResponse_Person)(nil),
		(*GetPersonResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_rawDesc = nil
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}
