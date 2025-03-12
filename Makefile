PROTO_DIRS := $(wildcard proto*)
PROTO_FILES := $(foreach dir,$(PROTO_DIRS),$(wildcard $(dir)/*.proto))
OUT_DIR := .

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
	go install github.com/mitchellh/protoc-gen-go-json@latest

proto:
	buf dep update
	buf generate