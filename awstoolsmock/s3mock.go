package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/nextunit-io/go-mock"
)

type s3MockStruct struct {
	PutObject *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.PutObjectInput
		OptFns []func(*s3.Options)
	},
		s3.PutObjectOutput,
	]
}

type S3Mock struct {
	Mock s3MockStruct
}

func GetS3Mock() *S3Mock {
	return &S3Mock{
		Mock: s3MockStruct{
			PutObject: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.PutObjectInput
				OptFns []func(*s3.Options)
			},
				s3.PutObjectOutput,
			](fmt.Errorf("PutObject general error")),
		},
	}
}

func (mock *S3Mock) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	mock.Mock.PutObject.AddInput(
		struct {
			Ctx    context.Context
			Params *s3.PutObjectInput
			OptFns []func(*s3.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.PutObject.GetNextResult()
}
