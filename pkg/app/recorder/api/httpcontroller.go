package api

import "github.com/gin-gonic/gin"

type HttpController struct {
	server *gin.Engine
	port   string
}

func NewHttpController(port string) HttpController {
	return HttpController{
		server: gin.Default(),
		port:   port,
	}
}

func (hc HttpController) ListenAndServe() error {
	return hc.server.Run(hc.port)
}
