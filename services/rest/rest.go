package rest

import (
	"context"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/app/config"
	"github.com/leungyauming/api/common/web"
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

func New(cfg *config.RestConfig) app.Service {
	srv := web.NewServer(cfg.BindAddr, cfg.BindPort)

	return &restService{webSrv: srv}
}
