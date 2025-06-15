package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	pb "github.com/sahilrana7582/orderX/pkg/generated/product"
	"github.com/sahilrana7582/product-service/configs"
	"github.com/sahilrana7582/product-service/internal/db"
	"github.com/sahilrana7582/product-service/internal/handler"
	"github.com/sahilrana7582/product-service/internal/repository/impl"
	"github.com/sahilrana7582/product-service/internal/utils"
	"google.golang.org/grpc"
)

type Application struct {
	Config      *configs.Config
	LoggerInfo  *log.Logger
	LoggerError *log.Logger
	DB          *sql.DB
}

var App *Application

func main() {
	infoLog, errLog := utils.InitLogger()

	cfg, err := configs.LoadConfig()
	if err != nil {
		errLog.Fatalf("could not load config: %v", err)
	}
	conn, err := db.InitDB(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode, infoLog)
	if err != nil {
		errLog.Println("Error connecting to DB:", err)
	}

	App = &Application{
		Config:      cfg,
		LoggerInfo:  infoLog,
		LoggerError: errLog,
		DB:          conn,
	}

	repo := impl.NewProductRepository(conn)

	productHandler := handler.NewProductHandler(repo)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, productHandler)

	infoLog.Println("gRPC server is running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
