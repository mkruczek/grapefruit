package repository

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"grapefruit/pkg/model"
	"testing"
)

func Test_ElasticsearchPingConnection(t *testing.T) {
	if _, err := testElasticSearchClient.Ping(); err != nil {
		t.Fatalf("can't ping elasticsearch: %s", err)
	}
}

func Test_ElasticSearchIntegration_Insert_Object(t *testing.T) {

	ctx := context.Background()

	//crete
	newObject := model.NewObject()
	newObject.Name = "myName"
	newObject.Value = 16

	_, err := testElasticSearchClient.InsertObject(ctx, newObject)
	if err != nil {
		t.Fatalf("expected not error for inserting object, got: %s", err)
	}

	//getAll
	objects, err := testElasticSearchClient.GetObjects(ctx)
	if err != nil {
		return
	}

	if len(objects) != 1 {
		t.Errorf("expected 1 object, got: %d", len(objects))
	}

	//get single by id
	got, err := testElasticSearchClient.GetObjectByID(ctx, newObject.ID)
	if err != nil {
		t.Fatalf("expected not error for getting object, got: %s", err)
	}

	if !cmp.Equal(got, newObject, cmpOpt) {
		t.Errorf("got other object then stored")
	}

	//delete by id
	err = testElasticSearchClient.DeleteObjectByID(ctx, newObject.ID)
	if err != nil {
		t.Fatalf("expected not error for deleting object, got: %s", err)
	}
}
