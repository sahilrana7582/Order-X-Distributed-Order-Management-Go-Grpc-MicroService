package db

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDB(host, port, user, password, dbName, sslMode string, logger *log.Logger) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Printf("Error opening DB connection: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logger.Printf("Error pinging DB: %v", err)
		return nil, err
	}

	logger.Println("DB connection established")
	return db, nil
}
