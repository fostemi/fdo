package main

import (
	"flag"
	"log"

	"github.com/fostemi/fdo/internal/api"
)

var (
  addr = flag.String("addr", ":8080", "http service address")
)

func main () {
  flag.Parse()

  api.CreateRoutes()
  err := api.ListenAndServe(*addr, nil)
  if err != nil {
    log.Fatal(err)
  }
}
