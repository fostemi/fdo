package bucket

import (
  "net/http"
  "context"

  // "github.com/fostemi/fdo/internal/common/utils"

	"github.com/gin-gonic/gin"

	"cloud.google.com/go/auth"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func Create(c *gin.Context) error {
  ctx := context.Background()
  creds := getCredsFromToken(c.Request)
  _, err := storage.NewClient(ctx, option.WithAuthCredentials(creds))
  if err != nil {
    return err
  }
  return nil
}

func getCredsFromToken(r *http.Request) *auth.Credentials{
  opts := auth.CredentialsOptions{
    TokenProvider: _,
  }
  creds := auth.NewCredentials(opts)
  return creds
}
