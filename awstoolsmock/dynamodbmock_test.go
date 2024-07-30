package awstoolsmock_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestDynamodbMockGetItem(t *testing.T) {
	t.Helper()
	dynamodbMock := awstoolsmock.GetDynamodbMock()
	ctx := context.TODO()

	t.Run("Testing GetItem", func(t *testing.T) {
		output := dynamodb.GetItemOutput{
			Item: map[string]types.AttributeValue{
				"Test": &types.AttributeValueMemberBOOL{
					Value: true,
				},
			},
		}

		dynamodbMock.Mock.GetItem.AddReturnValue(&output)
		dynamodbMock.Mock.GetItem.AddReturnValue(&output)
		dynamodbMock.Mock.GetItem.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := dynamodbMock.GetItem(ctx, &dynamodb.GetItemInput{
				TableName: aws.String("test-tablename"),
				Key: map[string]types.AttributeValue{
					"ID": &types.AttributeValueMemberS{
						Value: fmt.Sprintf("test-id-%d", i),
					},
				},
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := dynamodbMock.GetItem(ctx, &dynamodb.GetItemInput{
			TableName: aws.String("test-tablename"),
			Key: map[string]types.AttributeValue{
				"ID": &types.AttributeValueMemberS{
					Value: "test-id-error",
				},
			},
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("GetItem general error"), err)

		assert.Equal(t, 4, dynamodbMock.Mock.GetItem.HaveBeenCalled())
		for i := 0; i < 3; i++ {
			input := dynamodbMock.Mock.GetItem.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-id-%d", i), input.Params.Key["ID"].(*types.AttributeValueMemberS).Value)
			assert.Equal(t, "test-tablename", *input.Params.TableName)
		}

		input := dynamodbMock.Mock.GetItem.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-id-error", input.Params.Key["ID"].(*types.AttributeValueMemberS).Value)
		assert.Equal(t, "test-tablename", *input.Params.TableName)
	})
}

func TestDynamodbMockBatchWriteItem(t *testing.T) {
	t.Helper()
	dynamodbMock := awstoolsmock.GetDynamodbMock()
	ctx := context.TODO()

	t.Run("Testing BatchWriteItem", func(t *testing.T) {
		output := dynamodb.BatchWriteItemOutput{}

		dynamodbMock.Mock.BatchWriteItem.AddReturnValue(&output)
		dynamodbMock.Mock.BatchWriteItem.AddReturnValue(&output)
		dynamodbMock.Mock.BatchWriteItem.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := dynamodbMock.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
				RequestItems: map[string][]types.WriteRequest{
					"test": {
						{
							PutRequest: &types.PutRequest{},
						},
					},
				},
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := dynamodbMock.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				"testerror": {
					{
						DeleteRequest: &types.DeleteRequest{},
					},
				},
			},
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("BatchWriteItem general error"), err)

		assert.Equal(t, 4, dynamodbMock.Mock.BatchWriteItem.HaveBeenCalled())
		for i := 0; i < 3; i++ {
			input := dynamodbMock.Mock.BatchWriteItem.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, &types.PutRequest{}, input.Params.RequestItems["test"][0].PutRequest)
		}

		input := dynamodbMock.Mock.BatchWriteItem.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, &types.DeleteRequest{}, input.Params.RequestItems["testerror"][0].DeleteRequest)
	})
}

func TestDynamodbMockPutItem(t *testing.T) {
	t.Helper()
	dynamodbMock := awstoolsmock.GetDynamodbMock()
	ctx := context.TODO()

	t.Run("Testing PutItem", func(t *testing.T) {
		output := dynamodb.PutItemOutput{}

		dynamodbMock.Mock.PutItem.AddReturnValue(&output)
		dynamodbMock.Mock.PutItem.AddReturnValue(&output)
		dynamodbMock.Mock.PutItem.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := dynamodbMock.PutItem(ctx, &dynamodb.PutItemInput{
				TableName: aws.String("test-tablename"),
				Item: map[string]types.AttributeValue{
					"test": &types.AttributeValueMemberBOOL{
						Value: true,
					},
				},
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := dynamodbMock.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String("test-tablename"),
			Item: map[string]types.AttributeValue{
				"testerror": &types.AttributeValueMemberBOOL{
					Value: false,
				},
			},
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("PutItem general error"), err)

		assert.Equal(t, 4, dynamodbMock.Mock.PutItem.HaveBeenCalled())
		for i := 0; i < 3; i++ {
			input := dynamodbMock.Mock.PutItem.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, true, input.Params.Item["test"].(*types.AttributeValueMemberBOOL).Value)
			assert.Equal(t, "test-tablename", *input.Params.TableName)
		}

		input := dynamodbMock.Mock.PutItem.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, false, input.Params.Item["testerror"].(*types.AttributeValueMemberBOOL).Value)
		assert.Equal(t, "test-tablename", *input.Params.TableName)
	})
}

func TestDynamodbMockUpdateItem(t *testing.T) {
	t.Helper()
	dynamodbMock := awstoolsmock.GetDynamodbMock()
	ctx := context.TODO()

	t.Run("Testing UpdateItem", func(t *testing.T) {
		output := dynamodb.UpdateItemOutput{}

		dynamodbMock.Mock.UpdateItem.AddReturnValue(&output)
		dynamodbMock.Mock.UpdateItem.AddReturnValue(&output)
		dynamodbMock.Mock.UpdateItem.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := dynamodbMock.UpdateItem(ctx, &dynamodb.UpdateItemInput{
				TableName: aws.String("test-tablename"),
				Key: map[string]types.AttributeValue{
					"ID": &types.AttributeValueMemberS{
						Value: fmt.Sprintf("test-%d", i),
					},
				},
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := dynamodbMock.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName: aws.String("test-tablename"),
			Key: map[string]types.AttributeValue{
				"ID": &types.AttributeValueMemberS{
					Value: "test-error",
				},
			},
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("UpdateItem general error"), err)

		assert.Equal(t, 4, dynamodbMock.Mock.UpdateItem.HaveBeenCalled())
		for i := 0; i < 3; i++ {
			input := dynamodbMock.Mock.UpdateItem.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), input.Params.Key["ID"].(*types.AttributeValueMemberS).Value)
			assert.Equal(t, "test-tablename", *input.Params.TableName)
		}

		input := dynamodbMock.Mock.UpdateItem.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-error", input.Params.Key["ID"].(*types.AttributeValueMemberS).Value)
		assert.Equal(t, "test-tablename", *input.Params.TableName)
	})
}

func TestDynamodbMockDeleteItem(t *testing.T) {
	t.Helper()
	dynamodbMock := awstoolsmock.GetDynamodbMock()
	ctx := context.TODO()

	t.Run("Testing DeleteItem", func(t *testing.T) {
		output := dynamodb.DeleteItemOutput{}

		dynamodbMock.Mock.DeleteItem.AddReturnValue(&output)
		dynamodbMock.Mock.DeleteItem.AddReturnValue(&output)
		dynamodbMock.Mock.DeleteItem.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := dynamodbMock.DeleteItem(ctx, &dynamodb.DeleteItemInput{
				TableName: aws.String("test-tablename"),
				Key: map[string]types.AttributeValue{
					"ID": &types.AttributeValueMemberS{
						Value: fmt.Sprintf("test-%d", i),
					},
				},
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := dynamodbMock.DeleteItem(ctx, &dynamodb.DeleteItemInput{
			TableName: aws.String("test-tablename"),
			Key: map[string]types.AttributeValue{
				"ID": &types.AttributeValueMemberS{
					Value: "test-error",
				},
			},
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("DeleteItem general error"), err)

		assert.Equal(t, 4, dynamodbMock.Mock.DeleteItem.HaveBeenCalled())
		for i := 0; i < 3; i++ {
			input := dynamodbMock.Mock.DeleteItem.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), input.Params.Key["ID"].(*types.AttributeValueMemberS).Value)
			assert.Equal(t, "test-tablename", *input.Params.TableName)
		}

		input := dynamodbMock.Mock.DeleteItem.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-error", input.Params.Key["ID"].(*types.AttributeValueMemberS).Value)
		assert.Equal(t, "test-tablename", *input.Params.TableName)
	})
}

func TestDynamodbMockQuery(t *testing.T) {
	t.Helper()
	dynamodbMock := awstoolsmock.GetDynamodbMock()
	ctx := context.TODO()

	t.Run("Testing Query", func(t *testing.T) {
		output := dynamodb.QueryOutput{
			Items: []map[string]types.AttributeValue{
				{
					"TEST": &types.AttributeValueMemberBOOL{
						Value: true,
					},
				},
			},
		}

		dynamodbMock.Mock.Query.AddReturnValue(&output)
		dynamodbMock.Mock.Query.AddReturnValue(&output)
		dynamodbMock.Mock.Query.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := dynamodbMock.Query(ctx, &dynamodb.QueryInput{
				TableName: aws.String("test-tablename"),
				IndexName: aws.String(fmt.Sprintf("test-index-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := dynamodbMock.Query(ctx, &dynamodb.QueryInput{
			TableName: aws.String("test-tablename"),
			IndexName: aws.String("test-error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("Query general error"), err)

		assert.Equal(t, 4, dynamodbMock.Mock.Query.HaveBeenCalled())
		for i := 0; i < 3; i++ {
			input := dynamodbMock.Mock.Query.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-index-%d", i), *input.Params.IndexName)
			assert.Equal(t, "test-tablename", *input.Params.TableName)
		}

		input := dynamodbMock.Mock.Query.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-error", *input.Params.IndexName)
		assert.Equal(t, "test-tablename", *input.Params.TableName)
	})
}

func TestDynamodbMockScan(t *testing.T) {
	t.Helper()
	dynamodbMock := awstoolsmock.GetDynamodbMock()
	ctx := context.TODO()

	t.Run("Testing Scan", func(t *testing.T) {
		output := dynamodb.ScanOutput{
			Items: []map[string]types.AttributeValue{
				{
					"TEST": &types.AttributeValueMemberBOOL{
						Value: true,
					},
				},
			},
		}

		dynamodbMock.Mock.Scan.AddReturnValue(&output)
		dynamodbMock.Mock.Scan.AddReturnValue(&output)
		dynamodbMock.Mock.Scan.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := dynamodbMock.Scan(ctx, &dynamodb.ScanInput{
				TableName: aws.String("test-tablename"),
				IndexName: aws.String(fmt.Sprintf("test-index-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := dynamodbMock.Scan(ctx, &dynamodb.ScanInput{
			TableName: aws.String("test-tablename"),
			IndexName: aws.String("test-error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("Scan general error"), err)

		assert.Equal(t, 4, dynamodbMock.Mock.Scan.HaveBeenCalled())
		for i := 0; i < 3; i++ {
			input := dynamodbMock.Mock.Scan.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-index-%d", i), *input.Params.IndexName)
			assert.Equal(t, "test-tablename", *input.Params.TableName)
		}

		input := dynamodbMock.Mock.Scan.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-error", *input.Params.IndexName)
		assert.Equal(t, "test-tablename", *input.Params.TableName)
	})
}

func TestDynamodbMockDescribeTable(t *testing.T) {
	t.Helper()
	dynamodbMock := awstoolsmock.GetDynamodbMock()
	ctx := context.TODO()

	t.Run("Testing DescribeTable", func(t *testing.T) {
		output := dynamodb.DescribeTableOutput{
			Table: &types.TableDescription{
				ItemCount: aws.Int64(50),
			},
		}

		dynamodbMock.Mock.DescribeTable.AddReturnValue(&output)
		dynamodbMock.Mock.DescribeTable.AddReturnValue(&output)
		dynamodbMock.Mock.DescribeTable.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := dynamodbMock.DescribeTable(ctx, &dynamodb.DescribeTableInput{
				TableName: aws.String("test-tablename"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := dynamodbMock.DescribeTable(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String("test-tablename"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("DescribeTable general error"), err)

		assert.Equal(t, 4, dynamodbMock.Mock.DescribeTable.HaveBeenCalled())
		for i := 0; i < 3; i++ {
			input := dynamodbMock.Mock.DescribeTable.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, "test-tablename", *input.Params.TableName)
		}

		input := dynamodbMock.Mock.DescribeTable.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-tablename", *input.Params.TableName)
	})
}
