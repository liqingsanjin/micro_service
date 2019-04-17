#!/usr/bin/env bash
protoc -I ./api --go_out=plugins=grpc:./pkg/pb/ api/user.proto
GOOS=linux go build -o userservice cmd/userservice/main.go