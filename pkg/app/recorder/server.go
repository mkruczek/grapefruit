package recorder

import (
	"grapefruit/pkg/app/recorder/api"
	"grapefruit/pkg/config"
	"log"
)

type Service struct {
	httpClient HttpClient
	cfg        config.RecorderServer
}

func NewService(cfgPath string) (Service, error) {

	cfg := config.NewProvider(cfgPath)
	serverConfig, err := cfg.GetRecorderServer()
	if err != nil {
		return Service{}, err
	}

	hc := api.NewHttpServer(serverConfig.RecorderServerHTTPPort)
	hc.SetRouters()

	return Service{
		httpClient: hc,
	}, nil
}

func (s Service) Start() {
	log.Println("...starting server")

	if err := s.httpClient.ListenAndServe(); err != nil {
		log.Fatalf("can't start server:%s", err)
	}
}
