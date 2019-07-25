// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSmsChannelRequest
type GetSmsChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSmsChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSmsChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSmsChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSmsChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSmsChannelResponse
type GetSmsChannelOutput struct {
	_ struct{} `type:"structure" payload:"SMSChannelResponse"`

	// Provides information about the status and settings of the SMS channel for
	// an application.
	//
	// SMSChannelResponse is a required field
	SMSChannelResponse *SMSChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSmsChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSmsChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SMSChannelResponse != nil {
		v := s.SMSChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SMSChannelResponse", v, metadata)
	}
	return nil
}

const opGetSmsChannel = "GetSmsChannel"

// GetSmsChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of the SMS channel for
// an application.
//
//    // Example sending a request using GetSmsChannelRequest.
//    req := client.GetSmsChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSmsChannel
func (c *Client) GetSmsChannelRequest(input *GetSmsChannelInput) GetSmsChannelRequest {
	op := &aws.Operation{
		Name:       opGetSmsChannel,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/channels/sms",
	}

	if input == nil {
		input = &GetSmsChannelInput{}
	}

	req := c.newRequest(op, input, &GetSmsChannelOutput{})
	return GetSmsChannelRequest{Request: req, Input: input, Copy: c.GetSmsChannelRequest}
}

// GetSmsChannelRequest is the request type for the
// GetSmsChannel API operation.
type GetSmsChannelRequest struct {
	*aws.Request
	Input *GetSmsChannelInput
	Copy  func(*GetSmsChannelInput) GetSmsChannelRequest
}

// Send marshals and sends the GetSmsChannel API request.
func (r GetSmsChannelRequest) Send(ctx context.Context) (*GetSmsChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSmsChannelResponse{
		GetSmsChannelOutput: r.Request.Data.(*GetSmsChannelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSmsChannelResponse is the response type for the
// GetSmsChannel API operation.
type GetSmsChannelResponse struct {
	*GetSmsChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSmsChannel request.
func (r *GetSmsChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
