load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "page",
    srcs = [
        "page.go",
        "persist.go",
        "read.go",
        "write.go",
    ],
    importpath = "yiwei/database/page",
    deps = [
        "//database/persistence",
        "//proto:entity_go_proto",
        "@com_github_google_uuid//:go_default_library",
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)
