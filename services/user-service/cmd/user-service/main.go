package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sahilrana7582/user-service/configs"
	"github.com/sahilrana7582/user-service/internal/app"
	"google.golang.org/grpc"

	pb "github.com/sahilrana7582/orderX/pkg/generated/user"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	address := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, app.NewUserService())

	log.Printf("gRPC server listening on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
