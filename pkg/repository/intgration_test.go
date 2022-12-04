package repository

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"grapefruit/pkg/config"
	"log"
	"os"
	"testing"
	"time"
)

const (
	testConfigPath = "../../test_config.json"
)

var (
	testMongoClient         MongoClient
	testElasticSearchClient ElasticSearchClient
)

var cmpOpt = cmp.Comparer(func(x, y time.Time) bool {
	return (x.Sub(y) > -time.Millisecond) && (x.Sub(y) < time.Millisecond)
})

func TestMain(m *testing.M) {

	cfgProvider := config.NewProvider(testConfigPath)
	cfgMongo, err := cfgProvider.GetMongoDB()
	if err != nil {
		log.Fatalf("can't get config for mongo: %s", err)
	}

	testMongoClient, err = NewMongoClient(context.Background(), cfgMongo)
	if err != nil {
		log.Fatalf("can't connect to mongo db: %s", err)
	}

	cfgElasticsearch, err := cfgProvider.GetElasticsearch()

	testElasticSearchClient, err = NewElasticSearchClient(cfgElasticsearch)
	if err != nil {
		log.Fatalf("can't create elasticsearch client: %s", err)
	}

	os.Exit(m.Run())
}
