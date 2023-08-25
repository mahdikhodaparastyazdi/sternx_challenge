// cmd/cli/main.go
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"sternx-cli/application"
	"sternx-cli/persistence"
)

func main() {
	numRecords := flag.Int("num-records", 5, "Number of records to generate")
	flag.Parse()

	// Database connection parameters
	dsn := "user=postgres password=1234 host=localhost dbname=postgres sslmode=disable"

	// Open a database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tradeRepo := persistence.NewTradeRepository(db)
	tradeUseCase := application.NewTradeUseCase(tradeRepo)

	// Generate and insert random trades
	err = tradeUseCase.GenerateAndInsertTrades(*numRecords)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("")
}
