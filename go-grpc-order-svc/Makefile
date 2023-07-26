proto:
	protoc pkg/pb/order.proto --go_out=plugins=grpc:.
	protoc pkg/pb/product.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go