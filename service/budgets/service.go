// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Budgets provides the API operation methods for making requests to
// AWS Budgets. See this package's package overview docs
// for details on the service.
//
// Budgets methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Budgets struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*aws.Client)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// Service information constants
const (
	ServiceName = "budgets"   // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Budgets client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Budgets client from just a config.
//     svc := budgets.New(myConfig)
//
//     // Create a Budgets client with additional configuration
//     svc := budgets.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *Budgets {
	var signingName string
	signingRegion := aws.StringValue(config.Region)

	svc := &Budgets{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2016-10-20",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSBudgetServiceGateway",
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

// newRequest creates a new request for a Budgets operation and runs any
// custom request initialization.
func (c *Budgets) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
