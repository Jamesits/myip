#!/bin/bash
set -Eeuo pipefail
cd "$( dirname "${BASH_SOURCE[0]}" )"

export GOPATH=${GOPATH:-/tmp/gopath}
OUT_FILE=${OUT_FILE:-myip}

mkdir -p build
! mkdir -p "$GOPATH"

export GO111MODULE=on
go mod download
go mod verify

cd src
go build -ldflags "-s -w" -o "../build/$OUT_FILE"
