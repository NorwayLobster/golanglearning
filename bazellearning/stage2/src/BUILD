load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library")

cc_library(
    name = "hello-greet",
    srcs = ["hello-greet.cc"],
    hdrs = ["hello-greet.h"],
)

cc_binary(
    name = "hello-world",
    srcs = ["hello-world.cc"],
    deps = [
        "@//src:hello-greet",#complete label
        # "//src:hello-greet",
        # ":hello-greet",
        # "hello-greet",//the shortest one
    ],
)
