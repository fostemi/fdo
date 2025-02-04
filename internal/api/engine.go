package api

import (
  "fmt"
	"net/http"

  "github.com/fostemi/fdo/internal/common/bucket"
)

type Engine interface {

}

type Server struct {

}

func CreateRoutes() {
  // r := gin.Default()

  http.Handle("/healthz", http.HandlerFunc(Health))

  // r.POST("/create/bucket", bucket.Create)
  http.Handle("/create/bucket", http.HandlerFunc(bucket.Create))
}

func ListenAndServe(addr string, h http.Handler) (error) {
  err := http.ListenAndServe(addr, h)
  if err != nil {
    return err
  }
  return nil
}

func Health(w http.ResponseWriter, req *http.Request) {
  fmt.Fprint(w, "Healthy")
}
