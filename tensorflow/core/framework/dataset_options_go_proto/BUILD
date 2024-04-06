load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "dataset_options_go_proto",
    srcs = ["dataset_options.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/dataset_options_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/model_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

alias(
    name = "go_default_library",
    actual = ":dataset_options_go_proto",
    visibility = ["//visibility:public"],
)
