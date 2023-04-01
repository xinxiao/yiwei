load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "filter",
    srcs = [
        "parse.go",
    ],
    importpath = "yiwei/database/label/filter",
    deps = [
        "//database/label/filter/impl",
        "//proto:entity_go_proto",
    ],
)
