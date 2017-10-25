// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cognitoidentityiface provides an interface to enable mocking the Amazon Cognito Identity service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cognitoidentityiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
)

// CognitoIdentityAPI provides an interface to enable mocking the
// cognitoidentity.CognitoIdentity service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Cognito Identity.
//    func myFunc(svc cognitoidentityiface.CognitoIdentityAPI) bool {
//        // Make svc.CreateIdentityPool request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cognitoidentity.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCognitoIdentityClient struct {
//        cognitoidentityiface.CognitoIdentityAPI
//    }
//    func (m *mockCognitoIdentityClient) CreateIdentityPool(input *cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCognitoIdentityClient{}
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
type CognitoIdentityAPI interface {
	CreateIdentityPool(*cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error)
	CreateIdentityPoolWithContext(aws.Context, *cognitoidentity.CreateIdentityPoolInput, ...aws.Option) (*cognitoidentity.IdentityPool, error)
	CreateIdentityPoolRequest(*cognitoidentity.CreateIdentityPoolInput) (*aws.Request, *cognitoidentity.IdentityPool)

	DeleteIdentities(*cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error)
	DeleteIdentitiesWithContext(aws.Context, *cognitoidentity.DeleteIdentitiesInput, ...aws.Option) (*cognitoidentity.DeleteIdentitiesOutput, error)
	DeleteIdentitiesRequest(*cognitoidentity.DeleteIdentitiesInput) (*aws.Request, *cognitoidentity.DeleteIdentitiesOutput)

	DeleteIdentityPool(*cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error)
	DeleteIdentityPoolWithContext(aws.Context, *cognitoidentity.DeleteIdentityPoolInput, ...aws.Option) (*cognitoidentity.DeleteIdentityPoolOutput, error)
	DeleteIdentityPoolRequest(*cognitoidentity.DeleteIdentityPoolInput) (*aws.Request, *cognitoidentity.DeleteIdentityPoolOutput)

	DescribeIdentity(*cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error)
	DescribeIdentityWithContext(aws.Context, *cognitoidentity.DescribeIdentityInput, ...aws.Option) (*cognitoidentity.IdentityDescription, error)
	DescribeIdentityRequest(*cognitoidentity.DescribeIdentityInput) (*aws.Request, *cognitoidentity.IdentityDescription)

	DescribeIdentityPool(*cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error)
	DescribeIdentityPoolWithContext(aws.Context, *cognitoidentity.DescribeIdentityPoolInput, ...aws.Option) (*cognitoidentity.IdentityPool, error)
	DescribeIdentityPoolRequest(*cognitoidentity.DescribeIdentityPoolInput) (*aws.Request, *cognitoidentity.IdentityPool)

	GetCredentialsForIdentity(*cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error)
	GetCredentialsForIdentityWithContext(aws.Context, *cognitoidentity.GetCredentialsForIdentityInput, ...aws.Option) (*cognitoidentity.GetCredentialsForIdentityOutput, error)
	GetCredentialsForIdentityRequest(*cognitoidentity.GetCredentialsForIdentityInput) (*aws.Request, *cognitoidentity.GetCredentialsForIdentityOutput)

	GetId(*cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error)
	GetIdWithContext(aws.Context, *cognitoidentity.GetIdInput, ...aws.Option) (*cognitoidentity.GetIdOutput, error)
	GetIdRequest(*cognitoidentity.GetIdInput) (*aws.Request, *cognitoidentity.GetIdOutput)

	GetIdentityPoolRoles(*cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error)
	GetIdentityPoolRolesWithContext(aws.Context, *cognitoidentity.GetIdentityPoolRolesInput, ...aws.Option) (*cognitoidentity.GetIdentityPoolRolesOutput, error)
	GetIdentityPoolRolesRequest(*cognitoidentity.GetIdentityPoolRolesInput) (*aws.Request, *cognitoidentity.GetIdentityPoolRolesOutput)

	GetOpenIdToken(*cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error)
	GetOpenIdTokenWithContext(aws.Context, *cognitoidentity.GetOpenIdTokenInput, ...aws.Option) (*cognitoidentity.GetOpenIdTokenOutput, error)
	GetOpenIdTokenRequest(*cognitoidentity.GetOpenIdTokenInput) (*aws.Request, *cognitoidentity.GetOpenIdTokenOutput)

	GetOpenIdTokenForDeveloperIdentity(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error)
	GetOpenIdTokenForDeveloperIdentityWithContext(aws.Context, *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, ...aws.Option) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error)
	GetOpenIdTokenForDeveloperIdentityRequest(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*aws.Request, *cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput)

	ListIdentities(*cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error)
	ListIdentitiesWithContext(aws.Context, *cognitoidentity.ListIdentitiesInput, ...aws.Option) (*cognitoidentity.ListIdentitiesOutput, error)
	ListIdentitiesRequest(*cognitoidentity.ListIdentitiesInput) (*aws.Request, *cognitoidentity.ListIdentitiesOutput)

	ListIdentityPools(*cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error)
	ListIdentityPoolsWithContext(aws.Context, *cognitoidentity.ListIdentityPoolsInput, ...aws.Option) (*cognitoidentity.ListIdentityPoolsOutput, error)
	ListIdentityPoolsRequest(*cognitoidentity.ListIdentityPoolsInput) (*aws.Request, *cognitoidentity.ListIdentityPoolsOutput)

	LookupDeveloperIdentity(*cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error)
	LookupDeveloperIdentityWithContext(aws.Context, *cognitoidentity.LookupDeveloperIdentityInput, ...aws.Option) (*cognitoidentity.LookupDeveloperIdentityOutput, error)
	LookupDeveloperIdentityRequest(*cognitoidentity.LookupDeveloperIdentityInput) (*aws.Request, *cognitoidentity.LookupDeveloperIdentityOutput)

	MergeDeveloperIdentities(*cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error)
	MergeDeveloperIdentitiesWithContext(aws.Context, *cognitoidentity.MergeDeveloperIdentitiesInput, ...aws.Option) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error)
	MergeDeveloperIdentitiesRequest(*cognitoidentity.MergeDeveloperIdentitiesInput) (*aws.Request, *cognitoidentity.MergeDeveloperIdentitiesOutput)

	SetIdentityPoolRoles(*cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error)
	SetIdentityPoolRolesWithContext(aws.Context, *cognitoidentity.SetIdentityPoolRolesInput, ...aws.Option) (*cognitoidentity.SetIdentityPoolRolesOutput, error)
	SetIdentityPoolRolesRequest(*cognitoidentity.SetIdentityPoolRolesInput) (*aws.Request, *cognitoidentity.SetIdentityPoolRolesOutput)

	UnlinkDeveloperIdentity(*cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error)
	UnlinkDeveloperIdentityWithContext(aws.Context, *cognitoidentity.UnlinkDeveloperIdentityInput, ...aws.Option) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error)
	UnlinkDeveloperIdentityRequest(*cognitoidentity.UnlinkDeveloperIdentityInput) (*aws.Request, *cognitoidentity.UnlinkDeveloperIdentityOutput)

	UnlinkIdentity(*cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error)
	UnlinkIdentityWithContext(aws.Context, *cognitoidentity.UnlinkIdentityInput, ...aws.Option) (*cognitoidentity.UnlinkIdentityOutput, error)
	UnlinkIdentityRequest(*cognitoidentity.UnlinkIdentityInput) (*aws.Request, *cognitoidentity.UnlinkIdentityOutput)

	UpdateIdentityPool(*cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error)
	UpdateIdentityPoolWithContext(aws.Context, *cognitoidentity.IdentityPool, ...aws.Option) (*cognitoidentity.IdentityPool, error)
	UpdateIdentityPoolRequest(*cognitoidentity.IdentityPool) (*aws.Request, *cognitoidentity.IdentityPool)
}

var _ CognitoIdentityAPI = (*cognitoidentity.CognitoIdentity)(nil)
