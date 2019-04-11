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
go build -ldflags='-s -w' -o userService cmd/main.go
```
### 编译成 docker 镜像
```
```
