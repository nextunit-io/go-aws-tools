package awstools_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/nextunit-io/go-aws-tools/awstools"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestSetPresignerClient(t *testing.T) {
	client := awstoolsmock.GetPresignerMock()
	awstools.SetPresignInstance(client)
	assert.Equal(t, client, awstools.GetPresignInstance())
}

func TestWillSetupPresignerClientWithoutSettingIt(t *testing.T) {
	awstools.SetEventBridgeInstance(nil)

	assert.NotNil(t, awstools.GetPresignInstance())
}

func TestCanCreatePresignerInstanceByConfig(t *testing.T) {
	oldClient := awstools.GetPresignInstance()

	ctx := context.TODO()
	conf, _ := config.LoadDefaultConfig(ctx)

	awstools.CreatePresignInstance(conf)

	assert.NotEqual(t, oldClient, awstools.GetPresignInstance())
}
