package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	gomock "github.com/nextunit-io/go-mock"
)

type dynamodbMockStruct struct {
	GetItem *gomock.AwsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.GetItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.GetItemOutput,
	]
	BatchWriteItem *gomock.AwsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.BatchWriteItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.BatchWriteItemOutput,
	]
	PutItem *gomock.AwsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.PutItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.PutItemOutput,
	]
	UpdateItem *gomock.AwsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.UpdateItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.UpdateItemOutput,
	]
	Query *gomock.AwsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.QueryInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.QueryOutput,
	]
	Scan *gomock.AwsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.ScanInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.ScanOutput,
	]
	DeleteItem *gomock.AwsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.DeleteItemInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.DeleteItemOutput,
	]
	DescribeTable *gomock.AwsToolMock[struct {
		Ctx    context.Context
		Params *dynamodb.DescribeTableInput
		OptFns []func(*dynamodb.Options)
	},
		dynamodb.DescribeTableOutput,
	]
}

type DynamodbMock struct {
	Mock dynamodbMockStruct
}

func GetDynamodbMock() *DynamodbMock {
	return &DynamodbMock{
		Mock: dynamodbMockStruct{

			GetItem: gomock.GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.GetItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.GetItemOutput,
			](fmt.Errorf("GetItem general error")),

			BatchWriteItem: gomock.GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.BatchWriteItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.BatchWriteItemOutput,
			](fmt.Errorf("BatchWriteItem general error")),

			PutItem: gomock.GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.PutItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.PutItemOutput,
			](fmt.Errorf("PutItem general error")),

			UpdateItem: gomock.GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.UpdateItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.UpdateItemOutput,
			](fmt.Errorf("UpdateItem general error")),

			Query: gomock.GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.QueryInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.QueryOutput,
			](fmt.Errorf("Query general error")),

			Scan: gomock.GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.ScanInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.ScanOutput,
			](fmt.Errorf("Scan general error")),

			DeleteItem: gomock.GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.DeleteItemInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.DeleteItemOutput,
			](fmt.Errorf("DeleteItem general error")),

			DescribeTable: gomock.GetMock[struct {
				Ctx    context.Context
				Params *dynamodb.DescribeTableInput
				OptFns []func(*dynamodb.Options)
			},
				dynamodb.DescribeTableOutput,
			](fmt.Errorf("DescribeTable general error")),
		},
	}
}

func (d *DynamodbMock) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(options *dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	d.Mock.GetItem.AddInput(
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
func (d *DynamodbMock) BatchWriteItem(ctx context.Context, params *dynamodb.BatchWriteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error) {
	d.Mock.BatchWriteItem.AddInput(
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
func (d *DynamodbMock) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	d.Mock.PutItem.AddInput(
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
func (d *DynamodbMock) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	d.Mock.UpdateItem.AddInput(
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
func (d *DynamodbMock) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	d.Mock.Query.AddInput(
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
func (d *DynamodbMock) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	d.Mock.Scan.AddInput(
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

func (d *DynamodbMock) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	d.Mock.DeleteItem.AddInput(
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

func (d *DynamodbMock) DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	d.Mock.DescribeTable.AddInput(
		struct {
			Ctx    context.Context
			Params *dynamodb.DescribeTableInput
			OptFns []func(*dynamodb.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return d.Mock.DescribeTable.GetNextResult()
}
