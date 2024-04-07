#!/bin/sh

set -ex

WORKSPACE="$(git rev-parse --show-toplevel)/third_party/org_tensorflow/workspace.bzl"
TF_VERSION=${1}

if [ -z "${TF_VERSION}" ]; then
    printf 'ERROR: version required as the first positional argument.\n' >&2
    exit 1
fi

http_archive() {
    local TMPFILE
    TMPFILE=$(mktemp)
    curl -fSLs "${6}" > "${TMPFILE}"
  
    local HASH
    HASH=$(sha256sum "${TMPFILE}" | cut -d ' ' -f 1)
 
    rm -f "${TMPFILE}"
 
    cat <<EOF >> "${WORKSPACE}"

    http_archive(
        name = "${1}",
        build_file = "${2}",
        patch_args = [${3:+\"${3}\"}],
        patches = ${4},
        strip_prefix = "${5}",
        urls = ["${6}"],
        sha256 = "${HASH}",
    )
EOF
}


cat <<EOF > "${WORKSPACE}"
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@libtensorflow_defaults//:config.bzl", "LIBTENSORFLOW_PKG_URL")

def tf_repositories(ctx):
EOF

cat <<EOF >> "${WORKSPACE}"
    ###########################################################################
    # libtensorflow
    ###########################################################################
EOF

http_archive "libtensorflow_linux_x86_64_cpu" \
    "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD" \
    "-p1" \
    "[]" \
    "" \
    "https://storage.googleapis.com/tensorflow/versions/${TF_VERSION}/libtensorflow-cpu-linux-x86_64.tar.gz"

http_archive "libtensorflow_linux_x86_64_gpu" \
    "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD" \
    "-p1" \
    "[]" \
    "" \
    "https://storage.googleapis.com/tensorflow/versions/${TF_VERSION}/libtensorflow-gpu-linux-x86_64.tar.gz"

http_archive "libtensorflow_macos_x86_64_cpu" \
    "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD" \
    "-p1" \
    "[]" \
    "" \
    "https://storage.googleapis.com/tensorflow/versions/${TF_VERSION}/libtensorflow-cpu-darwin-x86_64.tar.gz"

http_archive "libtensorflow_macos_arm_64_cpu" \
    "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD" \
    "-p1" \
    "[]" \
    "" \
    "https://storage.googleapis.com/tensorflow/versions/${TF_VERSION}/libtensorflow-cpu-darwin-arm64.tar.gz"

cat <<EOF >> "${WORKSPACE}"

    http_archive(
        name = "libtensorflow_other_build",
        build_file = "@graft//third_party/org_tensorflow/libtensorflow:libtensorflow.BUILD",
        urls = [LIBTENSORFLOW_PKG_URL],
    )

    ###########################################################################
    # protos
    ###########################################################################
EOF

http_archive "libtensorflow_proto" \
    "@graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.BUILD" \
    "-p1" \
    "[
            \"@graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto.patch\",
            \"@graft//third_party/org_tensorflow/libtensorflow_proto:libtensorflow_proto_tsl.patch\",
        ]" \
    "libtensorflow-proto-${TF_VERSION}" \
    "https://github.com/wamuir/libtensorflow-proto/archive/refs/tags/v${TF_VERSION}.tar.gz"

cat <<EOF >> "${WORKSPACE}"

    ###########################################################################
    # base api def
    ###########################################################################
EOF

http_archive "tensorflow_base_api_def" \
    "@graft//third_party/org_tensorflow/tensorflow_base_api_def:tensorflow_base_api_def.BUILD" \
    "-p1" \
    "[]" \
    "tensorflow-${TF_VERSION}/tensorflow/core/api_def/base_api" \
    "https://github.com/tensorflow/tensorflow/archive/refs/tags/v${TF_VERSION}.tar.gz"

cat <<EOF >> "${WORKSPACE}"

    ###########################################################################
    # go bindings
    ###########################################################################
EOF

http_archive "tensorflow_go" \
    "@graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go.BUILD" \
    "-p1" \
    "[
            \"@graft//third_party/org_tensorflow/tensorflow_go:tensorflow_go_op.patch\",
        ]" \
    "tensorflow-${TF_VERSION}/tensorflow/go" \
    "https://github.com/tensorflow/tensorflow/archive/refs/tags/v${TF_VERSION}.tar.gz"

cat <<EOF >> "${WORKSPACE}"

download_tf_repositories = module_extension(
    implementation = tf_repositories,
)
EOF
