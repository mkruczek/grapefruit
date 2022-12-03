package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"grapefruit/pkg/config"
	"grapefruit/pkg/model"
	"log"
	"net"
	"net/http"
	"time"
)

type ElasticSearchClient struct {
	*elasticsearch.Client
}

func NewElasticSearchClient(cfg config.ElasticSearchDS) (ElasticSearchClient, error) {

	esCfg := elasticsearch.Config{
		Addresses: []string{
			cfg.ESUrl,
		},
		Username: cfg.ESUsername,
		Password: cfg.ESPassword,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	}

	client, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		return ElasticSearchClient{}, err
	}

	_, err = client.Ping()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	return ElasticSearchClient{
		client,
	}, nil
}

type elasticsearchObject struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Value   float64 `json:"value"`
	Created int64   `json:"created"`
}

func newElasticsearchObjectFromModel(m model.Object) elasticsearchObject {
	return elasticsearchObject{
		Id:      m.ID.String(),
		Name:    m.Name,
		Value:   m.Value,
		Created: m.Created.UnixMicro(),
	}
}

func (es ElasticSearchClient) InsertObject(ctx context.Context, ob model.Object) (model.Object, error) {

	dataJSON, err := json.Marshal(newElasticsearchObjectFromModel(ob))
	if err != nil {
		return model.Object{}, err
	}

	req := esapi.IndexRequest{
		Index:      "object",
		DocumentID: ob.ID.String(),
		Body:       bytes.NewReader(dataJSON),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		return model.Object{}, err
	}

	if err != nil {
		return model.Object{}, err
	}
	defer res.Body.Close()

	return ob, err
}
