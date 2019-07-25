// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSegmentsRequest
type GetSegmentsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s GetSegmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSegmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSegmentsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSegmentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "page-size", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSegmentsResponse
type GetSegmentsOutput struct {
	_ struct{} `type:"structure" payload:"SegmentsResponse"`

	// Provides information about all the segments that are associated with an application.
	//
	// SegmentsResponse is a required field
	SegmentsResponse *SegmentsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSegmentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSegmentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SegmentsResponse != nil {
		v := s.SegmentsResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SegmentsResponse", v, metadata)
	}
	return nil
}

const opGetSegments = "GetSegments"

// GetSegmentsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the configuration, dimension, and other settings
// for all the segments that are associated with an application.
//
//    // Example sending a request using GetSegmentsRequest.
//    req := client.GetSegmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSegments
func (c *Client) GetSegmentsRequest(input *GetSegmentsInput) GetSegmentsRequest {
	op := &aws.Operation{
		Name:       opGetSegments,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/segments",
	}

	if input == nil {
		input = &GetSegmentsInput{}
	}

	req := c.newRequest(op, input, &GetSegmentsOutput{})
	return GetSegmentsRequest{Request: req, Input: input, Copy: c.GetSegmentsRequest}
}

// GetSegmentsRequest is the request type for the
// GetSegments API operation.
type GetSegmentsRequest struct {
	*aws.Request
	Input *GetSegmentsInput
	Copy  func(*GetSegmentsInput) GetSegmentsRequest
}

// Send marshals and sends the GetSegments API request.
func (r GetSegmentsRequest) Send(ctx context.Context) (*GetSegmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSegmentsResponse{
		GetSegmentsOutput: r.Request.Data.(*GetSegmentsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSegmentsResponse is the response type for the
// GetSegments API operation.
type GetSegmentsResponse struct {
	*GetSegmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSegments request.
func (r *GetSegmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
