// connector.go
package drivers

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Configuraciones basicas para S3
type S3Config struct {
	AccessKey    string
	SecretKey    string
	Region       string
	Bucket       string
	Endpoint     string
	UsePathStyle bool
}

// ConnectS3 inicializa el cliente y devuelve el Driver listo
func ConnectS3(conf S3Config) (*S3Driver, error) {
	creds := credentials.NewStaticCredentialsProvider(conf.AccessKey, conf.SecretKey, "")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(conf.Region),
		config.WithCredentialsProvider(creds),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		if conf.Endpoint != "" {
			o.BaseEndpoint = aws.String(conf.Endpoint)
		}
		o.UsePathStyle = conf.UsePathStyle
	})

	return NewS3Driver(client, conf.Bucket), nil
}
