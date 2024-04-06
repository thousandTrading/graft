load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "reader_base_go_proto",
    srcs = ["reader_base.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/reader_base_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

alias(
    name = "go_default_library",
    actual = ":reader_base_go_proto",
    visibility = ["//visibility:public"],
)
