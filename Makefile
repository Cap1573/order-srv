
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	sudo docker run --rm -v $(PWD):$(PWD) -w $(PWD) cap1573/protoc -I ./ --go_out=./ --micro_out=./  ./proto/order/*.proto

.PHONY: build
build:
	go build -o order-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t order-srv:latest
