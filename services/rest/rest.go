package rest

import (
	"context"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/common"
	"github.com/leungyauming/api/common/web"
	"github.com/leungyauming/api/services/rest/controllers"
	"github.com/leungyauming/api/services/rest/controllers/v1/session"
	"github.com/leungyauming/api/services/rest/controllers/v1/user"
	"log"
	"net/http"
)

type restService struct {
	logger *log.Logger
	webSrv *web.Server
}

// Start implements app.Service interface
func (service *restService) Start() error {
	go func() {
		err := service.webSrv.Start()
		if err != http.ErrServerClosed {
			service.logger.Printf("failed to start web server -> %v", err)
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

	srv.Use(web.NewIPRateLimiter(deps.Config.App.Services.Rest.RateLimit).Middleware)

	v1Group := srv.Group("/v1")
	{
		sessionGroup := v1Group.Group("/session")
		{
			sessionGroup.Any("/login", session.NewLoginController(deps).Any)
		}

		userGroup := v1Group.Group("/user")
		{
			userGroup.Any("/register", user.NewRegisterController(deps).Any)
		}
	}

	srv.Any("/health", controllers.Health)
	srv.NoRoute(controllers.NoRoute)

	return &restService{
		logger: common.NewLogger("rest"),
		webSrv: srv,
	}
}
