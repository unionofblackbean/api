package rest

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/app/config"
	"github.com/leungyauming/api/common/web"
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

func New(cfg *config.RestConfig, dbPool *pgxpool.Pool) app.Service {
	srv := web.NewServer(cfg.BindAddr, cfg.BindPort)

	v1Group := srv.Group("/v1")
	{
		userGroup := v1Group.Group("/user")
		{
			loginController := user.NewLoginController(dbPool)
			userGroup.POST("/login", loginController.Post)

			registerController := user.NewRegisterController(dbPool)
			userGroup.POST("/register", registerController.Post)
		}
	}

	return &restService{webSrv: srv}
}
