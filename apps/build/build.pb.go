// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/build/pb/build.proto

package build

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

// 构建目标的类型
type TARGET_TYPE int32

const (
	// 容器镜像
	TARGET_TYPE_IMAGE TARGET_TYPE = 0
	// 包
	TARGET_TYPE_PKG TARGET_TYPE = 1
)

// Enum value maps for TARGET_TYPE.
var (
	TARGET_TYPE_name = map[int32]string{
		0: "IMAGE",
		1: "PKG",
	}
	TARGET_TYPE_value = map[string]int32{
		"IMAGE": 0,
		"PKG":   1,
	}
)

func (x TARGET_TYPE) Enum() *TARGET_TYPE {
	p := new(TARGET_TYPE)
	*p = x
	return p
}

func (x TARGET_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TARGET_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_build_pb_build_proto_enumTypes[0].Descriptor()
}

func (TARGET_TYPE) Type() protoreflect.EnumType {
	return &file_apps_build_pb_build_proto_enumTypes[0]
}

func (x TARGET_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TARGET_TYPE.Descriptor instead.
func (TARGET_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_build_pb_build_proto_rawDescGZIP(), []int{0}
}

type BuildConfigSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 清单
	Items []*BuildConfig `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *BuildConfigSet) Reset() {
	*x = BuildConfigSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_build_pb_build_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildConfigSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildConfigSet) ProtoMessage() {}

func (x *BuildConfigSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_build_pb_build_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildConfigSet.ProtoReflect.Descriptor instead.
func (*BuildConfigSet) Descriptor() ([]byte, []int) {
	return file_apps_build_pb_build_proto_rawDescGZIP(), []int{0}
}

func (x *BuildConfigSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BuildConfigSet) GetItems() []*BuildConfig {
	if x != nil {
		return x.Items
	}
	return nil
}

type BuildConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 构建配置Id
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 所属Domain
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 所属空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,5,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 创建信息
	// @gotags: bson:"spec" json:"spec"
	Spec *CreateBuildConfigRequest `protobuf:"bytes,6,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *BuildConfig) Reset() {
	*x = BuildConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_build_pb_build_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildConfig) ProtoMessage() {}

func (x *BuildConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apps_build_pb_build_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildConfig.ProtoReflect.Descriptor instead.
func (*BuildConfig) Descriptor() ([]byte, []int) {
	return file_apps_build_pb_build_proto_rawDescGZIP(), []int{1}
}

func (x *BuildConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BuildConfig) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *BuildConfig) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *BuildConfig) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *BuildConfig) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *BuildConfig) GetSpec() *CreateBuildConfigRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type CreateBuildConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 那个服务的构建
	// @gotags: bson:"service_id" json:"service_id"
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id" bson:"service_id"`
	// 触发条件
	// @gotags: bson:"condition" json:"condition"
	Condition *Trigger `protobuf:"bytes,2,opt,name=condition,proto3" json:"condition" bson:"condition"`
	// 构建的名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" bson:"name"`
	// 构建描述信息
	// @gotags: bson:"describe" json:"describe"
	Describe string `protobuf:"bytes,4,opt,name=describe,proto3" json:"describe" bson:"describe"`
	// 构建目标类型
	// @gotags: bson:"target_type" json:"target_type"
	TargetType TARGET_TYPE `protobuf:"varint,5,opt,name=target_type,json=targetType,proto3,enum=infraboard.mpaas.build.TARGET_TYPE" json:"target_type" bson:"target_type"`
	// 镜像构建配置
	// @gotags: bson:"image_build" json:"image_build"
	ImageBuild *ImageBuildConfig `protobuf:"bytes,6,opt,name=image_build,json=imageBuild,proto3" json:"image_build" bson:"image_build"`
	// 包构建配置
	// @gotags: bson:"pkg_build" json:"pkg_build"
	PkgBuild *PkgBuildConfig `protobuf:"bytes,7,opt,name=pkg_build,json=pkgBuild,proto3" json:"pkg_build" bson:"pkg_build"`
	// 构建配置标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *CreateBuildConfigRequest) Reset() {
	*x = CreateBuildConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_build_pb_build_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBuildConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBuildConfigRequest) ProtoMessage() {}

func (x *CreateBuildConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_build_pb_build_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBuildConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateBuildConfigRequest) Descriptor() ([]byte, []int) {
	return file_apps_build_pb_build_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBuildConfigRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *CreateBuildConfigRequest) GetCondition() *Trigger {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *CreateBuildConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBuildConfigRequest) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *CreateBuildConfigRequest) GetTargetType() TARGET_TYPE {
	if x != nil {
		return x.TargetType
	}
	return TARGET_TYPE_IMAGE
}

func (x *CreateBuildConfigRequest) GetImageBuild() *ImageBuildConfig {
	if x != nil {
		return x.ImageBuild
	}
	return nil
}

func (x *CreateBuildConfigRequest) GetPkgBuild() *PkgBuildConfig {
	if x != nil {
		return x.PkgBuild
	}
	return nil
}

func (x *CreateBuildConfigRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// 触发执行的条件
type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件名称,那些事件可以触发
	// @gotags: bson:"events" json:"events"
	Events []string `protobuf:"bytes,1,rep,name=events,proto3" json:"events" bson:"events"`
	// 分支名称, 那些分支可以触发, 支持正则
	// @gotags: bson:"branches" json:"branches"
	Branches []string `protobuf:"bytes,2,rep,name=branches,proto3" json:"branches" bson:"branches"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_build_pb_build_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_apps_build_pb_build_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_apps_build_pb_build_proto_rawDescGZIP(), []int{3}
}

func (x *Trigger) GetEvents() []string {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *Trigger) GetBranches() []string {
	if x != nil {
		return x.Branches
	}
	return nil
}

// 容器镜像构建配置
type ImageBuildConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dockerfile所在路径, 默认当前目录下 Dockerfile
	// @gotags: bson:"docker_file" json:"docker_file"
	DockerFile string `protobuf:"bytes,1,opt,name=docker_file,json=dockerFile,proto3" json:"docker_file" bson:"docker_file"`
	// 构建时注入的环境变量
	// @gotags: bson:"build_env_vars" json:"build_env_vars"
	BuildEnvVars map[string]string `protobuf:"bytes,2,rep,name=build_env_vars,json=buildEnvVars,proto3" json:"build_env_vars" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"build_env_vars"`
	// 其他信息
	// @gotags: bson:"extra" json:"extra"
	Extra map[string]string `protobuf:"bytes,9,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extra"`
}

func (x *ImageBuildConfig) Reset() {
	*x = ImageBuildConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_build_pb_build_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageBuildConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageBuildConfig) ProtoMessage() {}

func (x *ImageBuildConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apps_build_pb_build_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageBuildConfig.ProtoReflect.Descriptor instead.
func (*ImageBuildConfig) Descriptor() ([]byte, []int) {
	return file_apps_build_pb_build_proto_rawDescGZIP(), []int{4}
}

func (x *ImageBuildConfig) GetDockerFile() string {
	if x != nil {
		return x.DockerFile
	}
	return ""
}

func (x *ImageBuildConfig) GetBuildEnvVars() map[string]string {
	if x != nil {
		return x.BuildEnvVars
	}
	return nil
}

func (x *ImageBuildConfig) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

// 包构建配置
type PkgBuildConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 其他信息
	// @gotags: bson:"extra" json:"extra"
	Extra map[string]string `protobuf:"bytes,9,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extra"`
}

func (x *PkgBuildConfig) Reset() {
	*x = PkgBuildConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_build_pb_build_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PkgBuildConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PkgBuildConfig) ProtoMessage() {}

func (x *PkgBuildConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apps_build_pb_build_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PkgBuildConfig.ProtoReflect.Descriptor instead.
func (*PkgBuildConfig) Descriptor() ([]byte, []int) {
	return file_apps_build_pb_build_proto_rawDescGZIP(), []int{5}
}

func (x *PkgBuildConfig) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_apps_build_pb_build_proto protoreflect.FileDescriptor

var file_apps_build_pb_build_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x62, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x22, 0x61, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xd3, 0x01, 0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x8f, 0x04, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a,
	0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x43, 0x0a, 0x09, 0x70, 0x6b, 0x67, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x50, 0x6b, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x08, 0x70, 0x6b, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x54, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d,
	0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x22, 0xdb, 0x02,
	0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x60, 0x0a, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x65, 0x6e, 0x76,
	0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x76, 0x56, 0x61,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e,
	0x76, 0x56, 0x61, 0x72, 0x73, 0x12, 0x49, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x1a, 0x3f, 0x0a, 0x11, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x93, 0x01, 0x0a, 0x0e,
	0x50, 0x6b, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47,
	0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x50, 0x6b, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x2a, 0x21, 0x0a, 0x0b, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50,
	0x4b, 0x47, 0x10, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_build_pb_build_proto_rawDescOnce sync.Once
	file_apps_build_pb_build_proto_rawDescData = file_apps_build_pb_build_proto_rawDesc
)

func file_apps_build_pb_build_proto_rawDescGZIP() []byte {
	file_apps_build_pb_build_proto_rawDescOnce.Do(func() {
		file_apps_build_pb_build_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_build_pb_build_proto_rawDescData)
	})
	return file_apps_build_pb_build_proto_rawDescData
}

var file_apps_build_pb_build_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_build_pb_build_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_apps_build_pb_build_proto_goTypes = []interface{}{
	(TARGET_TYPE)(0),                 // 0: infraboard.mpaas.build.TARGET_TYPE
	(*BuildConfigSet)(nil),           // 1: infraboard.mpaas.build.BuildConfigSet
	(*BuildConfig)(nil),              // 2: infraboard.mpaas.build.BuildConfig
	(*CreateBuildConfigRequest)(nil), // 3: infraboard.mpaas.build.CreateBuildConfigRequest
	(*Trigger)(nil),                  // 4: infraboard.mpaas.build.Trigger
	(*ImageBuildConfig)(nil),         // 5: infraboard.mpaas.build.ImageBuildConfig
	(*PkgBuildConfig)(nil),           // 6: infraboard.mpaas.build.PkgBuildConfig
	nil,                              // 7: infraboard.mpaas.build.CreateBuildConfigRequest.LabelsEntry
	nil,                              // 8: infraboard.mpaas.build.ImageBuildConfig.BuildEnvVarsEntry
	nil,                              // 9: infraboard.mpaas.build.ImageBuildConfig.ExtraEntry
	nil,                              // 10: infraboard.mpaas.build.PkgBuildConfig.ExtraEntry
}
var file_apps_build_pb_build_proto_depIdxs = []int32{
	2,  // 0: infraboard.mpaas.build.BuildConfigSet.items:type_name -> infraboard.mpaas.build.BuildConfig
	3,  // 1: infraboard.mpaas.build.BuildConfig.spec:type_name -> infraboard.mpaas.build.CreateBuildConfigRequest
	4,  // 2: infraboard.mpaas.build.CreateBuildConfigRequest.condition:type_name -> infraboard.mpaas.build.Trigger
	0,  // 3: infraboard.mpaas.build.CreateBuildConfigRequest.target_type:type_name -> infraboard.mpaas.build.TARGET_TYPE
	5,  // 4: infraboard.mpaas.build.CreateBuildConfigRequest.image_build:type_name -> infraboard.mpaas.build.ImageBuildConfig
	6,  // 5: infraboard.mpaas.build.CreateBuildConfigRequest.pkg_build:type_name -> infraboard.mpaas.build.PkgBuildConfig
	7,  // 6: infraboard.mpaas.build.CreateBuildConfigRequest.labels:type_name -> infraboard.mpaas.build.CreateBuildConfigRequest.LabelsEntry
	8,  // 7: infraboard.mpaas.build.ImageBuildConfig.build_env_vars:type_name -> infraboard.mpaas.build.ImageBuildConfig.BuildEnvVarsEntry
	9,  // 8: infraboard.mpaas.build.ImageBuildConfig.extra:type_name -> infraboard.mpaas.build.ImageBuildConfig.ExtraEntry
	10, // 9: infraboard.mpaas.build.PkgBuildConfig.extra:type_name -> infraboard.mpaas.build.PkgBuildConfig.ExtraEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_apps_build_pb_build_proto_init() }
func file_apps_build_pb_build_proto_init() {
	if File_apps_build_pb_build_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_build_pb_build_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildConfigSet); i {
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
		file_apps_build_pb_build_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildConfig); i {
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
		file_apps_build_pb_build_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBuildConfigRequest); i {
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
		file_apps_build_pb_build_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger); i {
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
		file_apps_build_pb_build_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageBuildConfig); i {
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
		file_apps_build_pb_build_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PkgBuildConfig); i {
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
			RawDescriptor: file_apps_build_pb_build_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_build_pb_build_proto_goTypes,
		DependencyIndexes: file_apps_build_pb_build_proto_depIdxs,
		EnumInfos:         file_apps_build_pb_build_proto_enumTypes,
		MessageInfos:      file_apps_build_pb_build_proto_msgTypes,
	}.Build()
	File_apps_build_pb_build_proto = out.File
	file_apps_build_pb_build_proto_rawDesc = nil
	file_apps_build_pb_build_proto_goTypes = nil
	file_apps_build_pb_build_proto_depIdxs = nil
}
