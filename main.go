package main

import (
	_ "embed"
	"flag"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/app/config"
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

	err := saveDefaultConfig()
	if err != nil {
		log.Fatalf("failed to save default config -> %v", err)
	}

	cfg, err := config.ParseFile(configPath)
	if err != nil {
		log.Fatalf("failed to load config -> %v", err)
	}

	dbPool, err := initDbPool(cfg.DB)
	if err != nil {
		log.Fatalf("failed to initialize database connection pool -> %v", err)
	}

	if shouldInit {
		log.Println("initializing")

		err := initDb(dbPool)
		if err != nil {
			log.Fatalf("failed to initialize database -> %v", err)
		}

		log.Println("initialized")
	}

	deps := &app.Deps{
		Database: dbPool,
		Config:   cfg,
	}

	app_ := app.New(deps)
	app_.RegisterService(rest.New(deps))

	log.Println("starting services")
	errChan := make(chan error)
	app_.Start(errChan)
	log.Println("started all services")

	{
		shouldExit := make(chan os.Signal)
		signal.Notify(shouldExit, os.Interrupt)

		shouldShutdown := false
		for {
			if shouldShutdown {
				break
			}

			select {
			case err := <-errChan:
				if err != nil {
					log.Println(err)
				}
			case <-shouldExit:
				shouldShutdown = true
			}
		}
	}

	log.Println("shutting down services")
	errs := app_.Shutdown()
	for err := range errs {
		log.Printf("failed to shutdown services -> %v", err)
	}
	log.Println("shut down all services")
}
