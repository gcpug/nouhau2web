package main

import (
	"context"

	"cloud.google.com/go/storage"
)

type StorageService struct {
	gcs *storage.Client
}

func NewStorageService(ctx context.Context) (*StorageService, error) {
	gcs, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &StorageService{gcs: gcs}, nil
}

func (s *StorageService) NewWriter(ctx context.Context, bucket string, object string) *storage.Writer {
	return s.gcs.Bucket(bucket).Object(object).NewWriter(ctx)
}
