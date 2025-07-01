package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	gomock "github.com/nextunit-io/go-mock"
)

type cognitoMockStruct struct {
	CreateGroup *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.CreateGroupInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.CreateGroupOutput,
	]
	ListUsersInGroup *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.ListUsersInGroupInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.ListUsersInGroupOutput,
	]
	AdminAddUserToGroup *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.AdminAddUserToGroupInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.AdminAddUserToGroupOutput,
	]
	AdminListGroupsForUser *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.AdminListGroupsForUserInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.AdminListGroupsForUserOutput,
	]
	ListUsers *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.ListUsersInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.ListUsersOutput,
	]
	DescribeUserPool *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.DescribeUserPoolInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.DescribeUserPoolOutput,
	]
}

type CognitoMock struct {
	Mock cognitoMockStruct
}

func GetCognitoMock() *CognitoMock {
	return &CognitoMock{
		Mock: cognitoMockStruct{

			CreateGroup: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.CreateGroupInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.CreateGroupOutput,
			](fmt.Errorf("CreateGroup general error")),

			ListUsersInGroup: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.ListUsersInGroupInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.ListUsersInGroupOutput,
			](fmt.Errorf("ListUsersInGroup general error")),

			AdminAddUserToGroup: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.AdminAddUserToGroupInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.AdminAddUserToGroupOutput,
			](fmt.Errorf("AdminAddUserToGroup general error")),
			AdminListGroupsForUser: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.AdminListGroupsForUserInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.AdminListGroupsForUserOutput,
			](fmt.Errorf("AdminListGroupsForUser general error")),
			ListUsers: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.ListUsersInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.ListUsersOutput,
			](fmt.Errorf("ListUsers general error")),

			DescribeUserPool: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.DescribeUserPoolInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.DescribeUserPoolOutput,
			](fmt.Errorf("DescribeUserPool general error")),
		},
	}
}

func (c *CognitoMock) CreateGroup(ctx context.Context, params *cognitoidentityprovider.CreateGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateGroupOutput, error) {
	c.Mock.CreateGroup.AddInput(
		struct {
			Ctx    context.Context
			Params *cognitoidentityprovider.CreateGroupInput
			OptFns []func(*cognitoidentityprovider.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return c.Mock.CreateGroup.GetNextResult()
}

func (c *CognitoMock) ListUsersInGroup(ctx context.Context, params *cognitoidentityprovider.ListUsersInGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
	c.Mock.ListUsersInGroup.AddInput(
		struct {
			Ctx    context.Context
			Params *cognitoidentityprovider.ListUsersInGroupInput
			OptFns []func(*cognitoidentityprovider.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return c.Mock.ListUsersInGroup.GetNextResult()
}

func (c *CognitoMock) AdminAddUserToGroup(ctx context.Context, params *cognitoidentityprovider.AdminAddUserToGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error) {
	c.Mock.AdminAddUserToGroup.AddInput(
		struct {
			Ctx    context.Context
			Params *cognitoidentityprovider.AdminAddUserToGroupInput
			OptFns []func(*cognitoidentityprovider.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return c.Mock.AdminAddUserToGroup.GetNextResult()
}

func (c *CognitoMock) AdminListGroupsForUser(ctx context.Context, params *cognitoidentityprovider.AdminListGroupsForUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error) {
	c.Mock.AdminListGroupsForUser.AddInput(
		struct {
			Ctx    context.Context
			Params *cognitoidentityprovider.AdminListGroupsForUserInput
			OptFns []func(*cognitoidentityprovider.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return c.Mock.AdminListGroupsForUser.GetNextResult()
}

func (c *CognitoMock) ListUsers(ctx context.Context, params *cognitoidentityprovider.ListUsersInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error) {
	c.Mock.ListUsers.AddInput(
		struct {
			Ctx    context.Context
			Params *cognitoidentityprovider.ListUsersInput
			OptFns []func(*cognitoidentityprovider.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return c.Mock.ListUsers.GetNextResult()
}

func (c *CognitoMock) DescribeUserPool(ctx context.Context, params *cognitoidentityprovider.DescribeUserPoolInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
	c.Mock.DescribeUserPool.AddInput(
		struct {
			Ctx    context.Context
			Params *cognitoidentityprovider.DescribeUserPoolInput
			OptFns []func(*cognitoidentityprovider.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return c.Mock.DescribeUserPool.GetNextResult()
}
