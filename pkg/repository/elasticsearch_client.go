package repository

import (
	"github.com/elastic/go-elasticsearch/v8"
	"grapefruit/pkg/config"
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
