// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/admin/v2alpha/clusters.proto

package envoy_admin_v2alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Clusters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterStatuses []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
}

func (x *Clusters) Reset() {
	*x = Clusters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v2alpha_clusters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clusters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clusters) ProtoMessage() {}

func (x *Clusters) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v2alpha_clusters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clusters.ProtoReflect.Descriptor instead.
func (*Clusters) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v2alpha_clusters_proto_rawDescGZIP(), []int{0}
}

func (x *Clusters) GetClusterStatuses() []*ClusterStatus {
	if x != nil {
		return x.ClusterStatuses
	}
	return nil
}

type ClusterStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                                    string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AddedViaApi                             bool           `protobuf:"varint,2,opt,name=added_via_api,json=addedViaApi,proto3" json:"added_via_api,omitempty"`
	SuccessRateEjectionThreshold            *_type.Percent `protobuf:"bytes,3,opt,name=success_rate_ejection_threshold,json=successRateEjectionThreshold,proto3" json:"success_rate_ejection_threshold,omitempty"`
	HostStatuses                            []*HostStatus  `protobuf:"bytes,4,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	LocalOriginSuccessRateEjectionThreshold *_type.Percent `protobuf:"bytes,5,opt,name=local_origin_success_rate_ejection_threshold,json=localOriginSuccessRateEjectionThreshold,proto3" json:"local_origin_success_rate_ejection_threshold,omitempty"`
}

func (x *ClusterStatus) Reset() {
	*x = ClusterStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v2alpha_clusters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStatus) ProtoMessage() {}

func (x *ClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v2alpha_clusters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStatus.ProtoReflect.Descriptor instead.
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v2alpha_clusters_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterStatus) GetAddedViaApi() bool {
	if x != nil {
		return x.AddedViaApi
	}
	return false
}

func (x *ClusterStatus) GetSuccessRateEjectionThreshold() *_type.Percent {
	if x != nil {
		return x.SuccessRateEjectionThreshold
	}
	return nil
}

func (x *ClusterStatus) GetHostStatuses() []*HostStatus {
	if x != nil {
		return x.HostStatuses
	}
	return nil
}

func (x *ClusterStatus) GetLocalOriginSuccessRateEjectionThreshold() *_type.Percent {
	if x != nil {
		return x.LocalOriginSuccessRateEjectionThreshold
	}
	return nil
}

type HostStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address                *core.Address     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Stats                  []*SimpleMetric   `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	HealthStatus           *HostHealthStatus `protobuf:"bytes,3,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	SuccessRate            *_type.Percent    `protobuf:"bytes,4,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	Weight                 uint32            `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Hostname               string            `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Priority               uint32            `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	LocalOriginSuccessRate *_type.Percent    `protobuf:"bytes,8,opt,name=local_origin_success_rate,json=localOriginSuccessRate,proto3" json:"local_origin_success_rate,omitempty"`
	Locality               *core.Locality    `protobuf:"bytes,9,opt,name=locality,proto3" json:"locality,omitempty"`
}

func (x *HostStatus) Reset() {
	*x = HostStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v2alpha_clusters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostStatus) ProtoMessage() {}

func (x *HostStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v2alpha_clusters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostStatus.ProtoReflect.Descriptor instead.
func (*HostStatus) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v2alpha_clusters_proto_rawDescGZIP(), []int{2}
}

func (x *HostStatus) GetAddress() *core.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *HostStatus) GetStats() []*SimpleMetric {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *HostStatus) GetHealthStatus() *HostHealthStatus {
	if x != nil {
		return x.HealthStatus
	}
	return nil
}

func (x *HostStatus) GetSuccessRate() *_type.Percent {
	if x != nil {
		return x.SuccessRate
	}
	return nil
}

func (x *HostStatus) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *HostStatus) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *HostStatus) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *HostStatus) GetLocalOriginSuccessRate() *_type.Percent {
	if x != nil {
		return x.LocalOriginSuccessRate
	}
	return nil
}

func (x *HostStatus) GetLocality() *core.Locality {
	if x != nil {
		return x.Locality
	}
	return nil
}

type HostHealthStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailedActiveHealthCheck   bool              `protobuf:"varint,1,opt,name=failed_active_health_check,json=failedActiveHealthCheck,proto3" json:"failed_active_health_check,omitempty"`
	FailedOutlierCheck        bool              `protobuf:"varint,2,opt,name=failed_outlier_check,json=failedOutlierCheck,proto3" json:"failed_outlier_check,omitempty"`
	FailedActiveDegradedCheck bool              `protobuf:"varint,4,opt,name=failed_active_degraded_check,json=failedActiveDegradedCheck,proto3" json:"failed_active_degraded_check,omitempty"`
	PendingDynamicRemoval     bool              `protobuf:"varint,5,opt,name=pending_dynamic_removal,json=pendingDynamicRemoval,proto3" json:"pending_dynamic_removal,omitempty"`
	PendingActiveHc           bool              `protobuf:"varint,6,opt,name=pending_active_hc,json=pendingActiveHc,proto3" json:"pending_active_hc,omitempty"`
	EdsHealthStatus           core.HealthStatus `protobuf:"varint,3,opt,name=eds_health_status,json=edsHealthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"eds_health_status,omitempty"`
}

func (x *HostHealthStatus) Reset() {
	*x = HostHealthStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v2alpha_clusters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostHealthStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostHealthStatus) ProtoMessage() {}

func (x *HostHealthStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v2alpha_clusters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostHealthStatus.ProtoReflect.Descriptor instead.
func (*HostHealthStatus) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v2alpha_clusters_proto_rawDescGZIP(), []int{3}
}

func (x *HostHealthStatus) GetFailedActiveHealthCheck() bool {
	if x != nil {
		return x.FailedActiveHealthCheck
	}
	return false
}

func (x *HostHealthStatus) GetFailedOutlierCheck() bool {
	if x != nil {
		return x.FailedOutlierCheck
	}
	return false
}

func (x *HostHealthStatus) GetFailedActiveDegradedCheck() bool {
	if x != nil {
		return x.FailedActiveDegradedCheck
	}
	return false
}

func (x *HostHealthStatus) GetPendingDynamicRemoval() bool {
	if x != nil {
		return x.PendingDynamicRemoval
	}
	return false
}

func (x *HostHealthStatus) GetPendingActiveHc() bool {
	if x != nil {
		return x.PendingActiveHc
	}
	return false
}

func (x *HostHealthStatus) GetEdsHealthStatus() core.HealthStatus {
	if x != nil {
		return x.EdsHealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

var File_envoy_admin_v2alpha_clusters_proto protoreflect.FileDescriptor

var file_envoy_admin_v2alpha_clusters_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x08, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4d, 0x0a, 0x10, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0xdd, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x5f, 0x76, 0x69, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x65, 0x64, 0x56, 0x69, 0x61, 0x41, 0x70, 0x69, 0x12,
	0x5a, 0x0a, 0x1f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x1c, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x44, 0x0a, 0x0d, 0x68,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x12, 0x72, 0x0a, 0x2c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x27, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x61, 0x74, 0x65, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x22, 0xd8, 0x03, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x36, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x19, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52,
	0x16, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x22, 0xf3, 0x02, 0x0a, 0x10, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x66, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74,
	0x6c, 0x69, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x3f, 0x0a, 0x1c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x17, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x2a, 0x0a,
	0x11, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x68, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x48, 0x63, 0x12, 0x4b, 0x0a, 0x11, 0x65, 0x64, 0x73,
	0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x65, 0x64, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x3c, 0x0a, 0x21, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0d, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_admin_v2alpha_clusters_proto_rawDescOnce sync.Once
	file_envoy_admin_v2alpha_clusters_proto_rawDescData = file_envoy_admin_v2alpha_clusters_proto_rawDesc
)

func file_envoy_admin_v2alpha_clusters_proto_rawDescGZIP() []byte {
	file_envoy_admin_v2alpha_clusters_proto_rawDescOnce.Do(func() {
		file_envoy_admin_v2alpha_clusters_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_admin_v2alpha_clusters_proto_rawDescData)
	})
	return file_envoy_admin_v2alpha_clusters_proto_rawDescData
}

var file_envoy_admin_v2alpha_clusters_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_admin_v2alpha_clusters_proto_goTypes = []interface{}{
	(*Clusters)(nil),         // 0: envoy.admin.v2alpha.Clusters
	(*ClusterStatus)(nil),    // 1: envoy.admin.v2alpha.ClusterStatus
	(*HostStatus)(nil),       // 2: envoy.admin.v2alpha.HostStatus
	(*HostHealthStatus)(nil), // 3: envoy.admin.v2alpha.HostHealthStatus
	(*_type.Percent)(nil),    // 4: envoy.type.Percent
	(*core.Address)(nil),     // 5: envoy.api.v2.core.Address
	(*SimpleMetric)(nil),     // 6: envoy.admin.v2alpha.SimpleMetric
	(*core.Locality)(nil),    // 7: envoy.api.v2.core.Locality
	(core.HealthStatus)(0),   // 8: envoy.api.v2.core.HealthStatus
}
var file_envoy_admin_v2alpha_clusters_proto_depIdxs = []int32{
	1,  // 0: envoy.admin.v2alpha.Clusters.cluster_statuses:type_name -> envoy.admin.v2alpha.ClusterStatus
	4,  // 1: envoy.admin.v2alpha.ClusterStatus.success_rate_ejection_threshold:type_name -> envoy.type.Percent
	2,  // 2: envoy.admin.v2alpha.ClusterStatus.host_statuses:type_name -> envoy.admin.v2alpha.HostStatus
	4,  // 3: envoy.admin.v2alpha.ClusterStatus.local_origin_success_rate_ejection_threshold:type_name -> envoy.type.Percent
	5,  // 4: envoy.admin.v2alpha.HostStatus.address:type_name -> envoy.api.v2.core.Address
	6,  // 5: envoy.admin.v2alpha.HostStatus.stats:type_name -> envoy.admin.v2alpha.SimpleMetric
	3,  // 6: envoy.admin.v2alpha.HostStatus.health_status:type_name -> envoy.admin.v2alpha.HostHealthStatus
	4,  // 7: envoy.admin.v2alpha.HostStatus.success_rate:type_name -> envoy.type.Percent
	4,  // 8: envoy.admin.v2alpha.HostStatus.local_origin_success_rate:type_name -> envoy.type.Percent
	7,  // 9: envoy.admin.v2alpha.HostStatus.locality:type_name -> envoy.api.v2.core.Locality
	8,  // 10: envoy.admin.v2alpha.HostHealthStatus.eds_health_status:type_name -> envoy.api.v2.core.HealthStatus
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_envoy_admin_v2alpha_clusters_proto_init() }
func file_envoy_admin_v2alpha_clusters_proto_init() {
	if File_envoy_admin_v2alpha_clusters_proto != nil {
		return
	}
	file_envoy_admin_v2alpha_metrics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_admin_v2alpha_clusters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clusters); i {
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
		file_envoy_admin_v2alpha_clusters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterStatus); i {
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
		file_envoy_admin_v2alpha_clusters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostStatus); i {
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
		file_envoy_admin_v2alpha_clusters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostHealthStatus); i {
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
			RawDescriptor: file_envoy_admin_v2alpha_clusters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_admin_v2alpha_clusters_proto_goTypes,
		DependencyIndexes: file_envoy_admin_v2alpha_clusters_proto_depIdxs,
		MessageInfos:      file_envoy_admin_v2alpha_clusters_proto_msgTypes,
	}.Build()
	File_envoy_admin_v2alpha_clusters_proto = out.File
	file_envoy_admin_v2alpha_clusters_proto_rawDesc = nil
	file_envoy_admin_v2alpha_clusters_proto_goTypes = nil
	file_envoy_admin_v2alpha_clusters_proto_depIdxs = nil
}
