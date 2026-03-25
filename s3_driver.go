package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Driver struct {
	client *s3.Client
	bucket string
}

func NewS3Driver(client *s3.Client, bucket string) *S3Driver {
	return &S3Driver{client: client, bucket: bucket}
}

func (s *S3Driver) Upload(ctx context.Context, path string, file io.Reader) (string, error) {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	return path, nil // Retornas la ruta o URL
}

// ... implementar GetURL y Delete
