package drivers

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Driver struct {
	client   *s3.Client
	bucket   string
	endpoint string
}

func NewS3Driver(client *s3.Client, bucket string, endpoint string) *S3Driver {
	return &S3Driver{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
	}
}

func (s *S3Driver) Upload(ctx context.Context, path string, file io.Reader) (string, error) {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("error subiendo a S3: %w", err)
	}
	return path, nil // Retornas la ruta o URL
}

func (s *S3Driver) Delete(ctx context.Context, path string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})
	return err
}

func (s *S3Driver) Exists(ctx context.Context, path string) (bool, error) {
	_, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return false, nil // En S3, si HeadObject da error, usualmente es que no existe
	}
	return true, nil
}

func (s *S3Driver) GetURL(path string) string {
	// Si es S3 real, la URL suele ser: https://bucket.s3.region.amazonaws.com/path
	// Si es MinIO, depende de tu endpoint.
	// Por ahora, puedes devolver el path o construir la URL básica:
	if s.endpoint != "" {
		return fmt.Sprintf("%s/%s/%s", s.endpoint, s.bucket, path)
	}
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucket, path)
}
