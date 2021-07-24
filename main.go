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
	timer := common.NewTimer()

	logger.Print("initializing config file")
	timer.Start()
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
	timer.Stop()
	logger.Printf("initialized config file (%d ms)", timer.Duration().Milliseconds())

	logger.Print("initializing database connection pool")
	timer.Start()
	dbPool, err := initDbPool(cfg.App.DB)
	if err != nil {
		logger.Printf("failed to initialize database connection pool -> %v", err)
		return 1
	}
	timer.Stop()
	logger.Printf("initialized database connection pool (%d ms)", timer.Duration().Milliseconds())

	if shouldInitDb {
		logger.Println("initializing database")
		timer.Start()
		err := initDb(dbPool)
		if err != nil {
			logger.Printf("failed to initialize database -> %v", err)
			return 1
		}
		timer.Stop()
		logger.Printf("initialized database (%d ms)", timer.Duration().Milliseconds())
	}

	deps := &app.Deps{
		Database: dbPool,
		Config:   cfg,
	}

	app_ := app.New(deps)
	app_.RegisterService(rest.New(deps))

	logger.Println("starting services")
	timer.Start()
	err = app_.Start()
	if err != nil {
		logger.Printf("failed to start app -> %v", err)
		return 1
	}
	timer.Stop()
	logger.Printf("started all services (%d ms)", timer.Duration().Milliseconds())

	{
		shutdownCh := make(chan os.Signal)
		signal.Notify(shutdownCh, os.Interrupt)
		<-shutdownCh
	}

	logger.Println("shutting down services")
	timer.Start()
	errs := app_.Shutdown()
	timer.Stop()
	if len(errs) != 0 {
		for err := range errs {
			logger.Printf("failed to shutdown services -> %v", err)
		}
		return 1
	}
	logger.Printf("shut down all services (%d ms)", timer.Duration().Milliseconds())

	return 0
}

func main() {
	os.Exit(Main())
}
