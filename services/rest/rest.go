package rest

import (
	"context"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/common/web"
	"github.com/leungyauming/api/services/rest/controllers"
	"github.com/leungyauming/api/services/rest/controllers/v1/user"
)

type restService struct {
	webSrv *web.Server
}

// Start implements app.Service interface
func (service *restService) Start() error {
	return service.webSrv.Start()
}

// Shutdown implements app.Service interface
func (service *restService) Shutdown() error {
	return service.webSrv.Shutdown(context.Background())
}

func New(deps *app.Deps) app.Service {
	srv := web.NewServer(deps.Config.Rest.BindAddr, deps.Config.Rest.BindPort)

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
