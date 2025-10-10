package awstoolsmock_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestSSMMockDeleteParameter(t *testing.T) {
	t.Helper()
	ssmMock := awstoolsmock.GetSSMMock()
	ctx := context.TODO()

	t.Run("Testing DeleteParameter", func(t *testing.T) {
		output := ssm.DeleteParameterOutput{}

		ssmMock.Mock.DeleteParameter.AddReturnValue(&output)
		ssmMock.Mock.DeleteParameter.AddReturnValue(&output)
		ssmMock.Mock.DeleteParameter.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := ssmMock.DeleteParameter(ctx, &ssm.DeleteParameterInput{
				Name: aws.String(fmt.Sprintf("test-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.DeleteParameter(ctx, &ssm.DeleteParameterInput{
			Name: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("DeleteParameter general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.DeleteParameter.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.DeleteParameter.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Name)
		}

		input := ssmMock.Mock.DeleteParameter.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Name)
	})
}

func TestSSMMockOtherMethods(t *testing.T) {
	t.Helper()
	ssmMock := awstoolsmock.GetSSMMock()
	ctx := context.TODO()

	t.Run("Testing DeleteParameters", func(t *testing.T) {
		output := ssm.DeleteParametersOutput{}

		ssmMock.Mock.DeleteParameters.AddReturnValue(&output)
		ssmMock.Mock.DeleteParameters.AddReturnValue(&output)
		ssmMock.Mock.DeleteParameters.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			name := fmt.Sprintf("test-%d", i)
			o, err := ssmMock.DeleteParameters(ctx, &ssm.DeleteParametersInput{
				Names: []string{name},
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.DeleteParameters(ctx, &ssm.DeleteParametersInput{
			Names: []string{"error"},
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("DeleteParameters general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.DeleteParameters.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.DeleteParameters.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), input.Params.Names[0])
		}

		input := ssmMock.Mock.DeleteParameters.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", input.Params.Names[0])
	})

	t.Run("Testing DescribeParameters", func(t *testing.T) {
		output := ssm.DescribeParametersOutput{}

		ssmMock.Mock.DescribeParameters.AddReturnValue(&output)
		ssmMock.Mock.DescribeParameters.AddReturnValue(&output)
		ssmMock.Mock.DescribeParameters.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			max := int32(i + 1)
			o, err := ssmMock.DescribeParameters(ctx, &ssm.DescribeParametersInput{
				MaxResults: aws.Int32(max),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.DescribeParameters(ctx, &ssm.DescribeParametersInput{
			MaxResults: aws.Int32(999),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("DescribeParameters general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.DescribeParameters.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.DescribeParameters.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, int32(i+1), *input.Params.MaxResults)
		}

		input := ssmMock.Mock.DescribeParameters.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, int32(999), *input.Params.MaxResults)
	})

	t.Run("Testing GetParameter", func(t *testing.T) {
		output := ssm.GetParameterOutput{}

		ssmMock.Mock.GetParameter.AddReturnValue(&output)
		ssmMock.Mock.GetParameter.AddReturnValue(&output)
		ssmMock.Mock.GetParameter.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			name := fmt.Sprintf("test-%d", i)
			o, err := ssmMock.GetParameter(ctx, &ssm.GetParameterInput{
				Name: aws.String(name),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.GetParameter(ctx, &ssm.GetParameterInput{
			Name: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("GetParameter general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.GetParameter.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.GetParameter.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Name)
		}

		input := ssmMock.Mock.GetParameter.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Name)
	})

	t.Run("Testing GetParameterHistory", func(t *testing.T) {
		output := ssm.GetParameterHistoryOutput{}

		ssmMock.Mock.GetParameterHistory.AddReturnValue(&output)
		ssmMock.Mock.GetParameterHistory.AddReturnValue(&output)
		ssmMock.Mock.GetParameterHistory.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			name := fmt.Sprintf("test-%d", i)
			o, err := ssmMock.GetParameterHistory(ctx, &ssm.GetParameterHistoryInput{
				Name: aws.String(name),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.GetParameterHistory(ctx, &ssm.GetParameterHistoryInput{
			Name: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("GetParameterHistory general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.GetParameterHistory.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.GetParameterHistory.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Name)
		}

		input := ssmMock.Mock.GetParameterHistory.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Name)
	})

	t.Run("Testing GetParameters", func(t *testing.T) {
		output := ssm.GetParametersOutput{}

		ssmMock.Mock.GetParameters.AddReturnValue(&output)
		ssmMock.Mock.GetParameters.AddReturnValue(&output)
		ssmMock.Mock.GetParameters.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			name := fmt.Sprintf("test-%d", i)
			o, err := ssmMock.GetParameters(ctx, &ssm.GetParametersInput{
				Names: []string{name},
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.GetParameters(ctx, &ssm.GetParametersInput{
			Names: []string{"error"},
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("GetParameters general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.GetParameters.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.GetParameters.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), input.Params.Names[0])
		}

		input := ssmMock.Mock.GetParameters.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", input.Params.Names[0])
	})

	t.Run("Testing GetParametersByPath", func(t *testing.T) {
		output := ssm.GetParametersByPathOutput{}

		ssmMock.Mock.GetParametersByPath.AddReturnValue(&output)
		ssmMock.Mock.GetParametersByPath.AddReturnValue(&output)
		ssmMock.Mock.GetParametersByPath.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			path := fmt.Sprintf("/test/%d", i)
			o, err := ssmMock.GetParametersByPath(ctx, &ssm.GetParametersByPathInput{
				Path: aws.String(path),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.GetParametersByPath(ctx, &ssm.GetParametersByPathInput{
			Path: aws.String("/error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("GetParametersByPath general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.GetParametersByPath.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.GetParametersByPath.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("/test/%d", i), *input.Params.Path)
		}

		input := ssmMock.Mock.GetParametersByPath.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "/error", *input.Params.Path)
	})

	t.Run("Testing LabelParameterVersion", func(t *testing.T) {
		output := ssm.LabelParameterVersionOutput{}

		ssmMock.Mock.LabelParameterVersion.AddReturnValue(&output)
		ssmMock.Mock.LabelParameterVersion.AddReturnValue(&output)
		ssmMock.Mock.LabelParameterVersion.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			name := fmt.Sprintf("test-%d", i)
			o, err := ssmMock.LabelParameterVersion(ctx, &ssm.LabelParameterVersionInput{
				Name: aws.String(name),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.LabelParameterVersion(ctx, &ssm.LabelParameterVersionInput{
			Name: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("LabelParameterVersion general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.LabelParameterVersion.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.LabelParameterVersion.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Name)
		}

		input := ssmMock.Mock.LabelParameterVersion.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Name)
	})

	t.Run("Testing PutParameter", func(t *testing.T) {
		output := ssm.PutParameterOutput{}

		ssmMock.Mock.PutParameter.AddReturnValue(&output)
		ssmMock.Mock.PutParameter.AddReturnValue(&output)
		ssmMock.Mock.PutParameter.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			name := fmt.Sprintf("test-%d", i)
			value := fmt.Sprintf("value-%d", i)
			o, err := ssmMock.PutParameter(ctx, &ssm.PutParameterInput{
				Name:  aws.String(name),
				Value: aws.String(value),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.PutParameter(ctx, &ssm.PutParameterInput{
			Name:  aws.String("error"),
			Value: aws.String("err"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("PutParameter general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.PutParameter.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.PutParameter.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Name)
			assert.Equal(t, fmt.Sprintf("value-%d", i), *input.Params.Value)
		}

		input := ssmMock.Mock.PutParameter.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Name)
		assert.Equal(t, "err", *input.Params.Value)
	})

	t.Run("Testing UnlabelParameterVersion", func(t *testing.T) {
		output := ssm.UnlabelParameterVersionOutput{}

		ssmMock.Mock.UnlabelParameterVersion.AddReturnValue(&output)
		ssmMock.Mock.UnlabelParameterVersion.AddReturnValue(&output)
		ssmMock.Mock.UnlabelParameterVersion.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			name := fmt.Sprintf("test-%d", i)
			o, err := ssmMock.UnlabelParameterVersion(ctx, &ssm.UnlabelParameterVersionInput{
				Name: aws.String(name),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := ssmMock.UnlabelParameterVersion(ctx, &ssm.UnlabelParameterVersionInput{
			Name: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("UnlabelParameterVersion general error"), err)

		assert.Equal(t, 4, ssmMock.Mock.UnlabelParameterVersion.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := ssmMock.Mock.UnlabelParameterVersion.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Name)
		}

		input := ssmMock.Mock.UnlabelParameterVersion.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Name)
	})
}
