package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
  "github.com/fostemi/fdo/internal/common/bucket"
)

func Engine() *gin.Engine {
  r := gin.Default()

  r.GET("/healthz", func(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
      "message": "Healthy",
    })
  })

  r.POST("/create/bucket", bucket.Create)
  return r
}
