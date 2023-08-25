// persistence/trade_repository.go
package persistence

import (
	"database/sql"
	"sternx-cli/domain"
)

type TradeRepository interface {
	InsertTrade(trade domain.Trade) error
}

type tradeRepository struct {
	db *sql.DB
}

func NewTradeRepository(db *sql.DB) TradeRepository {
	return &tradeRepository{db: db}
}

func (tr *tradeRepository) InsertTrade(trade domain.Trade) error {
	query := `
		INSERT INTO Trade (InstrumentId, DateEn, Open, High, Low, Close)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := tr.db.Exec(query, trade.InstrumentID, trade.DateEn, trade.Open, trade.High, trade.Low, trade.Close)
	return err
}
