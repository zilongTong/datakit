//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: plugins/inputs/skywalking/v2/language-agent-v2/trace.proto

package language_agent_v2

import (
	common "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking/v2/common"
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

type UpstreamSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalTraceIds []*common.UniqueId `protobuf:"bytes,1,rep,name=globalTraceIds,proto3" json:"globalTraceIds,omitempty"`
	Segment        *SegmentObject     `protobuf:"bytes,2,opt,name=segment,proto3" json:"segment,omitempty"` // the byte array of TraceSegmentObject
}

func (x *UpstreamSegment) Reset() {
	*x = UpstreamSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSegment) ProtoMessage() {}

func (x *UpstreamSegment) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamSegment.ProtoReflect.Descriptor instead.
func (*UpstreamSegment) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamSegment) GetGlobalTraceIds() []*common.UniqueId {
	if x != nil {
		return x.GlobalTraceIds
	}
	return nil
}

func (x *UpstreamSegment) GetSegment() *SegmentObject {
	if x != nil {
		return x.Segment
	}
	return nil
}

type SegmentObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceSegmentId    *common.UniqueId `protobuf:"bytes,1,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	Spans             []*SpanObjectV2  `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	ServiceId         int32            `protobuf:"varint,3,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	ServiceInstanceId int32            `protobuf:"varint,4,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
	IsSizeLimited     bool             `protobuf:"varint,5,opt,name=isSizeLimited,proto3" json:"isSizeLimited,omitempty"`
}

func (x *SegmentObject) Reset() {
	*x = SegmentObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentObject) ProtoMessage() {}

func (x *SegmentObject) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentObject.ProtoReflect.Descriptor instead.
func (*SegmentObject) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescGZIP(), []int{1}
}

func (x *SegmentObject) GetTraceSegmentId() *common.UniqueId {
	if x != nil {
		return x.TraceSegmentId
	}
	return nil
}

func (x *SegmentObject) GetSpans() []*SpanObjectV2 {
	if x != nil {
		return x.Spans
	}
	return nil
}

func (x *SegmentObject) GetServiceId() int32 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *SegmentObject) GetServiceInstanceId() int32 {
	if x != nil {
		return x.ServiceInstanceId
	}
	return 0
}

func (x *SegmentObject) GetIsSizeLimited() bool {
	if x != nil {
		return x.IsSizeLimited
	}
	return false
}

type SegmentReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefType                 common.RefType   `protobuf:"varint,1,opt,name=refType,proto3,enum=skywalking.v2.RefType" json:"refType,omitempty"`
	ParentTraceSegmentId    *common.UniqueId `protobuf:"bytes,2,opt,name=parentTraceSegmentId,proto3" json:"parentTraceSegmentId,omitempty"`
	ParentSpanId            int32            `protobuf:"varint,3,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	ParentServiceInstanceId int32            `protobuf:"varint,4,opt,name=parentServiceInstanceId,proto3" json:"parentServiceInstanceId,omitempty"`
	NetworkAddress          string           `protobuf:"bytes,5,opt,name=networkAddress,proto3" json:"networkAddress,omitempty"`
	NetworkAddressId        int32            `protobuf:"varint,6,opt,name=networkAddressId,proto3" json:"networkAddressId,omitempty"`
	EntryServiceInstanceId  int32            `protobuf:"varint,7,opt,name=entryServiceInstanceId,proto3" json:"entryServiceInstanceId,omitempty"`
	EntryEndpoint           string           `protobuf:"bytes,8,opt,name=entryEndpoint,proto3" json:"entryEndpoint,omitempty"`
	EntryEndpointId         int32            `protobuf:"varint,9,opt,name=entryEndpointId,proto3" json:"entryEndpointId,omitempty"`
	ParentEndpoint          string           `protobuf:"bytes,10,opt,name=parentEndpoint,proto3" json:"parentEndpoint,omitempty"`
	ParentEndpointId        int32            `protobuf:"varint,11,opt,name=parentEndpointId,proto3" json:"parentEndpointId,omitempty"`
}

func (x *SegmentReference) Reset() {
	*x = SegmentReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentReference) ProtoMessage() {}

func (x *SegmentReference) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentReference.ProtoReflect.Descriptor instead.
func (*SegmentReference) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescGZIP(), []int{2}
}

func (x *SegmentReference) GetRefType() common.RefType {
	if x != nil {
		return x.RefType
	}
	return common.RefType(0)
}

func (x *SegmentReference) GetParentTraceSegmentId() *common.UniqueId {
	if x != nil {
		return x.ParentTraceSegmentId
	}
	return nil
}

func (x *SegmentReference) GetParentSpanId() int32 {
	if x != nil {
		return x.ParentSpanId
	}
	return 0
}

func (x *SegmentReference) GetParentServiceInstanceId() int32 {
	if x != nil {
		return x.ParentServiceInstanceId
	}
	return 0
}

func (x *SegmentReference) GetNetworkAddress() string {
	if x != nil {
		return x.NetworkAddress
	}
	return ""
}

func (x *SegmentReference) GetNetworkAddressId() int32 {
	if x != nil {
		return x.NetworkAddressId
	}
	return 0
}

func (x *SegmentReference) GetEntryServiceInstanceId() int32 {
	if x != nil {
		return x.EntryServiceInstanceId
	}
	return 0
}

func (x *SegmentReference) GetEntryEndpoint() string {
	if x != nil {
		return x.EntryEndpoint
	}
	return ""
}

func (x *SegmentReference) GetEntryEndpointId() int32 {
	if x != nil {
		return x.EntryEndpointId
	}
	return 0
}

func (x *SegmentReference) GetParentEndpoint() string {
	if x != nil {
		return x.ParentEndpoint
	}
	return ""
}

func (x *SegmentReference) GetParentEndpointId() int32 {
	if x != nil {
		return x.ParentEndpointId
	}
	return 0
}

type SpanObjectV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpanId          int32                        `protobuf:"varint,1,opt,name=spanId,proto3" json:"spanId,omitempty"`
	ParentSpanId    int32                        `protobuf:"varint,2,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	StartTime       int64                        `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime         int64                        `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Refs            []*SegmentReference          `protobuf:"bytes,5,rep,name=refs,proto3" json:"refs,omitempty"`
	OperationNameId int32                        `protobuf:"varint,6,opt,name=operationNameId,proto3" json:"operationNameId,omitempty"`
	OperationName   string                       `protobuf:"bytes,7,opt,name=operationName,proto3" json:"operationName,omitempty"`
	PeerId          int32                        `protobuf:"varint,8,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Peer            string                       `protobuf:"bytes,9,opt,name=peer,proto3" json:"peer,omitempty"`
	SpanType        common.SpanType              `protobuf:"varint,10,opt,name=spanType,proto3,enum=skywalking.v2.SpanType" json:"spanType,omitempty"`
	SpanLayer       common.SpanLayer             `protobuf:"varint,11,opt,name=spanLayer,proto3,enum=skywalking.v2.SpanLayer" json:"spanLayer,omitempty"`
	ComponentId     int32                        `protobuf:"varint,12,opt,name=componentId,proto3" json:"componentId,omitempty"`
	Component       string                       `protobuf:"bytes,13,opt,name=component,proto3" json:"component,omitempty"`
	IsError         bool                         `protobuf:"varint,14,opt,name=isError,proto3" json:"isError,omitempty"`
	Tags            []*common.KeyStringValuePair `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Logs            []*Log                       `protobuf:"bytes,16,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *SpanObjectV2) Reset() {
	*x = SpanObjectV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpanObjectV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpanObjectV2) ProtoMessage() {}

func (x *SpanObjectV2) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpanObjectV2.ProtoReflect.Descriptor instead.
func (*SpanObjectV2) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescGZIP(), []int{3}
}

func (x *SpanObjectV2) GetSpanId() int32 {
	if x != nil {
		return x.SpanId
	}
	return 0
}

func (x *SpanObjectV2) GetParentSpanId() int32 {
	if x != nil {
		return x.ParentSpanId
	}
	return 0
}

func (x *SpanObjectV2) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *SpanObjectV2) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *SpanObjectV2) GetRefs() []*SegmentReference {
	if x != nil {
		return x.Refs
	}
	return nil
}

func (x *SpanObjectV2) GetOperationNameId() int32 {
	if x != nil {
		return x.OperationNameId
	}
	return 0
}

func (x *SpanObjectV2) GetOperationName() string {
	if x != nil {
		return x.OperationName
	}
	return ""
}

func (x *SpanObjectV2) GetPeerId() int32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *SpanObjectV2) GetPeer() string {
	if x != nil {
		return x.Peer
	}
	return ""
}

func (x *SpanObjectV2) GetSpanType() common.SpanType {
	if x != nil {
		return x.SpanType
	}
	return common.SpanType(0)
}

func (x *SpanObjectV2) GetSpanLayer() common.SpanLayer {
	if x != nil {
		return x.SpanLayer
	}
	return common.SpanLayer(0)
}

func (x *SpanObjectV2) GetComponentId() int32 {
	if x != nil {
		return x.ComponentId
	}
	return 0
}

func (x *SpanObjectV2) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *SpanObjectV2) GetIsError() bool {
	if x != nil {
		return x.IsError
	}
	return false
}

func (x *SpanObjectV2) GetTags() []*common.KeyStringValuePair {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *SpanObjectV2) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64                        `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Data []*common.KeyStringValuePair `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescGZIP(), []int{4}
}

func (x *Log) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Log) GetData() []*common.KeyStringValuePair {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto protoreflect.FileDescriptor

var file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x76, 0x32,
	0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x30, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x77,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x32, 0x2e, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0xf5, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x49, 0x64, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56,
	0x32, 0x52, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x53,
	0x69, 0x7a, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x22, 0x9f, 0x04, 0x0a, 0x10, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32,
	0x2e, 0x52, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x72, 0x65, 0x66, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x4b, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e,
	0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x52, 0x14, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e,
	0x49, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x17, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x36, 0x0a, 0x16, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x16, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x28,
	0x0a, 0x0f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xd9, 0x04, 0x0a,
	0x0c, 0x53, 0x70, 0x61, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x32, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53,
	0x70, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x65, 0x66, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x04, 0x72, 0x65, 0x66, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x65,
	0x65, 0x72, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73,
	0x70, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x70, 0x61, 0x6e, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x6b, 0x79,
	0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x09, 0x73, 0x70, 0x61, 0x6e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x26, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x50, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x32, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x63, 0x0a, 0x19, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x12, 0x1e, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x28, 0x01, 0x42,
	0xb1, 0x01, 0x0a, 0x33, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73,
	0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x6a, 0x69, 0x61, 0x67, 0x6f, 0x75, 0x79, 0x75, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x61, 0x72, 0x65, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2d, 0x76, 0x32, 0xaa, 0x02, 0x1a, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescOnce sync.Once
	file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescData = file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDesc
)

func file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescGZIP() []byte {
	file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescOnce.Do(func() {
		file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescData)
	})
	return file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDescData
}

var file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_goTypes = []interface{}{
	(*UpstreamSegment)(nil),           // 0: skywalking.v2.UpstreamSegment
	(*SegmentObject)(nil),             // 1: skywalking.v2.SegmentObject
	(*SegmentReference)(nil),          // 2: skywalking.v2.SegmentReference
	(*SpanObjectV2)(nil),              // 3: skywalking.v2.SpanObjectV2
	(*Log)(nil),                       // 4: skywalking.v2.Log
	(*common.UniqueId)(nil),           // 5: skywalking.v2.UniqueId
	(common.RefType)(0),               // 6: skywalking.v2.RefType
	(common.SpanType)(0),              // 7: skywalking.v2.SpanType
	(common.SpanLayer)(0),             // 8: skywalking.v2.SpanLayer
	(*common.KeyStringValuePair)(nil), // 9: skywalking.v2.KeyStringValuePair
	(*common.Commands)(nil),           // 10: skywalking.v2.Commands
}
var file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_depIdxs = []int32{
	5,  // 0: skywalking.v2.UpstreamSegment.globalTraceIds:type_name -> skywalking.v2.UniqueId
	1,  // 1: skywalking.v2.UpstreamSegment.segment:type_name -> skywalking.v2.SegmentObject
	5,  // 2: skywalking.v2.SegmentObject.traceSegmentId:type_name -> skywalking.v2.UniqueId
	3,  // 3: skywalking.v2.SegmentObject.spans:type_name -> skywalking.v2.SpanObjectV2
	6,  // 4: skywalking.v2.SegmentReference.refType:type_name -> skywalking.v2.RefType
	5,  // 5: skywalking.v2.SegmentReference.parentTraceSegmentId:type_name -> skywalking.v2.UniqueId
	2,  // 6: skywalking.v2.SpanObjectV2.refs:type_name -> skywalking.v2.SegmentReference
	7,  // 7: skywalking.v2.SpanObjectV2.spanType:type_name -> skywalking.v2.SpanType
	8,  // 8: skywalking.v2.SpanObjectV2.spanLayer:type_name -> skywalking.v2.SpanLayer
	9,  // 9: skywalking.v2.SpanObjectV2.tags:type_name -> skywalking.v2.KeyStringValuePair
	4,  // 10: skywalking.v2.SpanObjectV2.logs:type_name -> skywalking.v2.Log
	9,  // 11: skywalking.v2.Log.data:type_name -> skywalking.v2.KeyStringValuePair
	0,  // 12: skywalking.v2.TraceSegmentReportService.collect:input_type -> skywalking.v2.UpstreamSegment
	10, // 13: skywalking.v2.TraceSegmentReportService.collect:output_type -> skywalking.v2.Commands
	13, // [13:14] is the sub-list for method output_type
	12, // [12:13] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_init() }
func file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_init() {
	if File_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamSegment); i {
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
		file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentObject); i {
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
		file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentReference); i {
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
		file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpanObjectV2); i {
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
		file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_goTypes,
		DependencyIndexes: file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_depIdxs,
		MessageInfos:      file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_msgTypes,
	}.Build()
	File_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto = out.File
	file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_rawDesc = nil
	file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_goTypes = nil
	file_plugins_inputs_skywalking_v2_language_agent_v2_trace_proto_depIdxs = nil
}
