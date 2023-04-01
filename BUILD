load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary")

gazelle(
    name = "gazelle",
    args = [
        "-from_file=go.mod",
        "-to_macro=go_repositories.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_binary(
    name = "yiwei",
    srcs = [
        "yiwei.go",
    ],
    importpath = "yiwei",
    deps = [
        "//data/persistence",
        "//data/series",
    ],
)
