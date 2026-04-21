#!/usr/bin/env sh

set -eu

if ! command -v atlas >/dev/null 2>&1; then
  echo "atlas CLI is required. Install it from https://atlasgo.io/getting-started" >&2
  exit 1
fi

if [ "$#" -lt 1 ]; then
  echo "usage: pnpm --filter api run migrate:diff -- <migration_name>" >&2
  exit 1
fi

MIGRATION_NAME="$1"
ATLAS_DEV_URL="${ATLAS_DEV_URL:-docker://postgres/15/dev?search_path=public}"

atlas migrate diff "$MIGRATION_NAME" \
  --dir "file://internal/infrastructure/persistence/migrations" \
  --to "ent://internal/infrastructure/persistence/ent/schema" \
  --dev-url "$ATLAS_DEV_URL"
