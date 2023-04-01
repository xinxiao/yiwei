load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

proto_library(
    name = "label_proto",
    srcs = [
        "label.proto",
    ],
)

proto_library(
    name = "entry_proto",
    srcs = [
        "entry.proto",
    ],
    deps = [
        ":label_proto",
    ],
)

proto_library(
    name = "page_proto",
    srcs = [
        "page.proto",
    ],
    deps = [
        ":entry_proto",
    ],
)

proto_library(
    name = "series_proto",
    srcs = [
        "series.proto",
    ],
    deps = [
        ":label_proto",
    ],
)

go_proto_library(
    name = "data_go_proto",
    importpath = "yiwei/proto",
    protos = [
        ":entry_proto",
        ":label_proto",
        ":page_proto",
        ":series_proto",
    ],
)
