package main

import (
	"context"

	"cloud.google.com/go/auth/credentials"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
  storageScope "cloud.google.com/go/storage/control/apiv2"
)

func JSONCredsStorageClient() (*storage.Client, error) {
  ctx := context.Background()

  creds, err := credentials.DetectDefault(&credentials.DetectOptions{
    Scopes: storageScope.DefaultAuthScopes(),
    CredentialsJSON: []byte("JSON creds"),
  })
  if err != nil {
    return nil, err
  }
  client, err := storage.NewClient(ctx, option.WithAuthCredentials(creds))
  if err != nil {
    return nil, err
  }

  return client, nil
}
