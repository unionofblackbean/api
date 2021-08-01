package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/unionofblackbean/api/app/config"
	"time"
)

func initPostgresPool(cfg *config.PostgresConfig) (*pgxpool.Pool, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancelFunc()

	pool, err := pgxpool.Connect(ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			cfg.Username, cfg.Password,
			cfg.Addr, cfg.Port,
			cfg.DBName))
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres database -> %v", err)
	}

	return pool, nil
}

//go:embed postgres_schema.sql
var postgresSchema string

func initPostgres(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), postgresSchema)
	return err
}
