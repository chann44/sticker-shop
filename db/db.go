package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/chann44/go-shop/config"
	_ "github.com/lib/pq"
)

func NewPostgresStorage(cfg config.Config) (*sql.DB, error) {
	// Corrected connection string for PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	// Initialize the database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// Verify the connection is working
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping PostgreSQL:", err)
	}

	return db, nil
}
