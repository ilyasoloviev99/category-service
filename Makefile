.PHONY: generate
generate:
	protoc --go_out=pkg api/category/category.proto \
        --go-grpc_out=pkg api/category/category.proto \
        api/category/category.proto

.PHONY: run
run:
	go run cmd/category_service/main.go
