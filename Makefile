TARGET_PROTO_FILES = $(shell find clientpingkratos -name "*.proto")

.PHONY: proto
proto:
	protoc --proto_path=./clientpingkratos \
		   --proto_path=./proto3ps \
		   --go_out=paths=source_relative:./clientpingkratos \
		   --go-http_out=paths=source_relative:./clientpingkratos \
		   --go-grpc_out=paths=source_relative:./clientpingkratos \
		   --openapi_out=fq_schema_naming=true,default_response=false,title=PING-KRATOS:. \
		   $(TARGET_PROTO_FILES)
