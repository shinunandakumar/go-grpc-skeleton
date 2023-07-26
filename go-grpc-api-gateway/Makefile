proto:
	protoc pkg/auth/pb/auth.proto --go_out=plugins=grpc:.
	protoc pkg/order/pb/order.proto --go_out=plugins=grpc:.
	protoc pkg/product/pb/product.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go