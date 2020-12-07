#!/bin/bash

GOOS=linux GOARCH=amd64 go build -tags netgo -gcflags "-N -l" struct.go
