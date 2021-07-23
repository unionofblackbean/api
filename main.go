package main

import (
	_ "embed"
	"flag"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/app/config"
	"github.com/leungyauming/api/common"
	"github.com/leungyauming/api/services/rest"
	"os"
	"os/signal"
)

var (
	shouldInitDb  bool
	shouldInitCfg bool
)

func init() {
	flag.BoolVar(&shouldInitDb, "init-db", false, "database initialization trigger")
	flag.BoolVar(&shouldInitCfg, "only-cfg", false, "only creates config trigger")
}

func Main() int {
	flag.Parse()

	logger := common.NewLogger("main")

	err := initConfigFile()
	if err != nil {
		logger.Printf("failed to save default config -> %v", err)
		return 1
	}
	if shouldInitCfg {
		return 0
	}

	cfg, err := config.ParseFile(configPath)
	if err != nil {
		logger.Printf("failed to load config -> %v", err)
		return 1
	}

	dbPool, err := initDbPool(cfg.App.DB)
	if err != nil {
		logger.Printf("failed to initialize database connection pool -> %v", err)
		return 1
	}

	if shouldInitDb {
		logger.Println("initializing database")
		err := initDb(dbPool)
		if err != nil {
			logger.Printf("failed to initialize database -> %v", err)
			return 1
		}
		logger.Println("initialized database")
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
		logger.Printf("failed to start app -> %v", err)
		return 1
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

	return 0
}

func main() {
	os.Exit(Main())
}
