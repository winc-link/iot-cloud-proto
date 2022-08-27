.PHONY: build clean test docker run


gen-product:
	protoc  -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./product/product.proto
