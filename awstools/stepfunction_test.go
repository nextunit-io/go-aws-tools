package awstools_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/nextunit-io/go-aws-tools/awstools"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestSetStepFunctionClient(t *testing.T) {
	client := awstoolsmock.GetStepFunctionMock()
	awstools.SetStepFunctionInstance(client)
	assert.Equal(t, client, awstools.GetStepFunctionInstance())
}

func TestWillSetupStepFunctionClientWithoutSettingIt(t *testing.T) {
	awstools.SetStepFunctionInstance(nil)

	assert.NotNil(t, awstools.GetStepFunctionInstance())
}

func TestCanCreateStepFunctionInstanceByConfig(t *testing.T) {
	oldClient := awstools.GetStepFunctionInstance()

	ctx := context.TODO()
	conf, _ := config.LoadDefaultConfig(ctx)

	awstools.CreateStepFunctionInstance(conf)

	assert.NotEqual(t, oldClient, awstools.GetStepFunctionInstance())
}
