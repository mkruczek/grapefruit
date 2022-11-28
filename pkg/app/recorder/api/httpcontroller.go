package api

import "github.com/gin-gonic/gin"

type HttpServer struct {
	server *gin.Engine
	port   string
}

func NewHttpServer(port string) HttpServer {
	return HttpServer{
		server: gin.Default(),
		port:   port,
	}
}

func (hc HttpServer) ListenAndServe() error {
	return hc.server.Run(hc.port)
}
