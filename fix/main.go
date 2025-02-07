package main

import (
	"context"
  "os"
	"fmt"
)

func main() {
  args := os.Args
  // for i, arg := range args {
  //   fmt.Printf("Arg %d: %s\n", i, arg)
  // }
  // return
  if len(args) != 2 {
    fmt.Println("Error: Too many arguments, only accepts one")
  }

  ctx := context.Background()

  projectID := "hydra-alert-dev-78gx"
  bucketName := "hydra-alert-dev-78gx-tst-00000"

  switch(args[1]) {
  case "create":
    err := createBucket(ctx, projectID, bucketName)
    if err != nil {
      fmt.Println("Error creating bucket: ", err)
    }
    break
  case "delete":
    err := deleteBucket(ctx, bucketName)
    if err != nil {
      fmt.Println("Error deleting bucket: ", err)
    }
    break
  default:
    fmt.Println("Error: Argument not supported.")
    break
  }

}
