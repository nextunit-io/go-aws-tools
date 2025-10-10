package awstools_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/nextunit-io/go-aws-tools/awstools"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestSetSSMClient(t *testing.T) {
	client := awstoolsmock.GetSSMMock()
	awstools.SetSSMInstance(client)
	assert.Equal(t, client, awstools.GetSSMInstance())
}

func TestWillSetupSSMClientWithoutSettingIt(t *testing.T) {
	awstools.SetSSMInstance(nil)

	assert.NotNil(t, awstools.GetSSMInstance())
}

func TestCanCreateSSMInstanceByConfig(t *testing.T) {
	oldClient := awstools.GetSSMInstance()

	ctx := context.TODO()
	conf, _ := config.LoadDefaultConfig(ctx)

	awstools.CreateSSMInstance(conf)

	assert.NotEqual(t, oldClient, awstools.GetSSMInstance())
}
