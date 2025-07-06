package awstools

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoInterface interface {
	AdminAddUserToGroup(ctx context.Context, params *cognitoidentityprovider.AdminAddUserToGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
	AdminCreateUser(ctx context.Context, params *cognitoidentityprovider.AdminCreateUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	AdminDeleteUser(ctx context.Context, params *cognitoidentityprovider.AdminDeleteUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
	AdminListGroupsForUser(ctx context.Context, params *cognitoidentityprovider.AdminListGroupsForUserInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
	CreateGroup(ctx context.Context, params *cognitoidentityprovider.CreateGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.CreateGroupOutput, error)
	DescribeUserPool(ctx context.Context, params *cognitoidentityprovider.DescribeUserPoolInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	ListUsers(ctx context.Context, params *cognitoidentityprovider.ListUsersInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error)
	ListUsersInGroup(ctx context.Context, params *cognitoidentityprovider.ListUsersInGroupInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
}

var cognitoInstance CognitoInterface

func CreateCognitoInstance(cfg aws.Config) {
	SetCognitoInstance(cognitoidentityprovider.NewFromConfig(cfg))
}

func GetCognitoInstance() CognitoInterface {
	if cognitoInstance == nil {
		ctx := context.TODO()

		conf, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			panic("cannot create a new cognito instance")
		}

		cognitoInstance = cognitoidentityprovider.NewFromConfig(conf)
	}

	return cognitoInstance
}

func SetCognitoInstance(client CognitoInterface) {
	cognitoInstance = client
}
