package(
    default_visibility = ["//visibility:private"],
)

ALL_PROTOS = glob(["**/*.proto"])

filegroup(
    name = "protos_all_proto_srcs",
    srcs = ALL_PROTOS,
    visibility = ["@graft//third_party/org_tensorflow:__pkg__"],
)
