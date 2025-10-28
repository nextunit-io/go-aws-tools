package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/nextunit-io/go-mock"
)

type s3MockStruct struct {
	GetObject *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.GetObjectInput
		OptFns []func(*s3.Options)
	},
		s3.GetObjectOutput,
	]
	HeadObject *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.HeadObjectInput
		OptFns []func(*s3.Options)
	},
		s3.HeadObjectOutput,
	]
	ListObjectsV2 *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.ListObjectsV2Input
		OptFns []func(*s3.Options)
	},
		s3.ListObjectsV2Output,
	]
	PutObject *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.PutObjectInput
		OptFns []func(*s3.Options)
	},
		s3.PutObjectOutput,
	]
	DeleteObject *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.DeleteObjectInput
		OptFns []func(*s3.Options)
	},
		s3.DeleteObjectOutput,
	]
}

type S3Mock struct {
	Mock s3MockStruct
}

func GetS3Mock() *S3Mock {
	return &S3Mock{
		Mock: s3MockStruct{
			GetObject: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.GetObjectInput
				OptFns []func(*s3.Options)
			},
				s3.GetObjectOutput,
			](fmt.Errorf("GetObject general error")),
			HeadObject: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.HeadObjectInput
				OptFns []func(*s3.Options)
			},
				s3.HeadObjectOutput,
			](fmt.Errorf("HeadObject general error")),
			ListObjectsV2: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.ListObjectsV2Input
				OptFns []func(*s3.Options)
			},
				s3.ListObjectsV2Output,
			](fmt.Errorf("ListObjectsV2 general error")),
			PutObject: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.PutObjectInput
				OptFns []func(*s3.Options)
			},
				s3.PutObjectOutput,
			](fmt.Errorf("PutObject general error")),
			DeleteObject: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.DeleteObjectInput
				OptFns []func(*s3.Options)
			},
				s3.DeleteObjectOutput,
			](fmt.Errorf("DeleteObject general error")),
		},
	}
}

func (mock *S3Mock) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	mock.Mock.GetObject.AddInput(
		struct {
			Ctx    context.Context
			Params *s3.GetObjectInput
			OptFns []func(*s3.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.GetObject.GetNextResult()
}

func (mock *S3Mock) HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	mock.Mock.HeadObject.AddInput(
		struct {
			Ctx    context.Context
			Params *s3.HeadObjectInput
			OptFns []func(*s3.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.HeadObject.GetNextResult()
}

func (mock *S3Mock) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	mock.Mock.ListObjectsV2.AddInput(
		struct {
			Ctx    context.Context
			Params *s3.ListObjectsV2Input
			OptFns []func(*s3.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.ListObjectsV2.GetNextResult()
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

func (mock *S3Mock) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	mock.Mock.DeleteObject.AddInput(
		struct {
			Ctx    context.Context
			Params *s3.DeleteObjectInput
			OptFns []func(*s3.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.DeleteObject.GetNextResult()
}
