package app

import (
	"github.com/unionofblackbean/api/app/config"
	"github.com/unionofblackbean/api/common"
	"log"
)

type App struct {
	logger   *log.Logger
	services []Service
	deps     *Deps
}

func New(deps *Deps) *App {
	return &App{
		logger: common.NewLogger("app"),
		deps:   deps,
	}
}

func (app *App) RegisterService(service Service) {
	app.services = append(app.services, service)
}

func (app *App) Start() error {
	for _, service := range app.services {
		if err := service.Start(); err != nil {
			if app.deps.Config.App.StartPolicy == config.StartPolicyExitOnError {
				return err
			}
			app.logger.Printf(`failed to start "%s" service -> %v`, service.Name(), err)
		}
	}

	return nil
}

func (app *App) Shutdown() []error {
	var errs []error
	for _, service := range app.services {
		if err := service.Shutdown(); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
