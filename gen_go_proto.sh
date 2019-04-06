#!/bin/sh

protoc -I kvrpc/ kvrpc/kvrpc.proto --go_out=plugins=grpc:kvrpc
