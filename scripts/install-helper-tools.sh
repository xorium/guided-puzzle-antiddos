#!/bin/env bash
set -e

# Fix for the 3rd party tools and binaries dir path.
if [ -z "${BIN_DIR}" ]; then BIN_DIR=$(pwd)/bin; fi

mkdir -p $BIN_DIR

if [[ ! -f "$BIN_DIR"/golangci-lint ]]; then
  echo "[*] Installing golangci-lint..."
  GOBIN="$BIN_DIR" go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

if [[ ! -f "$BIN_DIR"/goimports ]]; then
  echo "[*] Installing goimports..."
  GOBIN="$BIN_DIR" go install golang.org/x/tools/cmd/goimports@latest
fi

if [[ ! -f "$BIN_DIR"/godotenv ]]; then
  echo "[*] Installing godotenv..."
  GOBIN="$BIN_DIR" go install github.com/joho/godotenv/cmd/godotenv@latest
fi

if [[ ! -f "$BIN_DIR"/gofumpt ]]; then
  echo "[*] Installing gofumpt..."
  GOBIN="$BIN_DIR" go install mvdan.cc/gofumpt@latest
fi

chmod +x "$BIN_DIR"/*