package repository

import (
	"context"
	"testing"
)

func Test_MongoPingConnection(t *testing.T) {
	if err := testMongoClient.client.Ping(context.Background(), nil); err != nil {
		t.Fatalf("can't ping mongo db: %s", err)
	}
}
