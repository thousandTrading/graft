load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "event_go_proto",
    srcs = ["event.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/util/event_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/core/framework/summary_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

alias(
    name = "go_default_library",
    actual = ":event_go_proto",
    visibility = ["//visibility:public"],
)
