package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"grapefruit/pkg/config"
	"grapefruit/pkg/model"
	"time"
)

//https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver
type MongoClient struct {
	coll   *mongo.Collection
	client *mongo.Client
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

	//app for now working for single collection
	collection := client.Database(cfg.MODBName).Collection(cfg.MODBCollection)

	return MongoClient{
		client: client,
		coll:   collection,
	}, nil
}

type mongoObject struct {
	Id      string  `bson:"id"`
	Name    string  `bson:"name"`
	Value   float64 `bson:"value"`
	Created int64   `bson:"created"`
}

func newMongoObjectFromModel(m model.Object) mongoObject {
	return mongoObject{
		Id:      m.ID.String(),
		Name:    m.Name,
		Value:   m.Value,
		Created: m.Created.UnixMicro(),
	}
}

func (mo mongoObject) toModel() model.Object {
	return model.Object{
		ID:      uuid.MustParse(mo.Id),
		Name:    mo.Name,
		Value:   mo.Value,
		Created: time.UnixMicro(mo.Created),
	}
}

func (m MongoClient) CreateObject(ctx context.Context, newObject model.Object) error {
	mo := newMongoObjectFromModel(newObject)
	_, err := m.coll.InsertOne(ctx, mo)
	return err
}

func (m MongoClient) GetObject(ctx context.Context, id uuid.UUID) (model.Object, error) {
	filter := bson.D{{"id", id.String()}}
	sr := m.coll.FindOne(ctx, filter)

	var mo mongoObject
	err := sr.Decode(&mo)
	if err != nil {
		return model.Object{}, err
	}

	return mo.toModel(), nil
}

func (m MongoClient) DeleteObject(ctx context.Context, id uuid.UUID) error {
	filter := bson.D{{"id", id.String()}}
	dr, err := m.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if dr.DeletedCount != 1 {
		return fmt.Errorf("deleted more then single row")
	}

	return nil
}
