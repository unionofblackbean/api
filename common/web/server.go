package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	*gin.Engine

	httpSrv *http.Server
}

func NewServer(addr string, port uint16) *Server {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	return &Server{
		Engine: r,
		httpSrv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", addr, port),
			Handler: r,
		},
	}
}

// Start implements app.Service interface
func (srv *Server) Start() error {
	return srv.httpSrv.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.httpSrv.Shutdown(ctx)
}
