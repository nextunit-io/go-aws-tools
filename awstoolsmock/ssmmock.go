package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	gomock "github.com/nextunit-io/go-mock"
)

type ssmMockStruct struct {
	DeleteParameter *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.DeleteParameterInput
		OptFns []func(*ssm.Options)
	},
		ssm.DeleteParameterOutput,
	]

	DeleteParameters *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.DeleteParametersInput
		OptFns []func(*ssm.Options)
	},
		ssm.DeleteParametersOutput,
	]

	DescribeParameters *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.DescribeParametersInput
		OptFns []func(*ssm.Options)
	},
		ssm.DescribeParametersOutput,
	]

	GetParameter *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.GetParameterInput
		OptFns []func(*ssm.Options)
	},
		ssm.GetParameterOutput,
	]

	GetParameterHistory *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.GetParameterHistoryInput
		OptFns []func(*ssm.Options)
	},
		ssm.GetParameterHistoryOutput,
	]

	GetParameters *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.GetParametersInput
		OptFns []func(*ssm.Options)
	},
		ssm.GetParametersOutput,
	]

	GetParametersByPath *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.GetParametersByPathInput
		OptFns []func(*ssm.Options)
	},
		ssm.GetParametersByPathOutput,
	]

	LabelParameterVersion *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.LabelParameterVersionInput
		OptFns []func(*ssm.Options)
	},
		ssm.LabelParameterVersionOutput,
	]

	PutParameter *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.PutParameterInput
		OptFns []func(*ssm.Options)
	},
		ssm.PutParameterOutput,
	]

	UnlabelParameterVersion *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *ssm.UnlabelParameterVersionInput
		OptFns []func(*ssm.Options)
	},
		ssm.UnlabelParameterVersionOutput,
	]
}

type SSMMock struct {
	Mock ssmMockStruct
}

func GetSSMMock() *SSMMock {
	return &SSMMock{
		Mock: ssmMockStruct{
			DeleteParameter: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.DeleteParameterInput
				OptFns []func(*ssm.Options)
			},
				ssm.DeleteParameterOutput,
			](fmt.Errorf("DeleteParameter general error")),

			DeleteParameters: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.DeleteParametersInput
				OptFns []func(*ssm.Options)
			},
				ssm.DeleteParametersOutput,
			](fmt.Errorf("DeleteParameters general error")),

			DescribeParameters: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.DescribeParametersInput
				OptFns []func(*ssm.Options)
			},
				ssm.DescribeParametersOutput,
			](fmt.Errorf("DescribeParameters general error")),

			GetParameter: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.GetParameterInput
				OptFns []func(*ssm.Options)
			},
				ssm.GetParameterOutput,
			](fmt.Errorf("GetParameter general error")),

			GetParameterHistory: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.GetParameterHistoryInput
				OptFns []func(*ssm.Options)
			},
				ssm.GetParameterHistoryOutput,
			](fmt.Errorf("GetParameterHistory general error")),

			GetParameters: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.GetParametersInput
				OptFns []func(*ssm.Options)
			},
				ssm.GetParametersOutput,
			](fmt.Errorf("GetParameters general error")),

			GetParametersByPath: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.GetParametersByPathInput
				OptFns []func(*ssm.Options)
			},
				ssm.GetParametersByPathOutput,
			](fmt.Errorf("GetParametersByPath general error")),

			LabelParameterVersion: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.LabelParameterVersionInput
				OptFns []func(*ssm.Options)
			},
				ssm.LabelParameterVersionOutput,
			](fmt.Errorf("LabelParameterVersion general error")),

			PutParameter: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.PutParameterInput
				OptFns []func(*ssm.Options)
			},
				ssm.PutParameterOutput,
			](fmt.Errorf("PutParameter general error")),

			UnlabelParameterVersion: gomock.GetMock[struct {
				Ctx    context.Context
				Params *ssm.UnlabelParameterVersionInput
				OptFns []func(*ssm.Options)
			},
				ssm.UnlabelParameterVersionOutput,
			](fmt.Errorf("UnlabelParameterVersion general error")),
		},
	}
}

func (mock *SSMMock) DeleteParameter(ctx context.Context, params *ssm.DeleteParameterInput, optFns ...func(*ssm.Options)) (*ssm.DeleteParameterOutput, error) {
	mock.Mock.DeleteParameter.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.DeleteParameterInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.DeleteParameter.GetNextResult()
}

func (mock *SSMMock) DeleteParameters(ctx context.Context, params *ssm.DeleteParametersInput, optFns ...func(*ssm.Options)) (*ssm.DeleteParametersOutput, error) {
	mock.Mock.DeleteParameters.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.DeleteParametersInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.DeleteParameters.GetNextResult()
}

func (mock *SSMMock) DescribeParameters(ctx context.Context, params *ssm.DescribeParametersInput, optFns ...func(*ssm.Options)) (*ssm.DescribeParametersOutput, error) {
	mock.Mock.DescribeParameters.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.DescribeParametersInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.DescribeParameters.GetNextResult()
}

func (mock *SSMMock) GetParameter(ctx context.Context, params *ssm.GetParameterInput, optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error) {
	mock.Mock.GetParameter.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.GetParameterInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.GetParameter.GetNextResult()
}

func (mock *SSMMock) GetParameterHistory(ctx context.Context, params *ssm.GetParameterHistoryInput, optFns ...func(*ssm.Options)) (*ssm.GetParameterHistoryOutput, error) {
	mock.Mock.GetParameterHistory.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.GetParameterHistoryInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.GetParameterHistory.GetNextResult()
}

func (mock *SSMMock) GetParameters(ctx context.Context, params *ssm.GetParametersInput, optFns ...func(*ssm.Options)) (*ssm.GetParametersOutput, error) {
	mock.Mock.GetParameters.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.GetParametersInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.GetParameters.GetNextResult()
}

func (mock *SSMMock) GetParametersByPath(ctx context.Context, params *ssm.GetParametersByPathInput, optFns ...func(*ssm.Options)) (*ssm.GetParametersByPathOutput, error) {
	mock.Mock.GetParametersByPath.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.GetParametersByPathInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.GetParametersByPath.GetNextResult()
}

func (mock *SSMMock) LabelParameterVersion(ctx context.Context, params *ssm.LabelParameterVersionInput, optFns ...func(*ssm.Options)) (*ssm.LabelParameterVersionOutput, error) {
	mock.Mock.LabelParameterVersion.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.LabelParameterVersionInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.LabelParameterVersion.GetNextResult()
}

func (mock *SSMMock) PutParameter(ctx context.Context, params *ssm.PutParameterInput, optFns ...func(*ssm.Options)) (*ssm.PutParameterOutput, error) {
	mock.Mock.PutParameter.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.PutParameterInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.PutParameter.GetNextResult()
}

func (mock *SSMMock) UnlabelParameterVersion(ctx context.Context, params *ssm.UnlabelParameterVersionInput, optFns ...func(*ssm.Options)) (*ssm.UnlabelParameterVersionOutput, error) {
	mock.Mock.UnlabelParameterVersion.AddInput(
		struct {
			Ctx    context.Context
			Params *ssm.UnlabelParameterVersionInput
			OptFns []func(*ssm.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return mock.Mock.UnlabelParameterVersion.GetNextResult()
}
