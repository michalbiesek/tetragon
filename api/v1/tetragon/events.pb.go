// Copyright 2020 Authors of Hubble
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: tetragon/events.proto

package tetragon

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the type of a Tetragon event.
//
// NOTE: EventType constants must be in sync with the numbers used in the
// GetEventsResponse event oneof.
type EventType int32

const (
	EventType_UNDEF              EventType = 0
	EventType_PROCESS_EXEC       EventType = 1
	EventType_PROCESS_EXIT       EventType = 5
	EventType_PROCESS_KPROBE     EventType = 9
	EventType_PROCESS_TRACEPOINT EventType = 10
	EventType_TEST               EventType = 40000
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0:     "UNDEF",
		1:     "PROCESS_EXEC",
		5:     "PROCESS_EXIT",
		9:     "PROCESS_KPROBE",
		10:    "PROCESS_TRACEPOINT",
		40000: "TEST",
	}
	EventType_value = map[string]int32{
		"UNDEF":              0,
		"PROCESS_EXEC":       1,
		"PROCESS_EXIT":       5,
		"PROCESS_KPROBE":     9,
		"PROCESS_TRACEPOINT": 10,
		"TEST":               40000,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_tetragon_events_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_tetragon_events_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_tetragon_events_proto_rawDescGZIP(), []int{0}
}

// Determins the behaviour of a field filter
type FieldFilterAction int32

const (
	FieldFilterAction_INCLUDE FieldFilterAction = 0
	FieldFilterAction_EXCLUDE FieldFilterAction = 1
)

// Enum value maps for FieldFilterAction.
var (
	FieldFilterAction_name = map[int32]string{
		0: "INCLUDE",
		1: "EXCLUDE",
	}
	FieldFilterAction_value = map[string]int32{
		"INCLUDE": 0,
		"EXCLUDE": 1,
	}
)

func (x FieldFilterAction) Enum() *FieldFilterAction {
	p := new(FieldFilterAction)
	*p = x
	return p
}

func (x FieldFilterAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldFilterAction) Descriptor() protoreflect.EnumDescriptor {
	return file_tetragon_events_proto_enumTypes[1].Descriptor()
}

func (FieldFilterAction) Type() protoreflect.EnumType {
	return &file_tetragon_events_proto_enumTypes[1]
}

func (x FieldFilterAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldFilterAction.Descriptor instead.
func (FieldFilterAction) EnumDescriptor() ([]byte, []int) {
	return file_tetragon_events_proto_rawDescGZIP(), []int{1}
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BinaryRegex []string              `protobuf:"bytes,1,rep,name=binary_regex,json=binaryRegex,proto3" json:"binary_regex,omitempty"`
	Namespace   []string              `protobuf:"bytes,2,rep,name=namespace,proto3" json:"namespace,omitempty"`
	HealthCheck *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=health_check,json=healthCheck,proto3" json:"health_check,omitempty"`
	Pid         []uint32              `protobuf:"varint,4,rep,packed,name=pid,proto3" json:"pid,omitempty"`
	PidSet      []uint32              `protobuf:"varint,5,rep,packed,name=pid_set,json=pidSet,proto3" json:"pid_set,omitempty"`
	EventSet    []EventType           `protobuf:"varint,6,rep,packed,name=event_set,json=eventSet,proto3,enum=tetragon.EventType" json:"event_set,omitempty"`
	// Filter by process.pod.name field using RE2 regular expression syntax:
	// https://github.com/google/re2/wiki/Syntax
	PodRegex []string `protobuf:"bytes,7,rep,name=pod_regex,json=podRegex,proto3" json:"pod_regex,omitempty"`
	// Filter by process.arguments field using RE2 regular expression syntax:
	// https://github.com/google/re2/wiki/Syntax
	ArgumentsRegex []string `protobuf:"bytes,8,rep,name=arguments_regex,json=argumentsRegex,proto3" json:"arguments_regex,omitempty"`
	// Filter events by pod labels using Kubernetes label selector syntax:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	// Note that this filter never matches events without the pod field (i.e.
	// host process events).
	Labels []string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_tetragon_events_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetBinaryRegex() []string {
	if x != nil {
		return x.BinaryRegex
	}
	return nil
}

func (x *Filter) GetNamespace() []string {
	if x != nil {
		return x.Namespace
	}
	return nil
}

func (x *Filter) GetHealthCheck() *wrapperspb.BoolValue {
	if x != nil {
		return x.HealthCheck
	}
	return nil
}

func (x *Filter) GetPid() []uint32 {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *Filter) GetPidSet() []uint32 {
	if x != nil {
		return x.PidSet
	}
	return nil
}

func (x *Filter) GetEventSet() []EventType {
	if x != nil {
		return x.EventSet
	}
	return nil
}

func (x *Filter) GetPodRegex() []string {
	if x != nil {
		return x.PodRegex
	}
	return nil
}

func (x *Filter) GetArgumentsRegex() []string {
	if x != nil {
		return x.ArgumentsRegex
	}
	return nil
}

func (x *Filter) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type FieldFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event types to filter or undefined to filter over all event types.
	EventSet []EventType `protobuf:"varint,1,rep,packed,name=event_set,json=eventSet,proto3,enum=tetragon.EventType" json:"event_set,omitempty"`
	// Fields to include or exclude.
	Fields *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	// Whether to include or exclude fields.
	Action FieldFilterAction `protobuf:"varint,3,opt,name=action,proto3,enum=tetragon.FieldFilterAction" json:"action,omitempty"`
	// Whether or not the event set filter should be inverted.
	InvertEventSet *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=invert_event_set,json=invertEventSet,proto3" json:"invert_event_set,omitempty"`
}

func (x *FieldFilter) Reset() {
	*x = FieldFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldFilter) ProtoMessage() {}

func (x *FieldFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldFilter.ProtoReflect.Descriptor instead.
func (*FieldFilter) Descriptor() ([]byte, []int) {
	return file_tetragon_events_proto_rawDescGZIP(), []int{1}
}

func (x *FieldFilter) GetEventSet() []EventType {
	if x != nil {
		return x.EventSet
	}
	return nil
}

func (x *FieldFilter) GetFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *FieldFilter) GetAction() FieldFilterAction {
	if x != nil {
		return x.Action
	}
	return FieldFilterAction_INCLUDE
}

func (x *FieldFilter) GetInvertEventSet() *wrapperspb.BoolValue {
	if x != nil {
		return x.InvertEventSet
	}
	return nil
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// allow_list specifies a list of filters to apply to only return certain
	// events. If multiple filters are specified, at least one of them has to
	// match for an event to be included in the results.
	AllowList []*Filter `protobuf:"bytes,1,rep,name=allow_list,json=allowList,proto3" json:"allow_list,omitempty"`
	// deny_list specifies a list of filters to apply to exclude certain events
	// from the results. If multiple filters are specified, at least one of
	// them has to match for an event to be excluded.
	//
	// If both allow_list and deny_list are specified, the results contain the
	// set difference allow_list - deny_list.
	DenyList []*Filter `protobuf:"bytes,2,rep,name=deny_list,json=denyList,proto3" json:"deny_list,omitempty"`
	// aggregation_options configures aggregation options for this request.
	// If this field is not set, responses will not be aggregated.
	//
	// Note that currently only process_accept and process_connect events are
	// aggregated. Other events remain unaggregated.
	AggregationOptions *AggregationOptions `protobuf:"bytes,3,opt,name=aggregation_options,json=aggregationOptions,proto3" json:"aggregation_options,omitempty"`
	// Fields to include or exclude for events in the GetEventsResponse. Omitting this
	// field implies that all fields will be included. Exclusion always takes precedence
	// over inclusion in the case of conflicts.
	FieldFilters []*FieldFilter `protobuf:"bytes,4,rep,name=field_filters,json=fieldFilters,proto3" json:"field_filters,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_tetragon_events_proto_rawDescGZIP(), []int{2}
}

func (x *GetEventsRequest) GetAllowList() []*Filter {
	if x != nil {
		return x.AllowList
	}
	return nil
}

func (x *GetEventsRequest) GetDenyList() []*Filter {
	if x != nil {
		return x.DenyList
	}
	return nil
}

func (x *GetEventsRequest) GetAggregationOptions() *AggregationOptions {
	if x != nil {
		return x.AggregationOptions
	}
	return nil
}

func (x *GetEventsRequest) GetFieldFilters() []*FieldFilter {
	if x != nil {
		return x.FieldFilters
	}
	return nil
}

// AggregationOptions defines configuration options for aggregating events.
type AggregationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Aggregation window size. Defaults to 15 seconds if this field is not set.
	WindowSize *durationpb.Duration `protobuf:"bytes,1,opt,name=window_size,json=windowSize,proto3" json:"window_size,omitempty"`
	// Size of the buffer for the aggregator to receive incoming events. If the
	// buffer becomes full, the aggregator will log a warning and start dropping
	// incoming events.
	ChannelBufferSize uint64 `protobuf:"varint,2,opt,name=channel_buffer_size,json=channelBufferSize,proto3" json:"channel_buffer_size,omitempty"`
}

func (x *AggregationOptions) Reset() {
	*x = AggregationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregationOptions) ProtoMessage() {}

func (x *AggregationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregationOptions.ProtoReflect.Descriptor instead.
func (*AggregationOptions) Descriptor() ([]byte, []int) {
	return file_tetragon_events_proto_rawDescGZIP(), []int{3}
}

func (x *AggregationOptions) GetWindowSize() *durationpb.Duration {
	if x != nil {
		return x.WindowSize
	}
	return nil
}

func (x *AggregationOptions) GetChannelBufferSize() uint64 {
	if x != nil {
		return x.ChannelBufferSize
	}
	return 0
}

// AggregationInfo contains information about aggregation results.
type AggregationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total count of events in this aggregation time window.
	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AggregationInfo) Reset() {
	*x = AggregationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregationInfo) ProtoMessage() {}

func (x *AggregationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregationInfo.ProtoReflect.Descriptor instead.
func (*AggregationInfo) Descriptor() ([]byte, []int) {
	return file_tetragon_events_proto_rawDescGZIP(), []int{4}
}

func (x *AggregationInfo) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type-specific fields of an event.
	//
	// NOTE: Numbers must stay in sync with enum EventType.
	//
	// Types that are assignable to Event:
	//	*GetEventsResponse_ProcessExec
	//	*GetEventsResponse_ProcessExit
	//	*GetEventsResponse_ProcessKprobe
	//	*GetEventsResponse_ProcessTracepoint
	//	*GetEventsResponse_Test
	Event isGetEventsResponse_Event `protobuf_oneof:"event"`
	// Name of the node where this event was observed.
	NodeName string `protobuf:"bytes,1000,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Timestamp at which this event was observed.
	//
	// For an aggregated response, this field to set to the timestamp at which
	// the event was observed for the first time in a given aggregation time window.
	Time *timestamppb.Timestamp `protobuf:"bytes,1001,opt,name=time,proto3" json:"time,omitempty"`
	// aggregation_info contains information about aggregation results. This field
	// is set only for aggregated responses.
	AggregationInfo *AggregationInfo `protobuf:"bytes,1002,opt,name=aggregation_info,json=aggregationInfo,proto3" json:"aggregation_info,omitempty"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tetragon_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tetragon_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_tetragon_events_proto_rawDescGZIP(), []int{5}
}

func (m *GetEventsResponse) GetEvent() isGetEventsResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *GetEventsResponse) GetProcessExec() *ProcessExec {
	if x, ok := x.GetEvent().(*GetEventsResponse_ProcessExec); ok {
		return x.ProcessExec
	}
	return nil
}

func (x *GetEventsResponse) GetProcessExit() *ProcessExit {
	if x, ok := x.GetEvent().(*GetEventsResponse_ProcessExit); ok {
		return x.ProcessExit
	}
	return nil
}

func (x *GetEventsResponse) GetProcessKprobe() *ProcessKprobe {
	if x, ok := x.GetEvent().(*GetEventsResponse_ProcessKprobe); ok {
		return x.ProcessKprobe
	}
	return nil
}

func (x *GetEventsResponse) GetProcessTracepoint() *ProcessTracepoint {
	if x, ok := x.GetEvent().(*GetEventsResponse_ProcessTracepoint); ok {
		return x.ProcessTracepoint
	}
	return nil
}

func (x *GetEventsResponse) GetTest() *Test {
	if x, ok := x.GetEvent().(*GetEventsResponse_Test); ok {
		return x.Test
	}
	return nil
}

func (x *GetEventsResponse) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *GetEventsResponse) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *GetEventsResponse) GetAggregationInfo() *AggregationInfo {
	if x != nil {
		return x.AggregationInfo
	}
	return nil
}

type isGetEventsResponse_Event interface {
	isGetEventsResponse_Event()
}

type GetEventsResponse_ProcessExec struct {
	ProcessExec *ProcessExec `protobuf:"bytes,1,opt,name=process_exec,json=processExec,proto3,oneof"`
}

type GetEventsResponse_ProcessExit struct {
	ProcessExit *ProcessExit `protobuf:"bytes,5,opt,name=process_exit,json=processExit,proto3,oneof"`
}

type GetEventsResponse_ProcessKprobe struct {
	ProcessKprobe *ProcessKprobe `protobuf:"bytes,9,opt,name=process_kprobe,json=processKprobe,proto3,oneof"`
}

type GetEventsResponse_ProcessTracepoint struct {
	ProcessTracepoint *ProcessTracepoint `protobuf:"bytes,10,opt,name=process_tracepoint,json=processTracepoint,proto3,oneof"`
}

type GetEventsResponse_Test struct {
	Test *Test `protobuf:"bytes,40000,opt,name=test,proto3,oneof"`
}

func (*GetEventsResponse_ProcessExec) isGetEventsResponse_Event() {}

func (*GetEventsResponse_ProcessExit) isGetEventsResponse_Event() {}

func (*GetEventsResponse_ProcessKprobe) isGetEventsResponse_Event() {}

func (*GetEventsResponse_ProcessTracepoint) isGetEventsResponse_Event() {}

func (*GetEventsResponse_Test) isGetEventsResponse_Event() {}

var File_tetragon_events_proto protoreflect.FileDescriptor

var file_tetragon_events_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f,
	0x6e, 0x1a, 0x17, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x74, 0x72,
	0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x02,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x69,
	0x64, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x69, 0x64,
	0x53, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f,
	0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x64, 0x5f, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x64, 0x52, 0x65, 0x67,
	0x65, 0x78, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f,
	0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x22, 0xee, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f,
	0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x74, 0x72,
	0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44,
	0x0a, 0x10, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x74, 0x22, 0xfd, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x64, 0x65,
	0x6e, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x08, 0x64, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x13, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f,
	0x6e, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x12, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x77,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xe2, 0x03, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74,
	0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45,
	0x78, 0x65, 0x63, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78,
	0x65, 0x63, 0x12, 0x3a, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61,
	0x67, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x69, 0x74, 0x48,
	0x00, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x69, 0x74, 0x12, 0x40,
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f,
	0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x48,
	0x00, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x12, 0x4c, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74,
	0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x11, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x72, 0x61, 0x63, 0x65, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0xc0, 0xb8, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xe9, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x07, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2a, 0x72, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x10,
	0x05, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4b, 0x50, 0x52,
	0x4f, 0x42, 0x45, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x54, 0x52, 0x41, 0x43, 0x45, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x0a, 0x12, 0x0a, 0x0a,
	0x04, 0x54, 0x45, 0x53, 0x54, 0x10, 0xc0, 0xb8, 0x02, 0x2a, 0x2d, 0x0a, 0x11, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45,
	0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tetragon_events_proto_rawDescOnce sync.Once
	file_tetragon_events_proto_rawDescData = file_tetragon_events_proto_rawDesc
)

func file_tetragon_events_proto_rawDescGZIP() []byte {
	file_tetragon_events_proto_rawDescOnce.Do(func() {
		file_tetragon_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_tetragon_events_proto_rawDescData)
	})
	return file_tetragon_events_proto_rawDescData
}

var file_tetragon_events_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tetragon_events_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tetragon_events_proto_goTypes = []interface{}{
	(EventType)(0),                // 0: tetragon.EventType
	(FieldFilterAction)(0),        // 1: tetragon.FieldFilterAction
	(*Filter)(nil),                // 2: tetragon.Filter
	(*FieldFilter)(nil),           // 3: tetragon.FieldFilter
	(*GetEventsRequest)(nil),      // 4: tetragon.GetEventsRequest
	(*AggregationOptions)(nil),    // 5: tetragon.AggregationOptions
	(*AggregationInfo)(nil),       // 6: tetragon.AggregationInfo
	(*GetEventsResponse)(nil),     // 7: tetragon.GetEventsResponse
	(*wrapperspb.BoolValue)(nil),  // 8: google.protobuf.BoolValue
	(*fieldmaskpb.FieldMask)(nil), // 9: google.protobuf.FieldMask
	(*durationpb.Duration)(nil),   // 10: google.protobuf.Duration
	(*ProcessExec)(nil),           // 11: tetragon.ProcessExec
	(*ProcessExit)(nil),           // 12: tetragon.ProcessExit
	(*ProcessKprobe)(nil),         // 13: tetragon.ProcessKprobe
	(*ProcessTracepoint)(nil),     // 14: tetragon.ProcessTracepoint
	(*Test)(nil),                  // 15: tetragon.Test
	(*timestamppb.Timestamp)(nil), // 16: google.protobuf.Timestamp
}
var file_tetragon_events_proto_depIdxs = []int32{
	8,  // 0: tetragon.Filter.health_check:type_name -> google.protobuf.BoolValue
	0,  // 1: tetragon.Filter.event_set:type_name -> tetragon.EventType
	0,  // 2: tetragon.FieldFilter.event_set:type_name -> tetragon.EventType
	9,  // 3: tetragon.FieldFilter.fields:type_name -> google.protobuf.FieldMask
	1,  // 4: tetragon.FieldFilter.action:type_name -> tetragon.FieldFilterAction
	8,  // 5: tetragon.FieldFilter.invert_event_set:type_name -> google.protobuf.BoolValue
	2,  // 6: tetragon.GetEventsRequest.allow_list:type_name -> tetragon.Filter
	2,  // 7: tetragon.GetEventsRequest.deny_list:type_name -> tetragon.Filter
	5,  // 8: tetragon.GetEventsRequest.aggregation_options:type_name -> tetragon.AggregationOptions
	3,  // 9: tetragon.GetEventsRequest.field_filters:type_name -> tetragon.FieldFilter
	10, // 10: tetragon.AggregationOptions.window_size:type_name -> google.protobuf.Duration
	11, // 11: tetragon.GetEventsResponse.process_exec:type_name -> tetragon.ProcessExec
	12, // 12: tetragon.GetEventsResponse.process_exit:type_name -> tetragon.ProcessExit
	13, // 13: tetragon.GetEventsResponse.process_kprobe:type_name -> tetragon.ProcessKprobe
	14, // 14: tetragon.GetEventsResponse.process_tracepoint:type_name -> tetragon.ProcessTracepoint
	15, // 15: tetragon.GetEventsResponse.test:type_name -> tetragon.Test
	16, // 16: tetragon.GetEventsResponse.time:type_name -> google.protobuf.Timestamp
	6,  // 17: tetragon.GetEventsResponse.aggregation_info:type_name -> tetragon.AggregationInfo
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_tetragon_events_proto_init() }
func file_tetragon_events_proto_init() {
	if File_tetragon_events_proto != nil {
		return
	}
	file_tetragon_tetragon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tetragon_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_tetragon_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldFilter); i {
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
		file_tetragon_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
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
		file_tetragon_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregationOptions); i {
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
		file_tetragon_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregationInfo); i {
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
		file_tetragon_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsResponse); i {
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
	file_tetragon_events_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*GetEventsResponse_ProcessExec)(nil),
		(*GetEventsResponse_ProcessExit)(nil),
		(*GetEventsResponse_ProcessKprobe)(nil),
		(*GetEventsResponse_ProcessTracepoint)(nil),
		(*GetEventsResponse_Test)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tetragon_events_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tetragon_events_proto_goTypes,
		DependencyIndexes: file_tetragon_events_proto_depIdxs,
		EnumInfos:         file_tetragon_events_proto_enumTypes,
		MessageInfos:      file_tetragon_events_proto_msgTypes,
	}.Build()
	File_tetragon_events_proto = out.File
	file_tetragon_events_proto_rawDesc = nil
	file_tetragon_events_proto_goTypes = nil
	file_tetragon_events_proto_depIdxs = nil
}
