// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/api/v2/route.proto

package envoy_api_v2

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type RouteConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                            string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VirtualHosts                    []*route.VirtualHost      `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	Vhds                            *Vhds                     `protobuf:"bytes,9,opt,name=vhds,proto3" json:"vhds,omitempty"`
	InternalOnlyHeaders             []string                  `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	ResponseHeadersToAdd            []*core.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	ResponseHeadersToRemove         []string                  `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	RequestHeadersToAdd             []*core.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove          []string                  `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	MostSpecificHeaderMutationsWins bool                      `protobuf:"varint,10,opt,name=most_specific_header_mutations_wins,json=mostSpecificHeaderMutationsWins,proto3" json:"most_specific_header_mutations_wins,omitempty"`
	ValidateClusters                *wrappers.BoolValue       `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
}

func (x *RouteConfiguration) Reset() {
	*x = RouteConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteConfiguration) ProtoMessage() {}

func (x *RouteConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteConfiguration.ProtoReflect.Descriptor instead.
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_route_proto_rawDescGZIP(), []int{0}
}

func (x *RouteConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteConfiguration) GetVirtualHosts() []*route.VirtualHost {
	if x != nil {
		return x.VirtualHosts
	}
	return nil
}

func (x *RouteConfiguration) GetVhds() *Vhds {
	if x != nil {
		return x.Vhds
	}
	return nil
}

func (x *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if x != nil {
		return x.InternalOnlyHeaders
	}
	return nil
}

func (x *RouteConfiguration) GetResponseHeadersToAdd() []*core.HeaderValueOption {
	if x != nil {
		return x.ResponseHeadersToAdd
	}
	return nil
}

func (x *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if x != nil {
		return x.ResponseHeadersToRemove
	}
	return nil
}

func (x *RouteConfiguration) GetRequestHeadersToAdd() []*core.HeaderValueOption {
	if x != nil {
		return x.RequestHeadersToAdd
	}
	return nil
}

func (x *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if x != nil {
		return x.RequestHeadersToRemove
	}
	return nil
}

func (x *RouteConfiguration) GetMostSpecificHeaderMutationsWins() bool {
	if x != nil {
		return x.MostSpecificHeaderMutationsWins
	}
	return false
}

func (x *RouteConfiguration) GetValidateClusters() *wrappers.BoolValue {
	if x != nil {
		return x.ValidateClusters
	}
	return nil
}

type Vhds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigSource *core.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
}

func (x *Vhds) Reset() {
	*x = Vhds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vhds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vhds) ProtoMessage() {}

func (x *Vhds) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vhds.ProtoReflect.Descriptor instead.
func (*Vhds) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_route_proto_rawDescGZIP(), []int{1}
}

func (x *Vhds) GetConfigSource() *core.ConfigSource {
	if x != nil {
		return x.ConfigSource
	}
	return nil
}

var File_envoy_api_v2_route_proto protoreflect.FileDescriptor

var file_envoy_api_v2_route_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdd, 0x05, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x0c, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x04, 0x76, 0x68, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56,
	0x68, 0x64, 0x73, 0x52, 0x04, 0x76, 0x68, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x15, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x92, 0x01, 0x0a,
	0x22, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xc8, 0x01, 0x00, 0x52, 0x13, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x66, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x92, 0x01, 0x03, 0x10, 0xe8,
	0x07, 0x52, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x4d, 0x0a, 0x1a, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x10, 0xfa, 0x42, 0x0d,
	0x92, 0x01, 0x0a, 0x22, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xc8, 0x01, 0x00, 0x52, 0x17, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x64, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x92, 0x01, 0x03, 0x10, 0xe8, 0x07, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x4b, 0x0a, 0x19,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x10, 0xfa, 0x42, 0x0d, 0x92, 0x01, 0x0a, 0x22, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xc8, 0x01,
	0x00, 0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x4c, 0x0a, 0x23, 0x6d, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x77, 0x69, 0x6e, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1f, 0x6d, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x57, 0x69, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x22, 0x56, 0x0a, 0x04, 0x56, 0x68, 0x64, 0x73, 0x12, 0x4e, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x4f, 0x0a, 0x1a, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x17, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_api_v2_route_proto_rawDescOnce sync.Once
	file_envoy_api_v2_route_proto_rawDescData = file_envoy_api_v2_route_proto_rawDesc
)

func file_envoy_api_v2_route_proto_rawDescGZIP() []byte {
	file_envoy_api_v2_route_proto_rawDescOnce.Do(func() {
		file_envoy_api_v2_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_api_v2_route_proto_rawDescData)
	})
	return file_envoy_api_v2_route_proto_rawDescData
}

var file_envoy_api_v2_route_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_api_v2_route_proto_goTypes = []interface{}{
	(*RouteConfiguration)(nil),     // 0: envoy.api.v2.RouteConfiguration
	(*Vhds)(nil),                   // 1: envoy.api.v2.Vhds
	(*route.VirtualHost)(nil),      // 2: envoy.api.v2.route.VirtualHost
	(*core.HeaderValueOption)(nil), // 3: envoy.api.v2.core.HeaderValueOption
	(*wrappers.BoolValue)(nil),     // 4: google.protobuf.BoolValue
	(*core.ConfigSource)(nil),      // 5: envoy.api.v2.core.ConfigSource
}
var file_envoy_api_v2_route_proto_depIdxs = []int32{
	2, // 0: envoy.api.v2.RouteConfiguration.virtual_hosts:type_name -> envoy.api.v2.route.VirtualHost
	1, // 1: envoy.api.v2.RouteConfiguration.vhds:type_name -> envoy.api.v2.Vhds
	3, // 2: envoy.api.v2.RouteConfiguration.response_headers_to_add:type_name -> envoy.api.v2.core.HeaderValueOption
	3, // 3: envoy.api.v2.RouteConfiguration.request_headers_to_add:type_name -> envoy.api.v2.core.HeaderValueOption
	4, // 4: envoy.api.v2.RouteConfiguration.validate_clusters:type_name -> google.protobuf.BoolValue
	5, // 5: envoy.api.v2.Vhds.config_source:type_name -> envoy.api.v2.core.ConfigSource
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_api_v2_route_proto_init() }
func file_envoy_api_v2_route_proto_init() {
	if File_envoy_api_v2_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_api_v2_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteConfiguration); i {
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
		file_envoy_api_v2_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vhds); i {
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
			RawDescriptor: file_envoy_api_v2_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_api_v2_route_proto_goTypes,
		DependencyIndexes: file_envoy_api_v2_route_proto_depIdxs,
		MessageInfos:      file_envoy_api_v2_route_proto_msgTypes,
	}.Build()
	File_envoy_api_v2_route_proto = out.File
	file_envoy_api_v2_route_proto_rawDesc = nil
	file_envoy_api_v2_route_proto_goTypes = nil
	file_envoy_api_v2_route_proto_depIdxs = nil
}
