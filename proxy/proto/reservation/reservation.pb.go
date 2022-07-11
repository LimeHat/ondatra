// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: reservation.proto

package reservation

import (
	proto "github.com/openconfig/ondatra/proto"
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

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Devices map[string]*ResolvedDevice `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ates    map[string]*ResolvedDevice `protobuf:"bytes,3,rep,name=ates,proto3" json:"ates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Links   []*ResolvedLink            `protobuf:"bytes,4,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *Reservation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reservation) GetDevices() map[string]*ResolvedDevice {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *Reservation) GetAtes() map[string]*ResolvedDevice {
	if x != nil {
		return x.Ates
	}
	return nil
}

func (x *Reservation) GetLinks() []*ResolvedLink {
	if x != nil {
		return x.Links
	}
	return nil
}

// ResolvedDevice is a device after it has been reserved.
type ResolvedDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key of the device.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Resolvable name of the device.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Vendor of the device.
	Vendor proto.Device_Vendor `protobuf:"varint,3,opt,name=vendor,proto3,enum=ondatra.Device_Vendor" json:"vendor,omitempty"`
	// Hardware model of the device.
	HardwareModel string `protobuf:"bytes,4,opt,name=hardware_model,json=hardwareModel,proto3" json:"hardware_model,omitempty"`
	// Software version of the device.
	SoftwareVersion string                   `protobuf:"bytes,5,opt,name=software_version,json=softwareVersion,proto3" json:"software_version,omitempty"`
	Ports           map[string]*ResolvedPort `protobuf:"bytes,6,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Services provides a map for the service locations provided by the device.
	Services map[string]*Service `protobuf:"bytes,7,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResolvedDevice) Reset() {
	*x = ResolvedDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolvedDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolvedDevice) ProtoMessage() {}

func (x *ResolvedDevice) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolvedDevice.ProtoReflect.Descriptor instead.
func (*ResolvedDevice) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *ResolvedDevice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResolvedDevice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResolvedDevice) GetVendor() proto.Device_Vendor {
	if x != nil {
		return x.Vendor
	}
	return proto.Device_Vendor(0)
}

func (x *ResolvedDevice) GetHardwareModel() string {
	if x != nil {
		return x.HardwareModel
	}
	return ""
}

func (x *ResolvedDevice) GetSoftwareVersion() string {
	if x != nil {
		return x.SoftwareVersion
	}
	return ""
}

func (x *ResolvedDevice) GetPorts() map[string]*ResolvedPort {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *ResolvedDevice) GetServices() map[string]*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

// ResolvedPort is a port with a concrete assignment.
type ResolvedPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Speed enum values are the port speed in Gbps.
	Speed proto.Port_Speed `protobuf:"varint,2,opt,name=speed,proto3,enum=ondatra.Port_Speed" json:"speed,omitempty"`
	// Name is the resolved port name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ResolvedPort) Reset() {
	*x = ResolvedPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolvedPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolvedPort) ProtoMessage() {}

func (x *ResolvedPort) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolvedPort.ProtoReflect.Descriptor instead.
func (*ResolvedPort) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{2}
}

func (x *ResolvedPort) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResolvedPort) GetSpeed() proto.Port_Speed {
	if x != nil {
		return x.Speed
	}
	return proto.Port_Speed(0)
}

func (x *ResolvedPort) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A physical link between ports on DUTs or ATEs.
// The order does not matter: links are symmetrical.
// A given port may be specified in at most one link.
// If the port is not connected but defined in the Device.
// There will be no corresponding link.
type ResolvedLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"` // First port in the format "<device-id>:<port-id>".
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"` // Second port in the format "<device-id>:<port-id>".
}

func (x *ResolvedLink) Reset() {
	*x = ResolvedLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolvedLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolvedLink) ProtoMessage() {}

func (x *ResolvedLink) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolvedLink.ProtoReflect.Descriptor instead.
func (*ResolvedLink) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{3}
}

func (x *ResolvedLink) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *ResolvedLink) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

// Service provides a map of services provided by the DUT which can be mapped
// for use by ondatra to provide access to those endpoints.
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Endpoint is the endpoint type of the service.
	//
	// Types that are assignable to Endpoint:
	//	*Service_Grpc
	//	*Service_ProxiedGrpc
	//	*Service_Http
	//	*Service_HttpOverGrpc
	Endpoint isService_Endpoint `protobuf_oneof:"endpoint"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{4}
}

func (x *Service) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *Service) GetEndpoint() isService_Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (x *Service) GetGrpc() *GRPCEndpoint {
	if x, ok := x.GetEndpoint().(*Service_Grpc); ok {
		return x.Grpc
	}
	return nil
}

func (x *Service) GetProxiedGrpc() *ProxiedGRPCEndpoint {
	if x, ok := x.GetEndpoint().(*Service_ProxiedGrpc); ok {
		return x.ProxiedGrpc
	}
	return nil
}

func (x *Service) GetHttp() *HTTPEndpoint {
	if x, ok := x.GetEndpoint().(*Service_Http); ok {
		return x.Http
	}
	return nil
}

func (x *Service) GetHttpOverGrpc() *HTTPOverGRPCEndpoint {
	if x, ok := x.GetEndpoint().(*Service_HttpOverGrpc); ok {
		return x.HttpOverGrpc
	}
	return nil
}

type isService_Endpoint interface {
	isService_Endpoint()
}

type Service_Grpc struct {
	Grpc *GRPCEndpoint `protobuf:"bytes,100,opt,name=grpc,proto3,oneof"`
}

type Service_ProxiedGrpc struct {
	ProxiedGrpc *ProxiedGRPCEndpoint `protobuf:"bytes,101,opt,name=proxied_grpc,json=proxiedGrpc,proto3,oneof"`
}

type Service_Http struct {
	Http *HTTPEndpoint `protobuf:"bytes,102,opt,name=http,proto3,oneof"`
}

type Service_HttpOverGrpc struct {
	HttpOverGrpc *HTTPOverGRPCEndpoint `protobuf:"bytes,103,opt,name=http_over_grpc,json=httpOverGrpc,proto3,oneof"`
}

func (*Service_Grpc) isService_Endpoint() {}

func (*Service_ProxiedGrpc) isService_Endpoint() {}

func (*Service_Http) isService_Endpoint() {}

func (*Service_HttpOverGrpc) isService_Endpoint() {}

type GRPCEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address is the form <ip / dns>:<port>.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GRPCEndpoint) Reset() {
	*x = GRPCEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCEndpoint) ProtoMessage() {}

func (x *GRPCEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCEndpoint.ProtoReflect.Descriptor instead.
func (*GRPCEndpoint) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{5}
}

func (x *GRPCEndpoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// ProxiedGRPCEndpoint is an endpoint which will pass through a list of proxies.
type ProxiedGRPCEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address is the form <ip / dns>:<port>.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Proxy is the form <ip / dns>:<port>. Repeated for each proxy in the chain.
	Proxy []string `protobuf:"bytes,2,rep,name=proxy,proto3" json:"proxy,omitempty"`
}

func (x *ProxiedGRPCEndpoint) Reset() {
	*x = ProxiedGRPCEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxiedGRPCEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxiedGRPCEndpoint) ProtoMessage() {}

func (x *ProxiedGRPCEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxiedGRPCEndpoint.ProtoReflect.Descriptor instead.
func (*ProxiedGRPCEndpoint) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{6}
}

func (x *ProxiedGRPCEndpoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ProxiedGRPCEndpoint) GetProxy() []string {
	if x != nil {
		return x.Proxy
	}
	return nil
}

// HTTPEndpoint is a standard HTTP endpoint.
type HTTPEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address is the form <ip / dns>:<port>.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *HTTPEndpoint) Reset() {
	*x = HTTPEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPEndpoint) ProtoMessage() {}

func (x *HTTPEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPEndpoint.ProtoReflect.Descriptor instead.
func (*HTTPEndpoint) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{7}
}

func (x *HTTPEndpoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// HTTPOverGRPCEndpoint is a GRPC transport for HTTP proxied connections.
type HTTPOverGRPCEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address is the form <ip / dns>:<port>.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *HTTPOverGRPCEndpoint) Reset() {
	*x = HTTPOverGRPCEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPOverGRPCEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPOverGRPCEndpoint) ProtoMessage() {}

func (x *HTTPOverGRPCEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPOverGRPCEndpoint.ProtoReflect.Descriptor instead.
func (*HTTPOverGRPCEndpoint) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{8}
}

func (x *HTTPOverGRPCEndpoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_reservation_proto protoreflect.FileDescriptor

var file_reservation_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x74, 0x65, 0x73, 0x12, 0x2f,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x1a,
	0x57, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x54, 0x0a, 0x09, 0x41, 0x74, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe3,
	0x03, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x29, 0x0a, 0x10,
	0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x53, 0x0a, 0x0a,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x51, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x5d, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x62, 0x22,
	0x99, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x45, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x64, 0x47, 0x52, 0x50, 0x43, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x64, 0x47,
	0x72, 0x70, 0x63, 0x12, 0x2f, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x48, 0x54, 0x54, 0x50, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04,
	0x68, 0x74, 0x74, 0x70, 0x12, 0x49, 0x0a, 0x0e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6f, 0x76, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x4f,
	0x76, 0x65, 0x72, 0x47, 0x52, 0x50, 0x43, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x4f, 0x76, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x42,
	0x0a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x47,
	0x52, 0x50, 0x43, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x45, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x64,
	0x47, 0x52, 0x50, 0x43, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x28, 0x0a, 0x0c,
	0x48, 0x54, 0x54, 0x50, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x48, 0x54, 0x54, 0x50, 0x4f, 0x76,
	0x65, 0x72, 0x47, 0x52, 0x50, 0x43, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_reservation_proto_rawDescOnce sync.Once
	file_reservation_proto_rawDescData = file_reservation_proto_rawDesc
)

func file_reservation_proto_rawDescGZIP() []byte {
	file_reservation_proto_rawDescOnce.Do(func() {
		file_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservation_proto_rawDescData)
	})
	return file_reservation_proto_rawDescData
}

var file_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_reservation_proto_goTypes = []interface{}{
	(*Reservation)(nil),          // 0: reservation.Reservation
	(*ResolvedDevice)(nil),       // 1: reservation.ResolvedDevice
	(*ResolvedPort)(nil),         // 2: reservation.ResolvedPort
	(*ResolvedLink)(nil),         // 3: reservation.ResolvedLink
	(*Service)(nil),              // 4: reservation.Service
	(*GRPCEndpoint)(nil),         // 5: reservation.GRPCEndpoint
	(*ProxiedGRPCEndpoint)(nil),  // 6: reservation.ProxiedGRPCEndpoint
	(*HTTPEndpoint)(nil),         // 7: reservation.HTTPEndpoint
	(*HTTPOverGRPCEndpoint)(nil), // 8: reservation.HTTPOverGRPCEndpoint
	nil,                          // 9: reservation.Reservation.DevicesEntry
	nil,                          // 10: reservation.Reservation.AtesEntry
	nil,                          // 11: reservation.ResolvedDevice.PortsEntry
	nil,                          // 12: reservation.ResolvedDevice.ServicesEntry
	(proto.Device_Vendor)(0),     // 13: ondatra.Device.Vendor
	(proto.Port_Speed)(0),        // 14: ondatra.Port.Speed
}
var file_reservation_proto_depIdxs = []int32{
	9,  // 0: reservation.Reservation.devices:type_name -> reservation.Reservation.DevicesEntry
	10, // 1: reservation.Reservation.ates:type_name -> reservation.Reservation.AtesEntry
	3,  // 2: reservation.Reservation.links:type_name -> reservation.ResolvedLink
	13, // 3: reservation.ResolvedDevice.vendor:type_name -> ondatra.Device.Vendor
	11, // 4: reservation.ResolvedDevice.ports:type_name -> reservation.ResolvedDevice.PortsEntry
	12, // 5: reservation.ResolvedDevice.services:type_name -> reservation.ResolvedDevice.ServicesEntry
	14, // 6: reservation.ResolvedPort.speed:type_name -> ondatra.Port.Speed
	5,  // 7: reservation.Service.grpc:type_name -> reservation.GRPCEndpoint
	6,  // 8: reservation.Service.proxied_grpc:type_name -> reservation.ProxiedGRPCEndpoint
	7,  // 9: reservation.Service.http:type_name -> reservation.HTTPEndpoint
	8,  // 10: reservation.Service.http_over_grpc:type_name -> reservation.HTTPOverGRPCEndpoint
	1,  // 11: reservation.Reservation.DevicesEntry.value:type_name -> reservation.ResolvedDevice
	1,  // 12: reservation.Reservation.AtesEntry.value:type_name -> reservation.ResolvedDevice
	2,  // 13: reservation.ResolvedDevice.PortsEntry.value:type_name -> reservation.ResolvedPort
	4,  // 14: reservation.ResolvedDevice.ServicesEntry.value:type_name -> reservation.Service
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_reservation_proto_init() }
func file_reservation_proto_init() {
	if File_reservation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
		file_reservation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolvedDevice); i {
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
		file_reservation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolvedPort); i {
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
		file_reservation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolvedLink); i {
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
		file_reservation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_reservation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCEndpoint); i {
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
		file_reservation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxiedGRPCEndpoint); i {
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
		file_reservation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPEndpoint); i {
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
		file_reservation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPOverGRPCEndpoint); i {
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
	file_reservation_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Service_Grpc)(nil),
		(*Service_ProxiedGrpc)(nil),
		(*Service_Http)(nil),
		(*Service_HttpOverGrpc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reservation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reservation_proto_goTypes,
		DependencyIndexes: file_reservation_proto_depIdxs,
		MessageInfos:      file_reservation_proto_msgTypes,
	}.Build()
	File_reservation_proto = out.File
	file_reservation_proto_rawDesc = nil
	file_reservation_proto_goTypes = nil
	file_reservation_proto_depIdxs = nil
}
