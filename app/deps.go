package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leungyauming/api/app/config"
)

type Deps struct {
	Database *pgxpool.Pool
	Config   *config.Config
}
