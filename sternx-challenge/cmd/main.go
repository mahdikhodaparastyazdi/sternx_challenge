package main

import (
	"log"
	"os/signal"
	"syscall"

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
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)
	select {
	case err := <-serv.ListenAndServe():
		panic(err)
	case <-sigCh:
		logger.DebugLogger(" Crm-Biz Running").Msg("Shutdown service...")
		os.Exit(1)
	}
}
