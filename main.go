package main

import (
	"encoding/json"
	"fmt"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/app/config"
	"github.com/leungyauming/api/common/utils"
	"github.com/leungyauming/api/services/rest"
	"log"
	"os"
	"os/signal"
	"path/filepath"
)

var defaultConfig = config.Config{
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
var configPath = "config.json"

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

func main() {
	err := saveDefaultConfig()
	if err != nil {
		log.Fatalf("failed to save default config -> %v", err)
	}

	cfg, err := config.ParseFile(configPath)
	if err != nil {
		log.Fatalf("failed to load config -> %v", err)
	}

	app_ := app.New()
	app_.RegisterService(rest.New(cfg.Rest))

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
