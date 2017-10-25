// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cognitoidentityprovideriface provides an interface to enable mocking the Amazon Cognito Identity Provider service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cognitoidentityprovideriface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// CognitoIdentityProviderAPI provides an interface to enable mocking the
// cognitoidentityprovider.CognitoIdentityProvider service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Cognito Identity Provider.
//    func myFunc(svc cognitoidentityprovideriface.CognitoIdentityProviderAPI) bool {
//        // Make svc.AddCustomAttributes request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cognitoidentityprovider.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCognitoIdentityProviderClient struct {
//        cognitoidentityprovideriface.CognitoIdentityProviderAPI
//    }
//    func (m *mockCognitoIdentityProviderClient) AddCustomAttributes(input *cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCognitoIdentityProviderClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CognitoIdentityProviderAPI interface {
	AddCustomAttributes(*cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error)
	AddCustomAttributesWithContext(aws.Context, *cognitoidentityprovider.AddCustomAttributesInput, ...aws.Option) (*cognitoidentityprovider.AddCustomAttributesOutput, error)
	AddCustomAttributesRequest(*cognitoidentityprovider.AddCustomAttributesInput) (*aws.Request, *cognitoidentityprovider.AddCustomAttributesOutput)

	AdminAddUserToGroup(*cognitoidentityprovider.AdminAddUserToGroupInput) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
	AdminAddUserToGroupWithContext(aws.Context, *cognitoidentityprovider.AdminAddUserToGroupInput, ...aws.Option) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
	AdminAddUserToGroupRequest(*cognitoidentityprovider.AdminAddUserToGroupInput) (*aws.Request, *cognitoidentityprovider.AdminAddUserToGroupOutput)

	AdminConfirmSignUp(*cognitoidentityprovider.AdminConfirmSignUpInput) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error)
	AdminConfirmSignUpWithContext(aws.Context, *cognitoidentityprovider.AdminConfirmSignUpInput, ...aws.Option) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error)
	AdminConfirmSignUpRequest(*cognitoidentityprovider.AdminConfirmSignUpInput) (*aws.Request, *cognitoidentityprovider.AdminConfirmSignUpOutput)

	AdminCreateUser(*cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	AdminCreateUserWithContext(aws.Context, *cognitoidentityprovider.AdminCreateUserInput, ...aws.Option) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	AdminCreateUserRequest(*cognitoidentityprovider.AdminCreateUserInput) (*aws.Request, *cognitoidentityprovider.AdminCreateUserOutput)

	AdminDeleteUser(*cognitoidentityprovider.AdminDeleteUserInput) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
	AdminDeleteUserWithContext(aws.Context, *cognitoidentityprovider.AdminDeleteUserInput, ...aws.Option) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
	AdminDeleteUserRequest(*cognitoidentityprovider.AdminDeleteUserInput) (*aws.Request, *cognitoidentityprovider.AdminDeleteUserOutput)

	AdminDeleteUserAttributes(*cognitoidentityprovider.AdminDeleteUserAttributesInput) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error)
	AdminDeleteUserAttributesWithContext(aws.Context, *cognitoidentityprovider.AdminDeleteUserAttributesInput, ...aws.Option) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error)
	AdminDeleteUserAttributesRequest(*cognitoidentityprovider.AdminDeleteUserAttributesInput) (*aws.Request, *cognitoidentityprovider.AdminDeleteUserAttributesOutput)

	AdminDisableProviderForUser(*cognitoidentityprovider.AdminDisableProviderForUserInput) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error)
	AdminDisableProviderForUserWithContext(aws.Context, *cognitoidentityprovider.AdminDisableProviderForUserInput, ...aws.Option) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error)
	AdminDisableProviderForUserRequest(*cognitoidentityprovider.AdminDisableProviderForUserInput) (*aws.Request, *cognitoidentityprovider.AdminDisableProviderForUserOutput)

	AdminDisableUser(*cognitoidentityprovider.AdminDisableUserInput) (*cognitoidentityprovider.AdminDisableUserOutput, error)
	AdminDisableUserWithContext(aws.Context, *cognitoidentityprovider.AdminDisableUserInput, ...aws.Option) (*cognitoidentityprovider.AdminDisableUserOutput, error)
	AdminDisableUserRequest(*cognitoidentityprovider.AdminDisableUserInput) (*aws.Request, *cognitoidentityprovider.AdminDisableUserOutput)

	AdminEnableUser(*cognitoidentityprovider.AdminEnableUserInput) (*cognitoidentityprovider.AdminEnableUserOutput, error)
	AdminEnableUserWithContext(aws.Context, *cognitoidentityprovider.AdminEnableUserInput, ...aws.Option) (*cognitoidentityprovider.AdminEnableUserOutput, error)
	AdminEnableUserRequest(*cognitoidentityprovider.AdminEnableUserInput) (*aws.Request, *cognitoidentityprovider.AdminEnableUserOutput)

	AdminForgetDevice(*cognitoidentityprovider.AdminForgetDeviceInput) (*cognitoidentityprovider.AdminForgetDeviceOutput, error)
	AdminForgetDeviceWithContext(aws.Context, *cognitoidentityprovider.AdminForgetDeviceInput, ...aws.Option) (*cognitoidentityprovider.AdminForgetDeviceOutput, error)
	AdminForgetDeviceRequest(*cognitoidentityprovider.AdminForgetDeviceInput) (*aws.Request, *cognitoidentityprovider.AdminForgetDeviceOutput)

	AdminGetDevice(*cognitoidentityprovider.AdminGetDeviceInput) (*cognitoidentityprovider.AdminGetDeviceOutput, error)
	AdminGetDeviceWithContext(aws.Context, *cognitoidentityprovider.AdminGetDeviceInput, ...aws.Option) (*cognitoidentityprovider.AdminGetDeviceOutput, error)
	AdminGetDeviceRequest(*cognitoidentityprovider.AdminGetDeviceInput) (*aws.Request, *cognitoidentityprovider.AdminGetDeviceOutput)

	AdminGetUser(*cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error)
	AdminGetUserWithContext(aws.Context, *cognitoidentityprovider.AdminGetUserInput, ...aws.Option) (*cognitoidentityprovider.AdminGetUserOutput, error)
	AdminGetUserRequest(*cognitoidentityprovider.AdminGetUserInput) (*aws.Request, *cognitoidentityprovider.AdminGetUserOutput)

	AdminInitiateAuth(*cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	AdminInitiateAuthWithContext(aws.Context, *cognitoidentityprovider.AdminInitiateAuthInput, ...aws.Option) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	AdminInitiateAuthRequest(*cognitoidentityprovider.AdminInitiateAuthInput) (*aws.Request, *cognitoidentityprovider.AdminInitiateAuthOutput)

	AdminLinkProviderForUser(*cognitoidentityprovider.AdminLinkProviderForUserInput) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error)
	AdminLinkProviderForUserWithContext(aws.Context, *cognitoidentityprovider.AdminLinkProviderForUserInput, ...aws.Option) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error)
	AdminLinkProviderForUserRequest(*cognitoidentityprovider.AdminLinkProviderForUserInput) (*aws.Request, *cognitoidentityprovider.AdminLinkProviderForUserOutput)

	AdminListDevices(*cognitoidentityprovider.AdminListDevicesInput) (*cognitoidentityprovider.AdminListDevicesOutput, error)
	AdminListDevicesWithContext(aws.Context, *cognitoidentityprovider.AdminListDevicesInput, ...aws.Option) (*cognitoidentityprovider.AdminListDevicesOutput, error)
	AdminListDevicesRequest(*cognitoidentityprovider.AdminListDevicesInput) (*aws.Request, *cognitoidentityprovider.AdminListDevicesOutput)

	AdminListGroupsForUser(*cognitoidentityprovider.AdminListGroupsForUserInput) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
	AdminListGroupsForUserWithContext(aws.Context, *cognitoidentityprovider.AdminListGroupsForUserInput, ...aws.Option) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
	AdminListGroupsForUserRequest(*cognitoidentityprovider.AdminListGroupsForUserInput) (*aws.Request, *cognitoidentityprovider.AdminListGroupsForUserOutput)

	AdminRemoveUserFromGroup(*cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error)
	AdminRemoveUserFromGroupWithContext(aws.Context, *cognitoidentityprovider.AdminRemoveUserFromGroupInput, ...aws.Option) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error)
	AdminRemoveUserFromGroupRequest(*cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*aws.Request, *cognitoidentityprovider.AdminRemoveUserFromGroupOutput)

	AdminResetUserPassword(*cognitoidentityprovider.AdminResetUserPasswordInput) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error)
	AdminResetUserPasswordWithContext(aws.Context, *cognitoidentityprovider.AdminResetUserPasswordInput, ...aws.Option) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error)
	AdminResetUserPasswordRequest(*cognitoidentityprovider.AdminResetUserPasswordInput) (*aws.Request, *cognitoidentityprovider.AdminResetUserPasswordOutput)

	AdminRespondToAuthChallenge(*cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
	AdminRespondToAuthChallengeWithContext(aws.Context, *cognitoidentityprovider.AdminRespondToAuthChallengeInput, ...aws.Option) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
	AdminRespondToAuthChallengeRequest(*cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*aws.Request, *cognitoidentityprovider.AdminRespondToAuthChallengeOutput)

	AdminSetUserSettings(*cognitoidentityprovider.AdminSetUserSettingsInput) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error)
	AdminSetUserSettingsWithContext(aws.Context, *cognitoidentityprovider.AdminSetUserSettingsInput, ...aws.Option) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error)
	AdminSetUserSettingsRequest(*cognitoidentityprovider.AdminSetUserSettingsInput) (*aws.Request, *cognitoidentityprovider.AdminSetUserSettingsOutput)

	AdminUpdateDeviceStatus(*cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error)
	AdminUpdateDeviceStatusWithContext(aws.Context, *cognitoidentityprovider.AdminUpdateDeviceStatusInput, ...aws.Option) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error)
	AdminUpdateDeviceStatusRequest(*cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*aws.Request, *cognitoidentityprovider.AdminUpdateDeviceStatusOutput)

	AdminUpdateUserAttributes(*cognitoidentityprovider.AdminUpdateUserAttributesInput) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	AdminUpdateUserAttributesWithContext(aws.Context, *cognitoidentityprovider.AdminUpdateUserAttributesInput, ...aws.Option) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	AdminUpdateUserAttributesRequest(*cognitoidentityprovider.AdminUpdateUserAttributesInput) (*aws.Request, *cognitoidentityprovider.AdminUpdateUserAttributesOutput)

	AdminUserGlobalSignOut(*cognitoidentityprovider.AdminUserGlobalSignOutInput) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error)
	AdminUserGlobalSignOutWithContext(aws.Context, *cognitoidentityprovider.AdminUserGlobalSignOutInput, ...aws.Option) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error)
	AdminUserGlobalSignOutRequest(*cognitoidentityprovider.AdminUserGlobalSignOutInput) (*aws.Request, *cognitoidentityprovider.AdminUserGlobalSignOutOutput)

	ChangePassword(*cognitoidentityprovider.ChangePasswordInput) (*cognitoidentityprovider.ChangePasswordOutput, error)
	ChangePasswordWithContext(aws.Context, *cognitoidentityprovider.ChangePasswordInput, ...aws.Option) (*cognitoidentityprovider.ChangePasswordOutput, error)
	ChangePasswordRequest(*cognitoidentityprovider.ChangePasswordInput) (*aws.Request, *cognitoidentityprovider.ChangePasswordOutput)

	ConfirmDevice(*cognitoidentityprovider.ConfirmDeviceInput) (*cognitoidentityprovider.ConfirmDeviceOutput, error)
	ConfirmDeviceWithContext(aws.Context, *cognitoidentityprovider.ConfirmDeviceInput, ...aws.Option) (*cognitoidentityprovider.ConfirmDeviceOutput, error)
	ConfirmDeviceRequest(*cognitoidentityprovider.ConfirmDeviceInput) (*aws.Request, *cognitoidentityprovider.ConfirmDeviceOutput)

	ConfirmForgotPassword(*cognitoidentityprovider.ConfirmForgotPasswordInput) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
	ConfirmForgotPasswordWithContext(aws.Context, *cognitoidentityprovider.ConfirmForgotPasswordInput, ...aws.Option) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
	ConfirmForgotPasswordRequest(*cognitoidentityprovider.ConfirmForgotPasswordInput) (*aws.Request, *cognitoidentityprovider.ConfirmForgotPasswordOutput)

	ConfirmSignUp(*cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	ConfirmSignUpWithContext(aws.Context, *cognitoidentityprovider.ConfirmSignUpInput, ...aws.Option) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	ConfirmSignUpRequest(*cognitoidentityprovider.ConfirmSignUpInput) (*aws.Request, *cognitoidentityprovider.ConfirmSignUpOutput)

	CreateGroup(*cognitoidentityprovider.CreateGroupInput) (*cognitoidentityprovider.CreateGroupOutput, error)
	CreateGroupWithContext(aws.Context, *cognitoidentityprovider.CreateGroupInput, ...aws.Option) (*cognitoidentityprovider.CreateGroupOutput, error)
	CreateGroupRequest(*cognitoidentityprovider.CreateGroupInput) (*aws.Request, *cognitoidentityprovider.CreateGroupOutput)

	CreateIdentityProvider(*cognitoidentityprovider.CreateIdentityProviderInput) (*cognitoidentityprovider.CreateIdentityProviderOutput, error)
	CreateIdentityProviderWithContext(aws.Context, *cognitoidentityprovider.CreateIdentityProviderInput, ...aws.Option) (*cognitoidentityprovider.CreateIdentityProviderOutput, error)
	CreateIdentityProviderRequest(*cognitoidentityprovider.CreateIdentityProviderInput) (*aws.Request, *cognitoidentityprovider.CreateIdentityProviderOutput)

	CreateResourceServer(*cognitoidentityprovider.CreateResourceServerInput) (*cognitoidentityprovider.CreateResourceServerOutput, error)
	CreateResourceServerWithContext(aws.Context, *cognitoidentityprovider.CreateResourceServerInput, ...aws.Option) (*cognitoidentityprovider.CreateResourceServerOutput, error)
	CreateResourceServerRequest(*cognitoidentityprovider.CreateResourceServerInput) (*aws.Request, *cognitoidentityprovider.CreateResourceServerOutput)

	CreateUserImportJob(*cognitoidentityprovider.CreateUserImportJobInput) (*cognitoidentityprovider.CreateUserImportJobOutput, error)
	CreateUserImportJobWithContext(aws.Context, *cognitoidentityprovider.CreateUserImportJobInput, ...aws.Option) (*cognitoidentityprovider.CreateUserImportJobOutput, error)
	CreateUserImportJobRequest(*cognitoidentityprovider.CreateUserImportJobInput) (*aws.Request, *cognitoidentityprovider.CreateUserImportJobOutput)

	CreateUserPool(*cognitoidentityprovider.CreateUserPoolInput) (*cognitoidentityprovider.CreateUserPoolOutput, error)
	CreateUserPoolWithContext(aws.Context, *cognitoidentityprovider.CreateUserPoolInput, ...aws.Option) (*cognitoidentityprovider.CreateUserPoolOutput, error)
	CreateUserPoolRequest(*cognitoidentityprovider.CreateUserPoolInput) (*aws.Request, *cognitoidentityprovider.CreateUserPoolOutput)

	CreateUserPoolClient(*cognitoidentityprovider.CreateUserPoolClientInput) (*cognitoidentityprovider.CreateUserPoolClientOutput, error)
	CreateUserPoolClientWithContext(aws.Context, *cognitoidentityprovider.CreateUserPoolClientInput, ...aws.Option) (*cognitoidentityprovider.CreateUserPoolClientOutput, error)
	CreateUserPoolClientRequest(*cognitoidentityprovider.CreateUserPoolClientInput) (*aws.Request, *cognitoidentityprovider.CreateUserPoolClientOutput)

	CreateUserPoolDomain(*cognitoidentityprovider.CreateUserPoolDomainInput) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error)
	CreateUserPoolDomainWithContext(aws.Context, *cognitoidentityprovider.CreateUserPoolDomainInput, ...aws.Option) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error)
	CreateUserPoolDomainRequest(*cognitoidentityprovider.CreateUserPoolDomainInput) (*aws.Request, *cognitoidentityprovider.CreateUserPoolDomainOutput)

	DeleteGroup(*cognitoidentityprovider.DeleteGroupInput) (*cognitoidentityprovider.DeleteGroupOutput, error)
	DeleteGroupWithContext(aws.Context, *cognitoidentityprovider.DeleteGroupInput, ...aws.Option) (*cognitoidentityprovider.DeleteGroupOutput, error)
	DeleteGroupRequest(*cognitoidentityprovider.DeleteGroupInput) (*aws.Request, *cognitoidentityprovider.DeleteGroupOutput)

	DeleteIdentityProvider(*cognitoidentityprovider.DeleteIdentityProviderInput) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error)
	DeleteIdentityProviderWithContext(aws.Context, *cognitoidentityprovider.DeleteIdentityProviderInput, ...aws.Option) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error)
	DeleteIdentityProviderRequest(*cognitoidentityprovider.DeleteIdentityProviderInput) (*aws.Request, *cognitoidentityprovider.DeleteIdentityProviderOutput)

	DeleteResourceServer(*cognitoidentityprovider.DeleteResourceServerInput) (*cognitoidentityprovider.DeleteResourceServerOutput, error)
	DeleteResourceServerWithContext(aws.Context, *cognitoidentityprovider.DeleteResourceServerInput, ...aws.Option) (*cognitoidentityprovider.DeleteResourceServerOutput, error)
	DeleteResourceServerRequest(*cognitoidentityprovider.DeleteResourceServerInput) (*aws.Request, *cognitoidentityprovider.DeleteResourceServerOutput)

	DeleteUser(*cognitoidentityprovider.DeleteUserInput) (*cognitoidentityprovider.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *cognitoidentityprovider.DeleteUserInput, ...aws.Option) (*cognitoidentityprovider.DeleteUserOutput, error)
	DeleteUserRequest(*cognitoidentityprovider.DeleteUserInput) (*aws.Request, *cognitoidentityprovider.DeleteUserOutput)

	DeleteUserAttributes(*cognitoidentityprovider.DeleteUserAttributesInput) (*cognitoidentityprovider.DeleteUserAttributesOutput, error)
	DeleteUserAttributesWithContext(aws.Context, *cognitoidentityprovider.DeleteUserAttributesInput, ...aws.Option) (*cognitoidentityprovider.DeleteUserAttributesOutput, error)
	DeleteUserAttributesRequest(*cognitoidentityprovider.DeleteUserAttributesInput) (*aws.Request, *cognitoidentityprovider.DeleteUserAttributesOutput)

	DeleteUserPool(*cognitoidentityprovider.DeleteUserPoolInput) (*cognitoidentityprovider.DeleteUserPoolOutput, error)
	DeleteUserPoolWithContext(aws.Context, *cognitoidentityprovider.DeleteUserPoolInput, ...aws.Option) (*cognitoidentityprovider.DeleteUserPoolOutput, error)
	DeleteUserPoolRequest(*cognitoidentityprovider.DeleteUserPoolInput) (*aws.Request, *cognitoidentityprovider.DeleteUserPoolOutput)

	DeleteUserPoolClient(*cognitoidentityprovider.DeleteUserPoolClientInput) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error)
	DeleteUserPoolClientWithContext(aws.Context, *cognitoidentityprovider.DeleteUserPoolClientInput, ...aws.Option) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error)
	DeleteUserPoolClientRequest(*cognitoidentityprovider.DeleteUserPoolClientInput) (*aws.Request, *cognitoidentityprovider.DeleteUserPoolClientOutput)

	DeleteUserPoolDomain(*cognitoidentityprovider.DeleteUserPoolDomainInput) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error)
	DeleteUserPoolDomainWithContext(aws.Context, *cognitoidentityprovider.DeleteUserPoolDomainInput, ...aws.Option) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error)
	DeleteUserPoolDomainRequest(*cognitoidentityprovider.DeleteUserPoolDomainInput) (*aws.Request, *cognitoidentityprovider.DeleteUserPoolDomainOutput)

	DescribeIdentityProvider(*cognitoidentityprovider.DescribeIdentityProviderInput) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error)
	DescribeIdentityProviderWithContext(aws.Context, *cognitoidentityprovider.DescribeIdentityProviderInput, ...aws.Option) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error)
	DescribeIdentityProviderRequest(*cognitoidentityprovider.DescribeIdentityProviderInput) (*aws.Request, *cognitoidentityprovider.DescribeIdentityProviderOutput)

	DescribeResourceServer(*cognitoidentityprovider.DescribeResourceServerInput) (*cognitoidentityprovider.DescribeResourceServerOutput, error)
	DescribeResourceServerWithContext(aws.Context, *cognitoidentityprovider.DescribeResourceServerInput, ...aws.Option) (*cognitoidentityprovider.DescribeResourceServerOutput, error)
	DescribeResourceServerRequest(*cognitoidentityprovider.DescribeResourceServerInput) (*aws.Request, *cognitoidentityprovider.DescribeResourceServerOutput)

	DescribeUserImportJob(*cognitoidentityprovider.DescribeUserImportJobInput) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
	DescribeUserImportJobWithContext(aws.Context, *cognitoidentityprovider.DescribeUserImportJobInput, ...aws.Option) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
	DescribeUserImportJobRequest(*cognitoidentityprovider.DescribeUserImportJobInput) (*aws.Request, *cognitoidentityprovider.DescribeUserImportJobOutput)

	DescribeUserPool(*cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	DescribeUserPoolWithContext(aws.Context, *cognitoidentityprovider.DescribeUserPoolInput, ...aws.Option) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	DescribeUserPoolRequest(*cognitoidentityprovider.DescribeUserPoolInput) (*aws.Request, *cognitoidentityprovider.DescribeUserPoolOutput)

	DescribeUserPoolClient(*cognitoidentityprovider.DescribeUserPoolClientInput) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
	DescribeUserPoolClientWithContext(aws.Context, *cognitoidentityprovider.DescribeUserPoolClientInput, ...aws.Option) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
	DescribeUserPoolClientRequest(*cognitoidentityprovider.DescribeUserPoolClientInput) (*aws.Request, *cognitoidentityprovider.DescribeUserPoolClientOutput)

	DescribeUserPoolDomain(*cognitoidentityprovider.DescribeUserPoolDomainInput) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error)
	DescribeUserPoolDomainWithContext(aws.Context, *cognitoidentityprovider.DescribeUserPoolDomainInput, ...aws.Option) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error)
	DescribeUserPoolDomainRequest(*cognitoidentityprovider.DescribeUserPoolDomainInput) (*aws.Request, *cognitoidentityprovider.DescribeUserPoolDomainOutput)

	ForgetDevice(*cognitoidentityprovider.ForgetDeviceInput) (*cognitoidentityprovider.ForgetDeviceOutput, error)
	ForgetDeviceWithContext(aws.Context, *cognitoidentityprovider.ForgetDeviceInput, ...aws.Option) (*cognitoidentityprovider.ForgetDeviceOutput, error)
	ForgetDeviceRequest(*cognitoidentityprovider.ForgetDeviceInput) (*aws.Request, *cognitoidentityprovider.ForgetDeviceOutput)

	ForgotPassword(*cognitoidentityprovider.ForgotPasswordInput) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	ForgotPasswordWithContext(aws.Context, *cognitoidentityprovider.ForgotPasswordInput, ...aws.Option) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	ForgotPasswordRequest(*cognitoidentityprovider.ForgotPasswordInput) (*aws.Request, *cognitoidentityprovider.ForgotPasswordOutput)

	GetCSVHeader(*cognitoidentityprovider.GetCSVHeaderInput) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
	GetCSVHeaderWithContext(aws.Context, *cognitoidentityprovider.GetCSVHeaderInput, ...aws.Option) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
	GetCSVHeaderRequest(*cognitoidentityprovider.GetCSVHeaderInput) (*aws.Request, *cognitoidentityprovider.GetCSVHeaderOutput)

	GetDevice(*cognitoidentityprovider.GetDeviceInput) (*cognitoidentityprovider.GetDeviceOutput, error)
	GetDeviceWithContext(aws.Context, *cognitoidentityprovider.GetDeviceInput, ...aws.Option) (*cognitoidentityprovider.GetDeviceOutput, error)
	GetDeviceRequest(*cognitoidentityprovider.GetDeviceInput) (*aws.Request, *cognitoidentityprovider.GetDeviceOutput)

	GetGroup(*cognitoidentityprovider.GetGroupInput) (*cognitoidentityprovider.GetGroupOutput, error)
	GetGroupWithContext(aws.Context, *cognitoidentityprovider.GetGroupInput, ...aws.Option) (*cognitoidentityprovider.GetGroupOutput, error)
	GetGroupRequest(*cognitoidentityprovider.GetGroupInput) (*aws.Request, *cognitoidentityprovider.GetGroupOutput)

	GetIdentityProviderByIdentifier(*cognitoidentityprovider.GetIdentityProviderByIdentifierInput) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error)
	GetIdentityProviderByIdentifierWithContext(aws.Context, *cognitoidentityprovider.GetIdentityProviderByIdentifierInput, ...aws.Option) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error)
	GetIdentityProviderByIdentifierRequest(*cognitoidentityprovider.GetIdentityProviderByIdentifierInput) (*aws.Request, *cognitoidentityprovider.GetIdentityProviderByIdentifierOutput)

	GetUICustomization(*cognitoidentityprovider.GetUICustomizationInput) (*cognitoidentityprovider.GetUICustomizationOutput, error)
	GetUICustomizationWithContext(aws.Context, *cognitoidentityprovider.GetUICustomizationInput, ...aws.Option) (*cognitoidentityprovider.GetUICustomizationOutput, error)
	GetUICustomizationRequest(*cognitoidentityprovider.GetUICustomizationInput) (*aws.Request, *cognitoidentityprovider.GetUICustomizationOutput)

	GetUser(*cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error)
	GetUserWithContext(aws.Context, *cognitoidentityprovider.GetUserInput, ...aws.Option) (*cognitoidentityprovider.GetUserOutput, error)
	GetUserRequest(*cognitoidentityprovider.GetUserInput) (*aws.Request, *cognitoidentityprovider.GetUserOutput)

	GetUserAttributeVerificationCode(*cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
	GetUserAttributeVerificationCodeWithContext(aws.Context, *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, ...aws.Option) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
	GetUserAttributeVerificationCodeRequest(*cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*aws.Request, *cognitoidentityprovider.GetUserAttributeVerificationCodeOutput)

	GlobalSignOut(*cognitoidentityprovider.GlobalSignOutInput) (*cognitoidentityprovider.GlobalSignOutOutput, error)
	GlobalSignOutWithContext(aws.Context, *cognitoidentityprovider.GlobalSignOutInput, ...aws.Option) (*cognitoidentityprovider.GlobalSignOutOutput, error)
	GlobalSignOutRequest(*cognitoidentityprovider.GlobalSignOutInput) (*aws.Request, *cognitoidentityprovider.GlobalSignOutOutput)

	InitiateAuth(*cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error)
	InitiateAuthWithContext(aws.Context, *cognitoidentityprovider.InitiateAuthInput, ...aws.Option) (*cognitoidentityprovider.InitiateAuthOutput, error)
	InitiateAuthRequest(*cognitoidentityprovider.InitiateAuthInput) (*aws.Request, *cognitoidentityprovider.InitiateAuthOutput)

	ListDevices(*cognitoidentityprovider.ListDevicesInput) (*cognitoidentityprovider.ListDevicesOutput, error)
	ListDevicesWithContext(aws.Context, *cognitoidentityprovider.ListDevicesInput, ...aws.Option) (*cognitoidentityprovider.ListDevicesOutput, error)
	ListDevicesRequest(*cognitoidentityprovider.ListDevicesInput) (*aws.Request, *cognitoidentityprovider.ListDevicesOutput)

	ListGroups(*cognitoidentityprovider.ListGroupsInput) (*cognitoidentityprovider.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *cognitoidentityprovider.ListGroupsInput, ...aws.Option) (*cognitoidentityprovider.ListGroupsOutput, error)
	ListGroupsRequest(*cognitoidentityprovider.ListGroupsInput) (*aws.Request, *cognitoidentityprovider.ListGroupsOutput)

	ListIdentityProviders(*cognitoidentityprovider.ListIdentityProvidersInput) (*cognitoidentityprovider.ListIdentityProvidersOutput, error)
	ListIdentityProvidersWithContext(aws.Context, *cognitoidentityprovider.ListIdentityProvidersInput, ...aws.Option) (*cognitoidentityprovider.ListIdentityProvidersOutput, error)
	ListIdentityProvidersRequest(*cognitoidentityprovider.ListIdentityProvidersInput) (*aws.Request, *cognitoidentityprovider.ListIdentityProvidersOutput)

	ListResourceServers(*cognitoidentityprovider.ListResourceServersInput) (*cognitoidentityprovider.ListResourceServersOutput, error)
	ListResourceServersWithContext(aws.Context, *cognitoidentityprovider.ListResourceServersInput, ...aws.Option) (*cognitoidentityprovider.ListResourceServersOutput, error)
	ListResourceServersRequest(*cognitoidentityprovider.ListResourceServersInput) (*aws.Request, *cognitoidentityprovider.ListResourceServersOutput)

	ListUserImportJobs(*cognitoidentityprovider.ListUserImportJobsInput) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
	ListUserImportJobsWithContext(aws.Context, *cognitoidentityprovider.ListUserImportJobsInput, ...aws.Option) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
	ListUserImportJobsRequest(*cognitoidentityprovider.ListUserImportJobsInput) (*aws.Request, *cognitoidentityprovider.ListUserImportJobsOutput)

	ListUserPoolClients(*cognitoidentityprovider.ListUserPoolClientsInput) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
	ListUserPoolClientsWithContext(aws.Context, *cognitoidentityprovider.ListUserPoolClientsInput, ...aws.Option) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
	ListUserPoolClientsRequest(*cognitoidentityprovider.ListUserPoolClientsInput) (*aws.Request, *cognitoidentityprovider.ListUserPoolClientsOutput)

	ListUserPools(*cognitoidentityprovider.ListUserPoolsInput) (*cognitoidentityprovider.ListUserPoolsOutput, error)
	ListUserPoolsWithContext(aws.Context, *cognitoidentityprovider.ListUserPoolsInput, ...aws.Option) (*cognitoidentityprovider.ListUserPoolsOutput, error)
	ListUserPoolsRequest(*cognitoidentityprovider.ListUserPoolsInput) (*aws.Request, *cognitoidentityprovider.ListUserPoolsOutput)

	ListUsers(*cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *cognitoidentityprovider.ListUsersInput, ...aws.Option) (*cognitoidentityprovider.ListUsersOutput, error)
	ListUsersRequest(*cognitoidentityprovider.ListUsersInput) (*aws.Request, *cognitoidentityprovider.ListUsersOutput)

	ListUsersInGroup(*cognitoidentityprovider.ListUsersInGroupInput) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
	ListUsersInGroupWithContext(aws.Context, *cognitoidentityprovider.ListUsersInGroupInput, ...aws.Option) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
	ListUsersInGroupRequest(*cognitoidentityprovider.ListUsersInGroupInput) (*aws.Request, *cognitoidentityprovider.ListUsersInGroupOutput)

	ResendConfirmationCode(*cognitoidentityprovider.ResendConfirmationCodeInput) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	ResendConfirmationCodeWithContext(aws.Context, *cognitoidentityprovider.ResendConfirmationCodeInput, ...aws.Option) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	ResendConfirmationCodeRequest(*cognitoidentityprovider.ResendConfirmationCodeInput) (*aws.Request, *cognitoidentityprovider.ResendConfirmationCodeOutput)

	RespondToAuthChallenge(*cognitoidentityprovider.RespondToAuthChallengeInput) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error)
	RespondToAuthChallengeWithContext(aws.Context, *cognitoidentityprovider.RespondToAuthChallengeInput, ...aws.Option) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error)
	RespondToAuthChallengeRequest(*cognitoidentityprovider.RespondToAuthChallengeInput) (*aws.Request, *cognitoidentityprovider.RespondToAuthChallengeOutput)

	SetUICustomization(*cognitoidentityprovider.SetUICustomizationInput) (*cognitoidentityprovider.SetUICustomizationOutput, error)
	SetUICustomizationWithContext(aws.Context, *cognitoidentityprovider.SetUICustomizationInput, ...aws.Option) (*cognitoidentityprovider.SetUICustomizationOutput, error)
	SetUICustomizationRequest(*cognitoidentityprovider.SetUICustomizationInput) (*aws.Request, *cognitoidentityprovider.SetUICustomizationOutput)

	SetUserSettings(*cognitoidentityprovider.SetUserSettingsInput) (*cognitoidentityprovider.SetUserSettingsOutput, error)
	SetUserSettingsWithContext(aws.Context, *cognitoidentityprovider.SetUserSettingsInput, ...aws.Option) (*cognitoidentityprovider.SetUserSettingsOutput, error)
	SetUserSettingsRequest(*cognitoidentityprovider.SetUserSettingsInput) (*aws.Request, *cognitoidentityprovider.SetUserSettingsOutput)

	SignUp(*cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error)
	SignUpWithContext(aws.Context, *cognitoidentityprovider.SignUpInput, ...aws.Option) (*cognitoidentityprovider.SignUpOutput, error)
	SignUpRequest(*cognitoidentityprovider.SignUpInput) (*aws.Request, *cognitoidentityprovider.SignUpOutput)

	StartUserImportJob(*cognitoidentityprovider.StartUserImportJobInput) (*cognitoidentityprovider.StartUserImportJobOutput, error)
	StartUserImportJobWithContext(aws.Context, *cognitoidentityprovider.StartUserImportJobInput, ...aws.Option) (*cognitoidentityprovider.StartUserImportJobOutput, error)
	StartUserImportJobRequest(*cognitoidentityprovider.StartUserImportJobInput) (*aws.Request, *cognitoidentityprovider.StartUserImportJobOutput)

	StopUserImportJob(*cognitoidentityprovider.StopUserImportJobInput) (*cognitoidentityprovider.StopUserImportJobOutput, error)
	StopUserImportJobWithContext(aws.Context, *cognitoidentityprovider.StopUserImportJobInput, ...aws.Option) (*cognitoidentityprovider.StopUserImportJobOutput, error)
	StopUserImportJobRequest(*cognitoidentityprovider.StopUserImportJobInput) (*aws.Request, *cognitoidentityprovider.StopUserImportJobOutput)

	UpdateDeviceStatus(*cognitoidentityprovider.UpdateDeviceStatusInput) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error)
	UpdateDeviceStatusWithContext(aws.Context, *cognitoidentityprovider.UpdateDeviceStatusInput, ...aws.Option) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error)
	UpdateDeviceStatusRequest(*cognitoidentityprovider.UpdateDeviceStatusInput) (*aws.Request, *cognitoidentityprovider.UpdateDeviceStatusOutput)

	UpdateGroup(*cognitoidentityprovider.UpdateGroupInput) (*cognitoidentityprovider.UpdateGroupOutput, error)
	UpdateGroupWithContext(aws.Context, *cognitoidentityprovider.UpdateGroupInput, ...aws.Option) (*cognitoidentityprovider.UpdateGroupOutput, error)
	UpdateGroupRequest(*cognitoidentityprovider.UpdateGroupInput) (*aws.Request, *cognitoidentityprovider.UpdateGroupOutput)

	UpdateIdentityProvider(*cognitoidentityprovider.UpdateIdentityProviderInput) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error)
	UpdateIdentityProviderWithContext(aws.Context, *cognitoidentityprovider.UpdateIdentityProviderInput, ...aws.Option) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error)
	UpdateIdentityProviderRequest(*cognitoidentityprovider.UpdateIdentityProviderInput) (*aws.Request, *cognitoidentityprovider.UpdateIdentityProviderOutput)

	UpdateResourceServer(*cognitoidentityprovider.UpdateResourceServerInput) (*cognitoidentityprovider.UpdateResourceServerOutput, error)
	UpdateResourceServerWithContext(aws.Context, *cognitoidentityprovider.UpdateResourceServerInput, ...aws.Option) (*cognitoidentityprovider.UpdateResourceServerOutput, error)
	UpdateResourceServerRequest(*cognitoidentityprovider.UpdateResourceServerInput) (*aws.Request, *cognitoidentityprovider.UpdateResourceServerOutput)

	UpdateUserAttributes(*cognitoidentityprovider.UpdateUserAttributesInput) (*cognitoidentityprovider.UpdateUserAttributesOutput, error)
	UpdateUserAttributesWithContext(aws.Context, *cognitoidentityprovider.UpdateUserAttributesInput, ...aws.Option) (*cognitoidentityprovider.UpdateUserAttributesOutput, error)
	UpdateUserAttributesRequest(*cognitoidentityprovider.UpdateUserAttributesInput) (*aws.Request, *cognitoidentityprovider.UpdateUserAttributesOutput)

	UpdateUserPool(*cognitoidentityprovider.UpdateUserPoolInput) (*cognitoidentityprovider.UpdateUserPoolOutput, error)
	UpdateUserPoolWithContext(aws.Context, *cognitoidentityprovider.UpdateUserPoolInput, ...aws.Option) (*cognitoidentityprovider.UpdateUserPoolOutput, error)
	UpdateUserPoolRequest(*cognitoidentityprovider.UpdateUserPoolInput) (*aws.Request, *cognitoidentityprovider.UpdateUserPoolOutput)

	UpdateUserPoolClient(*cognitoidentityprovider.UpdateUserPoolClientInput) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error)
	UpdateUserPoolClientWithContext(aws.Context, *cognitoidentityprovider.UpdateUserPoolClientInput, ...aws.Option) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error)
	UpdateUserPoolClientRequest(*cognitoidentityprovider.UpdateUserPoolClientInput) (*aws.Request, *cognitoidentityprovider.UpdateUserPoolClientOutput)

	VerifyUserAttribute(*cognitoidentityprovider.VerifyUserAttributeInput) (*cognitoidentityprovider.VerifyUserAttributeOutput, error)
	VerifyUserAttributeWithContext(aws.Context, *cognitoidentityprovider.VerifyUserAttributeInput, ...aws.Option) (*cognitoidentityprovider.VerifyUserAttributeOutput, error)
	VerifyUserAttributeRequest(*cognitoidentityprovider.VerifyUserAttributeInput) (*aws.Request, *cognitoidentityprovider.VerifyUserAttributeOutput)
}

var _ CognitoIdentityProviderAPI = (*cognitoidentityprovider.CognitoIdentityProvider)(nil)
