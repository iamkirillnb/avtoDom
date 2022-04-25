#!/bin/bash

gen_proto() {
  echo "Generate golang code from Protobuf specs"
  # generate types and server interface in sub directories where exists ".url_proto"
  spec_file_mask="*.proto"
  spec_files=$(find ./url_proto -type f -path "${spec_file_mask}")
  for proto_file_rel_path in ${spec_files}; do
    protoc \
      --experimental_allow_proto3_optional \
      --go_out=. \
      --go_opt=paths=source_relative \
      --go-grpc_out=. \
      --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
        "${proto_file_rel_path}";
  done
}

gen_proto