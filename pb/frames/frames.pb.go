// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: frames.proto

package frames

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

type OPERATOR int32

const (
	OPERATOR_NONE OPERATOR = 0
	OPERATOR_WALK OPERATOR = 1
)

// Enum value maps for OPERATOR.
var (
	OPERATOR_name = map[int32]string{
		0: "NONE",
		1: "WALK",
	}
	OPERATOR_value = map[string]int32{
		"NONE": 0,
		"WALK": 1,
	}
)

func (x OPERATOR) Enum() *OPERATOR {
	p := new(OPERATOR)
	*p = x
	return p
}

func (x OPERATOR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OPERATOR) Descriptor() protoreflect.EnumDescriptor {
	return file_frames_proto_enumTypes[0].Descriptor()
}

func (OPERATOR) Type() protoreflect.EnumType {
	return &file_frames_proto_enumTypes[0]
}

func (x OPERATOR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OPERATOR.Descriptor instead.
func (OPERATOR) EnumDescriptor() ([]byte, []int) {
	return file_frames_proto_rawDescGZIP(), []int{0}
}

type DIRCTION int32

const (
	DIRCTION_UP    DIRCTION = 0
	DIRCTION_DOWN  DIRCTION = 1
	DIRCTION_LEFT  DIRCTION = 2
	DIRCTION_RIGHT DIRCTION = 3
)

// Enum value maps for DIRCTION.
var (
	DIRCTION_name = map[int32]string{
		0: "UP",
		1: "DOWN",
		2: "LEFT",
		3: "RIGHT",
	}
	DIRCTION_value = map[string]int32{
		"UP":    0,
		"DOWN":  1,
		"LEFT":  2,
		"RIGHT": 3,
	}
)

func (x DIRCTION) Enum() *DIRCTION {
	p := new(DIRCTION)
	*p = x
	return p
}

func (x DIRCTION) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DIRCTION) Descriptor() protoreflect.EnumDescriptor {
	return file_frames_proto_enumTypes[1].Descriptor()
}

func (DIRCTION) Type() protoreflect.EnumType {
	return &file_frames_proto_enumTypes[1]
}

func (x DIRCTION) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DIRCTION.Descriptor instead.
func (DIRCTION) EnumDescriptor() ([]byte, []int) {
	return file_frames_proto_rawDescGZIP(), []int{1}
}

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fid       int32    `protobuf:"varint,1,opt,name=fid,proto3" json:"fid,omitempty"`
	Timestamp int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Pid       int32    `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Opt       OPERATOR `protobuf:"varint,4,opt,name=opt,proto3,enum=OPERATOR" json:"opt,omitempty"`
	Dir       DIRCTION `protobuf:"varint,5,opt,name=dir,proto3,enum=DIRCTION" json:"dir,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frames_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_frames_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_frames_proto_rawDescGZIP(), []int{0}
}

func (x *Frame) GetFid() int32 {
	if x != nil {
		return x.Fid
	}
	return 0
}

func (x *Frame) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Frame) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Frame) GetOpt() OPERATOR {
	if x != nil {
		return x.Opt
	}
	return OPERATOR_NONE
}

func (x *Frame) GetDir() DIRCTION {
	if x != nil {
		return x.Dir
	}
	return DIRCTION_UP
}

type S2CFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fid       int32    `protobuf:"varint,1,opt,name=fid,proto3" json:"fid,omitempty"`
	Timestamp int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Pid       int32    `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Opt       OPERATOR `protobuf:"varint,4,opt,name=opt,proto3,enum=OPERATOR" json:"opt,omitempty"`
	Dir       DIRCTION `protobuf:"varint,5,opt,name=dir,proto3,enum=DIRCTION" json:"dir,omitempty"`
}

func (x *S2CFrame) Reset() {
	*x = S2CFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frames_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CFrame) ProtoMessage() {}

func (x *S2CFrame) ProtoReflect() protoreflect.Message {
	mi := &file_frames_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CFrame.ProtoReflect.Descriptor instead.
func (*S2CFrame) Descriptor() ([]byte, []int) {
	return file_frames_proto_rawDescGZIP(), []int{1}
}

func (x *S2CFrame) GetFid() int32 {
	if x != nil {
		return x.Fid
	}
	return 0
}

func (x *S2CFrame) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *S2CFrame) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *S2CFrame) GetOpt() OPERATOR {
	if x != nil {
		return x.Opt
	}
	return OPERATOR_NONE
}

func (x *S2CFrame) GetDir() DIRCTION {
	if x != nil {
		return x.Dir
	}
	return DIRCTION_UP
}

var File_frames_proto protoreflect.FileDescriptor

var file_frames_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83,
	0x01, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x66, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x03, 0x6f, 0x70,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x4f, 0x52, 0x52, 0x03, 0x6f, 0x70, 0x74, 0x12, 0x1b, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x44, 0x49, 0x52, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x52,
	0x03, 0x64, 0x69, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x53, 0x32, 0x43, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x66, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x70, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x03, 0x6f, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x09, 0x2e, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x52, 0x03, 0x6f, 0x70, 0x74,
	0x12, 0x1b, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e,
	0x44, 0x49, 0x52, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x52, 0x03, 0x64, 0x69, 0x72, 0x2a, 0x1e, 0x0a,
	0x08, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x4c, 0x4b, 0x10, 0x01, 0x2a, 0x31, 0x0a,
	0x08, 0x44, 0x49, 0x52, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c,
	0x45, 0x46, 0x54, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x03,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_frames_proto_rawDescOnce sync.Once
	file_frames_proto_rawDescData = file_frames_proto_rawDesc
)

func file_frames_proto_rawDescGZIP() []byte {
	file_frames_proto_rawDescOnce.Do(func() {
		file_frames_proto_rawDescData = protoimpl.X.CompressGZIP(file_frames_proto_rawDescData)
	})
	return file_frames_proto_rawDescData
}

var file_frames_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_frames_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_frames_proto_goTypes = []interface{}{
	(OPERATOR)(0),    // 0: OPERATOR
	(DIRCTION)(0),    // 1: DIRCTION
	(*Frame)(nil),    // 2: Frame
	(*S2CFrame)(nil), // 3: S2CFrame
}
var file_frames_proto_depIdxs = []int32{
	0, // 0: Frame.opt:type_name -> OPERATOR
	1, // 1: Frame.dir:type_name -> DIRCTION
	0, // 2: S2CFrame.opt:type_name -> OPERATOR
	1, // 3: S2CFrame.dir:type_name -> DIRCTION
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_frames_proto_init() }
func file_frames_proto_init() {
	if File_frames_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frames_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
		file_frames_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CFrame); i {
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
			RawDescriptor: file_frames_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frames_proto_goTypes,
		DependencyIndexes: file_frames_proto_depIdxs,
		EnumInfos:         file_frames_proto_enumTypes,
		MessageInfos:      file_frames_proto_msgTypes,
	}.Build()
	File_frames_proto = out.File
	file_frames_proto_rawDesc = nil
	file_frames_proto_goTypes = nil
	file_frames_proto_depIdxs = nil
}
