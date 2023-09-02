package service

import (
	"sternx-challenge/config"
	"sternx-challenge/internal/model"
	"sternx-challenge/internal/repository"
)

type TradeService interface {
	GetLatestTrades() ([]model.Trade, error)
}

type tradeService struct {
	tradeRepo repository.TradeRepository
	cfg       *config.Config
}

func NewTradeService(tradeRepo repository.TradeRepository, cfg *config.Config) TradeService {
	return &tradeService{
		tradeRepo: tradeRepo,
		cfg:       cfg,
	}
}

func (s *tradeService) GetLatestTrades() ([]model.Trade, error) {
	return s.tradeRepo.GetLatestTrades()
}
