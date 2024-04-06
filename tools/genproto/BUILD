package(default_visibility = ["//visibility:private"])

genrule(
    name = "go_protos",
    srcs = [
        "//tools/lib_package:libtensorflow_proto",
        "@com_google_protobuf//:well_known_type_protos",
    ],
    outs = ["go_protos.tar"],
    cmd = """
set -euo pipefail

# Create a temporary work directory.
TMP=$$(mktemp -d)
trap "rm -rf $${TMP}" EXIT

# Untar well-known protobufs.
mkdir -p $${TMP}/external/com_google_protobuf
tar cf - $(locations @com_google_protobuf//:well_known_type_protos) \
  | tar x --directory=$${TMP}/external/com_google_protobuf
WELL_KNOWN_PROTO_ROOT=$$(find $${TMP}/external/com_google_protobuf/ -name google -type d | head -n 1)

# Untar TensorFlow core protobufs.
mkdir -p $${TMP}/external/libtensorflow_proto
tar xf $(location //tools/lib_package:libtensorflow_proto) \
  --directory=$${TMP}/external/libtensorflow_proto
TF_PROTO_ROOT=$$(find $${TMP}/external/libtensorflow_proto/ -name tensorflow -type d | head -n 1)

# Generated Go code will be written to tensorflow/src.
mkdir -p $${TMP}/tensorflow/src

# Generate.
find $${TMP}/external/libtensorflow_proto -name "*.proto" \
  | xargs $(execpath @com_google_protobuf//:protoc) \
    --go_out=$${TMP}/tensorflow/src \
    --plugin=protoc-gen-go=$(execpath @org_golang_google_protobuf//cmd/protoc-gen-go) \
    --proto_path=$$(dirname $${WELL_KNOWN_PROTO_ROOT}) \
    --proto_path=$$(dirname $${TF_PROTO_ROOT})

# Tar generated Go code.
tar cf $@ --mtime='UTC 1970-01-01' -C $${TMP}/tensorflow src
""",
    tools = [
        "@com_google_protobuf//:protoc",
        "@org_golang_google_protobuf//cmd/protoc-gen-go",
    ],
    visibility = ["//tools:__subpackages__"],
)
