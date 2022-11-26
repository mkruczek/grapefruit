package config

import "github.com/tkanos/gonfig"

type Provider struct {
	cfgPath string
}

func NewProvider(cfgPath string) Provider {
	return Provider{cfgPath: cfgPath}
}

func (p Provider) GetMongoDB() (MongoDB, error) {
	cfg := MongoDB{}
	err := gonfig.GetConf(p.cfgPath, &cfg)
	if err != nil {
		return MongoDB{}, err
	}

	return cfg, nil
}

func (p Provider) GetElasticsearch() (ElasticSearchDS, error) {
	cfg := ElasticSearchDS{}
	err := gonfig.GetConf(p.cfgPath, &cfg)
	if err != nil {
		return ElasticSearchDS{}, err
	}

	return cfg, nil
}
