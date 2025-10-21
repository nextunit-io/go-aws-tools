package awstools

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Interface interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

var s3Instance S3Interface

func CreateS3Instance(cfg aws.Config) {
	SetS3Instance(s3.NewFromConfig(cfg))
}

func GetS3Instance() S3Interface {
	if s3Instance == nil {
		ctx := context.TODO()

		conf, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			panic("cannot create a new s3 instance")
		}

		s3Instance = s3.NewFromConfig(conf)
	}

	return s3Instance
}

func SetS3Instance(client S3Interface) {
	s3Instance = client
}
