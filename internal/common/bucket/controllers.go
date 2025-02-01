package bucket

import (
	"context"
  "fmt"
	"net/http"

	"github.com/fostemi/fdo/internal/common/utils"
	"github.com/gin-gonic/gin"

	"cloud.google.com/go/auth"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func Create(c *gin.Context) {
  ctx := context.Background()
  creds := getCredsFromToken(ctx, getToken(c.Request))
  _, err := storage.NewClient(ctx, option.WithAuthCredentials(creds))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err})
  }
  fmt.Println(creds)
  c.JSON(http.StatusOK, gin.H{"message": "Success getting token and auth to gcp"})
}

func getCredsFromToken(ctx context.Context, token string) *auth.Credentials{
  type tokenKey string
  key := tokenKey("Key")
  var tp auth.TokenProvider
  tp.Token(context.WithValue(ctx, key, token))
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
