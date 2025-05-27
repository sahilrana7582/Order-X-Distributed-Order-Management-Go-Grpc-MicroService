package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "github.com/sahilrana7582/orderX/pkg/generated/user"
	"github.com/sahilrana7582/user-service/configs"
	"github.com/sahilrana7582/user-service/internal/app"
	"github.com/sahilrana7582/user-service/internal/db"
	"github.com/sahilrana7582/user-service/internal/repository/postgres"
	"github.com/sahilrana7582/user-service/internal/utils"
	"google.golang.org/grpc"
)

type Application struct {
	Config     *configs.Config
	LoggerInfo *log.Logger
	LoggerErr  *log.Logger
	DB         *sql.DB
	UserSrv    *app.UserService
}

func main() {

	infoLog, errLog := utils.InitLogger()

	cfg, err := configs.LoadConfig()
	if err != nil {
		errLog.Println("Error loading config:", err)
		log.Fatalf("failed to load config: %v", err)
	}

	err = db.Connect(
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.SSLMode,
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	address := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userRepo := postgres.NewUserRepository(db.DB)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, app.NewUserService(userRepo))

	infoLog.Printf("gRPC server listening on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
