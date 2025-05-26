PROTO_SRC=./proto
PROTO_OUT=pkg/generated

.PHONY: proto-gen

proto-gen:
	protoc \
		--proto_path=$(PROTO_SRC) \
		--go_out=$(PROTO_OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative \
		$(PROTO_SRC)/user/user.proto
