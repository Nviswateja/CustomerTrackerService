// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: service/protos/Version2Protos/customerV2.proto

package protos

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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	PhoneNumber string `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_protos_Version2Protos_customerV2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_service_protos_Version2Protos_customerV2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_service_protos_Version2Protos_customerV2_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

// The request message containing the user's name.
type GetCustomerMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCustomerMessageRequest) Reset() {
	*x = GetCustomerMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_protos_Version2Protos_customerV2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerMessageRequest) ProtoMessage() {}

func (x *GetCustomerMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_protos_Version2Protos_customerV2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerMessageRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerMessageRequest) Descriptor() ([]byte, []int) {
	return file_service_protos_Version2Protos_customerV2_proto_rawDescGZIP(), []int{1}
}

type CustomerDetailsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customers []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
}

func (x *CustomerDetailsReply) Reset() {
	*x = CustomerDetailsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_protos_Version2Protos_customerV2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDetailsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDetailsReply) ProtoMessage() {}

func (x *CustomerDetailsReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_protos_Version2Protos_customerV2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDetailsReply.ProtoReflect.Descriptor instead.
func (*CustomerDetailsReply) Descriptor() ([]byte, []int) {
	return file_service_protos_Version2Protos_customerV2_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerDetailsReply) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

var File_service_protos_Version2Protos_customerV2_proto protoreflect.FileDescriptor

var file_service_protos_Version2Protos_customerV2_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x56, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x56, 0x32, 0x22, 0x40, 0x0a, 0x08,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x1b,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x14, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x56, 0x32, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x32, 0x73, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x56, 0x32, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x56, 0x32, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x3e, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x76, 0x69, 0x73, 0x77,
	0x61, 0x74, 0x65, 0x6a, 0x61, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_protos_Version2Protos_customerV2_proto_rawDescOnce sync.Once
	file_service_protos_Version2Protos_customerV2_proto_rawDescData = file_service_protos_Version2Protos_customerV2_proto_rawDesc
)

func file_service_protos_Version2Protos_customerV2_proto_rawDescGZIP() []byte {
	file_service_protos_Version2Protos_customerV2_proto_rawDescOnce.Do(func() {
		file_service_protos_Version2Protos_customerV2_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_protos_Version2Protos_customerV2_proto_rawDescData)
	})
	return file_service_protos_Version2Protos_customerV2_proto_rawDescData
}

var file_service_protos_Version2Protos_customerV2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_protos_Version2Protos_customerV2_proto_goTypes = []interface{}{
	(*Customer)(nil),                  // 0: customerV2.Customer
	(*GetCustomerMessageRequest)(nil), // 1: customerV2.GetCustomerMessageRequest
	(*CustomerDetailsReply)(nil),      // 2: customerV2.CustomerDetailsReply
}
var file_service_protos_Version2Protos_customerV2_proto_depIdxs = []int32{
	0, // 0: customerV2.CustomerDetailsReply.customers:type_name -> customerV2.Customer
	1, // 1: customerV2.CustomerService.GetCustomerWithName:input_type -> customerV2.GetCustomerMessageRequest
	2, // 2: customerV2.CustomerService.GetCustomerWithName:output_type -> customerV2.CustomerDetailsReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_protos_Version2Protos_customerV2_proto_init() }
func file_service_protos_Version2Protos_customerV2_proto_init() {
	if File_service_protos_Version2Protos_customerV2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_protos_Version2Protos_customerV2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_service_protos_Version2Protos_customerV2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerMessageRequest); i {
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
		file_service_protos_Version2Protos_customerV2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDetailsReply); i {
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
			RawDescriptor: file_service_protos_Version2Protos_customerV2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_protos_Version2Protos_customerV2_proto_goTypes,
		DependencyIndexes: file_service_protos_Version2Protos_customerV2_proto_depIdxs,
		MessageInfos:      file_service_protos_Version2Protos_customerV2_proto_msgTypes,
	}.Build()
	File_service_protos_Version2Protos_customerV2_proto = out.File
	file_service_protos_Version2Protos_customerV2_proto_rawDesc = nil
	file_service_protos_Version2Protos_customerV2_proto_goTypes = nil
	file_service_protos_Version2Protos_customerV2_proto_depIdxs = nil
}
