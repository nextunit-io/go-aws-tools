package awstools

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type PresignInterface interface {
	PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignDeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

var presignerInstance PresignInterface

func CreatePresignInstance(cfg aws.Config) {
	SetPresignInstance(s3.NewPresignClient(s3.NewFromConfig(cfg)))
}

func GetPresignInstance() PresignInterface {
	if presignerInstance == nil {
		ctx := context.TODO()

		conf, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			panic("cannot create a new s3 instance")
		}

		presignerInstance = s3.NewPresignClient(s3.NewFromConfig(conf))
	}

	return presignerInstance
}

func SetPresignInstance(client PresignInterface) {
	presignerInstance = client
}
