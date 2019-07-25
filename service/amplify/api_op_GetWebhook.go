// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for the get webhook request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetWebhookRequest
type GetWebhookInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for a webhook.
	//
	// WebhookId is a required field
	WebhookId *string `location:"uri" locationName:"webhookId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetWebhookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetWebhookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetWebhookInput"}

	if s.WebhookId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebhookId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetWebhookInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.WebhookId != nil {
		v := *s.WebhookId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "webhookId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for the get webhook request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetWebhookResult
type GetWebhookOutput struct {
	_ struct{} `type:"structure"`

	// Webhook structure.
	//
	// Webhook is a required field
	Webhook *Webhook `locationName:"webhook" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetWebhookOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetWebhookOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Webhook != nil {
		v := s.Webhook

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "webhook", v, metadata)
	}
	return nil
}

const opGetWebhook = "GetWebhook"

// GetWebhookRequest returns a request value for making API operation for
// AWS Amplify.
//
// Retrieves webhook info that corresponds to a webhookId.
//
//    // Example sending a request using GetWebhookRequest.
//    req := client.GetWebhookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetWebhook
func (c *Client) GetWebhookRequest(input *GetWebhookInput) GetWebhookRequest {
	op := &aws.Operation{
		Name:       opGetWebhook,
		HTTPMethod: "GET",
		HTTPPath:   "/webhooks/{webhookId}",
	}

	if input == nil {
		input = &GetWebhookInput{}
	}

	req := c.newRequest(op, input, &GetWebhookOutput{})
	return GetWebhookRequest{Request: req, Input: input, Copy: c.GetWebhookRequest}
}

// GetWebhookRequest is the request type for the
// GetWebhook API operation.
type GetWebhookRequest struct {
	*aws.Request
	Input *GetWebhookInput
	Copy  func(*GetWebhookInput) GetWebhookRequest
}

// Send marshals and sends the GetWebhook API request.
func (r GetWebhookRequest) Send(ctx context.Context) (*GetWebhookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetWebhookResponse{
		GetWebhookOutput: r.Request.Data.(*GetWebhookOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetWebhookResponse is the response type for the
// GetWebhook API operation.
type GetWebhookResponse struct {
	*GetWebhookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetWebhook request.
func (r *GetWebhookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
