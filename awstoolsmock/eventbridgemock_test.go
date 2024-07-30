package awstoolsmock_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestEventbridgeMockPutEvents(t *testing.T) {
	t.Helper()
	eventbridgeMock := awstoolsmock.GetEventbridgeMock()
	ctx := context.TODO()

	t.Run("Testing PutEvents", func(t *testing.T) {
		output := eventbridge.PutEventsOutput{
			Entries: []types.PutEventsResultEntry{
				{
					EventId: aws.String("test-output"),
				},
			},
		}

		eventbridgeMock.Mock.PutEvents.AddReturnValue(&output)
		eventbridgeMock.Mock.PutEvents.AddReturnValue(&output)
		eventbridgeMock.Mock.PutEvents.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := eventbridgeMock.PutEvents(ctx, &eventbridge.PutEventsInput{
				Entries: []types.PutEventsRequestEntry{
					{
						Detail: aws.String(fmt.Sprintf("test-%d", i)),
					},
				},
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := eventbridgeMock.PutEvents(ctx, &eventbridge.PutEventsInput{
			Entries: []types.PutEventsRequestEntry{
				{
					Detail: aws.String("error"),
				},
			},
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("PutEvents general error"), err)

		assert.Equal(t, 4, eventbridgeMock.Mock.PutEvents.HaveBeenCalled())
		for i := 0; i < 3; i++ {
			input := eventbridgeMock.Mock.PutEvents.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Entries[0].Detail)
		}

		input := eventbridgeMock.Mock.PutEvents.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Entries[0].Detail)
	})
}
