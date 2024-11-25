#!/bin/bash

PROTO_DIR="."
GO_OUT_DIR="."
PYTHON_OUT_DIR="./python/textgen"

mkdir -p ${PYTHON_OUT_DIR}

echo "Generating Go files..."
protoc --proto_path=${PROTO_DIR} \
    --go_out=${GO_OUT_DIR} --go-grpc_out=${GO_OUT_DIR} \
    service.proto

echo "Generating Python files..."
python3 -m grpc_tools.protoc -I=${PROTO_DIR} \
    --python_out=${PYTHON_OUT_DIR} --grpc_python_out=${PYTHON_OUT_DIR} \
    service.proto

echo "Generation complete!"