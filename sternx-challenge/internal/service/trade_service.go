package service

import (
	"sternx-challenge/internal/model"
	"sternx-challenge/internal/repository"
)

type TradeService interface {
	GetLatestTrades() ([]model.Trade, error)
}

type tradeService struct {
	tradeRepo repository.TradeRepository
}

func NewTradeService(tradeRepo repository.TradeRepository) TradeService {
	return &tradeService{tradeRepo: tradeRepo}
}

func (s *tradeService) GetLatestTrades() ([]model.Trade, error) {
	return s.tradeRepo.GetLatestTrades()
}
