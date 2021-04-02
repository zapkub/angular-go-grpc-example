#!/bin/bash
set -e
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

mkdir -p $DIR/.tmp/bin

protoc_url=https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-linux-x86_64.zip
gen_grpc_gateway_url=https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v2.3.0/protoc-gen-grpc-gateway-v2.3.0-linux-x86_64
gen_openapiv2_url=https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v2.3.0/protoc-gen-openapiv2-v2.3.0-linux-x86_64
gen_go=https://github.com/protocolbuffers/protobuf-go/releases/download/v1.26.0/protoc-gen-go.v1.26.0.linux.amd64.tar.gz
grpc_gateway_source=https://github.com/grpc-ecosystem/grpc-gateway/archive/refs/tags/v2.3.0.tar.gz

wget -nc -P $DIR/.tmp $protoc_url
wget -nc -P $DIR/.tmp -O $DIR/.tmp/bin/protoc-gen-grpc-gateway $gen_grpc_gateway_url
wget -nc -P $DIR/.tmp -O $DIR/.tmp/bin/protoc-gen-openapiv2 $gen_openapiv2_url
wget -nc -P $DIR/.tmp -O $DIR/.tmp/grpc-gateway-v2.3.0.tar.gz $grpc_gateway_source
wget -nc -P $DIR/.tmp $gen_go

echo "Extracting...."
unzip $DIR/.tmp/protoc-3.15.6-linux-x86_64.zip -d $DIR/.tmp
tar -xzf $DIR/.tmp/protoc-gen-go.v1.26.0.linux.amd64.tar.gz -C $DIR/.tmp/bin
tar -xzf $DIR/.tmp/grpc-gateway-v2.3.0.tar.gz -C $DIR/.tmp
echo "Done!"