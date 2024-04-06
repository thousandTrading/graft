// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.7
// source: tensorflow/core/framework/optimized_function_graph.proto

package optimized_function_graph_go_proto

import (
	graph_go_proto "github.com/wamuir/graft/tensorflow/core/framework/graph_go_proto"
	types_go_proto "github.com/wamuir/graft/tensorflow/core/framework/types_go_proto"
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

// Enum for distinguishing the origin where the proto is created.
//
// AOT: proto is created in ahead-of-time environment, which can be different
// from the environment where the graph is actually executed.
//
// JIT: proto is created in just-in-time execution, which has the same
// environment as the one the graph is actually executed.
type OptimizedFunctionGraph_OptimizationSource int32

const (
	OptimizedFunctionGraph_SOURCE_UNSPECIFIED OptimizedFunctionGraph_OptimizationSource = 0
	OptimizedFunctionGraph_AOT                OptimizedFunctionGraph_OptimizationSource = 1
	OptimizedFunctionGraph_JIT                OptimizedFunctionGraph_OptimizationSource = 2
)

// Enum value maps for OptimizedFunctionGraph_OptimizationSource.
var (
	OptimizedFunctionGraph_OptimizationSource_name = map[int32]string{
		0: "SOURCE_UNSPECIFIED",
		1: "AOT",
		2: "JIT",
	}
	OptimizedFunctionGraph_OptimizationSource_value = map[string]int32{
		"SOURCE_UNSPECIFIED": 0,
		"AOT":                1,
		"JIT":                2,
	}
)

func (x OptimizedFunctionGraph_OptimizationSource) Enum() *OptimizedFunctionGraph_OptimizationSource {
	p := new(OptimizedFunctionGraph_OptimizationSource)
	*p = x
	return p
}

func (x OptimizedFunctionGraph_OptimizationSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OptimizedFunctionGraph_OptimizationSource) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_framework_optimized_function_graph_proto_enumTypes[0].Descriptor()
}

func (OptimizedFunctionGraph_OptimizationSource) Type() protoreflect.EnumType {
	return &file_tensorflow_core_framework_optimized_function_graph_proto_enumTypes[0]
}

func (x OptimizedFunctionGraph_OptimizationSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OptimizedFunctionGraph_OptimizationSource.Descriptor instead.
func (OptimizedFunctionGraph_OptimizationSource) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_optimized_function_graph_proto_rawDescGZIP(), []int{0, 0}
}

// Optimized function graph after instantiation-related graph optimization
// passes (up till before graph partitioning). The first half of the proto is
// representing a GraphDef and the rest of the fields are extra information from
// graph optimizations.
type OptimizedFunctionGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Function name. It can be a human-readable SignatureDef's method name, or a
	// FunctionDef name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optimized function graph.
	FunctionGraph *graph_go_proto.GraphDef `protobuf:"bytes,2,opt,name=function_graph,json=functionGraph,proto3" json:"function_graph,omitempty"`
	// Maps from node name to control ret. This is an output from running TF/XLA
	// bridge.
	NodeNameToControlRet map[string]string `protobuf:"bytes,3,rep,name=node_name_to_control_ret,json=nodeNameToControlRet,proto3" json:"node_name_to_control_ret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Return node types of the function. This is an output of graph
	// preprocessing.
	RetTypes []types_go_proto.DataType `protobuf:"varint,4,rep,packed,name=ret_types,json=retTypes,proto3,enum=tensorflow.DataType" json:"ret_types,omitempty"`
	// Number of return nodes. This is an output of graph preprocessing.
	NumReturnNodes uint32 `protobuf:"varint,5,opt,name=num_return_nodes,json=numReturnNodes,proto3" json:"num_return_nodes,omitempty"`
	// Indicates the source environment where this proto is generated.
	Source *OptimizedFunctionGraph_OptimizationSource `protobuf:"varint,7,opt,name=source,proto3,enum=tensorflow.OptimizedFunctionGraph_OptimizationSource,oneof" json:"source,omitempty"`
	// Time (in microseconds) spent on running the graph optimization passes for
	// this function.
	OptimizationTimeUsecs *uint64 `protobuf:"varint,8,opt,name=optimization_time_usecs,json=optimizationTimeUsecs,proto3,oneof" json:"optimization_time_usecs,omitempty"`
}

func (x *OptimizedFunctionGraph) Reset() {
	*x = OptimizedFunctionGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_optimized_function_graph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptimizedFunctionGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptimizedFunctionGraph) ProtoMessage() {}

func (x *OptimizedFunctionGraph) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_optimized_function_graph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptimizedFunctionGraph.ProtoReflect.Descriptor instead.
func (*OptimizedFunctionGraph) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_optimized_function_graph_proto_rawDescGZIP(), []int{0}
}

func (x *OptimizedFunctionGraph) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OptimizedFunctionGraph) GetFunctionGraph() *graph_go_proto.GraphDef {
	if x != nil {
		return x.FunctionGraph
	}
	return nil
}

func (x *OptimizedFunctionGraph) GetNodeNameToControlRet() map[string]string {
	if x != nil {
		return x.NodeNameToControlRet
	}
	return nil
}

func (x *OptimizedFunctionGraph) GetRetTypes() []types_go_proto.DataType {
	if x != nil {
		return x.RetTypes
	}
	return nil
}

func (x *OptimizedFunctionGraph) GetNumReturnNodes() uint32 {
	if x != nil {
		return x.NumReturnNodes
	}
	return 0
}

func (x *OptimizedFunctionGraph) GetSource() OptimizedFunctionGraph_OptimizationSource {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return OptimizedFunctionGraph_SOURCE_UNSPECIFIED
}

func (x *OptimizedFunctionGraph) GetOptimizationTimeUsecs() uint64 {
	if x != nil && x.OptimizationTimeUsecs != nil {
		return *x.OptimizationTimeUsecs
	}
	return 0
}

var File_tensorflow_core_framework_optimized_function_graph_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_optimized_function_graph_proto_rawDesc = []byte{
	0x0a, 0x38, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6d, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x05, 0x0a, 0x16, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a,
	0x65, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65,
	0x66, 0x52, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x12, 0x74, 0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x6f,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x54,
	0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x14, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x52, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d,
	0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x17, 0x6f, 0x70, 0x74, 0x69, 0x6d,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x63, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x15, 0x6f, 0x70, 0x74, 0x69,
	0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x63,
	0x73, 0x88, 0x01, 0x01, 0x1a, 0x47, 0x0a, 0x19, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a,
	0x12, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x4f, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x49, 0x54, 0x10, 0x02, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75,
	0x73, 0x65, 0x63, 0x73, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x42, 0x61, 0x5a, 0x5f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6d, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_optimized_function_graph_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_optimized_function_graph_proto_rawDescData = file_tensorflow_core_framework_optimized_function_graph_proto_rawDesc
)

func file_tensorflow_core_framework_optimized_function_graph_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_optimized_function_graph_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_optimized_function_graph_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_optimized_function_graph_proto_rawDescData)
	})
	return file_tensorflow_core_framework_optimized_function_graph_proto_rawDescData
}

var file_tensorflow_core_framework_optimized_function_graph_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_framework_optimized_function_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_framework_optimized_function_graph_proto_goTypes = []interface{}{
	(OptimizedFunctionGraph_OptimizationSource)(0), // 0: tensorflow.OptimizedFunctionGraph.OptimizationSource
	(*OptimizedFunctionGraph)(nil),                 // 1: tensorflow.OptimizedFunctionGraph
	nil,                                            // 2: tensorflow.OptimizedFunctionGraph.NodeNameToControlRetEntry
	(*graph_go_proto.GraphDef)(nil),                // 3: tensorflow.GraphDef
	(types_go_proto.DataType)(0),                   // 4: tensorflow.DataType
}
var file_tensorflow_core_framework_optimized_function_graph_proto_depIdxs = []int32{
	3, // 0: tensorflow.OptimizedFunctionGraph.function_graph:type_name -> tensorflow.GraphDef
	2, // 1: tensorflow.OptimizedFunctionGraph.node_name_to_control_ret:type_name -> tensorflow.OptimizedFunctionGraph.NodeNameToControlRetEntry
	4, // 2: tensorflow.OptimizedFunctionGraph.ret_types:type_name -> tensorflow.DataType
	0, // 3: tensorflow.OptimizedFunctionGraph.source:type_name -> tensorflow.OptimizedFunctionGraph.OptimizationSource
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_optimized_function_graph_proto_init() }
func file_tensorflow_core_framework_optimized_function_graph_proto_init() {
	if File_tensorflow_core_framework_optimized_function_graph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_optimized_function_graph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptimizedFunctionGraph); i {
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
	file_tensorflow_core_framework_optimized_function_graph_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_framework_optimized_function_graph_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_optimized_function_graph_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_optimized_function_graph_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_framework_optimized_function_graph_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_framework_optimized_function_graph_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_optimized_function_graph_proto = out.File
	file_tensorflow_core_framework_optimized_function_graph_proto_rawDesc = nil
	file_tensorflow_core_framework_optimized_function_graph_proto_goTypes = nil
	file_tensorflow_core_framework_optimized_function_graph_proto_depIdxs = nil
}
