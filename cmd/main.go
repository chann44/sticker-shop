package main

import (
	"fmt"
	"log"

	"github.com/chann44/go-shop/cmd/api"
	"github.com/chann44/go-shop/config"
	"github.com/chann44/go-shop/db"
)

func main() {
	cfg := config.Envs

	db, err := db.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")

	server := api.NewApiServer(":8000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
