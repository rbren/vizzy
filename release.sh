#! /bin/bash
set -eo pipefail

docker buildx create --name vizzybuilder --use || true
docker buildx build --platform linux/amd64,linux/arm64 -t ghcr.io/rbren/vizzy --push .
