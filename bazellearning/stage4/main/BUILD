cc_library(
    name ="hellotime",
    srcs=["hellotime.cc"],
    hdrs=["hellotime.hpp"],
)
cc_binary(
    name ="hello-world",
    srcs=["helloworld.cc"],
    # hdrs=["helloworld.hpp"],
    deps=[
        ":hellotime",
        "//third_party/lib:hellogreet",
    ],
)
