package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/unionofblackbean/api/app/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Deps struct {
	Postgres *pgxpool.Pool
	Mongo    *mongo.Client

	Config *config.Config
}
