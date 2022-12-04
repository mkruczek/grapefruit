package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
	"grapefruit/pkg/config"
	"grapefruit/pkg/model"
	"log"
	"net"
	"net/http"
	"time"
)

//todo add layer to this and wrap all methods
type ElasticSearchClient struct {
	index string
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

	index := "objects" //for now, we have single index //todo move to config

	return ElasticSearchClient{
		index:  index,
		Client: client,
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

func (eo elasticsearchObject) toModel() model.Object {
	return model.Object{
		ID:      uuid.MustParse(eo.Id),
		Name:    eo.Name,
		Value:   eo.Value,
		Created: time.UnixMicro(eo.Created),
	}
}

func (es ElasticSearchClient) InsertObject(ctx context.Context, ob model.Object) (model.Object, error) {

	dataJSON, err := json.Marshal(newElasticsearchObjectFromModel(ob))
	if err != nil {
		return model.Object{}, err
	}

	req := esapi.IndexRequest{
		Index:      es.index,
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

func (es ElasticSearchClient) GetObjects(ctx context.Context) ([]model.Object, error) {

	var objects []model.Object

	req := esapi.SearchRequest{
		Index: []string{es.index},
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		eo, err := esResponseToEsObject(hit.(map[string]interface{}))
		if err != nil {
			return nil, err
		}

		objects = append(objects, eo.toModel())
	}

	return objects, nil
}

func (es ElasticSearchClient) GetObjectByID(ctx context.Context, id uuid.UUID) (model.Object, error) {

	req := esapi.GetRequest{
		Index:      es.index,
		DocumentID: id.String(),
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		return model.Object{}, err
	}

	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return model.Object{}, err
	}

	eo, err := esResponseToEsObject(r)
	if err != nil {
		return model.Object{}, err
	}

	return eo.toModel(), nil
}

func (es ElasticSearchClient) DeleteObjectByID(ctx context.Context, id uuid.UUID) error {

	req := esapi.DeleteRequest{
		Index:      es.index,
		DocumentID: id.String(),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func esResponseToEsObject(r map[string]interface{}) (elasticsearchObject, error) {
	var eo elasticsearchObject
	source := r["_source"]
	b, err := json.Marshal(source)
	if err != nil {
		return elasticsearchObject{}, err
	}

	if err := json.Unmarshal(b, &eo); err != nil {
		return elasticsearchObject{}, err
	}
	return eo, nil
}
