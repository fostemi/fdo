package utils

import (
  "net/http"
  "strings"
)

func ExtractToken(r *http.Request) string {
  keys := r.URL.Query()
  token := keys.Get("Bearer")
  if token != "" {
    return token
  }
  
  bearerToken := r.Header.Get("Authorization")
  split := strings.Split(bearerToken, " ")
  if len(split) == 2 {
    return split[1]
  }

  return ""
}
