package repository

import (
	"testing"
)

func Test_ElasticsearchPingConnection(t *testing.T) {
	if _, err := testElasticSearchClient.Ping(); err != nil {
		t.Fatalf("can't ping elasticsearch: %s", err)
	}
}
