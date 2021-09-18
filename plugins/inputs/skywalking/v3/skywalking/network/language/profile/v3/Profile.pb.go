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
// source: plugins/inputs/skywalking/v3/proto/profile/Profile.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	v3 "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking/v3/skywalking/network/common/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProfileTaskCommandQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// current sniffer information
	Service         string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance string `protobuf:"bytes,2,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	// last command timestamp
	LastCommandTime int64 `protobuf:"varint,3,opt,name=lastCommandTime,proto3" json:"lastCommandTime,omitempty"`
}

func (x *ProfileTaskCommandQuery) Reset() {
	*x = ProfileTaskCommandQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileTaskCommandQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileTaskCommandQuery) ProtoMessage() {}

func (x *ProfileTaskCommandQuery) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileTaskCommandQuery.ProtoReflect.Descriptor instead.
func (*ProfileTaskCommandQuery) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileTaskCommandQuery) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ProfileTaskCommandQuery) GetServiceInstance() string {
	if x != nil {
		return x.ServiceInstance
	}
	return ""
}

func (x *ProfileTaskCommandQuery) GetLastCommandTime() int64 {
	if x != nil {
		return x.LastCommandTime
	}
	return 0
}

// dumped thread snapshot
type ThreadSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// profile task id
	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	// dumped segment id
	TraceSegmentId string `protobuf:"bytes,2,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	// dump timestamp
	Time int64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	// snapshot dump sequence, start with zero
	Sequence int32 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// snapshot stack
	Stack *ThreadStack `protobuf:"bytes,5,opt,name=stack,proto3" json:"stack,omitempty"`
}

func (x *ThreadSnapshot) Reset() {
	*x = ThreadSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadSnapshot) ProtoMessage() {}

func (x *ThreadSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadSnapshot.ProtoReflect.Descriptor instead.
func (*ThreadSnapshot) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescGZIP(), []int{1}
}

func (x *ThreadSnapshot) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *ThreadSnapshot) GetTraceSegmentId() string {
	if x != nil {
		return x.TraceSegmentId
	}
	return ""
}

func (x *ThreadSnapshot) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ThreadSnapshot) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *ThreadSnapshot) GetStack() *ThreadStack {
	if x != nil {
		return x.Stack
	}
	return nil
}

type ThreadStack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack code signature list
	CodeSignatures []string `protobuf:"bytes,1,rep,name=codeSignatures,proto3" json:"codeSignatures,omitempty"`
}

func (x *ThreadStack) Reset() {
	*x = ThreadStack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadStack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadStack) ProtoMessage() {}

func (x *ThreadStack) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadStack.ProtoReflect.Descriptor instead.
func (*ThreadStack) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescGZIP(), []int{2}
}

func (x *ThreadStack) GetCodeSignatures() []string {
	if x != nil {
		return x.CodeSignatures
	}
	return nil
}

// profile task finished report
type ProfileTaskFinishReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// current sniffer information
	Service         string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance string `protobuf:"bytes,2,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	// profile task
	TaskId string `protobuf:"bytes,3,opt,name=taskId,proto3" json:"taskId,omitempty"`
}

func (x *ProfileTaskFinishReport) Reset() {
	*x = ProfileTaskFinishReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileTaskFinishReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileTaskFinishReport) ProtoMessage() {}

func (x *ProfileTaskFinishReport) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileTaskFinishReport.ProtoReflect.Descriptor instead.
func (*ProfileTaskFinishReport) Descriptor() ([]byte, []int) {
	return file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileTaskFinishReport) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ProfileTaskFinishReport) GetServiceInstance() string {
	if x != nil {
		return x.ServiceInstance
	}
	return ""
}

func (x *ProfileTaskFinishReport) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

var File_plugins_inputs_skywalking_v3_proto_profile_Profile_proto protoreflect.FileDescriptor

var file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDesc = []byte{
	0x0a, 0x38, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa4, 0x01, 0x0a,
	0x0e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x22, 0x35, 0x0a, 0x0b, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x64, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x75, 0x0a, 0x17, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x32, 0xb6, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x3d, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x09, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x12, 0x2f, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x12, 0x0f, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x1a, 0x09, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x28,
	0x01, 0x12, 0x37, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x18, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a,
	0x09, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x42, 0xcb, 0x01, 0x0a, 0x35, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x33, 0x50, 0x01, 0x5a, 0x70, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6a,
	0x69, 0x61, 0x67, 0x6f, 0x75, 0x79, 0x75, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x63, 0x61, 0x72, 0x65, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x33, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x33, 0xaa, 0x02, 0x1d, 0x53, 0x6b, 0x79, 0x57, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescOnce sync.Once
	file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescData = file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDesc
)

func file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescGZIP() []byte {
	file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescOnce.Do(func() {
		file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescData)
	})
	return file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDescData
}

var (
	file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
	file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_goTypes  = []interface{}{
		(*ProfileTaskCommandQuery)(nil), // 0: ProfileTaskCommandQuery
		(*ThreadSnapshot)(nil),          // 1: ThreadSnapshot
		(*ThreadStack)(nil),             // 2: ThreadStack
		(*ProfileTaskFinishReport)(nil), // 3: ProfileTaskFinishReport
		(*v3.Commands)(nil),             // 4: Commands
	}
)

var file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_depIdxs = []int32{
	2, // 0: ThreadSnapshot.stack:type_name -> ThreadStack
	0, // 1: ProfileTask.getProfileTaskCommands:input_type -> ProfileTaskCommandQuery
	1, // 2: ProfileTask.collectSnapshot:input_type -> ThreadSnapshot
	3, // 3: ProfileTask.reportTaskFinish:input_type -> ProfileTaskFinishReport
	4, // 4: ProfileTask.getProfileTaskCommands:output_type -> Commands
	4, // 5: ProfileTask.collectSnapshot:output_type -> Commands
	4, // 6: ProfileTask.reportTaskFinish:output_type -> Commands
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_init() }
func file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_init() {
	if File_plugins_inputs_skywalking_v3_proto_profile_Profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileTaskCommandQuery); i {
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
		file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadSnapshot); i {
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
		file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadStack); i {
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
		file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileTaskFinishReport); i {
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
			RawDescriptor: file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_goTypes,
		DependencyIndexes: file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_depIdxs,
		MessageInfos:      file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_msgTypes,
	}.Build()
	File_plugins_inputs_skywalking_v3_proto_profile_Profile_proto = out.File
	file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_rawDesc = nil
	file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_goTypes = nil
	file_plugins_inputs_skywalking_v3_proto_profile_Profile_proto_depIdxs = nil
}
