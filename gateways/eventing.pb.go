// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateways/eventing.proto

package gateways

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//*
// Represents an event source
type EventSource struct {
	// ID of the event source. internally generated.
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The event source name.
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// The event source configuration value.
	Data                 *string  `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSource) Reset()         { *m = EventSource{} }
func (m *EventSource) String() string { return proto.CompactTextString(m) }
func (*EventSource) ProtoMessage()    {}
func (*EventSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25325013aefc28a, []int{0}
}

func (m *EventSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSource.Unmarshal(m, b)
}
func (m *EventSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSource.Marshal(b, m, deterministic)
}
func (m *EventSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSource.Merge(m, src)
}
func (m *EventSource) XXX_Size() int {
	return xxx_messageInfo_EventSource.Size(m)
}
func (m *EventSource) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSource.DiscardUnknown(m)
}

var xxx_messageInfo_EventSource proto.InternalMessageInfo

func (m *EventSource) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *EventSource) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *EventSource) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

//*
// Represents an event
type Event struct {
	// The event source name.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The event payload.
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25325013aefc28a, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

//*
// Represents if an event source is valid or not
type ValidEventSource struct {
	// whether event source is valid
	IsValid *bool `protobuf:"varint,1,opt,name=isValid" json:"isValid,omitempty"`
	// reason if an event source is invalid
	Reason               *string  `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidEventSource) Reset()         { *m = ValidEventSource{} }
func (m *ValidEventSource) String() string { return proto.CompactTextString(m) }
func (*ValidEventSource) ProtoMessage()    {}
func (*ValidEventSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25325013aefc28a, []int{2}
}

func (m *ValidEventSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidEventSource.Unmarshal(m, b)
}
func (m *ValidEventSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidEventSource.Marshal(b, m, deterministic)
}
func (m *ValidEventSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidEventSource.Merge(m, src)
}
func (m *ValidEventSource) XXX_Size() int {
	return xxx_messageInfo_ValidEventSource.Size(m)
}
func (m *ValidEventSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidEventSource.DiscardUnknown(m)
}

var xxx_messageInfo_ValidEventSource proto.InternalMessageInfo

func (m *ValidEventSource) GetIsValid() bool {
	if m != nil && m.IsValid != nil {
		return *m.IsValid
	}
	return false
}

func (m *ValidEventSource) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*EventSource)(nil), "gateways.EventSource")
	proto.RegisterType((*Event)(nil), "gateways.Event")
	proto.RegisterType((*ValidEventSource)(nil), "gateways.ValidEventSource")
}

func init() { proto.RegisterFile("gateways/eventing.proto", fileDescriptor_c25325013aefc28a) }

var fileDescriptor_c25325013aefc28a = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0x2c, 0xd6, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x49, 0x28, 0xb9, 0x72, 0x71, 0xbb, 0x82, 0xe4, 0x82, 0xf3,
	0x4b, 0x8b, 0x92, 0x53, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0xc0, 0x22,
	0x60, 0x36, 0x48, 0x2c, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x19, 0x22, 0x06, 0x62, 0x2b, 0x99, 0x72,
	0xb1, 0x82, 0x8d, 0x81, 0x6b, 0x60, 0x44, 0xd2, 0x20, 0xc1, 0xc5, 0x5e, 0x90, 0x58, 0x99, 0x93,
	0x9f, 0x98, 0x02, 0x36, 0x87, 0x27, 0x08, 0xc6, 0x55, 0x72, 0xe1, 0x12, 0x08, 0x4b, 0xcc, 0xc9,
	0x4c, 0x41, 0x76, 0x82, 0x04, 0x17, 0x7b, 0x66, 0x31, 0x58, 0x14, 0x6c, 0x08, 0x47, 0x10, 0x8c,
	0x2b, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x9a, 0x58, 0x9c, 0x9f, 0x07, 0x75, 0x0e, 0x94, 0x67, 0x34,
	0x89, 0x91, 0x8b, 0xc3, 0x15, 0xea, 0x41, 0x21, 0x1b, 0x2e, 0x81, 0xe0, 0x92, 0xc4, 0xa2, 0x12,
	0x64, 0x23, 0x45, 0xf5, 0x60, 0xfe, 0xd5, 0x43, 0x12, 0x96, 0xe2, 0x47, 0x13, 0x36, 0x60, 0x14,
	0xf2, 0xe0, 0x12, 0x06, 0xdb, 0x95, 0x58, 0x92, 0x4a, 0x84, 0x01, 0x52, 0x08, 0x61, 0x74, 0x6f,
	0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0xac, 0x1f, 0xca, 0x3d, 0x7c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventingClient is the client API for Eventing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventingClient interface {
	// StartEventSource starts an event source and returns stream of events.
	StartEventSource(ctx context.Context, in *EventSource, opts ...grpc.CallOption) (Eventing_StartEventSourceClient, error)
	// ValidateEventSource validates an event source.
	ValidateEventSource(ctx context.Context, in *EventSource, opts ...grpc.CallOption) (*ValidEventSource, error)
}

type eventingClient struct {
	cc *grpc.ClientConn
}

func NewEventingClient(cc *grpc.ClientConn) EventingClient {
	return &eventingClient{cc}
}

func (c *eventingClient) StartEventSource(ctx context.Context, in *EventSource, opts ...grpc.CallOption) (Eventing_StartEventSourceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Eventing_serviceDesc.Streams[0], "/gateways.Eventing/StartEventSource", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventingStartEventSourceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Eventing_StartEventSourceClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventingStartEventSourceClient struct {
	grpc.ClientStream
}

func (x *eventingStartEventSourceClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventingClient) ValidateEventSource(ctx context.Context, in *EventSource, opts ...grpc.CallOption) (*ValidEventSource, error) {
	out := new(ValidEventSource)
	err := c.cc.Invoke(ctx, "/gateways.Eventing/ValidateEventSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventingServer is the server API for Eventing service.
type EventingServer interface {
	// StartEventSource starts an event source and returns stream of events.
	StartEventSource(*EventSource, Eventing_StartEventSourceServer) error
	// ValidateEventSource validates an event source.
	ValidateEventSource(context.Context, *EventSource) (*ValidEventSource, error)
}

func RegisterEventingServer(s *grpc.Server, srv EventingServer) {
	s.RegisterService(&_Eventing_serviceDesc, srv)
}

func _Eventing_StartEventSource_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventSource)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventingServer).StartEventSource(m, &eventingStartEventSourceServer{stream})
}

type Eventing_StartEventSourceServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventingStartEventSourceServer struct {
	grpc.ServerStream
}

func (x *eventingStartEventSourceServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Eventing_ValidateEventSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventSource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventingServer).ValidateEventSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateways.Eventing/ValidateEventSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventingServer).ValidateEventSource(ctx, req.(*EventSource))
	}
	return interceptor(ctx, in, info, handler)
}

var _Eventing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateways.Eventing",
	HandlerType: (*EventingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateEventSource",
			Handler:    _Eventing_ValidateEventSource_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartEventSource",
			Handler:       _Eventing_StartEventSource_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gateways/eventing.proto",
}
