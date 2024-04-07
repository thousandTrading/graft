load("@rules_pkg//:pkg.bzl", "pkg_tar")

package(
    default_visibility = ["//visibility:private"],
)

GO_SOURCE = glob(
    ["**/*"],
    exclude = [
        #"genop/**",
        #"op/generate.go",
        #"op/wrappers.go",
        "BUILD.bazel",
        "WORKSPACE",
    ],
) + ["//op:all_files"]

filegroup(
    name = "go_source",
    srcs = GO_SOURCE,
    visibility = ["@graft//third_party/org_tensorflow:__pkg__"],
)
