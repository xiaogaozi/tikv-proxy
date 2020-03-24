#!/usr/bin/env bash

mkdir -p pkg/serverpb
protoc -I ./proto --go_out=plugins=grpc:pkg/serverpb ./proto/serverpb.proto
