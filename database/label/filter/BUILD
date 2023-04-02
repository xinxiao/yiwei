load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "filter",
    srcs = [
        "filter.go",
        "parse.go",
    ],
    importpath = "yiwei/database/label/filter",
    deps = [
        "//database/label",
        "//database/label/filter/impl",
        "//proto:entity_go_proto",
    ],
)
