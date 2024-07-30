package awstools_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/nextunit-io/go-aws-tools/awstools"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestSetterIsWorking(t *testing.T) {
	testClient := awstoolsmock.GetDynamodbMock()

	awstools.SetDynamoDbInstance(testClient)
	assert.Equal(t, testClient, awstools.GetDynamoDbInstance())
}

func TestGetInstanceWithoutSetting(t *testing.T) {
	awstools.SetDynamoDbInstance(nil)

	assert.NotNil(t, awstools.GetDynamoDbInstance())
}

func TestGetInstanceWithOwnConfig(t *testing.T) {
	awstools.SetDynamoDbInstance(nil)

	ctx := context.TODO()
	conf, _ := config.LoadDefaultConfig(ctx)

	awstools.CreateDynamodbInstance(conf)

	assert.NotNil(t, awstools.GetDynamoDbInstance())
}
