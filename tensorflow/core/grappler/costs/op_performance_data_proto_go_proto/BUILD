load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "op_performance_data_proto_go_proto",
    srcs = ["op_performance_data.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/grappler/costs/op_performance_data_proto_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/attr_value_go_proto",
        "//tensorflow/core/framework/tensor_go_proto",
        "//tensorflow/core/framework/tensor_shape_go_proto",
        "//tensorflow/core/framework/types_go_proto",
        "//tensorflow/core/protobuf/for_core_protos_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

alias(
    name = "go_default_library",
    actual = ":op_performance_data_proto_go_proto",
    visibility = ["//visibility:public"],
)
