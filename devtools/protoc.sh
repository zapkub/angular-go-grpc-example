#!/bin/bash
set -e
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
workspacedir=$(pwd)

    
source $DIR/set_dev_env.sh
cd proto && \
protoc \
    -I . \
    -I $DIR/.tmp/googleapis \
    -I $DIR/.tmp/grpc-gateway-2.3.0 \
    -I $DIR/.tmp/grpc-gateway-2.3.0/protoc-gen-openapiv2/options \
    --go_out ../pkg/apis/ --go_opt paths=source_relative \
    --go-grpc_out ../pkg/apis --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ../pkg/apis --grpc-gateway_opt paths=source_relative \
    --openapiv2_out ../pkg/apis \
    ./bookstore.proto



sudo docker run -it --rm -v $workspacedir:/local swaggerapi/swagger-codegen-cli:2.4.18 generate \
    -i /local/pkg/apis/bookstore.swagger.json \
    -l typescript-angular \
    -o /local/web/mainsite/src/apis \
    -c /local/pkg/apis/swagger-options.json
sudo chown -R $(whoami) .