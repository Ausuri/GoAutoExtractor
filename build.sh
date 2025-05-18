#!/bin/bash

set -e

if [ ! -d "./bin" ]; then
  mkdir bin
fi

OUTPUT="./bin/MediaCompressionManager"

echo "Cleaning up old binary..."
rm -f "$OUTPUT"

echo "Building $OUTPUT at $(date +"%Y-%m-%d %H:%M:%S")..."
go build -o "$OUTPUT" .

echo "Build finished! ✅"
chmod +x $OUTPUT

if [ "$1" == "run" ]; then
    echo "Running $OUTPUT..."
    "$OUTPUT"
fi