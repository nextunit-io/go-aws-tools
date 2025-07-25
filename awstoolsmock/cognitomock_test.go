package awstoolsmock_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/nextunit-io/go-aws-tools/awstoolsmock"
	"github.com/stretchr/testify/assert"
)

func TestCognitoMockCreateGroup(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing CreateGroup", func(t *testing.T) {
		output := cognitoidentityprovider.CreateGroupOutput{
			Group: &types.GroupType{
				GroupName: aws.String("test-group-output"),
			},
		}

		cognitoMock.Mock.CreateGroup.AddReturnValue(&output)
		cognitoMock.Mock.CreateGroup.AddReturnValue(&output)
		cognitoMock.Mock.CreateGroup.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.CreateGroup(ctx, &cognitoidentityprovider.CreateGroupInput{
				GroupName:  aws.String(fmt.Sprintf("test-groupname-%d", i)),
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.CreateGroup(ctx, &cognitoidentityprovider.CreateGroupInput{
			GroupName:  aws.String("test-groupname-error"),
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("CreateGroup general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.CreateGroup.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.CreateGroup.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-groupname-%d", i), *input.Params.GroupName)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.CreateGroup.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-groupname-error", *input.Params.GroupName)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}

func TestCognitoMockListUsersInGroup(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing ListUsersInGroup", func(t *testing.T) {
		output := cognitoidentityprovider.ListUsersInGroupOutput{
			Users: []types.UserType{
				{
					Username: aws.String("test-username-1"),
				},
				{
					Username: aws.String("test-username-2"),
				},
			},
		}

		cognitoMock.Mock.ListUsersInGroup.AddReturnValue(&output)
		cognitoMock.Mock.ListUsersInGroup.AddReturnValue(&output)
		cognitoMock.Mock.ListUsersInGroup.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.ListUsersInGroup(ctx, &cognitoidentityprovider.ListUsersInGroupInput{
				GroupName:  aws.String(fmt.Sprintf("test-groupname-%d", i)),
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.ListUsersInGroup(ctx, &cognitoidentityprovider.ListUsersInGroupInput{
			GroupName:  aws.String("test-groupname-error"),
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("ListUsersInGroup general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.ListUsersInGroup.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.ListUsersInGroup.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-groupname-%d", i), *input.Params.GroupName)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.ListUsersInGroup.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-groupname-error", *input.Params.GroupName)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}

func TestCognitoMockAdminAddUserToGroup(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing AdminAddUserToGroup", func(t *testing.T) {
		output := cognitoidentityprovider.AdminAddUserToGroupOutput{}

		cognitoMock.Mock.AdminAddUserToGroup.AddReturnValue(&output)
		cognitoMock.Mock.AdminAddUserToGroup.AddReturnValue(&output)
		cognitoMock.Mock.AdminAddUserToGroup.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.AdminAddUserToGroup(ctx, &cognitoidentityprovider.AdminAddUserToGroupInput{
				GroupName:  aws.String(fmt.Sprintf("test-groupname-%d", i)),
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.AdminAddUserToGroup(ctx, &cognitoidentityprovider.AdminAddUserToGroupInput{
			GroupName:  aws.String("test-groupname-error"),
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("AdminAddUserToGroup general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.AdminAddUserToGroup.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.AdminAddUserToGroup.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-groupname-%d", i), *input.Params.GroupName)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.AdminAddUserToGroup.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-groupname-error", *input.Params.GroupName)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}

func TestCognitoMockAdminCreateUser(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing AdminCreateUser", func(t *testing.T) {
		output := cognitoidentityprovider.AdminCreateUserOutput{}

		cognitoMock.Mock.AdminCreateUser.AddReturnValue(&output)
		cognitoMock.Mock.AdminCreateUser.AddReturnValue(&output)
		cognitoMock.Mock.AdminCreateUser.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.AdminCreateUser(ctx, &cognitoidentityprovider.AdminCreateUserInput{
				Username:   aws.String(fmt.Sprintf("test-username-%d", i)),
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.AdminCreateUser(ctx, &cognitoidentityprovider.AdminCreateUserInput{
			Username:   aws.String("test-username-error"),
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("AdminCreateUser general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.AdminCreateUser.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.AdminCreateUser.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-username-%d", i), *input.Params.Username)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.AdminCreateUser.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-username-error", *input.Params.Username)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}

func TestCognitoMockAdminDeleteUser(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing AdminDeleteUser", func(t *testing.T) {
		output := cognitoidentityprovider.AdminDeleteUserOutput{}

		cognitoMock.Mock.AdminDeleteUser.AddReturnValue(&output)
		cognitoMock.Mock.AdminDeleteUser.AddReturnValue(&output)
		cognitoMock.Mock.AdminDeleteUser.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.AdminDeleteUser(ctx, &cognitoidentityprovider.AdminDeleteUserInput{
				Username:   aws.String(fmt.Sprintf("test-username-%d", i)),
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.AdminDeleteUser(ctx, &cognitoidentityprovider.AdminDeleteUserInput{
			Username:   aws.String("test-username-error"),
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("AdminDeleteUser general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.AdminDeleteUser.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.AdminDeleteUser.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-username-%d", i), *input.Params.Username)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.AdminDeleteUser.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-username-error", *input.Params.Username)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}

func TestCognitoMockAdminGetUser(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing AdminGetUser", func(t *testing.T) {
		output := cognitoidentityprovider.AdminGetUserOutput{}

		cognitoMock.Mock.AdminGetUser.AddReturnValue(&output)
		cognitoMock.Mock.AdminGetUser.AddReturnValue(&output)
		cognitoMock.Mock.AdminGetUser.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.AdminGetUser(ctx, &cognitoidentityprovider.AdminGetUserInput{
				Username:   aws.String(fmt.Sprintf("test-username-%d", i)),
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.AdminGetUser(ctx, &cognitoidentityprovider.AdminGetUserInput{
			Username:   aws.String("test-username-error"),
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("AdminGetUser general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.AdminGetUser.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.AdminGetUser.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-username-%d", i), *input.Params.Username)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.AdminGetUser.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-username-error", *input.Params.Username)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}

func TestCognitoMockAdminListGroupsForUser(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing AdminListGroupsForUser", func(t *testing.T) {
		output := cognitoidentityprovider.AdminListGroupsForUserOutput{}

		cognitoMock.Mock.AdminListGroupsForUser.AddReturnValue(&output)
		cognitoMock.Mock.AdminListGroupsForUser.AddReturnValue(&output)
		cognitoMock.Mock.AdminListGroupsForUser.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.AdminListGroupsForUser(ctx, &cognitoidentityprovider.AdminListGroupsForUserInput{
				Username:   aws.String(fmt.Sprintf("test-username-%d", i)),
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.AdminListGroupsForUser(ctx, &cognitoidentityprovider.AdminListGroupsForUserInput{
			Username:   aws.String("test-username-error"),
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("AdminListGroupsForUser general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.AdminListGroupsForUser.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.AdminListGroupsForUser.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-username-%d", i), *input.Params.Username)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.AdminListGroupsForUser.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-username-error", *input.Params.Username)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}

func TestCognitoMockAdminRemoveUserFromGroup(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing AdminRemoveUserFromGroup", func(t *testing.T) {
		output := cognitoidentityprovider.AdminRemoveUserFromGroupOutput{}

		cognitoMock.Mock.AdminRemoveUserFromGroup.AddReturnValue(&output)
		cognitoMock.Mock.AdminRemoveUserFromGroup.AddReturnValue(&output)
		cognitoMock.Mock.AdminRemoveUserFromGroup.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.AdminRemoveUserFromGroup(ctx, &cognitoidentityprovider.AdminRemoveUserFromGroupInput{
				Username:   aws.String(fmt.Sprintf("test-username-%d", i)),
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.AdminRemoveUserFromGroup(ctx, &cognitoidentityprovider.AdminRemoveUserFromGroupInput{
			Username:   aws.String("test-username-error"),
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("AdminRemoveUserFromGroup general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.AdminRemoveUserFromGroup.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.AdminRemoveUserFromGroup.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, fmt.Sprintf("test-username-%d", i), *input.Params.Username)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.AdminRemoveUserFromGroup.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-username-error", *input.Params.Username)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}

func TestCognitoMockListUsers(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing ListUsers", func(t *testing.T) {
		output := cognitoidentityprovider.ListUsersOutput{
			Users: []types.UserType{
				{
					Username: aws.String("test-username-1"),
				},
				{
					Username: aws.String("test-username-2"),
				},
			},
		}

		cognitoMock.Mock.ListUsers.AddReturnValue(&output)
		cognitoMock.Mock.ListUsers.AddReturnValue(&output)
		cognitoMock.Mock.ListUsers.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.ListUsers(ctx, &cognitoidentityprovider.ListUsersInput{
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.ListUsers(ctx, &cognitoidentityprovider.ListUsersInput{
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("ListUsers general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.ListUsers.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.ListUsers.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.ListUsers.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}

func TestCognitoMockDescribeUserPool(t *testing.T) {
	t.Helper()
	cognitoMock := awstoolsmock.GetCognitoMock()
	ctx := context.TODO()

	t.Run("Testing DescribeUserPool", func(t *testing.T) {
		output := cognitoidentityprovider.DescribeUserPoolOutput{}

		cognitoMock.Mock.DescribeUserPool.AddReturnValue(&output)
		cognitoMock.Mock.DescribeUserPool.AddReturnValue(&output)
		cognitoMock.Mock.DescribeUserPool.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := cognitoMock.DescribeUserPool(ctx, &cognitoidentityprovider.DescribeUserPoolInput{
				UserPoolId: aws.String("test-userpool-id"),
			})

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := cognitoMock.DescribeUserPool(ctx, &cognitoidentityprovider.DescribeUserPoolInput{
			UserPoolId: aws.String("test-userpool-id"),
		})

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("DescribeUserPool general error"), err)

		assert.Equal(t, 4, cognitoMock.Mock.DescribeUserPool.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := cognitoMock.Mock.DescribeUserPool.GetInput(i)
			assert.Equal(t, ctx, input.Ctx)
			assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
		}

		input := cognitoMock.Mock.DescribeUserPool.GetInput(3)
		assert.Equal(t, ctx, input.Ctx)
		assert.Equal(t, "test-userpool-id", *input.Params.UserPoolId)
	})
}
