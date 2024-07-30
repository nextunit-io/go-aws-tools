package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamodbMockStruct struct {
	GetItem *awsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.GetItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.GetItemOutput,
	]
	BatchWriteItem *awsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.BatchWriteItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.BatchWriteItemOutput,
	]
	PutItem *awsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.PutItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.PutItemOutput,
	]
	UpdateItem *awsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.UpdateItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.UpdateItemOutput,
	]
	Query *awsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.QueryInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.QueryOutput,
	]
	Scan *awsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.ScanInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.ScanOutput,
	]
	DeleteItem *awsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.DeleteItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.DeleteItemOutput,
	]
}

type dynamodbMock struct {
	Mock dynamodbMockStruct
}

func GetDynamodbMock() *dynamodbMock {
	return &dynamodbMock{
		Mock: dynamodbMockStruct{

			GetItem: GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.GetItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.GetItemOutput,
			](fmt.Errorf("GetItem general error")),

			BatchWriteItem: GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.BatchWriteItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.BatchWriteItemOutput,
			](fmt.Errorf("BatchWriteItem general error")),

			PutItem: GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.PutItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.PutItemOutput,
			](fmt.Errorf("PutItem general error")),

			UpdateItem: GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.UpdateItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.UpdateItemOutput,
			](fmt.Errorf("UpdateItem general error")),

			Query: GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.QueryInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.QueryOutput,
			](fmt.Errorf("Query general error")),

			Scan: GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.ScanInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.ScanOutput,
			](fmt.Errorf("Scan general error")),

			DeleteItem: GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.DeleteItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.DeleteItemOutput,
			](fmt.Errorf("DeleteItem general error")),
		},
	}
}

func (d *dynamodbMock) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(options *dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	d.Mock.GetItem.addInput(
		struct {
			Ctx    context.Context
			Params *dynamodb.GetItemInput
			OptFns []func(*dynamodb.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return d.Mock.GetItem.GetNextResult()
}
func (d *dynamodbMock) BatchWriteItem(ctx context.Context, params *dynamodb.BatchWriteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error) {
	d.Mock.BatchWriteItem.addInput(
		struct {
			Ctx    context.Context
			Params *dynamodb.BatchWriteItemInput
			OptFns []func(*dynamodb.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return d.Mock.BatchWriteItem.GetNextResult()
}
func (d *dynamodbMock) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	d.Mock.PutItem.addInput(
		struct {
			Ctx    context.Context
			Params *dynamodb.PutItemInput
			OptFns []func(*dynamodb.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return d.Mock.PutItem.GetNextResult()
}
func (d *dynamodbMock) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	d.Mock.UpdateItem.addInput(
		struct {
			Ctx    context.Context
			Params *dynamodb.UpdateItemInput
			OptFns []func(*dynamodb.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return d.Mock.UpdateItem.GetNextResult()
}
func (d *dynamodbMock) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	d.Mock.Query.addInput(
		struct {
			Ctx    context.Context
			Params *dynamodb.QueryInput
			OptFns []func(*dynamodb.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return d.Mock.Query.GetNextResult()
}
func (d *dynamodbMock) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	d.Mock.Scan.addInput(
		struct {
			Ctx    context.Context
			Params *dynamodb.ScanInput
			OptFns []func(*dynamodb.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return d.Mock.Scan.GetNextResult()
}
func (d *dynamodbMock) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	d.Mock.DeleteItem.addInput(
		struct {
			Ctx    context.Context
			Params *dynamodb.DeleteItemInput
			OptFns []func(*dynamodb.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return d.Mock.DeleteItem.GetNextResult()
}
