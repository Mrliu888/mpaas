// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/pipeline/pb/pipeline.proto

package pipeline

import (
	job "github.com/infraboard/mpaas/apps/job"
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

type PipelineSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 清单
	// @gotags: bson:"items" json:"items"
	Items []*Pipeline `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *PipelineSet) Reset() {
	*x = PipelineSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineSet) ProtoMessage() {}

func (x *PipelineSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineSet.ProtoReflect.Descriptor instead.
func (*PipelineSet) Descriptor() ([]byte, []int) {
	return file_apps_pipeline_pb_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PipelineSet) GetItems() []*Pipeline {
	if x != nil {
		return x.Items
	}
	return nil
}

// 流水线
type Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 唯一人
	// @gotags: bson:"update_by" json:"update_by"
	UpdateBy string `protobuf:"bytes,4,opt,name=update_by,json=updateBy,proto3" json:"update_by" bson:"update_by"`
	// 流水线定义
	// @gotags: bson:"spec" json:"spec"
	Spec *CreatePipelineRequest `protobuf:"bytes,10,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *Pipeline) Reset() {
	*x = Pipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipeline) ProtoMessage() {}

func (x *Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipeline.ProtoReflect.Descriptor instead.
func (*Pipeline) Descriptor() ([]byte, []int) {
	return file_apps_pipeline_pb_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *Pipeline) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pipeline) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Pipeline) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Pipeline) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Pipeline) GetSpec() *CreatePipelineRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type CreatePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 所属空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 创建人
	// @gotags: bson:"create_by" json:"create_by"
	CreateBy string `protobuf:"bytes,3,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 名称
	// @gotags: bson:"name" json:"name" validate:"required"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name" bson:"name" validate:"required"`
	// 描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description" bson:"description"`
	// 是否是模版, 用于快速继承模版参数进行修改, 模版不能用于执行
	// @gotags: bson:"is_template" json:"is_template"
	IsTemplate bool `protobuf:"varint,6,opt,name=is_template,json=isTemplate,proto3" json:"is_template" bson:"is_template"`
	// 具体编排阶段
	// @gotags: bson:"stages" json:"stages"
	Stages []*Stage `protobuf:"bytes,12,rep,name=stages,proto3" json:"stages" bson:"stages"`
	// 标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *CreatePipelineRequest) Reset() {
	*x = CreatePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePipelineRequest) ProtoMessage() {}

func (x *CreatePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePipelineRequest.ProtoReflect.Descriptor instead.
func (*CreatePipelineRequest) Descriptor() ([]byte, []int) {
	return file_apps_pipeline_pb_pipeline_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePipelineRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreatePipelineRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreatePipelineRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *CreatePipelineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePipelineRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePipelineRequest) GetIsTemplate() bool {
	if x != nil {
		return x.IsTemplate
	}
	return false
}

func (x *CreatePipelineRequest) GetStages() []*Stage {
	if x != nil {
		return x.Stages
	}
	return nil
}

func (x *CreatePipelineRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Stage todo
type Stage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 名称
	// @gotags: bson:"name" json:"name" validate:"required"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name" validate:"required"`
	// 是否并行, 如果并行执行 该Task里面的任务会同时执行, 否则串行
	// @gotags: bson:"is_parallel" json:"is_parallel"
	IsParallel bool `protobuf:"varint,2,opt,name=is_parallel,json=isParallel,proto3" json:"is_parallel" bson:"is_parallel"`
	// 运行时 全局参数, 会传递给该stage的每个Task
	// @gotags: bson:"with" json:"with"
	With []*job.RunParam `protobuf:"bytes,10,rep,name=with,proto3" json:"with" bson:"with"`
	// 需要执行的job
	// @gotags: bson:"jobs" json:"jobs"
	Jobs []*RunJobRequest `protobuf:"bytes,11,rep,name=jobs,proto3" json:"jobs" bson:"jobs"`
}

func (x *Stage) Reset() {
	*x = Stage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stage) ProtoMessage() {}

func (x *Stage) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stage.ProtoReflect.Descriptor instead.
func (*Stage) Descriptor() ([]byte, []int) {
	return file_apps_pipeline_pb_pipeline_proto_rawDescGZIP(), []int{3}
}

func (x *Stage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Stage) GetIsParallel() bool {
	if x != nil {
		return x.IsParallel
	}
	return false
}

func (x *Stage) GetWith() []*job.RunParam {
	if x != nil {
		return x.With
	}
	return nil
}

func (x *Stage) GetJobs() []*RunJobRequest {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type RunJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// task执行的域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// task执行的空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// task属于那个Pipeline 任务可, 以为空
	// @gotags: bson:"pipeline_task" json:"pipeline_task"
	PipelineTask string `protobuf:"bytes,3,opt,name=pipeline_task,json=pipelineTask,proto3" json:"pipeline_task" bson:"pipeline_task"`
	// job名称: name
	// @gotags: bson:"job" json:"job" validate:"required"
	Job string `protobuf:"bytes,4,opt,name=job,proto3" json:"job" bson:"job" validate:"required"`
	// job运行时参数
	// @gotags: bson:"params" json:"params"
	Params *job.VersionedRunParam `protobuf:"bytes,5,opt,name=params,proto3" json:"params" bson:"params"`
	// 忽略失败, 当Pipeline运行时, 不会因为执行失败而 中断Pipeline执行
	// @gotags: bson:"ignore_failed" json:"ignore_failed"
	IgnoreFailed bool `protobuf:"varint,6,opt,name=ignore_failed,json=ignoreFailed,proto3" json:"ignore_failed" bson:"ignore_failed"`
	// WebHook配置, 及时把job执行状态推送给外部, 比如各种机器人
	// @gotags: bson:"webhooks" json:"webhooks"
	Webhooks []*WebHook `protobuf:"bytes,7,rep,name=webhooks,proto3" json:"webhooks" bson:"webhooks"`
	// 任务标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *RunJobRequest) Reset() {
	*x = RunJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunJobRequest) ProtoMessage() {}

func (x *RunJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunJobRequest.ProtoReflect.Descriptor instead.
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return file_apps_pipeline_pb_pipeline_proto_rawDescGZIP(), []int{4}
}

func (x *RunJobRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RunJobRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RunJobRequest) GetPipelineTask() string {
	if x != nil {
		return x.PipelineTask
	}
	return ""
}

func (x *RunJobRequest) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *RunJobRequest) GetParams() *job.VersionedRunParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *RunJobRequest) GetIgnoreFailed() bool {
	if x != nil {
		return x.IgnoreFailed
	}
	return false
}

func (x *RunJobRequest) GetWebhooks() []*WebHook {
	if x != nil {
		return x.Webhooks
	}
	return nil
}

func (x *RunJobRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type WebHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// POST URL
	// @gotags: bson:"url" json:"url" validate:"required,url"
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url" bson:"url" validate:"required,url"`
	// 需要自定义添加的头, 用于身份认证
	// @gotags: bson:"header" json:"header"
	Header map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"header"`
	// 那些状态下触发
	// @gotags: bson:"events" json:"events"
	Events []string `protobuf:"bytes,3,rep,name=events,proto3" json:"events" bson:"events"`
	// 简单的描述信息
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description" bson:"description"`
	// 推送结果
	// @gotags: bson:"status" json:"status"
	Status *WebHookStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status" bson:"status"`
}

func (x *WebHook) Reset() {
	*x = WebHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebHook) ProtoMessage() {}

func (x *WebHook) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebHook.ProtoReflect.Descriptor instead.
func (*WebHook) Descriptor() ([]byte, []int) {
	return file_apps_pipeline_pb_pipeline_proto_rawDescGZIP(), []int{5}
}

func (x *WebHook) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WebHook) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *WebHook) GetEvents() []string {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *WebHook) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WebHook) GetStatus() *WebHookStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type WebHookStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 开始时间
	// @gotags: bson:"start_at" json:"start_at"
	StartAt int64 `protobuf:"varint,1,opt,name=start_at,json=startAt,proto3" json:"start_at" bson:"start_at"`
	// 耗时多久，单位毫秒
	// @gotags: bson:"cost" json:"cost"
	Cost int64 `protobuf:"varint,2,opt,name=cost,proto3" json:"cost" bson:"cost"`
	// 是否推送成功
	// @gotags: bson:"success" json:"success"
	Success bool `protobuf:"varint,3,opt,name=success,proto3" json:"success" bson:"success"`
	// 异常时的错误信息
	// @gotags: bson:"message" json:"message"
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message" bson:"message"`
}

func (x *WebHookStatus) Reset() {
	*x = WebHookStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebHookStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebHookStatus) ProtoMessage() {}

func (x *WebHookStatus) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pipeline_pb_pipeline_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebHookStatus.ProtoReflect.Descriptor instead.
func (*WebHookStatus) Descriptor() ([]byte, []int) {
	return file_apps_pipeline_pb_pipeline_proto_rawDescGZIP(), []int{6}
}

func (x *WebHookStatus) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *WebHookStatus) GetCost() int64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *WebHookStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *WebHookStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_apps_pipeline_pb_pipeline_proto protoreflect.FileDescriptor

var file_apps_pipeline_pb_pipeline_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x15, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0b, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0xb7, 0x01, 0x0a, 0x08, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x8c, 0x03,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xae, 0x01, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x12, 0x32, 0x0a, 0x04, 0x77,
	0x69, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x52, 0x75, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x04, 0x77, 0x69, 0x74, 0x68, 0x12,
	0x3c, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0xab, 0x03,
	0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f,
	0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x3f, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x52, 0x75, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x08, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x73, 0x12, 0x4c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x52,
	0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9a, 0x02, 0x0a, 0x07,
	0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x46, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x72, 0x0a, 0x0d, 0x57, 0x65, 0x62, 0x48,
	0x6f, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_pipeline_pb_pipeline_proto_rawDescOnce sync.Once
	file_apps_pipeline_pb_pipeline_proto_rawDescData = file_apps_pipeline_pb_pipeline_proto_rawDesc
)

func file_apps_pipeline_pb_pipeline_proto_rawDescGZIP() []byte {
	file_apps_pipeline_pb_pipeline_proto_rawDescOnce.Do(func() {
		file_apps_pipeline_pb_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_pipeline_pb_pipeline_proto_rawDescData)
	})
	return file_apps_pipeline_pb_pipeline_proto_rawDescData
}

var file_apps_pipeline_pb_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_apps_pipeline_pb_pipeline_proto_goTypes = []interface{}{
	(*PipelineSet)(nil),           // 0: infraboard.mpaas.pipeline.PipelineSet
	(*Pipeline)(nil),              // 1: infraboard.mpaas.pipeline.Pipeline
	(*CreatePipelineRequest)(nil), // 2: infraboard.mpaas.pipeline.CreatePipelineRequest
	(*Stage)(nil),                 // 3: infraboard.mpaas.pipeline.Stage
	(*RunJobRequest)(nil),         // 4: infraboard.mpaas.pipeline.RunJobRequest
	(*WebHook)(nil),               // 5: infraboard.mpaas.pipeline.WebHook
	(*WebHookStatus)(nil),         // 6: infraboard.mpaas.pipeline.WebHookStatus
	nil,                           // 7: infraboard.mpaas.pipeline.CreatePipelineRequest.LabelsEntry
	nil,                           // 8: infraboard.mpaas.pipeline.RunJobRequest.LabelsEntry
	nil,                           // 9: infraboard.mpaas.pipeline.WebHook.HeaderEntry
	(*job.RunParam)(nil),          // 10: infraboard.mpaas.job.RunParam
	(*job.VersionedRunParam)(nil), // 11: infraboard.mpaas.job.VersionedRunParam
}
var file_apps_pipeline_pb_pipeline_proto_depIdxs = []int32{
	1,  // 0: infraboard.mpaas.pipeline.PipelineSet.items:type_name -> infraboard.mpaas.pipeline.Pipeline
	2,  // 1: infraboard.mpaas.pipeline.Pipeline.spec:type_name -> infraboard.mpaas.pipeline.CreatePipelineRequest
	3,  // 2: infraboard.mpaas.pipeline.CreatePipelineRequest.stages:type_name -> infraboard.mpaas.pipeline.Stage
	7,  // 3: infraboard.mpaas.pipeline.CreatePipelineRequest.labels:type_name -> infraboard.mpaas.pipeline.CreatePipelineRequest.LabelsEntry
	10, // 4: infraboard.mpaas.pipeline.Stage.with:type_name -> infraboard.mpaas.job.RunParam
	4,  // 5: infraboard.mpaas.pipeline.Stage.jobs:type_name -> infraboard.mpaas.pipeline.RunJobRequest
	11, // 6: infraboard.mpaas.pipeline.RunJobRequest.params:type_name -> infraboard.mpaas.job.VersionedRunParam
	5,  // 7: infraboard.mpaas.pipeline.RunJobRequest.webhooks:type_name -> infraboard.mpaas.pipeline.WebHook
	8,  // 8: infraboard.mpaas.pipeline.RunJobRequest.labels:type_name -> infraboard.mpaas.pipeline.RunJobRequest.LabelsEntry
	9,  // 9: infraboard.mpaas.pipeline.WebHook.header:type_name -> infraboard.mpaas.pipeline.WebHook.HeaderEntry
	6,  // 10: infraboard.mpaas.pipeline.WebHook.status:type_name -> infraboard.mpaas.pipeline.WebHookStatus
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_apps_pipeline_pb_pipeline_proto_init() }
func file_apps_pipeline_pb_pipeline_proto_init() {
	if File_apps_pipeline_pb_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_pipeline_pb_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineSet); i {
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
		file_apps_pipeline_pb_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipeline); i {
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
		file_apps_pipeline_pb_pipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePipelineRequest); i {
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
		file_apps_pipeline_pb_pipeline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stage); i {
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
		file_apps_pipeline_pb_pipeline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunJobRequest); i {
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
		file_apps_pipeline_pb_pipeline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebHook); i {
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
		file_apps_pipeline_pb_pipeline_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebHookStatus); i {
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
			RawDescriptor: file_apps_pipeline_pb_pipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_pipeline_pb_pipeline_proto_goTypes,
		DependencyIndexes: file_apps_pipeline_pb_pipeline_proto_depIdxs,
		MessageInfos:      file_apps_pipeline_pb_pipeline_proto_msgTypes,
	}.Build()
	File_apps_pipeline_pb_pipeline_proto = out.File
	file_apps_pipeline_pb_pipeline_proto_rawDesc = nil
	file_apps_pipeline_pb_pipeline_proto_goTypes = nil
	file_apps_pipeline_pb_pipeline_proto_depIdxs = nil
}
