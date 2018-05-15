# 適当に設定を書く
# NAME	 := ecs-helper
# VERSION  := 0.1.0
# ...

BINARY_NAME=ecs-helper

all: setup build

cross: setup docker-cross-build

docker: build-docker

setup:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
build:
	go build -a -o $(BINARY_NAME)
docker-cross-build:
	for os in darwin linux; do \
		docker run --rm -it -e CGO_ENABLED=0 -e GOOS=$$os -e GOARCH=amd64 -v "$(PWD):/go/src/ecs-helper" -w /go/src/ecs-helper golang:latest go build -o bin/$(BINARY_NAME); \
		tar -czvf bin/$(BINARY_NAME)-$$os-amd64.tar.gz -C bin $(BINARY_NAME); \
		rm bin/$(BINARY_NAME); \
	done
build-docker:
	docker build -t $(IMAGE_TAG) .
