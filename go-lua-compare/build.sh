#!/bin/bash

set -x

GOGC=off CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -gcflags "-N -l" go-lua-compare.go
go build -gcflags "-N -l" lua_demo.go
#cp -f ./go-lua-compare ~/Desktop

#curl http://30.27.153.142:2020/tiger_test  -H "Host: www.tiger.com" -H "Cookie: test" -H "Connection: closed" -H "UID: 202021"

