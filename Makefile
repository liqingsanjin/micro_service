.PHONY: build

build:
	protoc -I ./api --go_out=plugins=grpc:./pkg/pb/ api/*.proto