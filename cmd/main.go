package main

import (
	"log"

	"github.com/MigueelLago/service-order/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	db := config.ConnectDB()
	_ = db
}
