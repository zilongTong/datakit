// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: span.proto

package ddtrace

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DDSpan struct {
	Service              string             `protobuf:"bytes,1,opt,name=service,proto3" json:"service" msg:"service"`
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name" msg:"name"`
	Resource             string             `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource" msg:"resource"`
	TraceID              uint64             `protobuf:"varint,4,opt,name=traceID,proto3" json:"trace_id" msg:"trace_id"`
	SpanID               uint64             `protobuf:"varint,5,opt,name=spanID,proto3" json:"span_id" msg:"span_id"`
	ParentID             uint64             `protobuf:"varint,6,opt,name=parentID,proto3" json:"parent_id" msg:"parent_id"`
	Start                int64              `protobuf:"varint,7,opt,name=start,proto3" json:"start" msg:"start"`
	Duration             int64              `protobuf:"varint,8,opt,name=duration,proto3" json:"duration" msg:"duration"`
	Error                int32              `protobuf:"varint,9,opt,name=error,proto3" json:"error" msg:"error"`
	Meta                 map[string]string  `protobuf:"bytes,10,rep,name=meta,proto3" json:"meta" msg:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metrics              map[string]float64 `protobuf:"bytes,11,rep,name=metrics,proto3" json:"metrics" msg:"metrics" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Type                 string             `protobuf:"bytes,12,opt,name=type,proto3" json:"type" msg:"type"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DDSpan) Reset()         { *m = DDSpan{} }
func (m *DDSpan) String() string { return proto.CompactTextString(m) }
func (*DDSpan) ProtoMessage()    {}
func (*DDSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc5f2b88b579999f, []int{0}
}
func (m *DDSpan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DDSpan.Unmarshal(m, b)
}
func (m *DDSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DDSpan.Marshal(b, m, deterministic)
}
func (m *DDSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DDSpan.Merge(m, src)
}
func (m *DDSpan) XXX_Size() int {
	return xxx_messageInfo_DDSpan.Size(m)
}
func (m *DDSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_DDSpan.DiscardUnknown(m)
}

var xxx_messageInfo_DDSpan proto.InternalMessageInfo

func (m *DDSpan) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *DDSpan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DDSpan) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *DDSpan) GetTraceID() uint64 {
	if m != nil {
		return m.TraceID
	}
	return 0
}

func (m *DDSpan) GetSpanID() uint64 {
	if m != nil {
		return m.SpanID
	}
	return 0
}

func (m *DDSpan) GetParentID() uint64 {
	if m != nil {
		return m.ParentID
	}
	return 0
}

func (m *DDSpan) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DDSpan) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *DDSpan) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *DDSpan) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DDSpan) GetMetrics() map[string]float64 {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *DDSpan) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*DDSpan)(nil), "ddtrace.DDSpan")
	proto.RegisterMapType((map[string]string)(nil), "ddtrace.DDSpan.MetaEntry")
	proto.RegisterMapType((map[string]float64)(nil), "ddtrace.DDSpan.MetricsEntry")
}

func init() { proto.RegisterFile("span.proto", fileDescriptor_fc5f2b88b579999f) }

var fileDescriptor_fc5f2b88b579999f = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x8e, 0xd3, 0x3e,
	0x10, 0xc7, 0x95, 0xed, 0x7f, 0x77, 0x7f, 0x3f, 0x90, 0x85, 0x90, 0xa9, 0x56, 0x24, 0xf2, 0xa9,
	0x42, 0xda, 0x44, 0x02, 0x04, 0xab, 0x8a, 0x53, 0x69, 0x0f, 0x3d, 0x70, 0x31, 0x9c, 0xb8, 0x20,
	0x37, 0x31, 0x21, 0x6c, 0x6b, 0x47, 0x8e, 0xbd, 0x52, 0xdf, 0x8a, 0x27, 0xca, 0x03, 0xf4, 0xd8,
	0x27, 0x40, 0x1e, 0x27, 0x66, 0x0f, 0x48, 0xdc, 0xf2, 0xfd, 0xce, 0x7c, 0x3c, 0x9e, 0xc9, 0x18,
	0xa1, 0xa6, 0xe6, 0x32, 0xad, 0xb5, 0x32, 0x0a, 0x4f, 0x8a, 0xc2, 0x68, 0x9e, 0x8b, 0xc5, 0x6d,
	0x59, 0x99, 0x1f, 0x76, 0x9f, 0xe6, 0xea, 0x98, 0x95, 0xaa, 0x54, 0x19, 0xc4, 0xf7, 0xf6, 0x3b,
	0x28, 0x10, 0xf0, 0xe5, 0x39, 0xfa, 0x6b, 0x8c, 0xc6, 0x9b, 0xcd, 0xe7, 0x9a, 0x4b, 0xfc, 0x0e,
	0x4d, 0x1a, 0xa1, 0x1f, 0xaa, 0x5c, 0x90, 0x28, 0x89, 0x96, 0xb3, 0xf5, 0xcd, 0xb9, 0x8d, 0x7b,
	0xeb, 0xd2, 0xc6, 0xff, 0x1d, 0x9b, 0x72, 0x45, 0x3b, 0x4d, 0x59, 0x1f, 0xc1, 0xaf, 0xd0, 0x50,
	0xf2, 0xa3, 0x20, 0x57, 0x00, 0x3d, 0x3f, 0xb7, 0x31, 0xe8, 0x4b, 0x1b, 0x23, 0x20, 0x9c, 0xa0,
	0x0c, 0x3c, 0xbc, 0x42, 0x53, 0x2d, 0x1a, 0x65, 0x75, 0x2e, 0xc8, 0x00, 0xf2, 0x5f, 0x9e, 0xdb,
	0x38, 0x78, 0x97, 0x36, 0xfe, 0x1f, 0x98, 0xde, 0xa0, 0x2c, 0xc4, 0xf0, 0x1d, 0x9a, 0x40, 0x8b,
	0xbb, 0x0d, 0x19, 0x26, 0xd1, 0x72, 0xe8, 0x51, 0xb0, 0xbe, 0x55, 0x45, 0x40, 0x7b, 0x83, 0xb2,
	0x3e, 0x1d, 0xbf, 0x45, 0x63, 0x37, 0xaa, 0xdd, 0x86, 0x8c, 0x00, 0xf4, 0x8d, 0xd5, 0x5c, 0x7a,
	0xae, 0x6b, 0xcc, 0x6b, 0xca, 0xba, 0x5c, 0xfc, 0x01, 0x4d, 0x6b, 0xae, 0x85, 0x34, 0xbb, 0x0d,
	0x19, 0x03, 0x97, 0x9c, 0xdb, 0x78, 0xe6, 0x3d, 0x4f, 0x3e, 0x01, 0x32, 0x38, 0x94, 0x05, 0x02,
	0xa7, 0x68, 0xd4, 0x18, 0xae, 0x0d, 0x99, 0x24, 0xd1, 0x72, 0xb0, 0x26, 0xe7, 0x36, 0xf6, 0xc6,
	0xa5, 0x8d, 0xe7, 0xbe, 0xa0, 0x53, 0x94, 0x79, 0xd7, 0x4d, 0xa6, 0xb0, 0x9a, 0x9b, 0x4a, 0x49,
	0x32, 0x05, 0x04, 0xda, 0xeb, 0xbd, 0xd0, 0x5e, 0x6f, 0x50, 0x16, 0x62, 0xae, 0x96, 0xd0, 0x5a,
	0x69, 0x32, 0x4b, 0xa2, 0xe5, 0xc8, 0xd7, 0x02, 0x23, 0xd4, 0x02, 0x45, 0x99, 0x77, 0xf1, 0x16,
	0x0d, 0x8f, 0xc2, 0x70, 0x82, 0x92, 0xc1, 0x72, 0xfe, 0xfa, 0x45, 0xda, 0xed, 0x4e, 0xea, 0x17,
	0x21, 0xfd, 0x24, 0x0c, 0xdf, 0x4a, 0xa3, 0x4f, 0xfe, 0x67, 0xba, 0xd4, 0xf0, 0x33, 0x9d, 0xa0,
	0x0c, 0x3c, 0xfc, 0x05, 0x4d, 0x8e, 0xc2, 0xe8, 0x2a, 0x6f, 0xc8, 0x1c, 0x4e, 0xba, 0xf9, 0xcb,
	0x49, 0x2e, 0xec, 0x0f, 0x83, 0xa9, 0x77, 0x40, 0x98, 0x7a, 0xa7, 0x29, 0xeb, 0x23, 0x6e, 0x9d,
	0xcc, 0xa9, 0x16, 0xe4, 0xfa, 0xcf, 0x3a, 0x39, 0x1d, 0x6e, 0xe0, 0x04, 0x65, 0xe0, 0x2d, 0xde,
	0xa3, 0x59, 0xb8, 0x2c, 0x7e, 0x8a, 0x06, 0xf7, 0xe2, 0xe4, 0x77, 0x97, 0xb9, 0x4f, 0xfc, 0x0c,
	0x8d, 0x1e, 0xf8, 0xc1, 0x76, 0xab, 0xc9, 0xbc, 0x58, 0x5d, 0xdd, 0x45, 0x8b, 0x15, 0xba, 0x7e,
	0x7c, 0xb7, 0x7f, 0xb1, 0xd1, 0x23, 0x76, 0xbd, 0xfd, 0xfa, 0xb1, 0xac, 0xcc, 0x81, 0xef, 0xd3,
	0x9f, 0x15, 0x2f, 0x95, 0x3d, 0x59, 0x09, 0xaf, 0x2d, 0x3f, 0x28, 0x5b, 0xe4, 0x5c, 0x8b, 0x5b,
	0xa3, 0xd4, 0xa1, 0xc9, 0x0a, 0x6e, 0xf8, 0x7d, 0x65, 0xb2, 0xfa, 0x60, 0xcb, 0x4a, 0x36, 0x59,
	0x25, 0x6b, 0x6b, 0x9a, 0xac, 0x1b, 0xd1, 0x7e, 0x0c, 0x0f, 0xf0, 0xcd, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x79, 0x81, 0x21, 0xfb, 0xc6, 0x03, 0x00, 0x00,
}
