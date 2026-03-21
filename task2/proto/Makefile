.PHONY: install-tools
install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: generate
generate:
	mkdir -p pkg/api/v1
	protoc -I api/v1 \
		--go_out=pkg/api/v1 \
	 	--go_opt=paths=source_relative \
		--go-grpc_out=pkg/api/v1 \
		--go-grpc_opt=paths=source_relative \
		api/v1/service.proto

.PHONY: all
all: install-tools generate
