package bucket

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fostemi/fdo/internal/common/utils"

	"cloud.google.com/go/auth"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func Create(w http.ResponseWriter, req *http.Request) {
  ctx := context.Background()
  creds := getCredsFromToken(getToken(req))
  _, err := storage.NewClient(ctx, option.WithAuthCredentials(creds))
  if err != nil {
    http.Error(w, fmt.Sprint("Error creating storage client with creds, error: ", err), http.StatusBadRequest)
  }
  fmt.Fprint(w, creds)
}

func getCredsFromToken(token string) *auth.Credentials{
  type tokenKey string
  key := tokenKey("Key")
  var tp auth.TokenProvider
  tp.Token(context.WithValue(context.Background(), key, token))
  opts := &auth.CredentialsOptions{
    TokenProvider: tp,
  }
  creds := auth.NewCredentials(opts)
  return creds
}

func getToken(r *http.Request) string {
  return utils.ExtractToken(r)
}

// func (tp *auth.TokenProvider) getToken(r *http.Request) auth.Token {
//   t := utils.ExtractToken(r)
//   token := &auth.Token{
//     Value: t,
//     Type: "Bearer",
//   }
//   return *token
// }
