package handler

import (
	"sternx-challenge/internal/service"

	"github.com/gofiber/fiber/v2"
)

type TradeHandler interface {
	HandleLatestTrades(ctx *fiber.Ctx) error
}

type tradeHandler struct {
	tradeService service.TradeService
}

func NewTradeHandler(tradeService service.TradeService) TradeHandler {
	return &tradeHandler{tradeService: tradeService}
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
