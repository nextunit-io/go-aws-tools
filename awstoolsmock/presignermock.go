package awstoolsmock

import (
	"context"
	"fmt"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/nextunit-io/go-mock"
)

type presignMockStruct struct {
	PresignGetObject *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.GetObjectInput
		OptFns []func(*s3.PresignOptions)
	},
		v4.PresignedHTTPRequest,
	]
	PresignDeleteObject *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.DeleteObjectInput
		OptFns []func(*s3.PresignOptions)
	},
		v4.PresignedHTTPRequest,
	]
	PresignPutObject *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.PutObjectInput
		OptFns []func(*s3.PresignOptions)
	},
		v4.PresignedHTTPRequest,
	]
	PresignPostObject *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *s3.PutObjectInput
		OptFns []func(*s3.PresignPostOptions)
	},
		v4.PresignedHTTPRequest,
	]
}

type PresignMock struct {
	Mock presignMockStruct
}

func GetPresignerMock() *PresignMock {
	return &PresignMock{
		Mock: presignMockStruct{
			PresignGetObject: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.GetObjectInput
				OptFns []func(*s3.PresignOptions)
			},
				v4.PresignedHTTPRequest,
			](fmt.Errorf("PresignGetObject general error")),
			PresignDeleteObject: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.DeleteObjectInput
				OptFns []func(*s3.PresignOptions)
			},
				v4.PresignedHTTPRequest,
			](fmt.Errorf("PresignDeleteObject general error")),
			PresignPutObject: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.PutObjectInput
				OptFns []func(*s3.PresignOptions)
			},
				v4.PresignedHTTPRequest,
			](fmt.Errorf("PresignPutObject general error")),
			PresignPostObject: gomock.GetMock[struct {
				Ctx    context.Context
				Params *s3.PutObjectInput
				OptFns []func(*s3.PresignPostOptions)
			},
				v4.PresignedHTTPRequest,
			](fmt.Errorf("PresignPostObject general error")),
		},
	}
}

func (p *PresignMock) PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	p.Mock.PresignGetObject.AddInput(
		struct {
			Ctx    context.Context
			Params *s3.GetObjectInput
			OptFns []func(*s3.PresignOptions)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return p.Mock.PresignGetObject.GetNextResult()
}

func (p *PresignMock) PresignDeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	p.Mock.PresignDeleteObject.AddInput(
		struct {
			Ctx    context.Context
			Params *s3.DeleteObjectInput
			OptFns []func(*s3.PresignOptions)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return p.Mock.PresignDeleteObject.GetNextResult()
}

func (p *PresignMock) PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	p.Mock.PresignPutObject.AddInput(
		struct {
			Ctx    context.Context
			Params *s3.PutObjectInput
			OptFns []func(*s3.PresignOptions)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return p.Mock.PresignPutObject.GetNextResult()
}

func (p *PresignMock) PresignPostObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignPostOptions)) (*v4.PresignedHTTPRequest, error) {
	p.Mock.PresignPostObject.AddInput(
		struct {
			Ctx    context.Context
			Params *s3.PutObjectInput
			OptFns []func(*s3.PresignPostOptions)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return p.Mock.PresignPostObject.GetNextResult()
}
