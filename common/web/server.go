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

	srv := new(Server)
	srv.Engine = r
	srv.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", addr, port),
		Handler: r,
	}

	return srv
}

func (srv *Server) Start() error {
	return srv.httpSrv.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.httpSrv.Shutdown(ctx)
}
