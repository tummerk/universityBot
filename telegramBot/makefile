PROTOC          = protoc
PROTOC_GEN_GO   = protoc-gen-go
PROTOC_GEN_GRPC = protoc-gen-go-grpc

API_DIR         = api
OUT_DIR         = pkg/api

generate:
	@echo ">> Generating gRPC code"
	$(PROTOC) \
	  --proto_path=$(API_DIR) \
	  --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
	  --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
	  $(shell find $(API_DIR) -name "*.proto")