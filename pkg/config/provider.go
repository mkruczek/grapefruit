package config

import "github.com/tkanos/gonfig"

type Provider struct {
	cfgPath string
}

func NewProvider(cfgPath string) Provider {
	return Provider{cfgPath: cfgPath}
}

func (p Provider) GetMongoCfg() (MongoDB, error) {
	mongoDB := MongoDB{}
	err := gonfig.GetConf(p.cfgPath, &mongoDB)
	if err != nil {
		return MongoDB{}, err
	}

	return mongoDB, nil
}
