package recorder

import (
	"grapefruit/pkg/app/recorder/api"
	"grapefruit/pkg/config"
	"log"
)

type Server struct {
	httpClient HttpClient
	cfg        config.RecorderServer
}

func NewServer(cfgPath string) (Server, error) {

	cfg := config.NewProvider(cfgPath)
	serverConfig, err := cfg.GetRecorderServer()
	if err != nil {
		return Server{}, err
	}

	hc := api.NewHttpController(serverConfig.RecorderServerHTTPPort)
	hc.SetRouters()

	return Server{
		httpClient: hc,
	}, nil
}

func (s Server) Start() {
	log.Println("...starting server")

	if err := s.httpClient.ListenAndServe(); err != nil {
		log.Fatalf("can't start server:%s", err)
	}
}
