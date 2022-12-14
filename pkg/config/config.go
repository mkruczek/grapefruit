package config

type RecorderServer struct {
	RecorderServerHTTPPort string
}

type ViewerServer struct {
	ViewerServerHTTPPort string
}

type MongoDB struct {
	MODBConnectionString string
	MODBName             string
	MODBCollection       string
}

type ElasticSearchDS struct {
	ESUrl      string
	ESUsername string
	ESPassword string
}
