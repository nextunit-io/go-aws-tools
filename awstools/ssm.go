package awstools

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMInterface interface {
	DeleteParameter(ctx context.Context, params *ssm.DeleteParameterInput, optFns ...func(*ssm.Options)) (*ssm.DeleteParameterOutput, error)
	DeleteParameters(ctx context.Context, params *ssm.DeleteParametersInput, optFns ...func(*ssm.Options)) (*ssm.DeleteParametersOutput, error)
	DescribeParameters(ctx context.Context, params *ssm.DescribeParametersInput, optFns ...func(*ssm.Options)) (*ssm.DescribeParametersOutput, error)
	GetParameter(ctx context.Context, params *ssm.GetParameterInput, optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
	GetParameterHistory(ctx context.Context, params *ssm.GetParameterHistoryInput, optFns ...func(*ssm.Options)) (*ssm.GetParameterHistoryOutput, error)
	GetParameters(ctx context.Context, params *ssm.GetParametersInput, optFns ...func(*ssm.Options)) (*ssm.GetParametersOutput, error)
	GetParametersByPath(ctx context.Context, params *ssm.GetParametersByPathInput, optFns ...func(*ssm.Options)) (*ssm.GetParametersByPathOutput, error)
	LabelParameterVersion(ctx context.Context, params *ssm.LabelParameterVersionInput, optFns ...func(*ssm.Options)) (*ssm.LabelParameterVersionOutput, error)
	PutParameter(ctx context.Context, params *ssm.PutParameterInput, optFns ...func(*ssm.Options)) (*ssm.PutParameterOutput, error)
	UnlabelParameterVersion(ctx context.Context, params *ssm.UnlabelParameterVersionInput, optFns ...func(*ssm.Options)) (*ssm.UnlabelParameterVersionOutput, error)
}

var ssmInstance SSMInterface

func CreateSSMInstance(cfg aws.Config) {
	SetSSMInstance(ssm.NewFromConfig(cfg))
}

func GetSSMInstance() SSMInterface {
	if ssmInstance == nil {
		ctx := context.TODO()

		conf, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			panic("cannot create a new ssm instance")
		}

		ssmInstance = ssm.NewFromConfig(conf)
	}

	return ssmInstance
}

func SetSSMInstance(client SSMInterface) {
	ssmInstance = client
}
