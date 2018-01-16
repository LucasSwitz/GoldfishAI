// Code generated by protoc-gen-go.
// source: goldfish_message.proto
// DO NOT EDIT!

/*
Package goldfish_message is a generated protocol buffer package.

It is generated from these files:
	goldfish_message.proto

It has these top-level messages:
	StateMessage
	RegistrationMessage
	ManagementMessage
	GoldFishMessage
*/
package goldfish_message

import "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type GoldFishMessage_Type int32

const (
	GoldFishMessage_STATE_UPDATE GoldFishMessage_Type = 0
	GoldFishMessage_REGISTRATION GoldFishMessage_Type = 1
	GoldFishMessage_MANAGMENT    GoldFishMessage_Type = 2
)

var GoldFishMessage_Type_name = map[int32]string{
	0: "STATE_UPDATE",
	1: "REGISTRATION",
	2: "MANAGMENT",
}
var GoldFishMessage_Type_value = map[string]int32{
	"STATE_UPDATE": 0,
	"REGISTRATION": 1,
	"MANAGMENT":    2,
}

func (x GoldFishMessage_Type) String() string {
	return proto.EnumName(GoldFishMessage_Type_name, int32(x))
}

type StateMessage struct {
}

func (m *StateMessage) Reset()         { *m = StateMessage{} }
func (m *StateMessage) String() string { return proto.CompactTextString(m) }
func (*StateMessage) ProtoMessage()    {}

type RegistrationMessage struct {
}

func (m *RegistrationMessage) Reset()         { *m = RegistrationMessage{} }
func (m *RegistrationMessage) String() string { return proto.CompactTextString(m) }
func (*RegistrationMessage) ProtoMessage()    {}

type ManagementMessage struct {
	Extra int32 `protobuf:"fixed32,1,opt,name=extra" json:"extra,omitempty"`
}

func (m *ManagementMessage) Reset()         { *m = ManagementMessage{} }
func (m *ManagementMessage) String() string { return proto.CompactTextString(m) }
func (*ManagementMessage) ProtoMessage()    {}

type GoldFishMessage struct {
	Type                GoldFishMessage_Type `protobuf:"varint,1,opt,name=type,enum=GoldFishMessage_Type" json:"type,omitempty"`
	StateMessage        *StateMessage        `protobuf:"bytes,2,opt,name=state_message" json:"state_message,omitempty"`
	RegistrationMessage *RegistrationMessage `protobuf:"bytes,3,opt,name=registration_message" json:"registration_message,omitempty"`
	ManagementMessage   *ManagementMessage   `protobuf:"bytes,4,opt,name=management_message" json:"management_message,omitempty"`
}

func (m *GoldFishMessage) Reset()         { *m = GoldFishMessage{} }
func (m *GoldFishMessage) String() string { return proto.CompactTextString(m) }
func (*GoldFishMessage) ProtoMessage()    {}

func (m *GoldFishMessage) GetStateMessage() *StateMessage {
	if m != nil {
		return m.StateMessage
	}
	return nil
}

func (m *GoldFishMessage) GetRegistrationMessage() *RegistrationMessage {
	if m != nil {
		return m.RegistrationMessage
	}
	return nil
}

func (m *GoldFishMessage) GetManagementMessage() *ManagementMessage {
	if m != nil {
		return m.ManagementMessage
	}
	return nil
}

func init() {
	proto.RegisterEnum("GoldFishMessage_Type", GoldFishMessage_Type_name, GoldFishMessage_Type_value)
}
