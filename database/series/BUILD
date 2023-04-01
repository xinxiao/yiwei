load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "series",
    srcs = [
        "persist.go",
        "read.go",
        "series.go",
        "write.go",
    ],
    importpath = "yiwei/database/series",
    deps = [
        "//database/label",
        "//database/page",
        "//database/persistence",
        "//database/series/index",
        "//proto:entity_go_proto",
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)
