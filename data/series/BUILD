load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "series",
    srcs = [
        "persistence.go",
        "series.go",
    ],
    importpath = "yiwei/data/series",
    deps = [
        "//data/persistence",
        "//data/page",
        "//proto:data_go_proto",
        "//data/series/index",
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)
