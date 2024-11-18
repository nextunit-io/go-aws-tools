package awstoolsmock_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestPresignerMockPresignGetObject(t *testing.T) {
	t.Helper()
	presignerMock := awstoolsmock.GetPresignerMock()
	ctx := context.TODO()

	t.Run("Testing PresignGetObject", func(t *testing.T) {
		output := v4.PresignedHTTPRequest{
			URL: "test-url-123",
		}

		presignerMock.Mock.PresignGetObject.AddReturnValue(&output)
		presignerMock.Mock.PresignGetObject.AddReturnValue(&output)
		presignerMock.Mock.PresignGetObject.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := presignerMock.PresignGetObject(ctx, &s3.GetObjectInput{
				Bucket: aws.String("test-bucket"),
				Key:    aws.String(fmt.Sprintf("file-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := presignerMock.PresignGetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String("test-bucket"),
			Key:    aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("PresignGetObject general error"), err)

		assert.Equal(t, 4, presignerMock.Mock.PresignGetObject.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := presignerMock.Mock.PresignGetObject.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, "test-bucket", *input.Params.Bucket)
			assert.Equal(t, fmt.Sprintf("file-%d", i), *input.Params.Key)
		}

		input := presignerMock.Mock.PresignGetObject.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-bucket", *input.Params.Bucket)
		assert.Equal(t, "error", *input.Params.Key)
	})
}

func TestPresignerMockPresignDeleteObject(t *testing.T) {
	t.Helper()
	presignerMock := awstoolsmock.GetPresignerMock()
	ctx := context.TODO()

	t.Run("Testing PresignDeleteObject", func(t *testing.T) {
		output := v4.PresignedHTTPRequest{
			URL: "test-url-123",
		}

		presignerMock.Mock.PresignDeleteObject.AddReturnValue(&output)
		presignerMock.Mock.PresignDeleteObject.AddReturnValue(&output)
		presignerMock.Mock.PresignDeleteObject.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := presignerMock.PresignDeleteObject(ctx, &s3.DeleteObjectInput{
				Bucket: aws.String("test-bucket"),
				Key:    aws.String(fmt.Sprintf("file-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := presignerMock.PresignDeleteObject(ctx, &s3.DeleteObjectInput{
			Bucket: aws.String("test-bucket"),
			Key:    aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("PresignDeleteObject general error"), err)

		assert.Equal(t, 4, presignerMock.Mock.PresignDeleteObject.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := presignerMock.Mock.PresignDeleteObject.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, "test-bucket", *input.Params.Bucket)
			assert.Equal(t, fmt.Sprintf("file-%d", i), *input.Params.Key)
		}

		input := presignerMock.Mock.PresignDeleteObject.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-bucket", *input.Params.Bucket)
		assert.Equal(t, "error", *input.Params.Key)
	})
}

func TestPresignerMockPresignPutObject(t *testing.T) {
	t.Helper()
	presignerMock := awstoolsmock.GetPresignerMock()
	ctx := context.TODO()

	t.Run("Testing PresignPutObject", func(t *testing.T) {
		output := v4.PresignedHTTPRequest{
			URL: "test-url-123",
		}

		presignerMock.Mock.PresignPutObject.AddReturnValue(&output)
		presignerMock.Mock.PresignPutObject.AddReturnValue(&output)
		presignerMock.Mock.PresignPutObject.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := presignerMock.PresignPutObject(ctx, &s3.PutObjectInput{
				Bucket: aws.String("test-bucket"),
				Key:    aws.String(fmt.Sprintf("file-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := presignerMock.PresignPutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String("test-bucket"),
			Key:    aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("PresignPutObject general error"), err)

		assert.Equal(t, 4, presignerMock.Mock.PresignPutObject.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := presignerMock.Mock.PresignPutObject.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, "test-bucket", *input.Params.Bucket)
			assert.Equal(t, fmt.Sprintf("file-%d", i), *input.Params.Key)
		}

		input := presignerMock.Mock.PresignPutObject.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-bucket", *input.Params.Bucket)
		assert.Equal(t, "error", *input.Params.Key)
	})
}
