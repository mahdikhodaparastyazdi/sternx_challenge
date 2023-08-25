// application/trade_usecase.go
package application

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"sternx-cli/domain"
	"sternx-cli/persistence"
)

var (
	instrumentIDs = []int{1, 2, 3} // Replace with actual instrument IDs
	usedDates     = make(map[string]map[int]struct{})
	mutex         sync.Mutex
)

type TradeUseCase interface {
	GenerateAndInsertTrades(numRecords int) error
}

type tradeUseCase struct {
	tradeRepo persistence.TradeRepository
}

func NewTradeUseCase(tr persistence.TradeRepository) TradeUseCase {
	return &tradeUseCase{tradeRepo: tr}
}

func (tu *tradeUseCase) GenerateAndInsertTrades(numRecords int) error {
	// Use WaitGroup to ensure all workers are done
	var wg sync.WaitGroup
	for i := 0; i < numRecords; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			tradeData := tu.generateRandomTradeData()
			err := tu.tradeRepo.InsertTrade(tradeData)
			if err != nil {
				fmt.Println("Error inserting trade:", err)
			}
		}()
	}
	wg.Wait()

	return nil
}

func (tu *tradeUseCase) generateRandomTradeData() domain.Trade {
	instrumentID := rand.Intn(len(instrumentIDs))
	dateEn := tu.generateUniqueDate(instrumentID)
	open := rand.Intn(1000) + 1000
	top := rand.Intn(1000) + 2000
	low := rand.Intn(1000)
	close := rand.Intn(1000) + low

	return domain.Trade{
		InstrumentID: instrumentID,
		DateEn:       dateEn,
		Open:         open,
		High:         top,
		Low:          low,
		Close:        close,
	}
}

func (tu *tradeUseCase) generateUniqueDate(instrumentID int) string {
	mutex.Lock()
	defer mutex.Unlock()

	for {
		date := generateRandomDate()
		if usedDates[date] == nil {
			usedDates[date] = make(map[int]struct{})
		}
		if _, exists := usedDates[date][instrumentID]; !exists {
			usedDates[date][instrumentID] = struct{}{}
			return date
		}
	}
}

func generateRandomDate() string {
	return time.Now().Add(-time.Duration(rand.Intn(365)) * 24 * time.Hour).Format("2006-01-02")
}
