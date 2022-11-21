#! /usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")/.."
TARGET_PATH="target"
mkdir -p "$TARGET_PATH"
PATH="$TARGET_PATH:$PATH"

go build -mod readonly -o "$TARGET_PATH"/ \
  google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
  github.com/grpc-ecosystem/grpc-health-probe

GO_PKG="lipsumgo"
PROTO_PATH="proto"
PROTO_PKG_PATH="proto/lipsumgo"

rm -f -- pkg/pb/*.pb.go pkg/pb/*.pb.gw.go
protoc \
  -I "$PROTO_PATH" \
  --go_out=. --go_opt=module="$GO_PKG" \
  --go-grpc_out=. --go-grpc_opt=module="$GO_PKG" \
  --grpc-gateway_out=. --grpc-gateway_opt=module="$GO_PKG" \
  --grpc-gateway_opt logtostderr=true \
  "$PROTO_PKG_PATH"/*.proto
