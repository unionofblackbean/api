package app

import "fmt"

type App struct {
	services []Service

	deps *Deps
}

func New(deps *Deps) *App {
	app := new(App)
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
			return fmt.Errorf(`failed to start "%s" service -> %v`, service.Name(), err)
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
