load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "database",
    srcs = [
        "database.go",
    ],
    importpath = "yiwei/database",
    deps = [
        "//database/label/filter",
        "//database/persistence",
        "//database/series",
        "//proto:entity_go_proto",
        "@org_golang_google_grpc//codes:go_default_library",
        "@org_golang_google_grpc//status:go_default_library",
        "@org_uber_go_zap//:go_default_library",
    ],
)
