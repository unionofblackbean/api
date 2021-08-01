package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/unionofblackbean/api/app/config"
)

type Deps struct {
	Postgres *pgxpool.Pool
	Config   *config.Config
}
