package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"sternx-challenge/config"
	"sternx-challenge/internal/database"
	"sternx-challenge/internal/handler"
	"sternx-challenge/internal/repository"
	"sternx-challenge/internal/server"
	"sternx-challenge/internal/service"

	"github.com/rs/zerolog"
)

var (
	configFile     string
	verbosityLevel int
	addr           string
)

func showHelp() {
	fmt.Printf("Usage:%s {params}\n", os.Args[0])
	fmt.Println("      -c {config file}")
	fmt.Println("      -v {0-10} (verbosity level, default 0)")
	fmt.Println("      -a {listen addr}")
}

func parse() bool {
	flag.StringVar(&configFile, "c", "config.yml", "config file")
	flag.StringVar(&addr, "a", "localhost", "address to use")
	flag.IntVar(&verbosityLevel, "v", -1, "verbosity level, higher value - more logs")
	flag.Parse()

	return config.Load(configFile)
}

func main() {

	if !parse() {
		showHelp()
		os.Exit(-1)
	}

	if verbosityLevel < 0 {
		verbosityLevel = config.Get().Log.Level
	}
	zerolog.SetGlobalLevel(zerolog.Level(verbosityLevel))

	cfg := config.Get()
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	tradeRepo := repository.NewTradeRepository(db, cfg)
	tradeService := service.NewTradeService(tradeRepo, cfg)
	tradeHandler := handler.NewTradeHandler(tradeService, cfg)

	server := server.NewServer(tradeHandler, cfg)
	log.Fatal(server.Run())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	<-sigCh
	log.Fatal("Shutdown service...")
	os.Exit(1)

}
