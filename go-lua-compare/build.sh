#!/bin/bash

set -x

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -gcflags "-N -l" go-lua-compare.go
cp -f ./go-lua-compare ~/Desktop

