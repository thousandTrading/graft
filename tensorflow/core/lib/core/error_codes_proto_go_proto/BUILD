load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "error_codes_proto_go_proto",
    srcs = ["error_codes.pb.go"],
    importpath = "github.com/wamuir/graft/tensorflow/core/lib/core/error_codes_proto_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//tensorflow/tsl/protobuf/for_core_protos_go_proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

alias(
    name = "go_default_library",
    actual = ":error_codes_proto_go_proto",
    visibility = ["//visibility:public"],
)
