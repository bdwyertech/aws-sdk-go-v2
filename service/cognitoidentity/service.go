// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// CognitoIdentity provides the API operation methods for making requests to
// Amazon Cognito Identity. See this package's package overview docs
// for details on the service.
//
// CognitoIdentity methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type CognitoIdentity struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*aws.Client)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// Service information constants
const (
	ServiceName = "cognito-identity" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName        // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the CognitoIdentity client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CognitoIdentity client from just a config.
//     svc := cognitoidentity.New(myConfig)
//
//     // Create a CognitoIdentity client with additional configuration
//     svc := cognitoidentity.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *CognitoIdentity {
	var signingName string
	signingRegion := aws.StringValue(config.Region)

	svc := &CognitoIdentity{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2014-06-30",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSCognitoIdentityService",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a CognitoIdentity operation and runs any
// custom request initialization.
func (c *CognitoIdentity) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
