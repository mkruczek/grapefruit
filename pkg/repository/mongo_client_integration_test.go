package repository

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"grapefruit/pkg/model"
	"testing"
	"time"
)

func Test_MongoPingConnection(t *testing.T) {
	if err := testMongoClient.client.Ping(context.Background(), nil); err != nil {
		t.Fatalf("can't ping mongo db: %s", err)
	}
}

func Test_MongoIntegration_Create_Get_Delete_Object(t *testing.T) {

	ctx := context.Background()

	//crete object
	newObject := model.NewObject()
	newObject.Name = "myName"
	newObject.Value = 16

	err := testMongoClient.CreateObject(ctx, newObject)
	if err != nil {
		t.Fatalf("expected not error for adding object, got: %s", err)
	}

	//get object
	got, err := testMongoClient.GetObject(ctx, newObject.ID)
	if err != nil {
		t.Fatalf("expected not error for getting object, got: %s", err)
	}

	cmpOpt := cmp.Comparer(func(x, y time.Time) bool {
		return (x.Sub(y) > -time.Millisecond) && (x.Sub(y) < time.Millisecond)
	})

	if !cmp.Equal(got, newObject, cmpOpt) {
		t.Fatalf("got other object then stored")
	}

	//delete
	if err := testMongoClient.DeleteObject(ctx, got.ID); err != nil {
		t.Fatalf("expected not error for deleting object, got: %s", err)
	}
}
