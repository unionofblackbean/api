package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type App struct {
	services []Service

	db *pgxpool.Pool
}

func New() *App {
	app := new(App)
	app.services = []Service{}

	return app
}

func (app *App) RegisterService(service Service) {
	app.services = append(app.services, service)
}

func (app *App) Start(errChan chan error) {
	for _, service := range app.services {
		go func(service Service) {
			errChan <- service.Start()
		}(service)
	}
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
