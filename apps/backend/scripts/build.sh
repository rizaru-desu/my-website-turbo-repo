#!/usr/bin/env sh

set -eu

HOST_GOOS="$(go env GOHOSTOS)"
HOST_GOARCH="$(go env GOHOSTARCH)"
TARGET_GOOS="${GOOS:-$HOST_GOOS}"
TARGET_GOARCH="${GOARCH:-$HOST_GOARCH}"

OUTPUT_DIR="${OUTPUT_DIR:-bin}"
OUTPUT_NAME="${OUTPUT_NAME:-server}"

if [ "$TARGET_GOOS" != "$HOST_GOOS" ] || [ "$TARGET_GOARCH" != "$HOST_GOARCH" ]; then
  OUTPUT_NAME="server-${TARGET_GOOS}-${TARGET_GOARCH}"
fi

mkdir -p "$OUTPUT_DIR"

go build -o "$OUTPUT_DIR/$OUTPUT_NAME" ./cmd/api/main.go
