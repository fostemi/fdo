package main

import (
  "fmt"
  "net/http"
  "log/slog"
  "flag"
)

var (
  addr = flag.String("addr", ":8080", "http service address")
)

func main () {
  flag.Parse()
  http.Handle("/healthz", http.HandlerFunc(health))
  err := http.ListenAndServe(*addr, nil)
  if err != nil {
    slog.Error("ListenAndServe:", slog.Any("error", err))
  }
}

func health(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Healthy")
}
