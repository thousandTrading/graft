load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@graft//tools/build_defs/repo:http.bzl", "http_embedded_archive")
load("@libtensorflow_defaults//:config.bzl", "LIBTENSORFLOW_PKG_URL")

def tf_repositories(ctx):
    ###########################################################################
    # libtensorflow
    ###########################################################################

    http_archive(
        name = "libtensorflow_linux_x86_64_cpu",
        build_file = "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/latest/libtensorflow-cpu-linux-x86_64.tar.gz"],
    )

    http_archive(
        name = "libtensorflow_linux_x86_64_gpu",
        build_file = "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/latest/libtensorflow-gpu-linux-x86_64.tar.gz"],
    )

    http_embedded_archive(
        name = "libtensorflow_macos_x86_64_cpu",
        build_file = "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/prod/tensorflow/release/macos/latest/macos_cpu_libtensorflow_binaries.tar.gz"],
        inner_archive = "lib_package/libtensorflow-cpu-darwin-x86_64.tar.gz",
    )

    http_embedded_archive(
        name = "libtensorflow_macos_arm_64_cpu",
        build_file = "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = ["https://storage.googleapis.com/libtensorflow-nightly/prod/tensorflow/release/macos/latest/macos_cpu_libtensorflow_binaries.tar.gz"],
        inner_archive = "lib_package/libtensorflow-cpu-darwin-arm64.tar.gz",
    )

    http_archive(
        name = "libtensorflow_other_build",
        build_file = "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = [LIBTENSORFLOW_PKG_URL],
    )

    ###########################################################################
    # protos
    ###########################################################################

    http_archive(
        name = "libtensorflow_proto",
        build_file = "@graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD",
        patch_args = ["-p1"],
        patches = [
            "@graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch",
            "@graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto_tsl.patch",
        ],
        strip_prefix = "libtensorflow-proto-nightly",
        urls = ["https://github.com/wamuir/libtensorflow-proto/archive/refs/heads/nightly.tar.gz"],
    )

    ###########################################################################
    # base api def
    ###########################################################################

    http_archive(
        name = "tensorflow_base_api_def",
        build_file = "@graft//third_party/org_tensorflow/tensorflow_base_api_def:tensorflow_base_api_def.BUILD",
        strip_prefix = "tensorflow-nightly/tensorflow/core/api_def/base_api",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/heads/nightly.tar.gz"],
    )

    ###########################################################################
    # go bindings
    ###########################################################################

    http_archive(
        name = "tensorflow_go",
        build_file = "@graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go.BUILD",
        strip_prefix = "tensorflow-nightly/tensorflow/go",
        urls = ["https://github.com/tensorflow/tensorflow/archive/refs/heads/nightly.tar.gz"],
        patch_args = ["-p1"],
        patches = ["@graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go_op.patch"],
    )

download_tf_repositories = module_extension(
    implementation = tf_repositories,
)
