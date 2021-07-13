package rest

import (
	"context"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/common/web"
	"github.com/leungyauming/api/services/rest/controllers"
	"github.com/leungyauming/api/services/rest/controllers/v1/user"
	"log"
	"net/http"
)

type restService struct {
	webSrv *web.Server
}

// Start implements app.Service interface
func (service *restService) Start() error {
	go func() {
		err := service.webSrv.Start()
		if err != http.ErrServerClosed {
			log.Printf("failed to start web server -> %v", err)
		}
	}()

	return nil
}

// Shutdown implements app.Service interface
func (service *restService) Shutdown() error {
	return service.webSrv.Shutdown(context.Background())
}

// Name implements app.Service interface
func (service *restService) Name() string {
	return "REST"
}

func New(deps *app.Deps) app.Service {
	srv := web.NewServer(deps.Config.App.Services.Rest.BindAddr, deps.Config.App.Services.Rest.BindPort)

	v1Group := srv.Group("/v1")
	{
		userGroup := v1Group.Group("/user")
		{
			loginController := user.NewLoginController(deps)
			userGroup.Any("/login", loginController.Any)

			registerController := user.NewRegisterController(deps)
			userGroup.Any("/register", registerController.Any)
		}
	}

	srv.NoRoute(controllers.NoRoute)

	return &restService{webSrv: srv}
}
