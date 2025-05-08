#!/bin/bash

set -e

echo "🔍 Searching for wire.go files..."

# Temukan semua file yang cocok
WIRE_FILES=$(find . -type f -name "wire.go")

if [ -z "$WIRE_FILES" ]; then
  echo "❌ No wire.go file found."
  exit 1
fi

for file in $WIRE_FILES; do
  DIR=$(dirname "$file")
  echo "⚙️ Running wire in $DIR ..."
  (cd "$DIR" && wire)
done

echo "✅ Done generating wire_gen.go for all wire.go files."
