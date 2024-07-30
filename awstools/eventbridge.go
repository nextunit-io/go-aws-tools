package awstools

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

type EventBridgeInterface interface {
	PutEvents(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)
}

var eventBridgeInstance EventBridgeInterface

func CreateEventBridgeInstance(cfg aws.Config) {
	SetEventBridgeInstance(eventbridge.NewFromConfig(cfg))
}

func GetEventBridgeInstance() EventBridgeInterface {
	if eventBridgeInstance == nil {
		ctx := context.TODO()

		conf, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			panic("cannot create a new event bridge instance")
		}

		eventBridgeInstance = eventbridge.NewFromConfig(conf)
	}

	return eventBridgeInstance
}

func SetEventBridgeInstance(client EventBridgeInterface) {
	eventBridgeInstance = client
}
