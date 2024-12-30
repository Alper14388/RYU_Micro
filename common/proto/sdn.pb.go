// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: proto/sdn.proto

package proto

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

// Flow Add Messages
type FlowAddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SwitchId      uint64                 `protobuf:"varint,1,opt,name=switch_id,json=switchId,proto3" json:"switch_id,omitempty"`
	InPort        uint32                 `protobuf:"varint,2,opt,name=in_port,json=inPort,proto3" json:"in_port,omitempty"`
	Src           string                 `protobuf:"bytes,3,opt,name=src,proto3" json:"src,omitempty"`
	Dst           string                 `protobuf:"bytes,4,opt,name=dst,proto3" json:"dst,omitempty"`
	OutPort       uint32                 `protobuf:"varint,5,opt,name=out_port,json=outPort,proto3" json:"out_port,omitempty"`
	Priority      uint32                 `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	HardTimeout   uint32                 `protobuf:"varint,7,opt,name=hard_timeout,json=hardTimeout,proto3" json:"hard_timeout,omitempty"`
	IdleTimeout   uint32                 `protobuf:"varint,8,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	BufferId      uint32                 `protobuf:"varint,9,opt,name=buffer_id,json=bufferId,proto3" json:"buffer_id,omitempty"`
	TableId       uint32                 `protobuf:"varint,10,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Flags         uint32                 `protobuf:"varint,11,opt,name=flags,proto3" json:"flags,omitempty"`
	Cookie        uint64                 `protobuf:"varint,12,opt,name=cookie,proto3" json:"cookie,omitempty"`
	CookieMask    uint64                 `protobuf:"varint,13,opt,name=cookie_mask,json=cookieMask,proto3" json:"cookie_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlowAddRequest) Reset() {
	*x = FlowAddRequest{}
	mi := &file_proto_sdn_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowAddRequest) ProtoMessage() {}

func (x *FlowAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sdn_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowAddRequest.ProtoReflect.Descriptor instead.
func (*FlowAddRequest) Descriptor() ([]byte, []int) {
	return file_proto_sdn_proto_rawDescGZIP(), []int{0}
}

func (x *FlowAddRequest) GetSwitchId() uint64 {
	if x != nil {
		return x.SwitchId
	}
	return 0
}

func (x *FlowAddRequest) GetInPort() uint32 {
	if x != nil {
		return x.InPort
	}
	return 0
}

func (x *FlowAddRequest) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *FlowAddRequest) GetDst() string {
	if x != nil {
		return x.Dst
	}
	return ""
}

func (x *FlowAddRequest) GetOutPort() uint32 {
	if x != nil {
		return x.OutPort
	}
	return 0
}

func (x *FlowAddRequest) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *FlowAddRequest) GetHardTimeout() uint32 {
	if x != nil {
		return x.HardTimeout
	}
	return 0
}

func (x *FlowAddRequest) GetIdleTimeout() uint32 {
	if x != nil {
		return x.IdleTimeout
	}
	return 0
}

func (x *FlowAddRequest) GetBufferId() uint32 {
	if x != nil {
		return x.BufferId
	}
	return 0
}

func (x *FlowAddRequest) GetTableId() uint32 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *FlowAddRequest) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *FlowAddRequest) GetCookie() uint64 {
	if x != nil {
		return x.Cookie
	}
	return 0
}

func (x *FlowAddRequest) GetCookieMask() uint64 {
	if x != nil {
		return x.CookieMask
	}
	return 0
}

type FlowAddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlowAddResponse) Reset() {
	*x = FlowAddResponse{}
	mi := &file_proto_sdn_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowAddResponse) ProtoMessage() {}

func (x *FlowAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sdn_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowAddResponse.ProtoReflect.Descriptor instead.
func (*FlowAddResponse) Descriptor() ([]byte, []int) {
	return file_proto_sdn_proto_rawDescGZIP(), []int{1}
}

func (x *FlowAddResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FlowAddResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Flow Mod Messages
type FlowModRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Command       uint32                 `protobuf:"varint,2,opt,name=command,proto3" json:"command,omitempty"`
	Flags         uint32                 `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
	TableId       uint32                 `protobuf:"varint,4,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Instructions  []*Instruction         `protobuf:"bytes,5,rep,name=instructions,proto3" json:"instructions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlowModRequest) Reset() {
	*x = FlowModRequest{}
	mi := &file_proto_sdn_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowModRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowModRequest) ProtoMessage() {}

func (x *FlowModRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sdn_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowModRequest.ProtoReflect.Descriptor instead.
func (*FlowModRequest) Descriptor() ([]byte, []int) {
	return file_proto_sdn_proto_rawDescGZIP(), []int{2}
}

func (x *FlowModRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FlowModRequest) GetCommand() uint32 {
	if x != nil {
		return x.Command
	}
	return 0
}

func (x *FlowModRequest) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *FlowModRequest) GetTableId() uint32 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *FlowModRequest) GetInstructions() []*Instruction {
	if x != nil {
		return x.Instructions
	}
	return nil
}

type FlowModResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorCode     uint32                 `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlowModResponse) Reset() {
	*x = FlowModResponse{}
	mi := &file_proto_sdn_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowModResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowModResponse) ProtoMessage() {}

func (x *FlowModResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sdn_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowModResponse.ProtoReflect.Descriptor instead.
func (*FlowModResponse) Descriptor() ([]byte, []int) {
	return file_proto_sdn_proto_rawDescGZIP(), []int{3}
}

func (x *FlowModResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FlowModResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FlowModResponse) GetErrorCode() uint32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

// Match Field Messages
type MatchField struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Class         uint32                 `protobuf:"varint,1,opt,name=class,proto3" json:"class,omitempty"`
	Field         uint32                 `protobuf:"varint,2,opt,name=field,proto3" json:"field,omitempty"`
	Value         []byte                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Mask          []byte                 `protobuf:"bytes,4,opt,name=mask,proto3" json:"mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchField) Reset() {
	*x = MatchField{}
	mi := &file_proto_sdn_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchField) ProtoMessage() {}

func (x *MatchField) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sdn_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchField.ProtoReflect.Descriptor instead.
func (*MatchField) Descriptor() ([]byte, []int) {
	return file_proto_sdn_proto_rawDescGZIP(), []int{4}
}

func (x *MatchField) GetClass() uint32 {
	if x != nil {
		return x.Class
	}
	return 0
}

func (x *MatchField) GetField() uint32 {
	if x != nil {
		return x.Field
	}
	return 0
}

func (x *MatchField) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *MatchField) GetMask() []byte {
	if x != nil {
		return x.Mask
	}
	return nil
}

// Action Messages
type Action struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          uint32                 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Port          uint32                 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	MaxLen        uint32                 `protobuf:"varint,3,opt,name=max_len,json=maxLen,proto3" json:"max_len,omitempty"`
	Data          []byte                 `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Action) Reset() {
	*x = Action{}
	mi := &file_proto_sdn_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sdn_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_proto_sdn_proto_rawDescGZIP(), []int{5}
}

func (x *Action) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Action) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Action) GetMaxLen() uint32 {
	if x != nil {
		return x.MaxLen
	}
	return 0
}

func (x *Action) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Instruction Messages
type Instruction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          uint32                 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`      // Instruction type
	Actions       []*Action              `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"` // For apply_actions
	Data          []byte                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`       // Additional instruction data
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Instruction) Reset() {
	*x = Instruction{}
	mi := &file_proto_sdn_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Instruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instruction) ProtoMessage() {}

func (x *Instruction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sdn_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instruction.ProtoReflect.Descriptor instead.
func (*Instruction) Descriptor() ([]byte, []int) {
	return file_proto_sdn_proto_rawDescGZIP(), []int{6}
}

func (x *Instruction) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Instruction) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Instruction) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PacketInRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SwitchId      uint64                 `protobuf:"varint,1,opt,name=switch_id,json=switchId,proto3" json:"switch_id,omitempty"`
	BufferId      uint32                 `protobuf:"varint,2,opt,name=buffer_id,json=bufferId,proto3" json:"buffer_id,omitempty"`
	Length        uint32                 `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	Reason        uint32                 `protobuf:"varint,4,opt,name=reason,proto3" json:"reason,omitempty"`
	TableId       uint32                 `protobuf:"varint,5,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Cookie        uint64                 `protobuf:"varint,6,opt,name=cookie,proto3" json:"cookie,omitempty"`
	MatchFields   []*MatchField          `protobuf:"bytes,7,rep,name=match_fields,json=matchFields,proto3" json:"match_fields,omitempty"`
	Data          []byte                 `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	TotalLen      uint32                 `protobuf:"varint,9,opt,name=total_len,json=totalLen,proto3" json:"total_len,omitempty"`
	InPort        uint32                 `protobuf:"varint,10,opt,name=in_port,json=inPort,proto3" json:"in_port,omitempty"`
	InPhyPort     uint32                 `protobuf:"varint,11,opt,name=in_phy_port,json=inPhyPort,proto3" json:"in_phy_port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PacketInRequest) Reset() {
	*x = PacketInRequest{}
	mi := &file_proto_sdn_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketInRequest) ProtoMessage() {}

func (x *PacketInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sdn_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketInRequest.ProtoReflect.Descriptor instead.
func (*PacketInRequest) Descriptor() ([]byte, []int) {
	return file_proto_sdn_proto_rawDescGZIP(), []int{7}
}

func (x *PacketInRequest) GetSwitchId() uint64 {
	if x != nil {
		return x.SwitchId
	}
	return 0
}

func (x *PacketInRequest) GetBufferId() uint32 {
	if x != nil {
		return x.BufferId
	}
	return 0
}

func (x *PacketInRequest) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *PacketInRequest) GetReason() uint32 {
	if x != nil {
		return x.Reason
	}
	return 0
}

func (x *PacketInRequest) GetTableId() uint32 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *PacketInRequest) GetCookie() uint64 {
	if x != nil {
		return x.Cookie
	}
	return 0
}

func (x *PacketInRequest) GetMatchFields() []*MatchField {
	if x != nil {
		return x.MatchFields
	}
	return nil
}

func (x *PacketInRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PacketInRequest) GetTotalLen() uint32 {
	if x != nil {
		return x.TotalLen
	}
	return 0
}

func (x *PacketInRequest) GetInPort() uint32 {
	if x != nil {
		return x.InPort
	}
	return 0
}

func (x *PacketInRequest) GetInPhyPort() uint32 {
	if x != nil {
		return x.InPhyPort
	}
	return 0
}

type PacketInResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorCode     uint32                 `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PacketInResponse) Reset() {
	*x = PacketInResponse{}
	mi := &file_proto_sdn_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketInResponse) ProtoMessage() {}

func (x *PacketInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sdn_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketInResponse.ProtoReflect.Descriptor instead.
func (*PacketInResponse) Descriptor() ([]byte, []int) {
	return file_proto_sdn_proto_rawDescGZIP(), []int{8}
}

func (x *PacketInResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PacketInResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PacketInResponse) GetErrorCode() uint32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

var File_proto_sdn_proto protoreflect.FileDescriptor

var file_proto_sdn_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x64, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x73, 0x64, 0x6e, 0x22, 0xee, 0x02, 0x0a, 0x0e, 0x46, 0x6c, 0x6f, 0x77, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72,
	0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61,
	0x72, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x68, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x45, 0x0a, 0x0f, 0x46, 0x6c, 0x6f, 0x77, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa5,
	0x01, 0x0a, 0x0e, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x64, 0x6e, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x64, 0x0a, 0x0f, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x6f,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x62, 0x0a, 0x0a,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b,
	0x22, 0x5d, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x5c, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x64, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcc, 0x02,
	0x0a, 0x0f, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x32,
	0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x64, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x6c, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x4c, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0b,
	0x69, 0x6e, 0x5f, 0x70, 0x68, 0x79, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x69, 0x6e, 0x50, 0x68, 0x79, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x65, 0x0a, 0x10,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x4f, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x12, 0x13, 0x2e, 0x73, 0x64, 0x6e, 0x2e, 0x46, 0x6c,
	0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73,
	0x64, 0x6e, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0x4e, 0x0a, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x12, 0x14, 0x2e, 0x73, 0x64, 0x6e, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x73, 0x64, 0x6e, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x47, 0x0a, 0x0d, 0x46, 0x6c, 0x6f, 0x77, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x46, 0x6c, 0x6f, 0x77,
	0x12, 0x13, 0x2e, 0x73, 0x64, 0x6e, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x64, 0x6e, 0x2e, 0x46, 0x6c, 0x6f, 0x77,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x73, 0x64, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_sdn_proto_rawDescOnce sync.Once
	file_proto_sdn_proto_rawDescData = file_proto_sdn_proto_rawDesc
)

func file_proto_sdn_proto_rawDescGZIP() []byte {
	file_proto_sdn_proto_rawDescOnce.Do(func() {
		file_proto_sdn_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sdn_proto_rawDescData)
	})
	return file_proto_sdn_proto_rawDescData
}

var file_proto_sdn_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_sdn_proto_goTypes = []any{
	(*FlowAddRequest)(nil),   // 0: sdn.FlowAddRequest
	(*FlowAddResponse)(nil),  // 1: sdn.FlowAddResponse
	(*FlowModRequest)(nil),   // 2: sdn.FlowModRequest
	(*FlowModResponse)(nil),  // 3: sdn.FlowModResponse
	(*MatchField)(nil),       // 4: sdn.MatchField
	(*Action)(nil),           // 5: sdn.Action
	(*Instruction)(nil),      // 6: sdn.Instruction
	(*PacketInRequest)(nil),  // 7: sdn.PacketInRequest
	(*PacketInResponse)(nil), // 8: sdn.PacketInResponse
}
var file_proto_sdn_proto_depIdxs = []int32{
	6, // 0: sdn.FlowModRequest.instructions:type_name -> sdn.Instruction
	5, // 1: sdn.Instruction.actions:type_name -> sdn.Action
	4, // 2: sdn.PacketInRequest.match_fields:type_name -> sdn.MatchField
	2, // 3: sdn.ConnectionManager.SendFlowMod:input_type -> sdn.FlowModRequest
	7, // 4: sdn.PacketHandler.HandlePacketIn:input_type -> sdn.PacketInRequest
	0, // 5: sdn.FlowOperation.AddFlow:input_type -> sdn.FlowAddRequest
	3, // 6: sdn.ConnectionManager.SendFlowMod:output_type -> sdn.FlowModResponse
	8, // 7: sdn.PacketHandler.HandlePacketIn:output_type -> sdn.PacketInResponse
	1, // 8: sdn.FlowOperation.AddFlow:output_type -> sdn.FlowAddResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_sdn_proto_init() }
func file_proto_sdn_proto_init() {
	if File_proto_sdn_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_sdn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_proto_sdn_proto_goTypes,
		DependencyIndexes: file_proto_sdn_proto_depIdxs,
		MessageInfos:      file_proto_sdn_proto_msgTypes,
	}.Build()
	File_proto_sdn_proto = out.File
	file_proto_sdn_proto_rawDesc = nil
	file_proto_sdn_proto_goTypes = nil
	file_proto_sdn_proto_depIdxs = nil
}
