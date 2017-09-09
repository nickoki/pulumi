// Code generated by protoc-gen-go.
// source: analyzer.proto
// DO NOT EDIT!

/*
Package lumirpc is a generated protocol buffer package.

It is generated from these files:
	analyzer.proto
	engine.proto
	languages.proto
	provider.proto

It has these top-level messages:
	AnalyzeRequest
	AnalyzeResponse
	AnalyzeFailure
	LogRequest
	RunRequest
	RunResponse
	NewResourceRequest
	NewResourceResponse
	ConfigureRequest
	CheckRequest
	CheckResponse
	CheckFailure
	DiffRequest
	DiffResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
*/
package lumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

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

type AnalyzeRequest struct {
	Type       string                  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *AnalyzeRequest) Reset()                    { *m = AnalyzeRequest{} }
func (m *AnalyzeRequest) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeRequest) ProtoMessage()               {}
func (*AnalyzeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AnalyzeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AnalyzeRequest) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type AnalyzeResponse struct {
	Failures []*AnalyzeFailure `protobuf:"bytes,1,rep,name=failures" json:"failures,omitempty"`
}

func (m *AnalyzeResponse) Reset()                    { *m = AnalyzeResponse{} }
func (m *AnalyzeResponse) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeResponse) ProtoMessage()               {}
func (*AnalyzeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AnalyzeResponse) GetFailures() []*AnalyzeFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

type AnalyzeFailure struct {
	Property string `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Reason   string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *AnalyzeFailure) Reset()                    { *m = AnalyzeFailure{} }
func (m *AnalyzeFailure) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeFailure) ProtoMessage()               {}
func (*AnalyzeFailure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AnalyzeFailure) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *AnalyzeFailure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*AnalyzeRequest)(nil), "lumirpc.AnalyzeRequest")
	proto.RegisterType((*AnalyzeResponse)(nil), "lumirpc.AnalyzeResponse")
	proto.RegisterType((*AnalyzeFailure)(nil), "lumirpc.AnalyzeFailure")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Analyzer service

type AnalyzerClient interface {
	// Analyze analyzes a single resource object, and returns any errors that it finds.
	Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
}

type analyzerClient struct {
	cc *grpc.ClientConn
}

func NewAnalyzerClient(cc *grpc.ClientConn) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := grpc.Invoke(ctx, "/lumirpc.Analyzer/Analyze", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Analyzer service

type AnalyzerServer interface {
	// Analyze analyzes a single resource object, and returns any errors that it finds.
	Analyze(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error)
}

func RegisterAnalyzerServer(s *grpc.Server, srv AnalyzerServer) {
	s.RegisterService(&_Analyzer_serviceDesc, srv)
}

func _Analyzer_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.Analyzer/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Analyze(ctx, req.(*AnalyzeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analyzer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lumirpc.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _Analyzer_Analyze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyzer.proto",
}

func init() { proto.RegisterFile("analyzer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0xb6, 0x24, 0xce, 0x15, 0x52, 0xb8, 0xa1, 0x35, 0xa6, 0x83, 0xf1, 0x94, 0x49,
	0x81, 0x64, 0xe8, 0x56, 0x28, 0x94, 0x0c, 0x1d, 0xd5, 0xb9, 0x83, 0x13, 0x2e, 0xc1, 0xe0, 0x5a,
	0xea, 0x49, 0x1a, 0xdc, 0x5f, 0x5f, 0x90, 0xae, 0xa2, 0xe0, 0xed, 0x4e, 0xef, 0xe9, 0xbb, 0x77,
	0x07, 0xeb, 0x6e, 0xec, 0x86, 0xe9, 0x87, 0x58, 0x59, 0x36, 0xde, 0xe0, 0x72, 0x08, 0x5f, 0x3d,
	0xdb, 0x53, 0xfd, 0x74, 0x31, 0xe6, 0x32, 0xd0, 0x36, 0x3e, 0x1f, 0xc3, 0x79, 0xeb, 0x3c, 0x87,
	0x93, 0x4f, 0xb6, 0xf6, 0x13, 0xd6, 0xaf, 0xe9, 0xa3, 0xa6, 0xef, 0x40, 0xce, 0x23, 0xc2, 0xad,
	0x9f, 0x2c, 0x55, 0x45, 0x53, 0x6c, 0x56, 0x3a, 0xd6, 0xf8, 0x0c, 0x60, 0xd9, 0x58, 0x62, 0xdf,
	0x93, 0xab, 0xae, 0x9b, 0x62, 0x73, 0xb7, 0x7b, 0x54, 0x09, 0xac, 0xfe, 0xc0, 0xea, 0x23, 0x82,
	0xf5, 0x3f, 0x6b, 0x7b, 0x80, 0xfb, 0x8c, 0x77, 0xd6, 0x8c, 0x8e, 0x70, 0x0f, 0xe5, 0xb9, 0xeb,
	0x87, 0xc0, 0xe4, 0xaa, 0xa2, 0xb9, 0x89, 0x24, 0xc9, 0xaa, 0xc4, 0x7b, 0x48, 0xba, 0xce, 0xc6,
	0xf6, 0x2d, 0xc7, 0x14, 0x0d, 0x6b, 0x28, 0x65, 0xce, 0x24, 0x51, 0x73, 0x8f, 0x0f, 0xb0, 0x60,
	0xea, 0x9c, 0x19, 0x63, 0xd4, 0x95, 0x96, 0x6e, 0xf7, 0x0e, 0xa5, 0x50, 0x18, 0x5f, 0x60, 0x29,
	0x35, 0xce, 0xe6, 0xcb, 0x29, 0xea, 0x6a, 0x2e, 0xa4, 0x25, 0xda, 0xab, 0xe3, 0x22, 0xae, 0xbd,
	0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xab, 0xff, 0x20, 0x89, 0x78, 0x01, 0x00, 0x00,
}