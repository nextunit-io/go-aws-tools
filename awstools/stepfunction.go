package awstools

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

type StepFunctionInterface interface {
	StartExecution(ctx context.Context, params *sfn.StartExecutionInput, optFns ...func(*sfn.Options)) (*sfn.StartExecutionOutput, error)
}

var stepFunctionInstance StepFunctionInterface

func CreateStepFunctionInstance(cfg aws.Config) {
	SetStepFunctionInstance(sfn.NewFromConfig(cfg))
}

func GetStepFunctionInstance() StepFunctionInterface {
	if stepFunctionInstance == nil {
		ctx := context.TODO()

		conf, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			panic("cannot create a new step function instance")
		}

		CreateStepFunctionInstance(conf)
	}

	return stepFunctionInstance
}

func SetStepFunctionInstance(client StepFunctionInterface) {
	stepFunctionInstance = client
}
