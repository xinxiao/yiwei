load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "label",
    srcs = [
        "convert.go",
        "env.go",
    ],
    importpath = "yiwei/database/label",
    deps = [
        "//proto:entity_go_proto",
    ],
)
