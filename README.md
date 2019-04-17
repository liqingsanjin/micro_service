# 用户服务

## 使用
### 下载依赖包
需要设置代理
```
export GO111MODULE=auto
go mod tidy
```
### 编译
```
protoc -I ./api --go_out=plugins=grpc:./pkg/pb/ api/user.proto

go build -ldflags='-s -w' -o userService cmd/userservice/main.go
```
### 编译成 docker 镜像
```
./build/ci/compile.sh

docker build -t userservice:1.0.0 -f build/deploy/Dockerfile .
```
