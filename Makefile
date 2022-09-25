BINARY := protoc-gen-connect-dart

TIMESTAMP := $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")
COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := -ldflags "-X main.Timestamp=${TIMESTAMP} -X main.Commit=${COMMIT}"

all: install

install:
	go install ${LDFLAGS} github.com/skadero/protoc-gen-connect-dart

run: install
	mkdir -p example/dart_client/rpc
	protoc --proto_path=./example/proto  --dart_out=./example/dart_client/rpc --connect-dart_out=./example/dart_client/rpc ./example/proto/haberdasher/haberdasher.proto ./example/proto/pinger/ping.proto


example: install
	mkdir -p example/dart_client/rpc && mkdir -p example/go_server/rpc && \
	protoc --proto_path=./example/proto  --go_out=paths=source_relative:example/go_server/rpc --connect-go_out=paths=source_relative:example/go_server/rpc ./example/proto/haberdasher/haberdasher.proto && \
	protoc --proto_path=./example/proto  --go_out=paths=source_relative:example/go_server/rpc --connect-go_out=paths=source_relative:example/go_server/rpc	./example/proto/pinger/ping.proto && \
	protoc --proto_path=./example/proto  --dart_out=./example/dart_client/rpc --connect-dart_out=./example/dart_client/rpc ./example/proto/haberdasher/haberdasher.proto ./example/proto/pinger/ping.proto
