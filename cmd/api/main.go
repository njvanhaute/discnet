package main

import (
	"flag"
	"log/slog"
	"os"
	"sync"

	"github.com/njvanhaute/discnet/internal/vcs"
)

type config struct {
	port int
	env  string
}
type application struct {
	config config
	logger *slog.Logger
	wg     sync.WaitGroup
}

var (
	version = vcs.Version()
)

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
