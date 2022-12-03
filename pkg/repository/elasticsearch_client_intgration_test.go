package repository

import (
	"context"
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
}
