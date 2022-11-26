package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"grapefruit/pkg/config"
)

//https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver
type MongoClient struct {
	client *mongo.Client
	*mongo.Collection
}

func NewMongoClient(ctx context.Context, cfg config.MongoDB) (MongoClient, error) {

	clientOptions := options.Client().ApplyURI(cfg.MODBConnectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return MongoClient{}, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return MongoClient{}, err
	}

	collection := client.Database("tasker").Collection("tasks")

	return MongoClient{
		client:     client,
		Collection: collection,
	}, nil
}
