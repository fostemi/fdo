package main

import (
  "github.com/fostemi/fdo/internal/api"
  "flag"
)

var (
  addr = flag.String("addr", ":8080", "http service address")
)

func main () {
  flag.Parse()

  e := api.Engine()
  e.Run(*addr)
}
