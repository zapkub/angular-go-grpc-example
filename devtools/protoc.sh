#!/bin/bash
set -e
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

    
source $DIR/set_dev_env.sh
cd proto && \
protoc \
    -I . \
    -I $DIR/.tmp/grpc-gateway-2.3.0 \
    -I $DIR/.tmp/grpc-gateway-2.3.0/protoc-gen-openapiv2/options \
    --go_out ../pkg/apis/ --go_opt paths=source_relative \
    --go-grpc_out ../pkg/apis --go-grpc_opt paths=source_relative \
    ./bookstore.proto