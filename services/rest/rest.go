package rest

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/unionofblackbean/api/app"
	"github.com/unionofblackbean/api/common"
	"github.com/unionofblackbean/api/common/responses"
	"github.com/unionofblackbean/api/common/web"
	"github.com/unionofblackbean/api/services/rest/v1"
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

func noRoute(ctx *gin.Context) {
	responses.SendErrorResponse(ctx,
		http.StatusNotFound,
		errors.New("unknown endpoint"))
}

func New(deps *app.Deps) app.Service {
	srv := web.NewServer(deps.Config.App.Services.Rest.BindAddr, deps.Config.App.Services.Rest.BindPort)

	srv.Use(web.NewIPRateLimiter(deps.Config.App.Services.Rest.RateLimit).Middleware)

	v1.RegisterEndpoints(srv, deps)

	srv.NoRoute(noRoute)

	return &restService{
		logger: common.NewLogger("rest"),
		webSrv: srv,
	}
}
