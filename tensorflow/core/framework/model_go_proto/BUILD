load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "model_go_proto",
    srcs = ["model.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/framework/model_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

alias(
    name = "go_default_library",
    actual = ":model_go_proto",
    visibility = ["//visibility:public"],
)
