package api

import "github.com/gin-gonic/gin"

type HttpController struct {
	server *gin.Engine
}

func NewHttpController() HttpController {
	return HttpController{
		server: gin.Default(),
	}
}

func (hc HttpController) ListenAndServe() error {
	return hc.server.Run()
}
