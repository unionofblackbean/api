package rest

import (
	"github.com/leungyauming/api/app/config"
	"github.com/leungyauming/api/common/web"
)

func New(cfg *config.RestConfig) *web.Server {
	srv := web.NewServer(cfg.BindAddr, cfg.BindPort)

	return srv
}
