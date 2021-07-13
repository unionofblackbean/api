package main

import (
	_ "embed"
	"flag"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/app/config"
	"github.com/leungyauming/api/common"
	"github.com/leungyauming/api/services/rest"
	"log"
	"os"
	"os/signal"
)

var shouldInit bool

func init() {
	flag.BoolVar(&shouldInit, "init", false, "initialization trigger")
}

func main() {
	flag.Parse()

	logger := common.NewLogger("main")

	err := saveDefaultConfig()
	if err != nil {
		logger.Fatalf("failed to save default config -> %v", err)
	}

	cfg, err := config.ParseFile(configPath)
	if err != nil {
		logger.Fatalf("failed to load config -> %v", err)
	}

	dbPool, err := initDbPool(cfg.App.DB)
	if err != nil {
		logger.Fatalf("failed to initialize database connection pool -> %v", err)
	}

	if shouldInit {
		logger.Println("initializing")
		err := initDb(dbPool)
		if err != nil {
			logger.Fatalf("failed to initialize database -> %v", err)
		}
		logger.Println("initialized")
	}

	deps := &app.Deps{
		Database: dbPool,
		Config:   cfg,
	}

	app_ := app.New(deps)
	app_.RegisterService(rest.New(deps))

	logger.Println("starting services")
	err = app_.Start()
	if err != nil {
		log.Fatalf("failed to start app -> %v", err)
	}
	logger.Println("started all services")

	{
		shutdownCh := make(chan os.Signal)
		signal.Notify(shutdownCh, os.Interrupt)
		<-shutdownCh
	}

	logger.Println("shutting down services")
	errs := app_.Shutdown()
	if len(errs) == 0 {
		logger.Println("shut down all services")
	} else {
		for err := range errs {
			logger.Printf("failed to shutdown services -> %v", err)
		}
	}
}
