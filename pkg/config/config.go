package config

type MongoDB struct {
	MODBConnectionString string
}

type ElasticSearchDS struct {
	ESUrl      string
	ESUsername string
	ESPassword string
}
