#!/usr/bin/env bash
mkdri pkg/pb
protoc -I ./api --go_out=plugins=grpc:./pkg/pb/ api/user.proto
GOOS=linux go build -ldflags="-s -w" -o userservice cmd/userservice/main.go