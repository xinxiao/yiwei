load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "yiwei",
    srcs = [
        "logger.go",
        "main.go",
        "server.go",
    ],
    importpath = "yiwei/server",
    deps = [
        "//database",
        "//proto:entity_go_proto",
        "@com_github_grpc_ecosystem_go_grpc_middleware//:go_default_library",
        "@com_github_grpc_ecosystem_go_grpc_middleware//logging/zap:go_default_library",
        "@com_github_grpc_ecosystem_go_grpc_middleware//recovery:go_default_library",
        "@com_github_grpc_ecosystem_go_grpc_prometheus//:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//reflection:go_default_library",
        "@org_uber_go_zap//:go_default_library",
    ],
)
