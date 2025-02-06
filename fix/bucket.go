package main

import (
  "time"
  "fmt"
  "context"

  "cloud.google.com/go/storage"
)

func createBucket(ctx context.Context, projectID, bucketName string) error {
  client, err := storage.NewClient(ctx)
  if err != nil {
    return err
  }
  defer client.Close()

  bucket := client.Bucket(bucketName)

  // this is just a child of original ctx
  ctx, cancel := context.WithTimeout(ctx, time.Second*10)
  defer cancel()
  if err := bucket.Create(ctx, projectID, nil); err != nil {
    return err
  }
  fmt.Println("Successfully created bucket")

  return nil
}

func generateCredsForBucket() {

}

func deleteBucket(ctx context.Context, bucketName string) error {
  client, err := storage.NewClient(ctx)
  if err != nil {
    return err
  }
  defer client.Close()

  bucket := client.Bucket(bucketName)

  if err := bucket.Delete(ctx); err != nil {
    return err
  }
  fmt.Println("Successfully deleted bucket")

  return nil
}

