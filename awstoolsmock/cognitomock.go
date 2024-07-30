package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type cognitoMockStruct struct {
	CreateGroup *awsToolMock[struct {
		ctx    context.Context
		params *cognitoidentityprovider.CreateGroupInput
		optFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.CreateGroupOutput,
	]
	ListUsersInGroup *awsToolMock[struct {
		ctx    context.Context
		params *cognitoidentityprovider.ListUsersInGroupInput
		optFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.ListUsersInGroupOutput,
	]
	AdminAddUserToGroup *awsToolMock[struct {
		ctx    context.Context
		params *cognitoidentityprovider.AdminAddUserToGroupInput
		optFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.AdminAddUserToGroupOutput,
	]
	ListUsers *awsToolMock[struct {
		ctx    context.Context
		params *cognitoidentityprovider.ListUsersInput
		optFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.ListUsersOutput,
	]
	DescribeUserPool *awsToolMock[struct {
		ctx    context.Context
		params *cognitoidentityprovider.DescribeUserPoolInput
		optFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.DescribeUserPoolOutput,
	]
}

type cognitoMock struct {
	Mock cognitoMockStruct
}

func GetCognitoMock() *cognitoMock {
	return &cognitoMock{
		Mock: cognitoMockStruct{

			CreateGroup: GetMock[struct {
				ctx    context.Context
				params *cognitoidentityprovider.CreateGroupInput
				optFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.CreateGroupOutput,
			](fmt.Errorf("CreateGroup general error")),

			ListUsersInGroup: GetMock[struct {
				ctx    context.Context
				params *cognitoidentityprovider.ListUsersInGroupInput
				optFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.ListUsersInGroupOutput,
			](fmt.Errorf("ListUsersInGroup general error")),
		},
	}
}

func (c *cognitoMock) CreateGroup(ctx context.Context, params *cognitoidentityprovider.CreateGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateGroupOutput, error) {
	c.Mock.CreateGroup.AddInput(
		struct {
			ctx    context.Context
			params *cognitoidentityprovider.CreateGroupInput
			optFns []func(*cognitoidentityprovider.Options)
		}{
			ctx:    ctx,
			params: params,
			optFns: optFns,
		},
	)

	return c.Mock.CreateGroup.GetNextResult()
}

func (c *cognitoMock) ListUsersInGroup(ctx context.Context, params *cognitoidentityprovider.ListUsersInGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
	c.Mock.ListUsersInGroup.AddInput(
		struct {
			ctx    context.Context
			params *cognitoidentityprovider.ListUsersInGroupInput
			optFns []func(*cognitoidentityprovider.Options)
		}{
			ctx:    ctx,
			params: params,
			optFns: optFns,
		},
	)

	return c.Mock.ListUsersInGroup.GetNextResult()
}

func (c *cognitoMock) AdminAddUserToGroup(ctx context.Context, params *cognitoidentityprovider.AdminAddUserToGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error) {
	c.Mock.AdminAddUserToGroup.AddInput(
		struct {
			ctx    context.Context
			params *cognitoidentityprovider.AdminAddUserToGroupInput
			optFns []func(*cognitoidentityprovider.Options)
		}{
			ctx:    ctx,
			params: params,
			optFns: optFns,
		},
	)

	return c.Mock.AdminAddUserToGroup.GetNextResult()
}

func (c *cognitoMock) ListUsers(ctx context.Context, params *cognitoidentityprovider.ListUsersInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error) {
	c.Mock.ListUsers.AddInput(
		struct {
			ctx    context.Context
			params *cognitoidentityprovider.ListUsersInput
			optFns []func(*cognitoidentityprovider.Options)
		}{
			ctx:    ctx,
			params: params,
			optFns: optFns,
		},
	)

	return c.Mock.ListUsers.GetNextResult()
}

func (c *cognitoMock) DescribeUserPool(ctx context.Context, params *cognitoidentityprovider.DescribeUserPoolInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
	c.Mock.DescribeUserPool.AddInput(
		struct {
			ctx    context.Context
			params *cognitoidentityprovider.DescribeUserPoolInput
			optFns []func(*cognitoidentityprovider.Options)
		}{
			ctx:    ctx,
			params: params,
			optFns: optFns,
		},
	)

	return c.Mock.DescribeUserPool.GetNextResult()
}
