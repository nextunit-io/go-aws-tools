package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	gomock "github.com/nextunit-io/go-mock"
)

type stepfunctionMockStruct struct {
	StartExecution *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *sfn.StartExecutionInput
		OptFns []func(*sfn.Options)
	},
		sfn.StartExecutionOutput,
	]
}

type StepFunctionMock struct {
	Mock stepfunctionMockStruct
}

func GetStepFunctionMock() *StepFunctionMock {
	return &StepFunctionMock{
		Mock: stepfunctionMockStruct{
			StartExecution: gomock.GetMock[struct {
				Ctx    context.Context
				Params *sfn.StartExecutionInput
				OptFns []func(*sfn.Options)
			},
				sfn.StartExecutionOutput,
			](fmt.Errorf("StartExceution general error")),
		},
	}
}

func (s *StepFunctionMock) StartExecution(ctx context.Context, params *sfn.StartExecutionInput, optFns ...func(*sfn.Options)) (*sfn.StartExecutionOutput, error) {
	s.Mock.StartExecution.AddInput(
		struct {
			Ctx    context.Context
			Params *sfn.StartExecutionInput
			OptFns []func(*sfn.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return s.Mock.StartExecution.GetNextResult()
}
