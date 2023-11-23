// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/operations.proto

package v1

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

// OperationEventType indicates whether the operation was created or updated
type OperationEventType int32

const (
	OperationEventType_EVENT_UNKNOWN OperationEventType = 0
	OperationEventType_EVENT_CREATED OperationEventType = 1
	OperationEventType_EVENT_UPDATED OperationEventType = 2
)

// Enum value maps for OperationEventType.
var (
	OperationEventType_name = map[int32]string{
		0: "EVENT_UNKNOWN",
		1: "EVENT_CREATED",
		2: "EVENT_UPDATED",
	}
	OperationEventType_value = map[string]int32{
		"EVENT_UNKNOWN": 0,
		"EVENT_CREATED": 1,
		"EVENT_UPDATED": 2,
	}
)

func (x OperationEventType) Enum() *OperationEventType {
	p := new(OperationEventType)
	*p = x
	return p
}

func (x OperationEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_operations_proto_enumTypes[0].Descriptor()
}

func (OperationEventType) Type() protoreflect.EnumType {
	return &file_v1_operations_proto_enumTypes[0]
}

func (x OperationEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationEventType.Descriptor instead.
func (OperationEventType) EnumDescriptor() ([]byte, []int) {
	return file_v1_operations_proto_rawDescGZIP(), []int{0}
}

type OperationStatus int32

const (
	OperationStatus_STATUS_UNKNOWN    OperationStatus = 0
	OperationStatus_STATUS_PENDING    OperationStatus = 1
	OperationStatus_STATUS_INPROGRESS OperationStatus = 2
	OperationStatus_STATUS_SUCCESS    OperationStatus = 3
	OperationStatus_STATUS_ERROR      OperationStatus = 4
)

// Enum value maps for OperationStatus.
var (
	OperationStatus_name = map[int32]string{
		0: "STATUS_UNKNOWN",
		1: "STATUS_PENDING",
		2: "STATUS_INPROGRESS",
		3: "STATUS_SUCCESS",
		4: "STATUS_ERROR",
	}
	OperationStatus_value = map[string]int32{
		"STATUS_UNKNOWN":    0,
		"STATUS_PENDING":    1,
		"STATUS_INPROGRESS": 2,
		"STATUS_SUCCESS":    3,
		"STATUS_ERROR":      4,
	}
)

func (x OperationStatus) Enum() *OperationStatus {
	p := new(OperationStatus)
	*p = x
	return p
}

func (x OperationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_operations_proto_enumTypes[1].Descriptor()
}

func (OperationStatus) Type() protoreflect.EnumType {
	return &file_v1_operations_proto_enumTypes[1]
}

func (x OperationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationStatus.Descriptor instead.
func (OperationStatus) EnumDescriptor() ([]byte, []int) {
	return file_v1_operations_proto_rawDescGZIP(), []int{1}
}

type OperationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *OperationList) Reset() {
	*x = OperationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_operations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationList) ProtoMessage() {}

func (x *OperationList) ProtoReflect() protoreflect.Message {
	mi := &file_v1_operations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationList.ProtoReflect.Descriptor instead.
func (*OperationList) Descriptor() ([]byte, []int) {
	return file_v1_operations_proto_rawDescGZIP(), []int{0}
}

func (x *OperationList) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: `storm:"id,increment"`
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" storm:"id,increment"`
	// repo id if associated with a repo (always true)
	// @gotags: `storm:"index"`
	RepoId string `protobuf:"bytes,2,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty" storm:"index"`
	// plan id if associated with a plan (always true)
	// @gotags: `storm:"index"`
	PlanId string `protobuf:"bytes,3,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty" storm:"index"`
	// normalized snapshot id (first 8 chars) if associated with a snapshot
	// @gotags: `storm:"index"`
	SnapshotId      string          `protobuf:"bytes,8,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty" storm:"index"`
	Status          OperationStatus `protobuf:"varint,4,opt,name=status,proto3,enum=v1.OperationStatus" json:"status,omitempty"`
	UnixTimeStartMs int64           `protobuf:"varint,5,opt,name=unix_time_start_ms,json=unixTimeStartMs,proto3" json:"unix_time_start_ms,omitempty"`
	// operation ended time in unix ms
	UnixTimeEndMs  int64  `protobuf:"varint,6,opt,name=unix_time_end_ms,json=unixTimeEndMs,proto3" json:"unix_time_end_ms,omitempty"`
	DisplayMessage string `protobuf:"bytes,7,opt,name=display_message,json=displayMessage,proto3" json:"display_message,omitempty"` // human readable context message (if any)
	// Types that are assignable to Op:
	//
	//	*Operation_OperationBackup
	//	*Operation_OperationIndexSnapshot
	Op isOperation_Op `protobuf_oneof:"op"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_operations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_v1_operations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_v1_operations_proto_rawDescGZIP(), []int{1}
}

func (x *Operation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Operation) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *Operation) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *Operation) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

func (x *Operation) GetStatus() OperationStatus {
	if x != nil {
		return x.Status
	}
	return OperationStatus_STATUS_UNKNOWN
}

func (x *Operation) GetUnixTimeStartMs() int64 {
	if x != nil {
		return x.UnixTimeStartMs
	}
	return 0
}

func (x *Operation) GetUnixTimeEndMs() int64 {
	if x != nil {
		return x.UnixTimeEndMs
	}
	return 0
}

func (x *Operation) GetDisplayMessage() string {
	if x != nil {
		return x.DisplayMessage
	}
	return ""
}

func (m *Operation) GetOp() isOperation_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (x *Operation) GetOperationBackup() *OperationBackup {
	if x, ok := x.GetOp().(*Operation_OperationBackup); ok {
		return x.OperationBackup
	}
	return nil
}

func (x *Operation) GetOperationIndexSnapshot() *OperationIndexSnapshot {
	if x, ok := x.GetOp().(*Operation_OperationIndexSnapshot); ok {
		return x.OperationIndexSnapshot
	}
	return nil
}

type isOperation_Op interface {
	isOperation_Op()
}

type Operation_OperationBackup struct {
	OperationBackup *OperationBackup `protobuf:"bytes,100,opt,name=operation_backup,json=operationBackup,proto3,oneof"`
}

type Operation_OperationIndexSnapshot struct {
	OperationIndexSnapshot *OperationIndexSnapshot `protobuf:"bytes,101,opt,name=operation_index_snapshot,json=operationIndexSnapshot,proto3,oneof"`
}

func (*Operation_OperationBackup) isOperation_Op() {}

func (*Operation_OperationIndexSnapshot) isOperation_Op() {}

// OperationEvent is used in the wireformat to stream operation changes to clients
type OperationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      OperationEventType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.OperationEventType" json:"type,omitempty"`
	Operation *Operation         `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *OperationEvent) Reset() {
	*x = OperationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_operations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationEvent) ProtoMessage() {}

func (x *OperationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_v1_operations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationEvent.ProtoReflect.Descriptor instead.
func (*OperationEvent) Descriptor() ([]byte, []int) {
	return file_v1_operations_proto_rawDescGZIP(), []int{2}
}

func (x *OperationEvent) GetType() OperationEventType {
	if x != nil {
		return x.Type
	}
	return OperationEventType_EVENT_UNKNOWN
}

func (x *OperationEvent) GetOperation() *Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type OperationBackup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastStatus *BackupProgressEntry `protobuf:"bytes,3,opt,name=last_status,json=lastStatus,proto3" json:"last_status,omitempty"`
}

func (x *OperationBackup) Reset() {
	*x = OperationBackup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_operations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationBackup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationBackup) ProtoMessage() {}

func (x *OperationBackup) ProtoReflect() protoreflect.Message {
	mi := &file_v1_operations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationBackup.ProtoReflect.Descriptor instead.
func (*OperationBackup) Descriptor() ([]byte, []int) {
	return file_v1_operations_proto_rawDescGZIP(), []int{3}
}

func (x *OperationBackup) GetLastStatus() *BackupProgressEntry {
	if x != nil {
		return x.LastStatus
	}
	return nil
}

// OperationIndexSnapshot tracks that a snapshot was detected by resticui.
type OperationIndexSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snapshot *ResticSnapshot `protobuf:"bytes,2,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (x *OperationIndexSnapshot) Reset() {
	*x = OperationIndexSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_operations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationIndexSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationIndexSnapshot) ProtoMessage() {}

func (x *OperationIndexSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_v1_operations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationIndexSnapshot.ProtoReflect.Descriptor instead.
func (*OperationIndexSnapshot) Descriptor() ([]byte, []int) {
	return file_v1_operations_proto_rawDescGZIP(), []int{4}
}

func (x *OperationIndexSnapshot) GetSnapshot() *ResticSnapshot {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

var File_v1_operations_proto protoreflect.FileDescriptor

var file_v1_operations_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0d, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xba, 0x03, 0x0a, 0x09, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x75, 0x6e, 0x69, 0x78,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6d, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4d, 0x73, 0x12, 0x27, 0x0a, 0x10, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x75, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x4d, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x48, 0x00, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x56, 0x0a, 0x18, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x48, 0x00, 0x52, 0x16, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x42, 0x04, 0x0a, 0x02, 0x6f, 0x70, 0x22, 0x69, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x38, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x48, 0x0a, 0x16, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x74, 0x69, 0x63, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2a, 0x4d, 0x0a, 0x12, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x11, 0x0a, 0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x76, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x61, 0x72, 0x65, 0x74, 0x68, 0x67, 0x65, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x69, 0x63, 0x75, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_operations_proto_rawDescOnce sync.Once
	file_v1_operations_proto_rawDescData = file_v1_operations_proto_rawDesc
)

func file_v1_operations_proto_rawDescGZIP() []byte {
	file_v1_operations_proto_rawDescOnce.Do(func() {
		file_v1_operations_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_operations_proto_rawDescData)
	})
	return file_v1_operations_proto_rawDescData
}

var file_v1_operations_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_operations_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_operations_proto_goTypes = []interface{}{
	(OperationEventType)(0),        // 0: v1.OperationEventType
	(OperationStatus)(0),           // 1: v1.OperationStatus
	(*OperationList)(nil),          // 2: v1.OperationList
	(*Operation)(nil),              // 3: v1.Operation
	(*OperationEvent)(nil),         // 4: v1.OperationEvent
	(*OperationBackup)(nil),        // 5: v1.OperationBackup
	(*OperationIndexSnapshot)(nil), // 6: v1.OperationIndexSnapshot
	(*BackupProgressEntry)(nil),    // 7: v1.BackupProgressEntry
	(*ResticSnapshot)(nil),         // 8: v1.ResticSnapshot
}
var file_v1_operations_proto_depIdxs = []int32{
	3, // 0: v1.OperationList.operations:type_name -> v1.Operation
	1, // 1: v1.Operation.status:type_name -> v1.OperationStatus
	5, // 2: v1.Operation.operation_backup:type_name -> v1.OperationBackup
	6, // 3: v1.Operation.operation_index_snapshot:type_name -> v1.OperationIndexSnapshot
	0, // 4: v1.OperationEvent.type:type_name -> v1.OperationEventType
	3, // 5: v1.OperationEvent.operation:type_name -> v1.Operation
	7, // 6: v1.OperationBackup.last_status:type_name -> v1.BackupProgressEntry
	8, // 7: v1.OperationIndexSnapshot.snapshot:type_name -> v1.ResticSnapshot
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_v1_operations_proto_init() }
func file_v1_operations_proto_init() {
	if File_v1_operations_proto != nil {
		return
	}
	file_v1_restic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_operations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationList); i {
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
		file_v1_operations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_v1_operations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationEvent); i {
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
		file_v1_operations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationBackup); i {
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
		file_v1_operations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationIndexSnapshot); i {
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
	file_v1_operations_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Operation_OperationBackup)(nil),
		(*Operation_OperationIndexSnapshot)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_operations_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_operations_proto_goTypes,
		DependencyIndexes: file_v1_operations_proto_depIdxs,
		EnumInfos:         file_v1_operations_proto_enumTypes,
		MessageInfos:      file_v1_operations_proto_msgTypes,
	}.Build()
	File_v1_operations_proto = out.File
	file_v1_operations_proto_rawDesc = nil
	file_v1_operations_proto_goTypes = nil
	file_v1_operations_proto_depIdxs = nil
}
