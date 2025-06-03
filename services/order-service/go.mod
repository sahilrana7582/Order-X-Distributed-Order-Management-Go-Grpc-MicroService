module github.com/sahilrana7582/order-service

go 1.23.3

require (
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
)

require (
	github.com/sahilrana7582/orderX/pkg/generated/order v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/grpc v1.72.2 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/sahilrana7582/orderX/pkg/generated/order => ../../pkg/generated/order
