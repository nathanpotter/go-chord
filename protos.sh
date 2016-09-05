#!/bin/bash

# This script builds all of the protobuf/grpc files
# need to figure out how to make import path correct for *pb.go files
# being imported as "common", need to be "github.com/nathanpotter/go-chord/protos/common"

protoc -I protos protos/common/common.proto --go_out=plugins=grpc:protos
protoc -I protos protos/node/node.proto --go_out=plugins=grpc:protos
protoc -I protos protos/supernode/supernode.proto --go_out=plugins=grpc:protos
