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
// source: plugins/inputs/skywalking/v3/proto/language-agent/CLRMetric.proto

package compile

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CLRMetricCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics         []*CLRMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Service         string       `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance string       `protobuf:"bytes,3,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
}

func (x *CLRMetricCollection) Reset() {
	*x = CLRMetricCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CLRMetricCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CLRMetricCollection) ProtoMessage() {}

func (x *CLRMetricCollection) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CLRMetricCollection.ProtoReflect.Descriptor instead.
func (*CLRMetricCollection) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescGZIP(), []int{0}
}

func (x *CLRMetricCollection) GetMetrics() []*CLRMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *CLRMetricCollection) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *CLRMetricCollection) GetServiceInstance() string {
	if x != nil {
		return x.ServiceInstance
	}
	return ""
}

type CLRMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time   int64      `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Cpu    *CPU       `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gc     *ClrGC     `protobuf:"bytes,3,opt,name=gc,proto3" json:"gc,omitempty"`
	Thread *ClrThread `protobuf:"bytes,4,opt,name=thread,proto3" json:"thread,omitempty"`
}

func (x *CLRMetric) Reset() {
	*x = CLRMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CLRMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CLRMetric) ProtoMessage() {}

func (x *CLRMetric) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CLRMetric.ProtoReflect.Descriptor instead.
func (*CLRMetric) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescGZIP(), []int{1}
}

func (x *CLRMetric) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *CLRMetric) GetCpu() *CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *CLRMetric) GetGc() *ClrGC {
	if x != nil {
		return x.Gc
	}
	return nil
}

func (x *CLRMetric) GetThread() *ClrThread {
	if x != nil {
		return x.Thread
	}
	return nil
}

type ClrGC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gen0CollectCount int64 `protobuf:"varint,1,opt,name=Gen0CollectCount,proto3" json:"Gen0CollectCount,omitempty"`
	Gen1CollectCount int64 `protobuf:"varint,2,opt,name=Gen1CollectCount,proto3" json:"Gen1CollectCount,omitempty"`
	Gen2CollectCount int64 `protobuf:"varint,3,opt,name=Gen2CollectCount,proto3" json:"Gen2CollectCount,omitempty"`
	HeapMemory       int64 `protobuf:"varint,4,opt,name=HeapMemory,proto3" json:"HeapMemory,omitempty"`
}

func (x *ClrGC) Reset() {
	*x = ClrGC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClrGC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClrGC) ProtoMessage() {}

func (x *ClrGC) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClrGC.ProtoReflect.Descriptor instead.
func (*ClrGC) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescGZIP(), []int{2}
}

func (x *ClrGC) GetGen0CollectCount() int64 {
	if x != nil {
		return x.Gen0CollectCount
	}
	return 0
}

func (x *ClrGC) GetGen1CollectCount() int64 {
	if x != nil {
		return x.Gen1CollectCount
	}
	return 0
}

func (x *ClrGC) GetGen2CollectCount() int64 {
	if x != nil {
		return x.Gen2CollectCount
	}
	return 0
}

func (x *ClrGC) GetHeapMemory() int64 {
	if x != nil {
		return x.HeapMemory
	}
	return 0
}

type ClrThread struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableCompletionPortThreads int32 `protobuf:"varint,1,opt,name=AvailableCompletionPortThreads,proto3" json:"AvailableCompletionPortThreads,omitempty"`
	AvailableWorkerThreads         int32 `protobuf:"varint,2,opt,name=AvailableWorkerThreads,proto3" json:"AvailableWorkerThreads,omitempty"`
	MaxCompletionPortThreads       int32 `protobuf:"varint,3,opt,name=MaxCompletionPortThreads,proto3" json:"MaxCompletionPortThreads,omitempty"`
	MaxWorkerThreads               int32 `protobuf:"varint,4,opt,name=MaxWorkerThreads,proto3" json:"MaxWorkerThreads,omitempty"`
}

func (x *ClrThread) Reset() {
	*x = ClrThread{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClrThread) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClrThread) ProtoMessage() {}

func (x *ClrThread) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClrThread.ProtoReflect.Descriptor instead.
func (*ClrThread) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescGZIP(), []int{3}
}

func (x *ClrThread) GetAvailableCompletionPortThreads() int32 {
	if x != nil {
		return x.AvailableCompletionPortThreads
	}
	return 0
}

func (x *ClrThread) GetAvailableWorkerThreads() int32 {
	if x != nil {
		return x.AvailableWorkerThreads
	}
	return 0
}

func (x *ClrThread) GetMaxCompletionPortThreads() int32 {
	if x != nil {
		return x.MaxCompletionPortThreads
	}
	return 0
}

func (x *ClrThread) GetMaxWorkerThreads() int32 {
	if x != nil {
		return x.MaxWorkerThreads
	}
	return 0
}

var File_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto protoreflect.FileDescriptor

var file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDesc = []byte{
	0x0a, 0x41, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x43, 0x4c, 0x52, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x1a, 0x36, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x43,
	0x4c, 0x52, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x4c, 0x52, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x09, 0x43,
	0x4c, 0x52, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x03,
	0x63, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6b, 0x79, 0x77,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x50, 0x55, 0x52, 0x03, 0x63,
	0x70, 0x75, 0x12, 0x24, 0x0a, 0x02, 0x67, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x6c, 0x72, 0x47, 0x43, 0x52, 0x02, 0x67, 0x63, 0x12, 0x30, 0x0a, 0x06, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x72, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x52, 0x06, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x05, 0x43,
	0x6c, 0x72, 0x47, 0x43, 0x12, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x30, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x47, 0x65, 0x6e, 0x30, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x31, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x47, 0x65, 0x6e, 0x31,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x47, 0x65, 0x6e, 0x32, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x47, 0x65, 0x6e, 0x32, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x48, 0x65, 0x61, 0x70,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x48, 0x65,
	0x61, 0x70, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0xf3, 0x01, 0x0a, 0x09, 0x43, 0x6c, 0x72,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x46, 0x0a, 0x1e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72,
	0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1e,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x36,
	0x0a, 0x16, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x4d, 0x61, 0x78, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x4d, 0x61,
	0x78, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x32, 0x62,
	0x0a, 0x16, 0x43, 0x4c, 0x52, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x4c, 0x52, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x22, 0x00, 0x42, 0xaa, 0x01, 0x0a, 0x33, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6a, 0x69, 0x61, 0x67, 0x6f, 0x75, 0x79, 0x75, 0x6e, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x61, 0x72, 0x65, 0x2d, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0xaa,
	0x02, 0x1d, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescOnce sync.Once
	file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescData = file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDesc
)

func file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescGZIP() []byte {
	file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescOnce.Do(func() {
		file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescData)
	})
	return file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDescData
}

var (
	file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
	file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_goTypes  = []interface{}{
		(*CLRMetricCollection)(nil), // 0: skywalking.v3.CLRMetricCollection
		(*CLRMetric)(nil),           // 1: skywalking.v3.CLRMetric
		(*ClrGC)(nil),               // 2: skywalking.v3.ClrGC
		(*ClrThread)(nil),           // 3: skywalking.v3.ClrThread
		(*CPU)(nil),                 // 4: skywalking.v3.CPU
		(*Commands)(nil),            // 5: skywalking.v3.Commands
	}
)

var file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_depIdxs = []int32{
	1, // 0: skywalking.v3.CLRMetricCollection.metrics:type_name -> skywalking.v3.CLRMetric
	4, // 1: skywalking.v3.CLRMetric.cpu:type_name -> skywalking.v3.CPU
	2, // 2: skywalking.v3.CLRMetric.gc:type_name -> skywalking.v3.ClrGC
	3, // 3: skywalking.v3.CLRMetric.thread:type_name -> skywalking.v3.ClrThread
	0, // 4: skywalking.v3.CLRMetricReportService.collect:input_type -> skywalking.v3.CLRMetricCollection
	5, // 5: skywalking.v3.CLRMetricReportService.collect:output_type -> skywalking.v3.Commands
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_init() }
func file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_init() {
	if File_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto != nil {
		return
	}
	file_plugins_inputs_skywalking_v3_proto_common_Common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CLRMetricCollection); i {
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
		file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CLRMetric); i {
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
		file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClrGC); i {
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
		file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClrThread); i {
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
			RawDescriptor: file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_goTypes,
		DependencyIndexes: file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_depIdxs,
		MessageInfos:      file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_msgTypes,
	}.Build()
	File_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto = out.File
	file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_rawDesc = nil
	file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_goTypes = nil
	file_plugins_inputs_skywalking_v3_proto_language_agent_CLRMetric_proto_depIdxs = nil
}
