#!/usr/bin/env sh

set -eu

if ! command -v atlas >/dev/null 2>&1; then
  echo "atlas CLI is required. Install it from https://atlasgo.io/getting-started" >&2
  exit 1
fi

if [ -f ".env" ]; then
  set -a
  . ./.env
  set +a
fi

if [ "${DATABASE_URL:-}" = "" ]; then
  echo "DATABASE_URL is required" >&2
  exit 1
fi

atlas migrate apply \
  --dir "file://internal/infrastructure/persistence/migrations" \
  --url "$DATABASE_URL"
