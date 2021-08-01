package main

import (
	"context"
	"fmt"
	"github.com/unionofblackbean/api/app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func initMongoClient(cfg *config.MongoConfig) (*mongo.Client, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(
			fmt.Sprintf("mongodb://%s:%s@%s:%d",
				cfg.Username, cfg.Password,
				cfg.Addr, cfg.Port)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongo database -> %v", err)
	}

	return client, nil
}
