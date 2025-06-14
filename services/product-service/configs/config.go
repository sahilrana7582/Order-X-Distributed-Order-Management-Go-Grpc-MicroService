package configs

import (
	"github.com/joho/godotenv"
	"github.com/sahilrana7582/product-service/internal/utils"
)

type Config struct {
	DB    DBConfig
	Kafka KafkaConfig
	GRPC  GRPCConfig
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
		return nil, err
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
			GroupID: utils.GetEnv("KAFKA_GROUP_ID", "product-service-group"),
		},
		GRPC: GRPCConfig{
			Host: utils.GetEnv("GRPC_HOST", "localhost"),
			Port: utils.GetEnv("GRPC_PORT", 50053),
		},
	}

	return cfg, nil
}
