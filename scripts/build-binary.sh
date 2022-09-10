#!/bin/env bash
set -e

export GOOS=${GOOS:-linux}
export CGO_ENABLED=${CGO_ENABLED:-0}
export GO111MODULE=${GO111MODULE:-on}

go mod tidy -v

ldflags="-X 'build.Version=${GIT_TAG}' -X 'build.Date=${BUILD_DATE}' -X 'build.GIT_SHA=${GIT_COMMIT}'"

for app_name in $(ls ${TARGETS_PATH}); do
  relative_bin_dir=$(basename $BIN_DIR)
  relative_target_dir=$(basename ${TARGETS_PATH})

  echo "building ./$relative_bin_dir/$app_name from ./$relative_target_dir/$app_name/*.go"
  go build \
    -tags="${BUILD_TAGS}" -a -o "$BIN_DIR/$apo_name" \
    -ldflags "$ldflags" "${TARGETS_PATH}/$app_name"/*.go
done;