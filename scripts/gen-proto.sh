#!/bin/bash

CURRENT_DIR=$(pwd)

rm -rf ./genproto/*

for module in $(find $CURRENT_DIR/protos/* -type d); do
    protoc -I $CURRENT_DIR/protos/ \
           --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/ \
            $module/*.proto;

done;
