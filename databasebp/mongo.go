package databasebp

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(uri string, db string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("databasebp: failed to create mongodb conn: %w", err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("databasebp: failed to ping mongodb: %w", err)
	}

	return client, nil
}
