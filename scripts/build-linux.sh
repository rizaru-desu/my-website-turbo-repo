#!/usr/bin/env sh

set -eu

TARGET_ARCH="${1:-}"

if [ -z "$TARGET_ARCH" ]; then
  echo "Usage: sh ./scripts/build-linux.sh <amd64|arm64>" >&2
  exit 1
fi

case "$TARGET_ARCH" in
  amd64|arm64)
    ;;
  *)
    echo "Unsupported Linux architecture: $TARGET_ARCH" >&2
    exit 1
    ;;
esac

pnpm --filter admin run build
pnpm --filter public run build

CGO_ENABLED=0 GOOS=linux GOARCH="$TARGET_ARCH" pnpm --filter api run build
