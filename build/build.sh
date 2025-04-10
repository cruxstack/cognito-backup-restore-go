#!/usr/bin/env bash

RELEASE_VERSION="${1:-v0.0.0}"
PLATFORMS="linux:amd64 linux:arm64 darwin:amd64 darwin:arm64 windows:amd64"

mkdir -p dist
for platform in ${PLATFORMS}; do
  GOOS=${platform%:*}
  GOARCH=${platform#*:}
  base_name="cbr_${RELEASE_VERSION}_${GOOS}_${GOARCH}"
  if [ "$GOOS" = "windows" ]; then
    base_name="${base_name}.exe"
  fi
  env GOOS=$GOOS GOARCH=$GOARCH go build -o "dist/${base_name}" .
done
