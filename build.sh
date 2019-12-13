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

pushd src
go build -ldflags "-s -w" -o "../build/$OUT_FILE"
popd

# upx
if command -v upx; then
        ! upx "build/$OUT_FILE"
else
        echo "UPX not installed, compression skipped"
fi

ls -lh "build/$OUT_FILE"

# set exit code even if the previous command fails
exit 0
