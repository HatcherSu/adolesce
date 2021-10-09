# Build all by default, even if it's not first
.DEFAULT_GOAL := all
.PHONY: all
all: init wire build-app build-timer

# ==============================================================================
# Variable
PROJECT_NAME=adolesce
PROJECT_PATH=${shell pwd}
GOPATH=$(shell go env GOPATH)
API_PROTO_FILES=$(shell find api -name *.proto)
INTERNAL_PROTO_FILES=${shell find internal -name *.proto}

# ==============================================================================
# Includes

include scripts/make-rules/common.mk


# ==============================================================================
# Targets

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

.PHONY:	tidy
tidy:
	@$(GO) mod tidy

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
	./bin/api_server

.PHONY: run-timer
run-timer:
	./bin/app_timer



