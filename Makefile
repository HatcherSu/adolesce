PROJECT_NAME=cloud_callback
PROJECT_PATH=${shell pwd}
GOPATH=$(shell go env GOPATH)
API_PROTO_FILES=$(shell find api -name *.proto)
INTERNAL_PROTO_FILES=${shell find internal -name *.proto}


.PHONY:	init
# init env
init:
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u github.com/google/wire/cmd/wire
	go get -u github.com/gin-gonic/gin
	go get -u github.com/lestrrat/go-file-rotatelogs
	go get -u github.com/go-redis/redis/v8

## todo make test\generate\help

# generate wire
.PHONY: wire
wire:
	@cd cmd && go run github.com/google/wire/cmd/wire

.PHONY: build-app
build-app:
	@mkdir -p bin/ && go build -o  ./bin/cloud_callback ./cmd/api-server/

.PHONY: build-timer
build-timer:
	@mkdir -p bin/ && go build -o  ./bin/app_timer ./cmd/app-timer/

.PHONY: run-app
run-app:
	./bin/cloud_callback

.PHONY: run-timer
run-timer:
	./bin/app_timer

.PHONY: all
all:
	make init
	make wire
	make build
	make run

