// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: spire/api/types/entry.proto

package types

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

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Globally unique ID for the entry.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The SPIFFE ID of the identity described by this entry.
	SpiffeId *SPIFFEID `protobuf:"bytes,2,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Who the entry is delegated to. If the entry describes a node, this is
	// set to the SPIFFE ID of the SPIRE server of the trust domain (e.g.
	// spiffe://example.org/spire/server). Otherwise, it will be set to a node
	// SPIFFE ID.
	ParentId *SPIFFEID `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// The selectors which identify which entities match this entry. If this is
	// an entry for a node, these selectors represent selectors produced by
	// node attestation. Otherwise, these selectors represent those produced by
	// workload attestation.
	Selectors []*Selector `protobuf:"bytes,4,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// The time to live for identities issued for this entry (in seconds).
	Ttl int32 `protobuf:"varint,5,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// The names of trust domains the identity described by this entry
	// federates with.
	FederatesWith []string `protobuf:"bytes,6,rep,name=federates_with,json=federatesWith,proto3" json:"federates_with,omitempty"`
	// Whether or not the identity described by this entry is an administrative
	// workload. Administrative workloads are granted additional access to
	// various managerial server APIs, such as entry registration.
	Admin bool `protobuf:"varint,7,opt,name=admin,proto3" json:"admin,omitempty"`
	// Whether or not the identity described by this entry represents a
	// downstream SPIRE server. Downstream SPIRE servers have additional access
	// to various signing APIs, such as those used to sign X.509 CA
	// certificates and publish JWT signing keys.
	Downstream bool `protobuf:"varint,8,opt,name=downstream,proto3" json:"downstream,omitempty"`
	// When the entry expires (seconds since Unix epoch).
	ExpiresAt int64 `protobuf:"varint,9,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// A list of DNS names associated with the identity described by this entry.
	DnsNames []string `protobuf:"bytes,10,rep,name=dns_names,json=dnsNames,proto3" json:"dns_names,omitempty"`
	// Revision number is bumped every time the entry is updated
	RevisionNumber int64 `protobuf:"varint,11,opt,name=revision_number,json=revisionNumber,proto3" json:"revision_number,omitempty"`
	// Determines if the issued identity is exportable to a store
	StoreSvid bool `protobuf:"varint,12,opt,name=store_svid,json=storeSvid,proto3" json:"store_svid,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_types_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_types_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_spire_api_types_entry_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Entry) GetSpiffeId() *SPIFFEID {
	if x != nil {
		return x.SpiffeId
	}
	return nil
}

func (x *Entry) GetParentId() *SPIFFEID {
	if x != nil {
		return x.ParentId
	}
	return nil
}

func (x *Entry) GetSelectors() []*Selector {
	if x != nil {
		return x.Selectors
	}
	return nil
}

func (x *Entry) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *Entry) GetFederatesWith() []string {
	if x != nil {
		return x.FederatesWith
	}
	return nil
}

func (x *Entry) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *Entry) GetDownstream() bool {
	if x != nil {
		return x.Downstream
	}
	return false
}

func (x *Entry) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *Entry) GetDnsNames() []string {
	if x != nil {
		return x.DnsNames
	}
	return nil
}

func (x *Entry) GetRevisionNumber() int64 {
	if x != nil {
		return x.RevisionNumber
	}
	return 0
}

func (x *Entry) GetStoreSvid() bool {
	if x != nil {
		return x.StoreSvid
	}
	return false
}

// Field mask for Entry fields
type EntryMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// spiffe_id field mask
	SpiffeId bool `protobuf:"varint,2,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// parent_id field mask
	ParentId bool `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// selectors field mask
	Selectors bool `protobuf:"varint,4,opt,name=selectors,proto3" json:"selectors,omitempty"`
	// ttl field mask
	Ttl bool `protobuf:"varint,5,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// federates_with field mask
	FederatesWith bool `protobuf:"varint,6,opt,name=federates_with,json=federatesWith,proto3" json:"federates_with,omitempty"`
	// admin field mask
	Admin bool `protobuf:"varint,7,opt,name=admin,proto3" json:"admin,omitempty"`
	// downstream field mask
	Downstream bool `protobuf:"varint,8,opt,name=downstream,proto3" json:"downstream,omitempty"`
	// expires_at field mask
	ExpiresAt bool `protobuf:"varint,9,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// dns_names field mask
	DnsNames bool `protobuf:"varint,10,opt,name=dns_names,json=dnsNames,proto3" json:"dns_names,omitempty"`
	// revision_number field mask
	RevisionNumber bool `protobuf:"varint,11,opt,name=revision_number,json=revisionNumber,proto3" json:"revision_number,omitempty"`
	// store_svid field mask
	StoreSvid bool `protobuf:"varint,12,opt,name=store_svid,json=storeSvid,proto3" json:"store_svid,omitempty"`
}

func (x *EntryMask) Reset() {
	*x = EntryMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_types_entry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryMask) ProtoMessage() {}

func (x *EntryMask) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_types_entry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryMask.ProtoReflect.Descriptor instead.
func (*EntryMask) Descriptor() ([]byte, []int) {
	return file_spire_api_types_entry_proto_rawDescGZIP(), []int{1}
}

func (x *EntryMask) GetSpiffeId() bool {
	if x != nil {
		return x.SpiffeId
	}
	return false
}

func (x *EntryMask) GetParentId() bool {
	if x != nil {
		return x.ParentId
	}
	return false
}

func (x *EntryMask) GetSelectors() bool {
	if x != nil {
		return x.Selectors
	}
	return false
}

func (x *EntryMask) GetTtl() bool {
	if x != nil {
		return x.Ttl
	}
	return false
}

func (x *EntryMask) GetFederatesWith() bool {
	if x != nil {
		return x.FederatesWith
	}
	return false
}

func (x *EntryMask) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *EntryMask) GetDownstream() bool {
	if x != nil {
		return x.Downstream
	}
	return false
}

func (x *EntryMask) GetExpiresAt() bool {
	if x != nil {
		return x.ExpiresAt
	}
	return false
}

func (x *EntryMask) GetDnsNames() bool {
	if x != nil {
		return x.DnsNames
	}
	return false
}

func (x *EntryMask) GetRevisionNumber() bool {
	if x != nil {
		return x.RevisionNumber
	}
	return false
}

func (x *EntryMask) GetStoreSvid() bool {
	if x != nil {
		return x.StoreSvid
	}
	return false
}

var File_spire_api_types_entry_proto protoreflect.FileDescriptor

var file_spire_api_types_entry_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3,
	0x03, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x70, 0x69, 0x66,
	0x66, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x50,
	0x49, 0x46, 0x46, 0x45, 0x49, 0x44, 0x52, 0x08, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x49, 0x64,
	0x12, 0x36, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x49, 0x44, 0x52, 0x08,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x74, 0x74, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73,
	0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x6e, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73,
	0x76, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x53, 0x76, 0x69, 0x64, 0x22, 0xd6, 0x02, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x25, 0x0a, 0x0e,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x77,
	0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64,
	0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6e, 0x73, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x6e, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x76, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x76, 0x69, 0x64, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66,
	0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_types_entry_proto_rawDescOnce sync.Once
	file_spire_api_types_entry_proto_rawDescData = file_spire_api_types_entry_proto_rawDesc
)

func file_spire_api_types_entry_proto_rawDescGZIP() []byte {
	file_spire_api_types_entry_proto_rawDescOnce.Do(func() {
		file_spire_api_types_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_types_entry_proto_rawDescData)
	})
	return file_spire_api_types_entry_proto_rawDescData
}

var file_spire_api_types_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spire_api_types_entry_proto_goTypes = []interface{}{
	(*Entry)(nil),     // 0: spire.api.types.Entry
	(*EntryMask)(nil), // 1: spire.api.types.EntryMask
	(*SPIFFEID)(nil),  // 2: spire.api.types.SPIFFEID
	(*Selector)(nil),  // 3: spire.api.types.Selector
}
var file_spire_api_types_entry_proto_depIdxs = []int32{
	2, // 0: spire.api.types.Entry.spiffe_id:type_name -> spire.api.types.SPIFFEID
	2, // 1: spire.api.types.Entry.parent_id:type_name -> spire.api.types.SPIFFEID
	3, // 2: spire.api.types.Entry.selectors:type_name -> spire.api.types.Selector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spire_api_types_entry_proto_init() }
func file_spire_api_types_entry_proto_init() {
	if File_spire_api_types_entry_proto != nil {
		return
	}
	file_spire_api_types_selector_proto_init()
	file_spire_api_types_spiffeid_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spire_api_types_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_spire_api_types_entry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryMask); i {
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
			RawDescriptor: file_spire_api_types_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spire_api_types_entry_proto_goTypes,
		DependencyIndexes: file_spire_api_types_entry_proto_depIdxs,
		MessageInfos:      file_spire_api_types_entry_proto_msgTypes,
	}.Build()
	File_spire_api_types_entry_proto = out.File
	file_spire_api_types_entry_proto_rawDesc = nil
	file_spire_api_types_entry_proto_goTypes = nil
	file_spire_api_types_entry_proto_depIdxs = nil
}