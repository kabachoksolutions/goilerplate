package storagebp

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// AWSStorage represents an AWS S3 storage instance.
type AWSStorage struct {
	Client *s3.Client
}

type Region string

const (
	RegionWNAM Region = "wnam" // Western North America
	RegionENAM Region = "enam" // Eastern North America
	RegionWEUR Region = "weur" // Western Europe
	RegionEEUR Region = "eeur" // Eastern Europe
	RegionAPAC Region = "apac" // Asia-Pacific
)

func NewAWSStorage(accessKeyID, accessKeySecret, url string, region Region) (*AWSStorage, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: url,
			}, nil
		},
	)

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, ""),
		),
		config.WithRegion(string(region)),
	)
	if err != nil {
		return nil, fmt.Errorf("storagebp: failed to load default config: %w", err)
	}

	return &AWSStorage{
		Client: s3.NewFromConfig(cfg),
	}, nil
}

func (s *AWSStorage) DecodeBase64ToBytes(data string) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(strings.Split(data, "base64,")[1])
	if err != nil {
		return nil, fmt.Errorf("storagebp: failed to decode base64 to bytes: %w", err)
	}

	return decodedData, nil
}

func (s *AWSStorage) Upload(ctx context.Context, name string, tag string, data []byte, bucket string) error {
	var (
		fileBytes = bytes.NewReader(data)
		fileType  = http.DetectContentType(data)
		fileSize  = int64(len(data))
		filePath  = tag + name
	)

	input := &s3.PutObjectInput{
		Body:          fileBytes,
		Bucket:        aws.String(bucket),
		Key:           aws.String(filePath),
		ContentType:   aws.String(fileType),
		ContentLength: &fileSize,
	}

	if _, err := s.Client.PutObject(ctx, input); err != nil {
		return fmt.Errorf("storagebp: failed to put object: %w", err)
	}

	return nil
}

func (s *AWSStorage) Delete(ctx context.Context, fileName string, fileTag string, bucket string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileTag + fileName),
	}

	if _, err := s.Client.DeleteObject(context.Background(), input); err != nil {
		return fmt.Errorf("storagebp: failed to delete object: %w", err)
	}

	return nil
}
