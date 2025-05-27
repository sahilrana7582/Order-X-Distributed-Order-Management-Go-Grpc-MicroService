package configs

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sahilrana7582/user-service/internal/utils"
)

type Config struct {
	DB       DBConfig
	Kafka    KafkaConfig
	GRPC     GRPCConfig
	LogLevel string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type KafkaConfig struct {
	Brokers []string
	GroupID string
}

type GRPCConfig struct {
	Host string
	Port int
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: no .env file found or error reading it: %v", err)
	}

	cfg := &Config{
		DB: DBConfig{
			Host:     utils.GetEnv("DB_HOST", "localhost"),
			Port:     utils.GetEnv("DB_PORT", "5432"),
			User:     utils.GetEnv("DB_USER", "postgres"),
			Password: utils.GetEnv("DB_PASSWORD", "1234"),
			Name:     utils.GetEnv("DB_NAME", "orderX"),
			SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
		},
		Kafka: KafkaConfig{
			Brokers: utils.GetEnv("KAFKA_BROKERS", []string{"localhost:9092"}),
			GroupID: utils.GetEnv("KAFKA_GROUP_ID", "user-service-group"),
		},
		GRPC: GRPCConfig{
			Host: utils.GetEnv("GRPC_HOST", "localhost"),
			Port: utils.GetEnv("GRPC_PORT", 50051),
		},
		LogLevel: utils.GetEnv("LOG_LEVEL", "info"),
	}

	if cfg.DB.Host == "" || cfg.DB.User == "" || cfg.DB.Password == "" || cfg.DB.Name == "" {
		return nil, fmt.Errorf("database configuration incomplete")
	}
	if len(cfg.Kafka.Brokers) == 0 || (len(cfg.Kafka.Brokers) == 1 && cfg.Kafka.Brokers[0] == "") || cfg.Kafka.GroupID == "" {
		return nil, fmt.Errorf("kafka configuration incomplete")
	}
	if cfg.GRPC.Port == 0 {
		return nil, fmt.Errorf("grpc port must be set")
	}

	return cfg, nil

}
