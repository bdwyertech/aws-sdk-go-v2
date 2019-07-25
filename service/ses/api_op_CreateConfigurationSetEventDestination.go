// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to create a configuration set event destination. A configuration
// set event destination, which can be either Amazon CloudWatch or Amazon Kinesis
// Firehose, describes an AWS service in which Amazon SES publishes the email
// sending events associated with a configuration set. For information about
// using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateConfigurationSetEventDestinationRequest
type CreateConfigurationSetEventDestinationInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that the event destination should be associated
	// with.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// An object that describes the AWS service that email sending event information
	// will be published to.
	//
	// EventDestination is a required field
	EventDestination *EventDestination `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateConfigurationSetEventDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConfigurationSetEventDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConfigurationSetEventDestinationInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if s.EventDestination == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventDestination"))
	}
	if s.EventDestination != nil {
		if err := s.EventDestination.Validate(); err != nil {
			invalidParams.AddNested("EventDestination", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateConfigurationSetEventDestinationResponse
type CreateConfigurationSetEventDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateConfigurationSetEventDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateConfigurationSetEventDestination = "CreateConfigurationSetEventDestination"

// CreateConfigurationSetEventDestinationRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates a configuration set event destination.
//
// When you create or update an event destination, you must provide one, and
// only one, destination. The destination can be CloudWatch, Amazon Kinesis
// Firehose, or Amazon Simple Notification Service (Amazon SNS).
//
// An event destination is the AWS service to which Amazon SES publishes the
// email sending events associated with a configuration set. For information
// about using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using CreateConfigurationSetEventDestinationRequest.
//    req := client.CreateConfigurationSetEventDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateConfigurationSetEventDestination
func (c *Client) CreateConfigurationSetEventDestinationRequest(input *CreateConfigurationSetEventDestinationInput) CreateConfigurationSetEventDestinationRequest {
	op := &aws.Operation{
		Name:       opCreateConfigurationSetEventDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateConfigurationSetEventDestinationInput{}
	}

	req := c.newRequest(op, input, &CreateConfigurationSetEventDestinationOutput{})
	return CreateConfigurationSetEventDestinationRequest{Request: req, Input: input, Copy: c.CreateConfigurationSetEventDestinationRequest}
}

// CreateConfigurationSetEventDestinationRequest is the request type for the
// CreateConfigurationSetEventDestination API operation.
type CreateConfigurationSetEventDestinationRequest struct {
	*aws.Request
	Input *CreateConfigurationSetEventDestinationInput
	Copy  func(*CreateConfigurationSetEventDestinationInput) CreateConfigurationSetEventDestinationRequest
}

// Send marshals and sends the CreateConfigurationSetEventDestination API request.
func (r CreateConfigurationSetEventDestinationRequest) Send(ctx context.Context) (*CreateConfigurationSetEventDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConfigurationSetEventDestinationResponse{
		CreateConfigurationSetEventDestinationOutput: r.Request.Data.(*CreateConfigurationSetEventDestinationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConfigurationSetEventDestinationResponse is the response type for the
// CreateConfigurationSetEventDestination API operation.
type CreateConfigurationSetEventDestinationResponse struct {
	*CreateConfigurationSetEventDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConfigurationSetEventDestination request.
func (r *CreateConfigurationSetEventDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
