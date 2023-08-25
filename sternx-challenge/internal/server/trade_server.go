package server

import (
	"log"
	"sternx-challenge/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type Server interface {
	Run() error
}

type server struct {
	tradeHandler handler.TradeHandler
}

func NewServer(tradeHandler handler.TradeHandler) Server {
	return &server{tradeHandler: tradeHandler}
}

func (s *server) Run() error {
	app := fiber.New()

	app.Get("/latest-trades", s.tradeHandler.HandleLatestTrades)

	err := app.Listen(":8088")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
