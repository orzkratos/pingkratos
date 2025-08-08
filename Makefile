TARGET_PROTO_FILES=$(shell find pingpong -name *.proto)

.PHONY: ping-pong
# generate internal proto
ping-pong:
	protoc --proto_path=./pingpong \
		   --proto_path=./proto3ps \
		   --go_out=paths=source_relative:./pingpong \
		   --go-http_out=paths=source_relative:./pingpong \
		   --go-grpc_out=paths=source_relative:./pingpong \
		   --openapi_out=fq_schema_naming=true,default_response=false,title=DEMO-TITLE:. \
		   $(TARGET_PROTO_FILES)
