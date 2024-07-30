package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

type eventbridgeMockStruct struct {
	PutEvents *awsToolMock[struct {
		Ctx    context.Context
		Params *eventbridge.PutEventsInput
		OptFns []func(*eventbridge.Options)
	},
		eventbridge.PutEventsOutput,
	]
}

type eventbridgeMock struct {
	Mock eventbridgeMockStruct
}

func GetEventbridgeMock() *eventbridgeMock {
	return &eventbridgeMock{
		Mock: eventbridgeMockStruct{
			PutEvents: GetMock[struct {
				Ctx    context.Context
				Params *eventbridge.PutEventsInput
				OptFns []func(*eventbridge.Options)
			},
				eventbridge.PutEventsOutput,
			](fmt.Errorf("PutEvents general error")),
		},
	}
}

func (e *eventbridgeMock) PutEvents(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error) {
	e.Mock.PutEvents.addInput(
		struct {
			Ctx    context.Context
			Params *eventbridge.PutEventsInput
			OptFns []func(*eventbridge.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return e.Mock.PutEvents.GetNextResult()
}
