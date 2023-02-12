// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/cluster/pb/cluster.proto

package cluster

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

// Cluster k8s集群相关信息
type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 集群状态
	// @gotags: json:"meta" bson:"meta"
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:"meta"`
	// 基础信息
	// @gotags: json:"spec" bson:"spec"
	Spec *CreateClusterRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:"spec"`
	// 集群状态
	// @gotags: json:"status" bson:"status"
	Status *Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status" bson:"status"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cluster_pb_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cluster_pb_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_apps_cluster_pb_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Cluster) GetSpec() *CreateClusterRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Cluster) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一ID
	// @gotags: json:"id" bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 录入时间
	// @gotags: json:"create_at" bson:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: json:"update_at" bson:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 更新人
	// @gotags: json:"update_by" bson:"update_by"
	UpdateBy string `protobuf:"bytes,4,opt,name=update_by,json=updateBy,proto3" json:"update_by" bson:"update_by"`
	// 集群相关信息
	// @gotags: json:"server_info" bson:"server_info"
	ServerInfo *ServerInfo `protobuf:"bytes,8,opt,name=server_info,json=serverInfo,proto3" json:"server_info" bson:"server_info"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cluster_pb_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cluster_pb_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_apps_cluster_pb_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Meta) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Meta) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Meta) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Meta) GetServerInfo() *ServerInfo {
	if x != nil {
		return x.ServerInfo
	}
	return nil
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// k8s的地址
	// @gotags: json:"server" bson:"server"
	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server" bson:"server"`
	// k8s版本
	// @gotags: json:"version" bson:"version"
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version" bson:"version"`
	// 连接用户
	// @gotags: json:"auth_user" bson:"auth_user"
	AuthUser string `protobuf:"bytes,3,opt,name=auth_user,json=authUser,proto3" json:"auth_user" bson:"auth_user"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cluster_pb_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cluster_pb_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_apps_cluster_pb_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *ServerInfo) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *ServerInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServerInfo) GetAuthUser() string {
	if x != nil {
		return x.AuthUser
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 检查时间
	// @gotags: json:"check_at" bson:"check_at"
	CheckAt int64 `protobuf:"varint,1,opt,name=check_at,json=checkAt,proto3" json:"check_at" bson:"check_at"`
	// API Server是否正常
	// @gotags: json:"is_alive" bson:"is_alive"
	IsAlive bool `protobuf:"varint,2,opt,name=is_alive,json=isAlive,proto3" json:"is_alive" bson:"is_alive"`
	// 异常消息
	// @gotags: json:"message" bson:"message"
	Message string `protobuf:"bytes,10,opt,name=message,proto3" json:"message" bson:"message"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cluster_pb_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cluster_pb_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_apps_cluster_pb_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetCheckAt() int64 {
	if x != nil {
		return x.CheckAt
	}
	return 0
}

func (x *Status) GetIsAlive() bool {
	if x != nil {
		return x.IsAlive
	}
	return false
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 集群所属域
	// @gotags: json:"domain" form:"domain" bson:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" form:"domain" bson:"domain"`
	// 集群所属空间
	// @gotags: json:"namespace" form:"namespace" bson:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" form:"namespace" bson:"namespace"`
	// 创建人
	// @gotags: json:"create_by" form:"create_by" bson:"create_by"
	CreateBy string `protobuf:"bytes,3,opt,name=create_by,json=createBy,proto3" json:"create_by" form:"create_by" bson:"create_by"`
	// 集群提供商
	// @gotags: json:"vendor" bson:"vendor" form:"vendor" validate:"required"
	Vendor string `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor" bson:"vendor" form:"vendor" validate:"required"`
	// 集群所处地域
	// @gotags: json:"region" bson:"region" form:"region" validate:"required"
	Region string `protobuf:"bytes,5,opt,name=region,proto3" json:"region" bson:"region" form:"region" validate:"required"`
	// 名称
	// @gotags: json:"name" bson:"name" form:"name" validate:"required"
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name" bson:"name" form:"name" validate:"required"`
	// 集群客户端访问凭证
	// @gotags: json:"kube_config" bson:"kube_config" form:"kube_config" validate:"required"
	KubeConfig string `protobuf:"bytes,7,opt,name=kube_config,json=kubeConfig,proto3" json:"kube_config" bson:"kube_config" form:"kube_config" validate:"required"`
	// 集群描述
	// @gotags: json:"description" form:"description" bson:"description"
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description" form:"description" bson:"description"`
	// 集群标签, env=prod
	// @gotags: json:"lables" form:"lables" bson:"lables"
	Lables map[string]string `protobuf:"bytes,9,rep,name=lables,proto3" json:"lables" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" form:"lables" bson:"lables"`
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cluster_pb_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cluster_pb_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return file_apps_cluster_pb_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *CreateClusterRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateClusterRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateClusterRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *CreateClusterRequest) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *CreateClusterRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreateClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClusterRequest) GetKubeConfig() string {
	if x != nil {
		return x.KubeConfig
	}
	return ""
}

func (x *CreateClusterRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateClusterRequest) GetLables() map[string]string {
	if x != nil {
		return x.Lables
	}
	return nil
}

// ClusterSet todo
type ClusterSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页时，返回总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 一页的数据
	// @gotags: json:"items"
	Items []*Cluster `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *ClusterSet) Reset() {
	*x = ClusterSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cluster_pb_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterSet) ProtoMessage() {}

func (x *ClusterSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cluster_pb_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterSet.ProtoReflect.Descriptor instead.
func (*ClusterSet) Descriptor() ([]byte, []int) {
	return file_apps_cluster_pb_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *ClusterSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ClusterSet) GetItems() []*Cluster {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_cluster_pb_cluster_proto protoreflect.FileDescriptor

var file_apps_cluster_pb_cluster_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x38, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x45, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5b,
	0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x22, 0x58, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xff, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a, 0x0a, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_cluster_pb_cluster_proto_rawDescOnce sync.Once
	file_apps_cluster_pb_cluster_proto_rawDescData = file_apps_cluster_pb_cluster_proto_rawDesc
)

func file_apps_cluster_pb_cluster_proto_rawDescGZIP() []byte {
	file_apps_cluster_pb_cluster_proto_rawDescOnce.Do(func() {
		file_apps_cluster_pb_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_cluster_pb_cluster_proto_rawDescData)
	})
	return file_apps_cluster_pb_cluster_proto_rawDescData
}

var file_apps_cluster_pb_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apps_cluster_pb_cluster_proto_goTypes = []interface{}{
	(*Cluster)(nil),              // 0: infraboard.mpaas.cluster.Cluster
	(*Meta)(nil),                 // 1: infraboard.mpaas.cluster.Meta
	(*ServerInfo)(nil),           // 2: infraboard.mpaas.cluster.ServerInfo
	(*Status)(nil),               // 3: infraboard.mpaas.cluster.Status
	(*CreateClusterRequest)(nil), // 4: infraboard.mpaas.cluster.CreateClusterRequest
	(*ClusterSet)(nil),           // 5: infraboard.mpaas.cluster.ClusterSet
	nil,                          // 6: infraboard.mpaas.cluster.CreateClusterRequest.LablesEntry
}
var file_apps_cluster_pb_cluster_proto_depIdxs = []int32{
	1, // 0: infraboard.mpaas.cluster.Cluster.meta:type_name -> infraboard.mpaas.cluster.Meta
	4, // 1: infraboard.mpaas.cluster.Cluster.spec:type_name -> infraboard.mpaas.cluster.CreateClusterRequest
	3, // 2: infraboard.mpaas.cluster.Cluster.status:type_name -> infraboard.mpaas.cluster.Status
	2, // 3: infraboard.mpaas.cluster.Meta.server_info:type_name -> infraboard.mpaas.cluster.ServerInfo
	6, // 4: infraboard.mpaas.cluster.CreateClusterRequest.lables:type_name -> infraboard.mpaas.cluster.CreateClusterRequest.LablesEntry
	0, // 5: infraboard.mpaas.cluster.ClusterSet.items:type_name -> infraboard.mpaas.cluster.Cluster
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apps_cluster_pb_cluster_proto_init() }
func file_apps_cluster_pb_cluster_proto_init() {
	if File_apps_cluster_pb_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_cluster_pb_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_apps_cluster_pb_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_apps_cluster_pb_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_apps_cluster_pb_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_apps_cluster_pb_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterRequest); i {
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
		file_apps_cluster_pb_cluster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterSet); i {
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
			RawDescriptor: file_apps_cluster_pb_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_cluster_pb_cluster_proto_goTypes,
		DependencyIndexes: file_apps_cluster_pb_cluster_proto_depIdxs,
		MessageInfos:      file_apps_cluster_pb_cluster_proto_msgTypes,
	}.Build()
	File_apps_cluster_pb_cluster_proto = out.File
	file_apps_cluster_pb_cluster_proto_rawDesc = nil
	file_apps_cluster_pb_cluster_proto_goTypes = nil
	file_apps_cluster_pb_cluster_proto_depIdxs = nil
}
