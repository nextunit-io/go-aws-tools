package awstoolsmock_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestS3MockGetObject(t *testing.T) {
	t.Helper()
	s3Mock := awstoolsmock.GetS3Mock()
	ctx := context.TODO()

	t.Run("Testing GetObject", func(t *testing.T) {
		output := s3.GetObjectOutput{}

		s3Mock.Mock.GetObject.AddReturnValue(&output)
		s3Mock.Mock.GetObject.AddReturnValue(&output)
		s3Mock.Mock.GetObject.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := s3Mock.GetObject(ctx, &s3.GetObjectInput{
				Key: aws.String(fmt.Sprintf("test-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := s3Mock.GetObject(ctx, &s3.GetObjectInput{
			Key: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("GetObject general error"), err)

		assert.Equal(t, 4, s3Mock.Mock.GetObject.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := s3Mock.Mock.GetObject.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Key)
		}

		input := s3Mock.Mock.GetObject.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Key)
	})
}

func TestS3MockPutObject(t *testing.T) {
	t.Helper()
	s3Mock := awstoolsmock.GetS3Mock()
	ctx := context.TODO()

	t.Run("Testing PutObject", func(t *testing.T) {
		output := s3.PutObjectOutput{}

		s3Mock.Mock.PutObject.AddReturnValue(&output)
		s3Mock.Mock.PutObject.AddReturnValue(&output)
		s3Mock.Mock.PutObject.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := s3Mock.PutObject(ctx, &s3.PutObjectInput{
				Key: aws.String(fmt.Sprintf("test-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := s3Mock.PutObject(ctx, &s3.PutObjectInput{
			Key: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("PutObject general error"), err)

		assert.Equal(t, 4, s3Mock.Mock.PutObject.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := s3Mock.Mock.PutObject.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Key)
		}

		input := s3Mock.Mock.PutObject.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Key)
	})
}
