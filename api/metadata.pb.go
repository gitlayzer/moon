// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: metadata.proto

package api

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

// 指标标签值列表
type MetricLabelValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *MetricLabelValues) Reset() {
	*x = MetricLabelValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricLabelValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricLabelValues) ProtoMessage() {}

func (x *MetricLabelValues) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricLabelValues.ProtoReflect.Descriptor instead.
func (*MetricLabelValues) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *MetricLabelValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// 查询到的数据详情， 用与元数据构建
type MetricDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 指标名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 帮助信息
	Help string `protobuf:"bytes,2,opt,name=help,proto3" json:"help,omitempty"`
	// 类型
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// 标签集合
	Labels map[string]*MetricLabelValues `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 指标单位
	Unit string `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (x *MetricDetail) Reset() {
	*x = MetricDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDetail) ProtoMessage() {}

func (x *MetricDetail) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDetail.ProtoReflect.Descriptor instead.
func (*MetricDetail) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *MetricDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricDetail) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *MetricDetail) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MetricDetail) GetLabels() map[string]*MetricLabelValues {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MetricDetail) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

// Metric类型数据查询结果
type MetricQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 标签集合
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 结果类型
	ResultType string `protobuf:"bytes,2,opt,name=resultType,proto3" json:"resultType,omitempty"`
	// 结果值(图表)
	Values []*MetricQueryValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	// 结果值（单数据）
	Value *MetricQueryValue `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MetricQueryResult) Reset() {
	*x = MetricQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricQueryResult) ProtoMessage() {}

func (x *MetricQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricQueryResult.ProtoReflect.Descriptor instead.
func (*MetricQueryResult) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *MetricQueryResult) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MetricQueryResult) GetResultType() string {
	if x != nil {
		return x.ResultType
	}
	return ""
}

func (x *MetricQueryResult) GetValues() []*MetricQueryValue {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *MetricQueryResult) GetValue() *MetricQueryValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type MetricQueryValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 值
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// 时间戳
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MetricQueryValue) Reset() {
	*x = MetricQueryValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricQueryValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricQueryValue) ProtoMessage() {}

func (x *MetricQueryValue) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricQueryValue.ProtoReflect.Descriptor instead.
func (*MetricQueryValue) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *MetricQueryValue) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *MetricQueryValue) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_metadata_proto protoreflect.FileDescriptor

var file_metadata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2b, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0xe8, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x35, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x1a, 0x51, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x86, 0x02,
	0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2d, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x2b,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x46, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x2b,
	0x0a, 0x03, 0x61, 0x70, 0x69, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x64, 0x65, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x6f, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_metadata_proto_rawDescOnce sync.Once
	file_metadata_proto_rawDescData = file_metadata_proto_rawDesc
)

func file_metadata_proto_rawDescGZIP() []byte {
	file_metadata_proto_rawDescOnce.Do(func() {
		file_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_proto_rawDescData)
	})
	return file_metadata_proto_rawDescData
}

var file_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metadata_proto_goTypes = []interface{}{
	(*MetricLabelValues)(nil), // 0: api.MetricLabelValues
	(*MetricDetail)(nil),      // 1: api.MetricDetail
	(*MetricQueryResult)(nil), // 2: api.MetricQueryResult
	(*MetricQueryValue)(nil),  // 3: api.MetricQueryValue
	nil,                       // 4: api.MetricDetail.LabelsEntry
	nil,                       // 5: api.MetricQueryResult.LabelsEntry
}
var file_metadata_proto_depIdxs = []int32{
	4, // 0: api.MetricDetail.labels:type_name -> api.MetricDetail.LabelsEntry
	5, // 1: api.MetricQueryResult.labels:type_name -> api.MetricQueryResult.LabelsEntry
	3, // 2: api.MetricQueryResult.values:type_name -> api.MetricQueryValue
	3, // 3: api.MetricQueryResult.value:type_name -> api.MetricQueryValue
	0, // 4: api.MetricDetail.LabelsEntry.value:type_name -> api.MetricLabelValues
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_metadata_proto_init() }
func file_metadata_proto_init() {
	if File_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricLabelValues); i {
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
		file_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricDetail); i {
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
		file_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricQueryResult); i {
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
		file_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricQueryValue); i {
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
			RawDescriptor: file_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metadata_proto_goTypes,
		DependencyIndexes: file_metadata_proto_depIdxs,
		MessageInfos:      file_metadata_proto_msgTypes,
	}.Build()
	File_metadata_proto = out.File
	file_metadata_proto_rawDesc = nil
	file_metadata_proto_goTypes = nil
	file_metadata_proto_depIdxs = nil
}
