# 相关命令
    ldd hellocpp: 
    ldd helloG
    nm hellocpp
    nm helloG
    GOROOT
    GOPATH
    go build
    go install: 1. 先build 2.然后把elf复制到$GOPATH/bin
    go run
    CGO_ENABLED="1": 
    默认情况下，Go的runtime环境变量CGO_ENABLED=1，即默认开始cgo，允许你在Go代码中调用C代码，Go的pre-compiled标准库的.a文件也是在这种情况下编译出来的
   [CGO](https://johng.cn/cgo-enabled-affect-go-static-compile/)
    CGO_CFLAGS
    GOOS=linux
    GOARCH=amd64
