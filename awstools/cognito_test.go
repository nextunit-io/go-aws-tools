package awstools_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/nextunit-io/go-aws-tools/awstools"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestCognitoSetterIsWorking(t *testing.T) {
	testClient := awstoolsmock.GetCognitoMock()

	awstools.SetCognitoInstance(testClient)
	assert.Equal(t, testClient, awstools.GetCognitoInstance())
}

func TestGetCognitoInstanceWithoutSetting(t *testing.T) {
	awstools.SetCognitoInstance(nil)

	assert.NotNil(t, awstools.GetCognitoInstance())
}

func TestGetCognitoInstanceWithOwnConfig(t *testing.T) {
	awstools.SetCognitoInstance(nil)

	ctx := context.TODO()
	conf, _ := config.LoadDefaultConfig(ctx)

	awstools.CreateCognitoInstance(conf)

	assert.NotNil(t, awstools.GetCognitoInstance())
}
