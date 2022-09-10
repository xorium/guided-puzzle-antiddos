#!/bin/env bash
set -e

export GOOS=${GOOS:-linux}
export CGO_ENABLED=${CGO_ENABLED:-0}
export GO111MODULE=${GO111MODULE:-on}

go mod tidy -v

ldflags="-X 'build.Version=${GIT_TAG}' -X 'build.Date=${BUILD_DATE}' -X 'build.GIT_SHA=${GIT_COMMIT}'"

for app_path in "${TARGETS_PATH}/*/main.go"; do
  dir_path=$(dirname $app_path)
  bin_name=$(basename $dir_path)

  echo "building $BIN_DIR/${TARGET} from $dir_path/*.go"
  go build -tags="${BUILD_TAGS}" -a -o "$BIN_DIR/$bin_name" -ldflags "$ldflags" "$dir_path"/*.go
done;