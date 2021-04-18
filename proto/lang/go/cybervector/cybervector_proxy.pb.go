// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: cybervector_proxy.proto

package cybervector

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

type MessageType int32

const (
	MessageType_Unknown       MessageType = 0
	MessageType_Subscribed    MessageType = 1
	MessageType_Unsubscribed  MessageType = 2
	MessageType_KeepAlive     MessageType = 3
	MessageType_ProcessIntent MessageType = 4
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "Unknown",
		1: "Subscribed",
		2: "Unsubscribed",
		3: "KeepAlive",
		4: "ProcessIntent",
	}
	MessageType_value = map[string]int32{
		"Unknown":       0,
		"Subscribed":    1,
		"Unsubscribed":  2,
		"KeepAlive":     3,
		"ProcessIntent": 4,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_cybervector_proxy_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_cybervector_proxy_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_cybervector_proxy_proto_rawDescGZIP(), []int{0}
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// keep alive stream interval in seconds,
	// requested from client, default 30 seconds
	KeepAlive int64 `protobuf:"varint,1,opt,name=keep_alive,json=keepAlive,proto3" json:"keep_alive,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cybervector_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cybervector_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_cybervector_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetKeepAlive() int64 {
	if x != nil {
		return x.KeepAlive
	}
	return 0
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cybervector_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cybervector_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_cybervector_proxy_proto_rawDescGZIP(), []int{1}
}

type ProxyMessaage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Proxy Message Type enum
	MessageType MessageType `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=cybervector.MessageType" json:"message_type,omitempty"`
	// Proxy Message Json data
	MessageData string `protobuf:"bytes,2,opt,name=message_data,json=messageData,proto3" json:"message_data,omitempty"`
	// The named intent to respond with.
	IntentName string `protobuf:"bytes,3,opt,name=intent_name,json=intentName,proto3" json:"intent_name,omitempty"`
	// A list of parameters as parsed by the escape pod
	Parameters map[string]string `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProxyMessaage) Reset() {
	*x = ProxyMessaage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cybervector_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyMessaage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyMessaage) ProtoMessage() {}

func (x *ProxyMessaage) ProtoReflect() protoreflect.Message {
	mi := &file_cybervector_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyMessaage.ProtoReflect.Descriptor instead.
func (*ProxyMessaage) Descriptor() ([]byte, []int) {
	return file_cybervector_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyMessaage) GetMessageType() MessageType {
	if x != nil {
		return x.MessageType
	}
	return MessageType_Unknown
}

func (x *ProxyMessaage) GetMessageData() string {
	if x != nil {
		return x.MessageData
	}
	return ""
}

func (x *ProxyMessaage) GetIntentName() string {
	if x != nil {
		return x.IntentName
	}
	return ""
}

func (x *ProxyMessaage) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

var File_cybervector_proxy_proto protoreflect.FileDescriptor

var file_cybervector_proxy_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x79, 0x62, 0x65, 0x72, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x79, 0x62, 0x65, 0x72,
	0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x31, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65,
	0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x9b, 0x02, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x61, 0x67,
	0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d,
	0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x5e, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4b,
	0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x10, 0x04, 0x32, 0xb3, 0x01,
	0x0a, 0x17, 0x43, 0x79, 0x62, 0x65, 0x72, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x0b, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x1f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x79, 0x62, 0x33, 0x72, 0x64, 0x6f, 0x67, 0x2f, 0x65, 0x73, 0x63, 0x61, 0x70,
	0x65, 0x2d, 0x70, 0x6f, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x79, 0x62, 0x65, 0x72, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cybervector_proxy_proto_rawDescOnce sync.Once
	file_cybervector_proxy_proto_rawDescData = file_cybervector_proxy_proto_rawDesc
)

func file_cybervector_proxy_proto_rawDescGZIP() []byte {
	file_cybervector_proxy_proto_rawDescOnce.Do(func() {
		file_cybervector_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_cybervector_proxy_proto_rawDescData)
	})
	return file_cybervector_proxy_proto_rawDescData
}

var file_cybervector_proxy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cybervector_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cybervector_proxy_proto_goTypes = []interface{}{
	(MessageType)(0),           // 0: cybervector.MessageType
	(*SubscribeRequest)(nil),   // 1: cybervector.SubscribeRequest
	(*UnsubscribeRequest)(nil), // 2: cybervector.UnsubscribeRequest
	(*ProxyMessaage)(nil),      // 3: cybervector.ProxyMessaage
	nil,                        // 4: cybervector.ProxyMessaage.ParametersEntry
}
var file_cybervector_proxy_proto_depIdxs = []int32{
	0, // 0: cybervector.ProxyMessaage.message_type:type_name -> cybervector.MessageType
	4, // 1: cybervector.ProxyMessaage.parameters:type_name -> cybervector.ProxyMessaage.ParametersEntry
	1, // 2: cybervector.CyberVectorProxyService.Subscribe:input_type -> cybervector.SubscribeRequest
	2, // 3: cybervector.CyberVectorProxyService.UnSubscribe:input_type -> cybervector.UnsubscribeRequest
	3, // 4: cybervector.CyberVectorProxyService.Subscribe:output_type -> cybervector.ProxyMessaage
	3, // 5: cybervector.CyberVectorProxyService.UnSubscribe:output_type -> cybervector.ProxyMessaage
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cybervector_proxy_proto_init() }
func file_cybervector_proxy_proto_init() {
	if File_cybervector_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cybervector_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_cybervector_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
		file_cybervector_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyMessaage); i {
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
			RawDescriptor: file_cybervector_proxy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cybervector_proxy_proto_goTypes,
		DependencyIndexes: file_cybervector_proxy_proto_depIdxs,
		EnumInfos:         file_cybervector_proxy_proto_enumTypes,
		MessageInfos:      file_cybervector_proxy_proto_msgTypes,
	}.Build()
	File_cybervector_proxy_proto = out.File
	file_cybervector_proxy_proto_rawDesc = nil
	file_cybervector_proxy_proto_goTypes = nil
	file_cybervector_proxy_proto_depIdxs = nil
}
