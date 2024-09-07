package awstoolsmock_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestStepFunctionMockStartExecution(t *testing.T) {
	t.Helper()
	stepFunctionMock := awstoolsmock.GetStepFunctionMock()
	ctx := context.TODO()

	t.Run("Testing StartExecution", func(t *testing.T) {
		output := sfn.StartExecutionOutput{
			ExecutionArn: aws.String("test-execution-arn"),
		}

		stepFunctionMock.Mock.StartExecution.AddReturnValue(&output)
		stepFunctionMock.Mock.StartExecution.AddReturnValue(&output)
		stepFunctionMock.Mock.StartExecution.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := stepFunctionMock.StartExecution(ctx, &sfn.StartExecutionInput{
				StateMachineArn: aws.String(fmt.Sprintf("test-statemachine-arn-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := stepFunctionMock.StartExecution(ctx, &sfn.StartExecutionInput{
			StateMachineArn: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("StartExceution general error"), err)

		assert.Equal(t, 4, stepFunctionMock.Mock.StartExecution.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := stepFunctionMock.Mock.StartExecution.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-statemachine-arn-%d", i), *input.Params.StateMachineArn)
		}

		input := stepFunctionMock.Mock.StartExecution.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.StateMachineArn)
	})
}
