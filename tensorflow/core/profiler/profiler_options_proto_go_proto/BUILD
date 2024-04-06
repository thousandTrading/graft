load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "profiler_options_proto_go_proto",
    srcs = ["profiler_options.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/profiler/profiler_options_proto_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/tsl/profiler/protobuf/for_profiler_protos_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

alias(
    name = "go_default_library",
    actual = ":profiler_options_proto_go_proto",
    visibility = ["//visibility:public"],
)
