// Code generated by protoc-gen-go.
// source: eventium.proto
// DO NOT EDIT!

/*
Package eventiumpb is a generated protocol buffer package.

It is generated from these files:
	eventium.proto

It has these top-level messages:
	GetEventByIDRequest
	PostEventRequest
	PostEventResponse
	Location
	Eventum
*/
package eventiumpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// =============================== MESSAGES
type GetEventByIDRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetEventByIDRequest) Reset()                    { *m = GetEventByIDRequest{} }
func (m *GetEventByIDRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEventByIDRequest) ProtoMessage()               {}
func (*GetEventByIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetEventByIDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PostEventRequest struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Event     *Eventum                   `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
}

func (m *PostEventRequest) Reset()                    { *m = PostEventRequest{} }
func (m *PostEventRequest) String() string            { return proto.CompactTextString(m) }
func (*PostEventRequest) ProtoMessage()               {}
func (*PostEventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PostEventRequest) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PostEventRequest) GetEvent() *Eventum {
	if m != nil {
		return m.Event
	}
	return nil
}

type PostEventResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	EventID bool `protobuf:"varint,2,opt,name=eventID" json:"eventID,omitempty"`
}

func (m *PostEventResponse) Reset()                    { *m = PostEventResponse{} }
func (m *PostEventResponse) String() string            { return proto.CompactTextString(m) }
func (*PostEventResponse) ProtoMessage()               {}
func (*PostEventResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PostEventResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PostEventResponse) GetEventID() bool {
	if m != nil {
		return m.EventID
	}
	return false
}

// =============================== OBJECTS
type Location struct {
	X float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Location) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Location) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Eventum struct {
	Location  *Location                  `protobuf:"bytes,16,opt,name=location" json:"location,omitempty"`
	Begin     *google_protobuf.Timestamp `protobuf:"bytes,17,opt,name=begin" json:"begin,omitempty"`
	End       *google_protobuf.Timestamp `protobuf:"bytes,18,opt,name=end" json:"end,omitempty"`
	Headline  string                     `protobuf:"bytes,19,opt,name=headline" json:"headline,omitempty"`
	MediaData []byte                     `protobuf:"bytes,20,opt,name=mediaData,proto3" json:"mediaData,omitempty"`
	Tags      []string                   `protobuf:"bytes,21,rep,name=tags" json:"tags,omitempty"`
}

func (m *Eventum) Reset()                    { *m = Eventum{} }
func (m *Eventum) String() string            { return proto.CompactTextString(m) }
func (*Eventum) ProtoMessage()               {}
func (*Eventum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Eventum) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Eventum) GetBegin() *google_protobuf.Timestamp {
	if m != nil {
		return m.Begin
	}
	return nil
}

func (m *Eventum) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Eventum) GetHeadline() string {
	if m != nil {
		return m.Headline
	}
	return ""
}

func (m *Eventum) GetMediaData() []byte {
	if m != nil {
		return m.MediaData
	}
	return nil
}

func (m *Eventum) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*GetEventByIDRequest)(nil), "eventiumpb.GetEventByIDRequest")
	proto.RegisterType((*PostEventRequest)(nil), "eventiumpb.PostEventRequest")
	proto.RegisterType((*PostEventResponse)(nil), "eventiumpb.PostEventResponse")
	proto.RegisterType((*Location)(nil), "eventiumpb.Location")
	proto.RegisterType((*Eventum)(nil), "eventiumpb.Eventum")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Eventium service

type EventiumClient interface {
	PostEvent(ctx context.Context, in *PostEventRequest, opts ...grpc.CallOption) (*PostEventResponse, error)
	GetEventByID(ctx context.Context, in *GetEventByIDRequest, opts ...grpc.CallOption) (*Eventum, error)
}

type eventiumClient struct {
	cc *grpc.ClientConn
}

func NewEventiumClient(cc *grpc.ClientConn) EventiumClient {
	return &eventiumClient{cc}
}

func (c *eventiumClient) PostEvent(ctx context.Context, in *PostEventRequest, opts ...grpc.CallOption) (*PostEventResponse, error) {
	out := new(PostEventResponse)
	err := grpc.Invoke(ctx, "/eventiumpb.Eventium/PostEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventiumClient) GetEventByID(ctx context.Context, in *GetEventByIDRequest, opts ...grpc.CallOption) (*Eventum, error) {
	out := new(Eventum)
	err := grpc.Invoke(ctx, "/eventiumpb.Eventium/GetEventByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Eventium service

type EventiumServer interface {
	PostEvent(context.Context, *PostEventRequest) (*PostEventResponse, error)
	GetEventByID(context.Context, *GetEventByIDRequest) (*Eventum, error)
}

func RegisterEventiumServer(s *grpc.Server, srv EventiumServer) {
	s.RegisterService(&_Eventium_serviceDesc, srv)
}

func _Eventium_PostEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventiumServer).PostEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventiumpb.Eventium/PostEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventiumServer).PostEvent(ctx, req.(*PostEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Eventium_GetEventByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventiumServer).GetEventByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventiumpb.Eventium/GetEventByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventiumServer).GetEventByID(ctx, req.(*GetEventByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Eventium_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventiumpb.Eventium",
	HandlerType: (*EventiumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostEvent",
			Handler:    _Eventium_PostEvent_Handler,
		},
		{
			MethodName: "GetEventByID",
			Handler:    _Eventium_GetEventByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventium.proto",
}

func init() { proto.RegisterFile("eventium.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xe1, 0x8a, 0xda, 0x40,
	0x10, 0xc7, 0xdd, 0x58, 0x6b, 0x32, 0x15, 0xd1, 0xd1, 0xc2, 0x12, 0x2c, 0x4a, 0xa0, 0xc5, 0x42,
	0x89, 0x62, 0xbf, 0xf4, 0x73, 0xd1, 0x8a, 0xa5, 0x1f, 0xca, 0xd2, 0x17, 0x48, 0xcc, 0x36, 0x5d,
	0x30, 0xd9, 0xf4, 0x76, 0x73, 0xa7, 0x2f, 0x73, 0xef, 0x77, 0x6f, 0x71, 0x64, 0x63, 0x62, 0x38,
	0x3c, 0xfc, 0x96, 0xc9, 0xfc, 0xfe, 0x33, 0xff, 0x9d, 0x19, 0xe8, 0xf3, 0x7b, 0x9e, 0x6a, 0x91,
	0x27, 0x7e, 0x76, 0x27, 0xb5, 0x44, 0xa8, 0xe2, 0x2c, 0x74, 0xa7, 0xb1, 0x94, 0xf1, 0x81, 0x2f,
	0x4c, 0x26, 0xcc, 0xff, 0x2e, 0xb4, 0x48, 0xb8, 0xd2, 0x41, 0x92, 0x95, 0xb0, 0xf7, 0x11, 0x46,
	0x5b, 0xae, 0x37, 0x85, 0xe2, 0xfb, 0x69, 0xb7, 0x66, 0xfc, 0x7f, 0xce, 0x95, 0xc6, 0x3e, 0x58,
	0x22, 0xa2, 0x64, 0x46, 0xe6, 0x0e, 0xb3, 0x44, 0xe4, 0x3d, 0xc0, 0xe0, 0xb7, 0x54, 0x25, 0x57,
	0x31, 0xdf, 0xc0, 0xa9, 0xab, 0x19, 0xf4, 0xdd, 0xca, 0xf5, 0xcb, 0x7e, 0x7e, 0xd5, 0xcf, 0xff,
	0x53, 0x11, 0xec, 0x02, 0xe3, 0x67, 0xe8, 0x18, 0x8f, 0xd4, 0x32, 0xaa, 0x91, 0x7f, 0x71, 0xec,
	0x9b, 0x16, 0x79, 0xc2, 0x4a, 0xc2, 0xdb, 0xc2, 0xb0, 0xd1, 0x58, 0x65, 0x32, 0x55, 0x1c, 0x29,
	0x74, 0x55, 0xbe, 0xdf, 0x73, 0xa5, 0x4c, 0x5f, 0x9b, 0x55, 0x61, 0x91, 0x31, 0xba, 0xdd, 0xda,
	0xd4, 0xb6, 0x59, 0x15, 0x7a, 0x9f, 0xc0, 0xfe, 0x25, 0xf7, 0x81, 0x16, 0x32, 0xc5, 0x1e, 0x90,
	0xa3, 0x51, 0x5a, 0x8c, 0x1c, 0x8b, 0xe8, 0x64, 0x68, 0x8b, 0x91, 0x93, 0xf7, 0x44, 0xa0, 0x7b,
	0xf6, 0x80, 0x4b, 0xb0, 0x0f, 0x67, 0x0d, 0x1d, 0x18, 0xab, 0xe3, 0xa6, 0xd5, 0xaa, 0x1e, 0xab,
	0x29, 0x5c, 0x42, 0x27, 0xe4, 0xb1, 0x48, 0xe9, 0xf0, 0xe6, 0x3c, 0x4a, 0x10, 0xbf, 0x40, 0x9b,
	0xa7, 0x11, 0xc5, 0x9b, 0x7c, 0x81, 0xa1, 0x0b, 0xf6, 0x3f, 0x1e, 0x44, 0x07, 0x91, 0x72, 0x3a,
	0x32, 0xdb, 0xa9, 0x63, 0x9c, 0x80, 0x93, 0xf0, 0x48, 0x04, 0xeb, 0x40, 0x07, 0x74, 0x3c, 0x23,
	0xf3, 0x1e, 0xbb, 0xfc, 0x40, 0x84, 0x37, 0x3a, 0x88, 0x15, 0x7d, 0x3f, 0x6b, 0xcf, 0x1d, 0x66,
	0xbe, 0x57, 0x8f, 0x04, 0xec, 0xcd, 0xf9, 0x3d, 0xf8, 0x13, 0x9c, 0x7a, 0xd2, 0x38, 0x69, 0xbe,
	0xf3, 0xe5, 0xe6, 0xdd, 0x0f, 0xaf, 0x64, 0xcb, 0xf5, 0x78, 0x2d, 0xfc, 0x01, 0xbd, 0xe6, 0x55,
	0xe1, 0xb4, 0x29, 0xb8, 0x72, 0x6f, 0xee, 0xb5, 0x13, 0xf0, 0x5a, 0xe1, 0x5b, 0x33, 0x87, 0xaf,
	0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x8c, 0x1f, 0x35, 0xe3, 0x02, 0x00, 0x00,
}