load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "log_memory_go_proto",
    srcs = ["log_memory.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/log_memory_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/tensor_description_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

alias(
    name = "go_default_library",
    actual = ":log_memory_go_proto",
    visibility = ["//visibility:public"],
)
