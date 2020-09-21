#!/bin/bash

GOOS=linux GOARCH=amd64 go tool compile -S x.go        
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "-N -l" interface-pointer-test.go
#go build -gcflags -S x.go
