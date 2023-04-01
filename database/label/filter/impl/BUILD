load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "impl",
    srcs = glob(["*.go"]),
    importpath = "yiwei/database/label/filter/impl",
    deps = [
        "//database/label",
        "@org_golang_x_exp//slices:go_default_library",
    ],
)
