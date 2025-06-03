package main

import (
	"database/sql"
	"log"
	"net"
	"sync"

	_ "github.com/lib/pq"
	"github.com/sahilrana7582/order-service/configs"
	"github.com/sahilrana7582/order-service/internal/app"
	"github.com/sahilrana7582/order-service/internal/db"
	"github.com/sahilrana7582/order-service/internal/utils"
	pb "github.com/sahilrana7582/orderX/pkg/generated/order"
	"google.golang.org/grpc"
)

type Application struct {
	Config      *configs.Config
	LoggerInfo  *log.Logger
	LoggerError *log.Logger
	DB          *sql.DB

	mu sync.RWMutex
}

func NewApplication() (*Application, error) {
	infoLog, errLog := utils.InitLogger()

	cfg, err := configs.LoadConfig()
	if err != nil {
		errLog.Println("Error loading config:", err)
		return nil, err
	}

	conn, err := db.InitDB(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode, infoLog)
	if err != nil {
		errLog.Println("Error connecting to DB:", err)
		return nil, err
	}

	app := &Application{
		Config:      cfg,
		LoggerInfo:  infoLog,
		LoggerError: errLog,
		DB:          conn,
	}

	app.LoggerInfo.Println("Application initialized with config:", cfg)
	return app, nil
}

func (app *Application) Info(msg string, args ...interface{}) {
	app.LoggerInfo.Printf(msg, args...)
}

func (app *Application) Error(msg string, args ...interface{}) {
	app.LoggerError.Printf(msg, args...)
}

func (app *Application) Close() error {
	app.LoggerInfo.Println("Closing application resources...")
	if app.DB != nil {
		return app.DB.Close()
	}
	return nil
}

func main() {
	application, err := NewApplication()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer application.Close()

	application.Info("Starting application...")

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, app.NewOrderService())

	application.Info("Application started successfully")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		application.Error("Failed to listen on port 50052: %v", err)
		log.Fatalf("failed to listen: %v", err)
	}
	application.Info("gRPC server listening on port 50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
