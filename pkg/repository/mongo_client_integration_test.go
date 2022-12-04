package repository

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"grapefruit/pkg/model"
	"testing"
)

func Test_MongoPingConnection(t *testing.T) {
	if err := testMongoClient.client.Ping(context.Background(), nil); err != nil {
		t.Fatalf("can't ping mongo db: %s", err)
	}
}

func Test_MongoIntegration_Create_Get_Update_Delete_Object(t *testing.T) {

	ctx := context.Background()

	//crete
	newObject := model.NewObject()
	newObject.Name = "myName"
	newObject.Value = 16

	err := testMongoClient.CreateObject(ctx, newObject)
	if err != nil {
		t.Fatalf("expected not error for adding object, got: %s", err)
	}

	//get
	got, err := testMongoClient.GetObject(ctx, newObject.ID)
	if err != nil {
		t.Fatalf("expected not error for getting object, got: %s", err)
	}

	if !cmp.Equal(got, newObject, cmpOpt) {
		t.Fatalf("got other object then stored")
	}

	//update
	toUpdate := newObject
	updateName := "updatedName"
	toUpdate.Name = updateName
	updatedValue := 34.2
	toUpdate.Value = updatedValue
	_, err = testMongoClient.UpdateObject(ctx, toUpdate)
	if err != nil {
		t.Fatalf("expected not error for updateing object, got: %s", err)
	}

	got, err = testMongoClient.GetObject(ctx, newObject.ID)
	if err != nil {
		t.Fatalf("expected not error for getting updated object, got: %s", err)
	}

	if got.Name != updateName || got.Value != updatedValue {
		t.Errorf("updated object has wrong data.")
	}

	//delete
	if err := testMongoClient.DeleteObject(ctx, got.ID); err != nil {
		t.Fatalf("expected not error for deleting object, got: %s", err)
	}
}
