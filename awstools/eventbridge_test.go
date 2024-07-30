package awstools_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/nextunit-io/go-aws-tools/awstools"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestSetEventBridgeClient(t *testing.T) {
	client := awstoolsmock.GetEventbridgeMock()
	awstools.SetEventBridgeInstance(client)
	assert.Equal(t, client, awstools.GetEventBridgeInstance())
}

func TestWillSetupEventbridgeClientWithoutSettingIt(t *testing.T) {
	awstools.SetEventBridgeInstance(nil)

	assert.NotNil(t, awstools.GetEventBridgeInstance())
}

func TestCanCreateEventBridgeInstanceByConfig(t *testing.T) {
	oldClient := awstools.GetEventBridgeInstance()

	ctx := context.TODO()
	conf, _ := config.LoadDefaultConfig(ctx)

	awstools.CreateEventBridgeInstance(conf)

	assert.NotEqual(t, oldClient, awstools.GetEventBridgeInstance())
}
