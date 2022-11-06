package client

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClient(host, port, database string) (*mongo.Database, error) {
	ctx := context.TODO()

	mongoDbUrl := fmt.Sprintf("mongodb://%s:%s", host, port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUrl))

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping to database: %s", err)
	}

	return client.Database(database), nil
}
