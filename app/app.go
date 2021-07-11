package app

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
