package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthz() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "i am healthy")
	}
}
