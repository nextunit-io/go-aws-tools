package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	gomock "github.com/nextunit-io/go-mock"
)

type eventbridgeMockStruct struct {
	PutEvents *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *eventbridge.PutEventsInput
		OptFns []func(*eventbridge.Options)
	},
		eventbridge.PutEventsOutput,
	]
}

type EventbridgeMock struct {
	Mock eventbridgeMockStruct
}

func GetEventbridgeMock() *EventbridgeMock {
	return &EventbridgeMock{
		Mock: eventbridgeMockStruct{
			PutEvents: gomock.GetMock[struct {
				Ctx    context.Context
				Params *eventbridge.PutEventsInput
				OptFns []func(*eventbridge.Options)
			},
				eventbridge.PutEventsOutput,
			](fmt.Errorf("PutEvents general error")),
		},
	}
}

func (e *EventbridgeMock) PutEvents(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error) {
	e.Mock.PutEvents.AddInput(
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
