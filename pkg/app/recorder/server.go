package recorder

import (
	"grapefruit/pkg/app/recorder/api"
	"log"
)

type Server struct {
	httpClient HttpClient
}

func NewServer() Server {
	hc := api.NewHttpController()
	hc.SetRouters()

	return Server{
		httpClient: hc,
	}
}

func (s Server) Start() {
	log.Println("...starting server")

	if err := s.httpClient.ListenAndServe(); err != nil {
		log.Fatalf("can't start server:%s", err)
	}
}
