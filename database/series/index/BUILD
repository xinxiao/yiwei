load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "index",
    srcs = glob(["*.go"]),
    importpath = "yiwei/database/series/index",
    deps = [
        "//proto:entity_go_proto",
    ],
)
