// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/task/pb/pipeline_task.proto

package task

import (
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
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

type PipelineTaskSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 清单
	// @gotags: bson:"items" json:"items"
	Items []*PipelineTask `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *PipelineTaskSet) Reset() {
	*x = PipelineTaskSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_pipeline_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineTaskSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineTaskSet) ProtoMessage() {}

func (x *PipelineTaskSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_pipeline_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineTaskSet.ProtoReflect.Descriptor instead.
func (*PipelineTaskSet) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_pipeline_task_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineTaskSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PipelineTaskSet) GetItems() []*PipelineTask {
	if x != nil {
		return x.Items
	}
	return nil
}

type PipelineTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 部署Id
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 关联Job
	// @gotags: bson:"pipeline" json:"pipeline"
	Pipeline *pipeline.Pipeline `protobuf:"bytes,3,opt,name=pipeline,proto3" json:"pipeline" bson:"pipeline"`
	// 任务当前状态
	// @gotags: bson:"status" json:"status"
	Status *PipelineTaskStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status" bson:"status"`
}

func (x *PipelineTask) Reset() {
	*x = PipelineTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_pipeline_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineTask) ProtoMessage() {}

func (x *PipelineTask) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_pipeline_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineTask.ProtoReflect.Descriptor instead.
func (*PipelineTask) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_pipeline_task_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PipelineTask) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *PipelineTask) GetPipeline() *pipeline.Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

func (x *PipelineTask) GetStatus() *PipelineTaskStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type PipelineTaskStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务当前状态
	// @gotags: bson:"stage" json:"stage"
	Stage STAGE `protobuf:"varint,1,opt,name=stage,proto3,enum=infraboard.mpaas.task.STAGE" json:"stage" bson:"stage"`
	// 任务开始时间
	// @gotags: bson:"start_at" json:"start_at"
	StartAt int64 `protobuf:"varint,2,opt,name=start_at,json=startAt,proto3" json:"start_at" bson:"start_at"`
	// 任务更新时间, 当pipeline job task运行结束时更新
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 任务结束时间
	// @gotags: bson:"end_at" json:"end_at"
	EndAt int64 `protobuf:"varint,4,opt,name=end_at,json=endAt,proto3" json:"end_at" bson:"end_at"`
	// 状态描述
	// @gotags: bson:"message" json:"message"
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message" bson:"message"`
	// 任务状态详细描述, 用于Debug
	// @gotags: bson:"stage_status" json:"stage_status"
	StageStatus []*StageStatus `protobuf:"bytes,6,rep,name=stage_status,json=stageStatus,proto3" json:"stage_status" bson:"stage_status"`
}

func (x *PipelineTaskStatus) Reset() {
	*x = PipelineTaskStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_pipeline_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineTaskStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineTaskStatus) ProtoMessage() {}

func (x *PipelineTaskStatus) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_pipeline_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineTaskStatus.ProtoReflect.Descriptor instead.
func (*PipelineTaskStatus) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_pipeline_task_proto_rawDescGZIP(), []int{2}
}

func (x *PipelineTaskStatus) GetStage() STAGE {
	if x != nil {
		return x.Stage
	}
	return STAGE_PENDDING
}

func (x *PipelineTaskStatus) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *PipelineTaskStatus) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *PipelineTaskStatus) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *PipelineTaskStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PipelineTaskStatus) GetStageStatus() []*StageStatus {
	if x != nil {
		return x.StageStatus
	}
	return nil
}

type StageStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stage名称
	// @gotags: bson:"name" json:"name" validate:"required"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name" validate:"required"`
	// stage当前状态, 通过动态计算task状态的获得
	// @gotags: bson:"stage" json:"stage"
	Stage STAGE `protobuf:"varint,2,opt,name=stage,proto3,enum=infraboard.mpaas.task.STAGE" json:"stage" bson:"stage"`
	// 任务状态
	// @gotags: bson:"job_tasks" json:"job_tasks"
	JobTasks []*JobTask `protobuf:"bytes,3,rep,name=job_tasks,json=jobTasks,proto3" json:"job_tasks" bson:"job_tasks"`
}

func (x *StageStatus) Reset() {
	*x = StageStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_pipeline_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StageStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageStatus) ProtoMessage() {}

func (x *StageStatus) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_pipeline_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageStatus.ProtoReflect.Descriptor instead.
func (*StageStatus) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_pipeline_task_proto_rawDescGZIP(), []int{3}
}

func (x *StageStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StageStatus) GetStage() STAGE {
	if x != nil {
		return x.Stage
	}
	return STAGE_PENDDING
}

func (x *StageStatus) GetJobTasks() []*JobTask {
	if x != nil {
		return x.JobTasks
	}
	return nil
}

var File_apps_task_pb_pipeline_task_proto protoreflect.FileDescriptor

var file_apps_task_pb_pipeline_task_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0f, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x39, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x0c,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xf8, 0x01,
	0x0a, 0x12, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x54, 0x41, 0x47,
	0x45, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x45, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x53, 0x54, 0x41, 0x47, 0x45, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x3b, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_task_pb_pipeline_task_proto_rawDescOnce sync.Once
	file_apps_task_pb_pipeline_task_proto_rawDescData = file_apps_task_pb_pipeline_task_proto_rawDesc
)

func file_apps_task_pb_pipeline_task_proto_rawDescGZIP() []byte {
	file_apps_task_pb_pipeline_task_proto_rawDescOnce.Do(func() {
		file_apps_task_pb_pipeline_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_task_pb_pipeline_task_proto_rawDescData)
	})
	return file_apps_task_pb_pipeline_task_proto_rawDescData
}

var file_apps_task_pb_pipeline_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_task_pb_pipeline_task_proto_goTypes = []interface{}{
	(*PipelineTaskSet)(nil),    // 0: infraboard.mpaas.task.PipelineTaskSet
	(*PipelineTask)(nil),       // 1: infraboard.mpaas.task.PipelineTask
	(*PipelineTaskStatus)(nil), // 2: infraboard.mpaas.task.PipelineTaskStatus
	(*StageStatus)(nil),        // 3: infraboard.mpaas.task.StageStatus
	(*pipeline.Pipeline)(nil),  // 4: infraboard.mpaas.pipeline.Pipeline
	(STAGE)(0),                 // 5: infraboard.mpaas.task.STAGE
	(*JobTask)(nil),            // 6: infraboard.mpaas.task.JobTask
}
var file_apps_task_pb_pipeline_task_proto_depIdxs = []int32{
	1, // 0: infraboard.mpaas.task.PipelineTaskSet.items:type_name -> infraboard.mpaas.task.PipelineTask
	4, // 1: infraboard.mpaas.task.PipelineTask.pipeline:type_name -> infraboard.mpaas.pipeline.Pipeline
	2, // 2: infraboard.mpaas.task.PipelineTask.status:type_name -> infraboard.mpaas.task.PipelineTaskStatus
	5, // 3: infraboard.mpaas.task.PipelineTaskStatus.stage:type_name -> infraboard.mpaas.task.STAGE
	3, // 4: infraboard.mpaas.task.PipelineTaskStatus.stage_status:type_name -> infraboard.mpaas.task.StageStatus
	5, // 5: infraboard.mpaas.task.StageStatus.stage:type_name -> infraboard.mpaas.task.STAGE
	6, // 6: infraboard.mpaas.task.StageStatus.job_tasks:type_name -> infraboard.mpaas.task.JobTask
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_apps_task_pb_pipeline_task_proto_init() }
func file_apps_task_pb_pipeline_task_proto_init() {
	if File_apps_task_pb_pipeline_task_proto != nil {
		return
	}
	file_apps_task_pb_job_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_task_pb_pipeline_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineTaskSet); i {
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
		file_apps_task_pb_pipeline_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineTask); i {
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
		file_apps_task_pb_pipeline_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineTaskStatus); i {
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
		file_apps_task_pb_pipeline_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StageStatus); i {
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
			RawDescriptor: file_apps_task_pb_pipeline_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_task_pb_pipeline_task_proto_goTypes,
		DependencyIndexes: file_apps_task_pb_pipeline_task_proto_depIdxs,
		MessageInfos:      file_apps_task_pb_pipeline_task_proto_msgTypes,
	}.Build()
	File_apps_task_pb_pipeline_task_proto = out.File
	file_apps_task_pb_pipeline_task_proto_rawDesc = nil
	file_apps_task_pb_pipeline_task_proto_goTypes = nil
	file_apps_task_pb_pipeline_task_proto_depIdxs = nil
}
