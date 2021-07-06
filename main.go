package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/app/config"
	"github.com/leungyauming/api/common/utils"
	"github.com/leungyauming/api/services/rest"
	"log"
	"os"
	"os/signal"
	"path/filepath"
)

var (
	defaultConfig = config.Config{
		Rest: &config.RestConfig{
			BindAddr: "127.0.0.1",
			BindPort: 8080,
		},
		DB: &config.DBConfig{
			Addr:     "127.0.0.1",
			Port:     5432,
			Username: "api",
			Password: "api",
			DBName:   "api",
		},
	}
	configPath = "config.json"
)

func saveDefaultConfig() error {
	execDir, err := utils.ExecutableDirectory()
	if err != nil {
		return fmt.Errorf("failed to get executable directory -> %v", err)
	}

	fullConfigPath := filepath.Join(execDir, configPath)
	if !utils.IsFileExists(fullConfigPath) {
		configFile, err := os.Create(fullConfigPath)
		if err != nil {
			return fmt.Errorf("failed to create config file -> %v", err)
		}

		configBytes, err := json.MarshalIndent(&defaultConfig, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal default config struct -> %v", err)
		}

		_, err = configFile.Write(configBytes)
		if err != nil {
			return fmt.Errorf("failed to write default config to file -> %v", err)
		}
	}

	return nil
}

func initDbPool(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			cfg.Username, cfg.Password,
			cfg.Addr, cfg.Port,
			cfg.DBName))
	if err != nil {
		return nil, fmt.Errorf("failed to connect database -> %v", err)
	}

	return pool, nil
}

//go:embed schema.sql
var dbSchema string

func initDb(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), dbSchema)
	return err
}

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

	app_ := app.New()
	app_.RegisterService(rest.New(cfg.Rest, dbPool))

	log.Println("starting services")
	errChan := make(chan error)
	app_.Start(errChan)
	log.Println("started all services")

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

	log.Println("shutting down services")
	errs := app_.Shutdown()
	for err := range errs {
		log.Printf("failed to shutdown services -> %v", err)
	}
	log.Println("shut down all services")
}
