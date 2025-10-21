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
func TestS3MockListObjectsV2(t *testing.T) {
	t.Helper()
	s3Mock := awstoolsmock.GetS3Mock()
	ctx := context.TODO()

	t.Run("Testing ListObjectsV2", func(t *testing.T) {
		output := s3.ListObjectsV2Output{}

		s3Mock.Mock.ListObjectsV2.AddReturnValue(&output)
		s3Mock.Mock.ListObjectsV2.AddReturnValue(&output)
		s3Mock.Mock.ListObjectsV2.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := s3Mock.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
				Bucket: aws.String(fmt.Sprintf("test-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := s3Mock.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("ListObjectsV2 general error"), err)

		assert.Equal(t, 4, s3Mock.Mock.ListObjectsV2.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := s3Mock.Mock.ListObjectsV2.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Bucket)
		}

		input := s3Mock.Mock.ListObjectsV2.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Bucket)
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

func TestS3MockDeleteObject(t *testing.T) {
	t.Helper()
	s3Mock := awstoolsmock.GetS3Mock()
	ctx := context.TODO()

	t.Run("Testing DeleteObject", func(t *testing.T) {
		output := s3.DeleteObjectOutput{}

		s3Mock.Mock.DeleteObject.AddReturnValue(&output)
		s3Mock.Mock.DeleteObject.AddReturnValue(&output)
		s3Mock.Mock.DeleteObject.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := s3Mock.DeleteObject(ctx, &s3.DeleteObjectInput{
				Key: aws.String(fmt.Sprintf("test-%d", i)),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := s3Mock.DeleteObject(ctx, &s3.DeleteObjectInput{
			Key: aws.String("error"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("DeleteObject general error"), err)

		assert.Equal(t, 4, s3Mock.Mock.DeleteObject.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := s3Mock.Mock.DeleteObject.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-%d", i), *input.Params.Key)
		}

		input := s3Mock.Mock.DeleteObject.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "error", *input.Params.Key)
	})
}
