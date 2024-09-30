package awstools_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/nextunit-io/go-aws-tools/awstools"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestSetS3Client(t *testing.T) {
	client := awstoolsmock.GetS3Mock()
	awstools.SetS3Instance(client)
	assert.Equal(t, client, awstools.GetS3Instance())
}

func TestWillSetupS3ClientWithoutSettingIt(t *testing.T) {
	awstools.SetEventBridgeInstance(nil)

	assert.NotNil(t, awstools.GetS3Instance())
}

func TestCanCreateS3InstanceByConfig(t *testing.T) {
	oldClient := awstools.GetS3Instance()

	ctx := context.TODO()
	conf, _ := config.LoadDefaultConfig(ctx)

	awstools.CreateS3Instance(conf)

	assert.NotEqual(t, oldClient, awstools.GetS3Instance())
}
