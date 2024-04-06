load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "graph_transfer_info_go_proto",
    srcs = ["graph_transfer_info.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/graph_transfer_info_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/types_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

alias(
    name = "go_default_library",
    actual = ":graph_transfer_info_go_proto",
    visibility = ["//visibility:public"],
)
