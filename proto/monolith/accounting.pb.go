// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: monolith/accounting.proto

package monolith

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity  *Keygraph           `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Peers     []*Account_Peer     `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	Providers []*Account_Provider `protobuf:"bytes,3,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monolith_accounting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_monolith_accounting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_monolith_accounting_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetIdentity() *Keygraph {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Account) GetPeers() []*Account_Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *Account) GetProviders() []*Account_Provider {
	if x != nil {
		return x.Providers
	}
	return nil
}

type RegisterPeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerInfo *PeerInfo `protobuf:"bytes,1,opt,name=peer_info,json=peerInfo,proto3" json:"peer_info,omitempty"`
	Topics   []string  `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *RegisterPeerRequest) Reset() {
	*x = RegisterPeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monolith_accounting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPeerRequest) ProtoMessage() {}

func (x *RegisterPeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monolith_accounting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPeerRequest.ProtoReflect.Descriptor instead.
func (*RegisterPeerRequest) Descriptor() ([]byte, []int) {
	return file_monolith_accounting_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterPeerRequest) GetPeerInfo() *PeerInfo {
	if x != nil {
		return x.PeerInfo
	}
	return nil
}

func (x *RegisterPeerRequest) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type RegisterPeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer *Account_Peer `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *RegisterPeerResponse) Reset() {
	*x = RegisterPeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monolith_accounting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPeerResponse) ProtoMessage() {}

func (x *RegisterPeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monolith_accounting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPeerResponse.ProtoReflect.Descriptor instead.
func (*RegisterPeerResponse) Descriptor() ([]byte, []int) {
	return file_monolith_accounting_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterPeerResponse) GetPeer() *Account_Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

type Account_Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerInfo *PeerInfo `protobuf:"bytes,1,opt,name=peer_info,json=peerInfo,proto3" json:"peer_info,omitempty"`
	Topics   []string  `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *Account_Peer) Reset() {
	*x = Account_Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monolith_accounting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account_Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account_Peer) ProtoMessage() {}

func (x *Account_Peer) ProtoReflect() protoreflect.Message {
	mi := &file_monolith_accounting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account_Peer.ProtoReflect.Descriptor instead.
func (*Account_Peer) Descriptor() ([]byte, []int) {
	return file_monolith_accounting_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Account_Peer) GetPeerInfo() *PeerInfo {
	if x != nil {
		return x.PeerInfo
	}
	return nil
}

func (x *Account_Peer) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type Account_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Methods     []string `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *Account_Service) Reset() {
	*x = Account_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monolith_accounting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account_Service) ProtoMessage() {}

func (x *Account_Service) ProtoReflect() protoreflect.Message {
	mi := &file_monolith_accounting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account_Service.ProtoReflect.Descriptor instead.
func (*Account_Service) Descriptor() ([]byte, []int) {
	return file_monolith_accounting_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Account_Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account_Service) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Account_Service) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

type Account_Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *Keygraph          `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Peers    []*Account_Peer    `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	Services []*Account_Service `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *Account_Provider) Reset() {
	*x = Account_Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monolith_accounting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account_Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account_Provider) ProtoMessage() {}

func (x *Account_Provider) ProtoReflect() protoreflect.Message {
	mi := &file_monolith_accounting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account_Provider.ProtoReflect.Descriptor instead.
func (*Account_Provider) Descriptor() ([]byte, []int) {
	return file_monolith_accounting_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Account_Provider) GetIdentity() *Keygraph {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Account_Provider) GetPeers() []*Account_Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *Account_Provider) GetServices() []*Account_Service {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_monolith_accounting_proto protoreflect.FileDescriptor

var file_monolith_accounting_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x6f, 0x6e,
	0x6f, 0x6c, 0x69, 0x74, 0x68, 0x1a, 0x14, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x03, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f,
	0x6c, 0x69, 0x74, 0x68, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74,
	0x68, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c,
	0x69, 0x74, 0x68, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x1a,
	0x4f, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x6e,
	0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x70, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x1a, 0x59, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x9f, 0x01, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x6e,
	0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69,
	0x74, 0x68, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c,
	0x69, 0x74, 0x68, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x5e, 0x0a,
	0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69,
	0x74, 0x68, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x65, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0x42, 0x0a,
	0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65,
	0x72, 0x32, 0x62, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x83, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f,
	0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x42, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x72, 0x65, 0x61, 0x6d, 0x64,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0xa2, 0x02, 0x03,
	0x4d, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x4d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0xca, 0x02,
	0x08, 0x4d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0xe2, 0x02, 0x14, 0x4d, 0x6f, 0x6e, 0x6f,
	0x6c, 0x69, 0x74, 0x68, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x08, 0x4d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_monolith_accounting_proto_rawDescOnce sync.Once
	file_monolith_accounting_proto_rawDescData = file_monolith_accounting_proto_rawDesc
)

func file_monolith_accounting_proto_rawDescGZIP() []byte {
	file_monolith_accounting_proto_rawDescOnce.Do(func() {
		file_monolith_accounting_proto_rawDescData = protoimpl.X.CompressGZIP(file_monolith_accounting_proto_rawDescData)
	})
	return file_monolith_accounting_proto_rawDescData
}

var file_monolith_accounting_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_monolith_accounting_proto_goTypes = []interface{}{
	(*Account)(nil),              // 0: monolith.Account
	(*RegisterPeerRequest)(nil),  // 1: monolith.RegisterPeerRequest
	(*RegisterPeerResponse)(nil), // 2: monolith.RegisterPeerResponse
	(*Account_Peer)(nil),         // 3: monolith.Account.Peer
	(*Account_Service)(nil),      // 4: monolith.Account.Service
	(*Account_Provider)(nil),     // 5: monolith.Account.Provider
	(*Keygraph)(nil),             // 6: monolith.Keygraph
	(*PeerInfo)(nil),             // 7: monolith.PeerInfo
}
var file_monolith_accounting_proto_depIdxs = []int32{
	6,  // 0: monolith.Account.identity:type_name -> monolith.Keygraph
	3,  // 1: monolith.Account.peers:type_name -> monolith.Account.Peer
	5,  // 2: monolith.Account.providers:type_name -> monolith.Account.Provider
	7,  // 3: monolith.RegisterPeerRequest.peer_info:type_name -> monolith.PeerInfo
	3,  // 4: monolith.RegisterPeerResponse.peer:type_name -> monolith.Account.Peer
	7,  // 5: monolith.Account.Peer.peer_info:type_name -> monolith.PeerInfo
	6,  // 6: monolith.Account.Provider.identity:type_name -> monolith.Keygraph
	3,  // 7: monolith.Account.Provider.peers:type_name -> monolith.Account.Peer
	4,  // 8: monolith.Account.Provider.services:type_name -> monolith.Account.Service
	1,  // 9: monolith.AccountingService.RegisterPeer:input_type -> monolith.RegisterPeerRequest
	2,  // 10: monolith.AccountingService.RegisterPeer:output_type -> monolith.RegisterPeerResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_monolith_accounting_proto_init() }
func file_monolith_accounting_proto_init() {
	if File_monolith_accounting_proto != nil {
		return
	}
	file_monolith_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_monolith_accounting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_monolith_accounting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPeerRequest); i {
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
		file_monolith_accounting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPeerResponse); i {
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
		file_monolith_accounting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account_Peer); i {
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
		file_monolith_accounting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account_Service); i {
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
		file_monolith_accounting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account_Provider); i {
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
			RawDescriptor: file_monolith_accounting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monolith_accounting_proto_goTypes,
		DependencyIndexes: file_monolith_accounting_proto_depIdxs,
		MessageInfos:      file_monolith_accounting_proto_msgTypes,
	}.Build()
	File_monolith_accounting_proto = out.File
	file_monolith_accounting_proto_rawDesc = nil
	file_monolith_accounting_proto_goTypes = nil
	file_monolith_accounting_proto_depIdxs = nil
}
