// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: token/v1/token.proto

package v1

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

// The request message containing the user's name.
type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick    string `protobuf:"bytes,1,opt,name=tick,proto3" json:"tick,omitempty"`
	TokenId uint64 `protobuf:"varint,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	P       string `protobuf:"bytes,3,opt,name=p,proto3" json:"p,omitempty"`
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_v1_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_v1_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_v1_token_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokenRequest) GetTick() string {
	if x != nil {
		return x.Tick
	}
	return ""
}

func (x *GetTokenRequest) GetTokenId() uint64 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

func (x *GetTokenRequest) GetP() string {
	if x != nil {
		return x.P
	}
	return ""
}

type TokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *TokenMessage `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TokenReply) Reset() {
	*x = TokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_v1_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReply) ProtoMessage() {}

func (x *TokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_token_v1_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReply.ProtoReflect.Descriptor instead.
func (*TokenReply) Descriptor() ([]byte, []int) {
	return file_token_v1_token_proto_rawDescGZIP(), []int{1}
}

func (x *TokenReply) GetData() *TokenMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

// The response message containing the token
type TokenMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P              string                 `protobuf:"bytes,1,opt,name=p,proto3" json:"p,omitempty"`
	Tick           string                 `protobuf:"bytes,2,opt,name=tick,proto3" json:"tick,omitempty"`
	TokenId        uint64                 `protobuf:"varint,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	TxHash         string                 `protobuf:"bytes,4,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	BlockHeight    uint64                 `protobuf:"varint,5,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	BlockTime      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	Address        string                 `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	InscriptionId  int64                  `protobuf:"varint,8,opt,name=inscription_id,json=inscriptionId,proto3" json:"inscription_id,omitempty"`
	InscriptionUid string                 `protobuf:"bytes,9,opt,name=inscription_uid,json=inscriptionUid,proto3" json:"inscription_uid,omitempty"`
}

func (x *TokenMessage) Reset() {
	*x = TokenMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_v1_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMessage) ProtoMessage() {}

func (x *TokenMessage) ProtoReflect() protoreflect.Message {
	mi := &file_token_v1_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMessage.ProtoReflect.Descriptor instead.
func (*TokenMessage) Descriptor() ([]byte, []int) {
	return file_token_v1_token_proto_rawDescGZIP(), []int{2}
}

func (x *TokenMessage) GetP() string {
	if x != nil {
		return x.P
	}
	return ""
}

func (x *TokenMessage) GetTick() string {
	if x != nil {
		return x.Tick
	}
	return ""
}

func (x *TokenMessage) GetTokenId() uint64 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

func (x *TokenMessage) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *TokenMessage) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *TokenMessage) GetBlockTime() *timestamppb.Timestamp {
	if x != nil {
		return x.BlockTime
	}
	return nil
}

func (x *TokenMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TokenMessage) GetInscriptionId() int64 {
	if x != nil {
		return x.InscriptionId
	}
	return 0
}

func (x *TokenMessage) GetInscriptionUid() string {
	if x != nil {
		return x.InscriptionUid
	}
	return ""
}

type ListTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P       string `protobuf:"bytes,1,opt,name=p,proto3" json:"p,omitempty"`
	Tick    string `protobuf:"bytes,2,opt,name=tick,proto3" json:"tick,omitempty"`
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Limit   uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  uint64 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListTokenRequest) Reset() {
	*x = ListTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_v1_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTokenRequest) ProtoMessage() {}

func (x *ListTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_v1_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTokenRequest.ProtoReflect.Descriptor instead.
func (*ListTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_v1_token_proto_rawDescGZIP(), []int{3}
}

func (x *ListTokenRequest) GetP() string {
	if x != nil {
		return x.P
	}
	return ""
}

func (x *ListTokenRequest) GetTick() string {
	if x != nil {
		return x.Tick
	}
	return ""
}

func (x *ListTokenRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListTokenRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTokenRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []*TokenMessage `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Paging *Paging         `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *ListTokenReply) Reset() {
	*x = ListTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_v1_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTokenReply) ProtoMessage() {}

func (x *ListTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_token_v1_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTokenReply.ProtoReflect.Descriptor instead.
func (*ListTokenReply) Descriptor() ([]byte, []int) {
	return file_token_v1_token_proto_rawDescGZIP(), []int{4}
}

func (x *ListTokenReply) GetData() []*TokenMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListTokenReply) GetPaging() *Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type Paging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount uint64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Count      uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Paging) Reset() {
	*x = Paging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_v1_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paging) ProtoMessage() {}

func (x *Paging) ProtoReflect() protoreflect.Message {
	mi := &file_token_v1_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paging.ProtoReflect.Descriptor instead.
func (*Paging) Descriptor() ([]byte, []int) {
	return file_token_v1_token_proto_rawDescGZIP(), []int{5}
}

func (x *Paging) GetTotalCount() uint64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *Paging) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_token_v1_token_proto protoreflect.FileDescriptor

var file_token_v1_token_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x64, 0x12, 0x0c, 0x0a, 0x01, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x70, 0x22,
	0x38, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xac, 0x02, 0x0a, 0x0c, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x69, 0x64, 0x22, 0x7d, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x66, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22,
	0x3f, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0xc2, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x61, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12,
	0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x7b, 0x74, 0x69, 0x63,
	0x6b, 0x7d, 0x2f, 0x7b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x56, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x5d, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x73,
	0x68, 0x61, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x2d, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_v1_token_proto_rawDescOnce sync.Once
	file_token_v1_token_proto_rawDescData = file_token_v1_token_proto_rawDesc
)

func file_token_v1_token_proto_rawDescGZIP() []byte {
	file_token_v1_token_proto_rawDescOnce.Do(func() {
		file_token_v1_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_v1_token_proto_rawDescData)
	})
	return file_token_v1_token_proto_rawDescData
}

var file_token_v1_token_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_token_v1_token_proto_goTypes = []interface{}{
	(*GetTokenRequest)(nil),       // 0: token.v1.GetTokenRequest
	(*TokenReply)(nil),            // 1: token.v1.TokenReply
	(*TokenMessage)(nil),          // 2: token.v1.TokenMessage
	(*ListTokenRequest)(nil),      // 3: token.v1.ListTokenRequest
	(*ListTokenReply)(nil),        // 4: token.v1.ListTokenReply
	(*Paging)(nil),                // 5: token.v1.Paging
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_token_v1_token_proto_depIdxs = []int32{
	2, // 0: token.v1.TokenReply.data:type_name -> token.v1.TokenMessage
	6, // 1: token.v1.TokenMessage.block_time:type_name -> google.protobuf.Timestamp
	2, // 2: token.v1.ListTokenReply.data:type_name -> token.v1.TokenMessage
	5, // 3: token.v1.ListTokenReply.paging:type_name -> token.v1.Paging
	0, // 4: token.v1.Token.GetToken:input_type -> token.v1.GetTokenRequest
	3, // 5: token.v1.Token.ListTokens:input_type -> token.v1.ListTokenRequest
	1, // 6: token.v1.Token.GetToken:output_type -> token.v1.TokenReply
	4, // 7: token.v1.Token.ListTokens:output_type -> token.v1.ListTokenReply
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_token_v1_token_proto_init() }
func file_token_v1_token_proto_init() {
	if File_token_v1_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_v1_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenRequest); i {
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
		file_token_v1_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReply); i {
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
		file_token_v1_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenMessage); i {
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
		file_token_v1_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTokenRequest); i {
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
		file_token_v1_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTokenReply); i {
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
		file_token_v1_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paging); i {
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
			RawDescriptor: file_token_v1_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_v1_token_proto_goTypes,
		DependencyIndexes: file_token_v1_token_proto_depIdxs,
		MessageInfos:      file_token_v1_token_proto_msgTypes,
	}.Build()
	File_token_v1_token_proto = out.File
	file_token_v1_token_proto_rawDesc = nil
	file_token_v1_token_proto_goTypes = nil
	file_token_v1_token_proto_depIdxs = nil
}
