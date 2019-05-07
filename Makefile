.PHONY: build, dockerBuild, start, build

build:
	protoc -I ./api --go_out=plugins=grpc:./pkg/pb/ api/*.proto
	bash build/ci/compile_staticservice.sh
	bash build/ci/compile_institutionservice.sh
	docker build -t staticservice:1.0.0 --target staticservice -f build/deploy/Dockerfile_static .
	docker build -t institutionservice:1.0.0 --target institutionservice -f build/deploy/Dockerfile_institution .

start:
	docker-compose -f deployments/testDocker-compose.yml --project-directory ./ up

clear:
	docker-compose -f deployments/testDocker-compose.yml --project-directory ./ down
