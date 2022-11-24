package repository

import (
	"context"
	"grapefruit/pkg/config"
	"log"
	"os"
	"testing"
)

const (
	testConfigPath = "../../test_config.json"
)

var (
	testRepository MongoRepository
)

func TestMain(m *testing.M) {

	cfgProvider := config.NewProvider(testConfigPath)
	cfgMongo, err := cfgProvider.GetMongoCfg()
	if err != nil {
		log.Fatalf("can't get config for mongo: %s", err)
	}

	testRepository, err = NewMongoClient(context.Background(), cfgMongo)
	if err != nil {
		log.Fatalf("can't connect to mongo db: %s", err)
	}

	os.Exit(m.Run())
}

func Test_MongoPingConnection(t *testing.T) {
	if err := testRepository.client.Ping(context.Background(), nil); err != nil {
		t.Fatalf("can't ping mongo db: %s", err)
	}
}
