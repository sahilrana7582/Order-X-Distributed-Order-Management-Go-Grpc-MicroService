package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/sahilrana7582/product-service/configs"
	"github.com/sahilrana7582/product-service/internal/db"
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

	grpcServer := grpc.NewServer()

	infoLog.Println("Application started successfully")

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	infoLog.Println("gRPC server listening on port 50053")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}

}
