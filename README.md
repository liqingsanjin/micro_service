# 用户服务

## 使用
### 下载依赖包
需要设置代理
```
go get -v -u github.com/golang/protobuf/protoc-gen-go
go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

export GO111MODULE=auto
go mod tidy
```
### 编译
```
mkdir pkg/pb

protoc -I ./api --go_out=plugins=grpc:./pkg/pb/ api/*.proto

go build -ldflags='-s -w' -o userService cmd/userservice/main.go
```
### 编译成 docker 镜像
```
./build/ci/compile_userservice.sh

docker build -t userservice:1.0.0 -f build/deploy/Dockerfile_user .
```


# api gateway

## 使用
### 下载依赖包
需要设置代理
```
go get -v -u github.com/golang/protobuf/protoc-gen-go
go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

export GO111MODULE=auto
go mod tidy
```

### 编译
```
mkdir pkg/pb

protoc -I ./api -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./pkg/pb/ api/user.proto
 
protoc -I ./api -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:./pkg/pb/ api/user.proto

go build -ldflags='-s -w' -o apigateway cmd/apigateway/main.go
```
### 编译成 docker 镜像
```
./build/ci/compile_apigateway.sh

docker build -t apigateway:1.0.0 -f build/deploy/Dockerfile_gateway .
```

# 商户服务
### 编译
```
./build/ci/compile_merchantservice.sh
```
### 编译成 docker 镜像
```
./build/ci/compile_merchantservice.sh

docker build -t merchantservice:1.0.0 -f build/deploy/Dockerfile_merchant .
```
