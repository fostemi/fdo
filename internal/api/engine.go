package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
  r := gin.Default()
  r.GET("/healthz", func(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
      "message": "Healthy",
    })
  })
  return r
}
