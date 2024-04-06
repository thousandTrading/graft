load("@gazelle//:def.bzl", "gazelle")

package(
    default_visibility = ["//visibility:private"],
)

# gazelle:prefix github.com/wamuir/graft
# gazelle:build_file_name BUILD
# gazelle:exclude tensorflow/op/wrappers.go
# gazelle:exclude tools/graft_package
# gazelle:exclude tools/lib_package
# gazelle:go_naming_convention import_alias
gazelle(name = "gazelle")
