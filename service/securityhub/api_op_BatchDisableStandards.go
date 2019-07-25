// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/BatchDisableStandardsRequest
type BatchDisableStandardsInput struct {
	_ struct{} `type:"structure"`

	// The ARNs of the standards subscriptions to disable.
	//
	// StandardsSubscriptionArns is a required field
	StandardsSubscriptionArns []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDisableStandardsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDisableStandardsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDisableStandardsInput"}

	if s.StandardsSubscriptionArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("StandardsSubscriptionArns"))
	}
	if s.StandardsSubscriptionArns != nil && len(s.StandardsSubscriptionArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StandardsSubscriptionArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchDisableStandardsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.StandardsSubscriptionArns != nil {
		v := s.StandardsSubscriptionArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "StandardsSubscriptionArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/BatchDisableStandardsResponse
type BatchDisableStandardsOutput struct {
	_ struct{} `type:"structure"`

	// The details of the standards subscriptions that were disabled.
	StandardsSubscriptions []StandardsSubscription `type:"list"`
}

// String returns the string representation
func (s BatchDisableStandardsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchDisableStandardsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.StandardsSubscriptions != nil {
		v := s.StandardsSubscriptions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "StandardsSubscriptions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchDisableStandards = "BatchDisableStandards"

// BatchDisableStandardsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Disables the standards specified by the provided StandardsSubscriptionArns.
// For more information, see Standards Supported in AWS Security Hub (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards.html).
//
//    // Example sending a request using BatchDisableStandardsRequest.
//    req := client.BatchDisableStandardsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/BatchDisableStandards
func (c *Client) BatchDisableStandardsRequest(input *BatchDisableStandardsInput) BatchDisableStandardsRequest {
	op := &aws.Operation{
		Name:       opBatchDisableStandards,
		HTTPMethod: "POST",
		HTTPPath:   "/standards/deregister",
	}

	if input == nil {
		input = &BatchDisableStandardsInput{}
	}

	req := c.newRequest(op, input, &BatchDisableStandardsOutput{})
	return BatchDisableStandardsRequest{Request: req, Input: input, Copy: c.BatchDisableStandardsRequest}
}

// BatchDisableStandardsRequest is the request type for the
// BatchDisableStandards API operation.
type BatchDisableStandardsRequest struct {
	*aws.Request
	Input *BatchDisableStandardsInput
	Copy  func(*BatchDisableStandardsInput) BatchDisableStandardsRequest
}

// Send marshals and sends the BatchDisableStandards API request.
func (r BatchDisableStandardsRequest) Send(ctx context.Context) (*BatchDisableStandardsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDisableStandardsResponse{
		BatchDisableStandardsOutput: r.Request.Data.(*BatchDisableStandardsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDisableStandardsResponse is the response type for the
// BatchDisableStandards API operation.
type BatchDisableStandardsResponse struct {
	*BatchDisableStandardsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDisableStandards request.
func (r *BatchDisableStandardsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
