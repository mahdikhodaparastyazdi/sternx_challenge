package main

import (
	"log"

	"sternx-challenge/internal/database"
	"sternx-challenge/internal/handler"
	"sternx-challenge/internal/repository"
	"sternx-challenge/internal/server"
	"sternx-challenge/internal/service"
)

func main() {
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	tradeRepo := repository.NewTradeRepository(db)
	tradeService := service.NewTradeService(tradeRepo)
	tradeHandler := handler.NewTradeHandler(tradeService)

	server := server.NewServer(tradeHandler)
	log.Fatal(server.Run())
}
