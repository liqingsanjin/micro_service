#!/usr/bin/env bash
mkdir pkg/pb
protoc -I ./api --go_out=plugins=grpc:./pkg/pb/ api/*.proto
GOOS=linux go build -ldflags="-s -w" -o userservice cmd/userservice/main.go