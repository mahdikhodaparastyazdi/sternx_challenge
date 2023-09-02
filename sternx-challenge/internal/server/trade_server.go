package server

import (
	"log"
	"sternx-challenge/config"
	"sternx-challenge/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Server interface {
	Run() error
}

type server struct {
	tradeHandler handler.TradeHandler
	cfg          *config.Config
}

func NewServer(tradeHandler handler.TradeHandler, cfg *config.Config) Server {
	return &server{
		tradeHandler: tradeHandler,
		cfg:          cfg,
	}
}

func (s *server) Run() error {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "{\"level\":\"info\",\"task\":\"stern-x\",\"error\":\"${status}\",\"time\":\"${time}\",\"message\":\"${locals:requestid} ${latency} ${method} ${path}\"}\n",
		TimeFormat: "2006-01-02T15:04:05-0700",
	}))
	api := app.Group("/api", s.tradeHandler.Middleware)
	v1 := api.Group("/v1", s.tradeHandler.Middleware)

	v1.Get("/latest-trades", s.tradeHandler.HandleLatestTrades)
	err := app.Listen(s.cfg.Server.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
