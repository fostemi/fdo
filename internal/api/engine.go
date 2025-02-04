package api

import (
  "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
  "github.com/fostemi/fdo/internal/common/bucket"
)

type Engine interface {

}

type Server struct {

}

func  CreateRoutes() {
  r := gin.Default()

  http.Handle("/healthz", http.HandlerFunc(Health))

  r.POST("/create/bucket", bucket.Create)
}

func Health(w http.ResponseWriter, req *http.Request) {
  fmt.Fprint(w, "Healthy")
}
