package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leungyauming/api/app/config"
)

func initDbPool(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			cfg.Username, cfg.Password,
			cfg.Addr, cfg.Port,
			cfg.DBName))
	if err != nil {
		return nil, fmt.Errorf("failed to connect database -> %v", err)
	}

	return pool, nil
}

//go:embed schema.sql
var dbSchema string

func initDb(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), dbSchema)
	return err
}
