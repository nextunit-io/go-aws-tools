package awstoolsmock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	gomock "github.com/nextunit-io/go-mock"
)

type cognitoMockStruct struct {
	AdminAddUserToGroup *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.AdminAddUserToGroupInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.AdminAddUserToGroupOutput,
	]
	AdminCreateUser *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.AdminCreateUserInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.AdminCreateUserOutput,
	]
	AdminDeleteUser *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.AdminDeleteUserInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.AdminDeleteUserOutput,
	]
	AdminGetUser *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.AdminGetUserInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.AdminGetUserOutput,
	]
	AdminListGroupsForUser *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.AdminListGroupsForUserInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.AdminListGroupsForUserOutput,
	]
	CreateGroup *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.CreateGroupInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.CreateGroupOutput,
	]
	DescribeUserPool *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.DescribeUserPoolInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.DescribeUserPoolOutput,
	]
	ListUsers *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.ListUsersInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.ListUsersOutput,
	]
	ListUsersInGroup *gomock.ToolMock[struct {
		Ctx    context.Context
		Params *cognitoidentityprovider.ListUsersInGroupInput
		OptFns []func(*cognitoidentityprovider.Options)
	},
		cognitoidentityprovider.ListUsersInGroupOutput,
	]
}

type CognitoMock struct {
	Mock cognitoMockStruct
}

func GetCognitoMock() *CognitoMock {
	return &CognitoMock{
		Mock: cognitoMockStruct{

			AdminAddUserToGroup: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.AdminAddUserToGroupInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.AdminAddUserToGroupOutput,
			](fmt.Errorf("AdminAddUserToGroup general error")),

			AdminCreateUser: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.AdminCreateUserInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.AdminCreateUserOutput,
			](fmt.Errorf("AdminCreateUser general error")),

			AdminDeleteUser: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.AdminDeleteUserInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.AdminDeleteUserOutput,
			](fmt.Errorf("AdminDeleteUser general error")),

			AdminGetUser: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.AdminGetUserInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.AdminGetUserOutput,
			](fmt.Errorf("AdminGetUser general error")),

			AdminListGroupsForUser: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.AdminListGroupsForUserInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.AdminListGroupsForUserOutput,
			](fmt.Errorf("AdminListGroupsForUser general error")),

			CreateGroup: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.CreateGroupInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.CreateGroupOutput,
			](fmt.Errorf("CreateGroup general error")),

			DescribeUserPool: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.DescribeUserPoolInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.DescribeUserPoolOutput,
			](fmt.Errorf("DescribeUserPool general error")),

			ListUsers: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.ListUsersInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.ListUsersOutput,
			](fmt.Errorf("ListUsers general error")),

			ListUsersInGroup: gomock.GetMock[struct {
				Ctx    context.Context
				Params *cognitoidentityprovider.ListUsersInGroupInput
				OptFns []func(*cognitoidentityprovider.Options)
			},
				cognitoidentityprovider.ListUsersInGroupOutput,
			](fmt.Errorf("ListUsersInGroup general error")),
		},
	}
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

func (c *CognitoMock) AdminCreateUser(ctx context.Context, params *cognitoidentityprovider.AdminCreateUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
	c.Mock.AdminCreateUser.AddInput(
		struct {
			Ctx    context.Context
			Params *cognitoidentityprovider.AdminCreateUserInput
			OptFns []func(*cognitoidentityprovider.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return c.Mock.AdminCreateUser.GetNextResult()
}

func (c *CognitoMock) AdminDeleteUser(ctx context.Context, params *cognitoidentityprovider.AdminDeleteUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {
	c.Mock.AdminDeleteUser.AddInput(
		struct {
			Ctx    context.Context
			Params *cognitoidentityprovider.AdminDeleteUserInput
			OptFns []func(*cognitoidentityprovider.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return c.Mock.AdminDeleteUser.GetNextResult()
}

func (c *CognitoMock) AdminGetUser(ctx context.Context, params *cognitoidentityprovider.AdminGetUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	c.Mock.AdminGetUser.AddInput(
		struct {
			Ctx    context.Context
			Params *cognitoidentityprovider.AdminGetUserInput
			OptFns []func(*cognitoidentityprovider.Options)
		}{
			Ctx:    ctx,
			Params: params,
			OptFns: optFns,
		},
	)

	return c.Mock.AdminGetUser.GetNextResult()
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
