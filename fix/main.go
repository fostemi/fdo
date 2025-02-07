package main

import (
	"context"
	"flag"
	"fmt"
)

func main() {
  flag.Parse()
  ctx := context.Background()

  projectID := "hydra-alert-dev-78gx"
  bucketName := "hydra-alert-dev-78gx-tst-00000"

  err := createBucket(ctx, projectID, bucketName)
  if err != nil {
    fmt.Println("Error creating bucket: ", err)
  }
}
