package app

import (
	"github.com/leungyauming/api/app/config"
	"github.com/leungyauming/api/common"
	"log"
)

type App struct {
	logger   *log.Logger
	services []Service
	deps     *Deps
}

func New(deps *Deps) *App {
	app := new(App)
	app.logger = common.NewLogger("app")
	app.services = []Service{}
	app.deps = deps

	return app
}

func (app *App) RegisterService(service Service) {
	app.services = append(app.services, service)
}

func (app *App) Start() error {
	for _, service := range app.services {
		if err := service.Start(); err != nil {
			switch app.deps.Config.App.StartPolicy {
			case config.StartPolicyNeverExit:
				app.logger.Printf(`failed to start "%s" service -> %v`, service.Name(), err)
			case config.StartPolicyExitOnError:
				return err
			}
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
