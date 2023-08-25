package repository

import (
	"fmt"
	"sternx-challenge/internal/database"
	"sternx-challenge/internal/model"
)

type TradeRepository interface {
	GetLatestTrades() ([]model.Trade, error)
}

type tradeRepository struct {
	db database.DB
}

func NewTradeRepository(db database.DB) TradeRepository {
	return &tradeRepository{db: db}
}

func (r *tradeRepository) GetLatestTrades() ([]model.Trade, error) {
	query := `SELECT
		I.Name AS Symbol, T.DateEn AS LastTransactionDate,
		T.Open, T.High, T.Low, T.Close
	FROM
		Instrument I
		JOIN Trade T ON I.Id = T.InstrumentId
		JOIN (
			SELECT
				InstrumentId, MAX(DateEn) AS MaxDate
			FROM
				Trade
			GROUP BY
				InstrumentId
		) MaxTradeDates ON T.InstrumentId = MaxTradeDates.InstrumentId AND T.DateEn = MaxTradeDates.MaxDate;`
	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var trades []model.Trade
	for rows.Next() {
		var trade model.Trade
		err := rows.Scan(
			&trade.Symbol,
			&trade.LastTransactionDate,
			&trade.Open,
			&trade.High,
			&trade.Low,
			&trade.Close,
		)
		if err != nil {
			return nil, err
		}
		trades = append(trades, trade)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trades, nil
}
