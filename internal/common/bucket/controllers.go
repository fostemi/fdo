package bucket

import (
	"context"
	"net/http"

	// "github.com/fostemi/fdo/internal/common/utils"

	"github.com/fostemi/fdo/internal/common/utils"
	"github.com/gin-gonic/gin"

	"cloud.google.com/go/auth"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func Create(c *gin.Context) error {
  ctx := context.Background()
  creds := getCredsFromToken(ctx, getToken(c.Request))
  _, err := storage.NewClient(ctx, option.WithAuthCredentials(creds))
  if err != nil {
    return err
  }
  return nil
}

func getCredsFromToken(ctx context.Context, token string) *auth.Credentials{
  t := auth.Token{
    Value: token,
    Type: "Bearer",
  }
  key := auth.Token{}
  var tp auth.TokenProvider
  tp.Token(context.WithValue(ctx, key, t))
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
