package(
    default_visibility = ["//visibility:private"],
)

ALL_PBTXT = glob(["**/*.pbtxt"])

filegroup(
    name = "all_pbtxt",
    srcs = ALL_PBTXT,
    visibility = ["@graft//third_party/org_tensorflow:__pkg__"],
)
