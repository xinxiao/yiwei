load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    args = [
        "-from_file=go.mod",
        "-to_macro=go_repositories.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
