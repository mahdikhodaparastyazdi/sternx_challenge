package handler

import (
	"sternx-challenge/config"
	"sternx-challenge/internal/service"

	"github.com/gofiber/fiber/v2"
)

type TradeHandler interface {
	HandleLatestTrades(ctx *fiber.Ctx) error
	Middleware(c *fiber.Ctx) error
}

type tradeHandler struct {
	tradeService service.TradeService
	cfg          *config.Config
}

func NewTradeHandler(tradeService service.TradeService, cfg *config.Config) TradeHandler {
	return &tradeHandler{
		tradeService: tradeService,
		cfg:          cfg,
	}
}

func (h *tradeHandler) HandleLatestTrades(ctx *fiber.Ctx) error {
	trades, err := h.tradeService.GetLatestTrades()
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).SendString("Failed to get latest trades")
		return err
	}
	ctx.JSON(trades)
	return nil
}
func (h *tradeHandler) Middleware(c *fiber.Ctx) error {
	return c.Next()
}
