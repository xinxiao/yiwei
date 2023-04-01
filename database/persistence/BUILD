load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "persistence",
    srcs = [
        "path.go",
        "proto.go",
    ],
    importpath = "yiwei/database/persistence",
    deps = [
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)
