load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_proto_grpc//go:defs.bzl", "go_grpc_library")

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

proto_library(
    name = "database_proto",
    srcs = [
        "database.proto",
    ],
    deps = [
        ":label_proto",
        ":series_proto",
    ],
)

go_grpc_library(
    name = "entity_go_proto",
    importpath = "yiwei/proto",
    protos = [
        ":database_proto",
        ":entry_proto",
        ":label_proto",
        ":page_proto",
        ":series_proto",
    ],
)
