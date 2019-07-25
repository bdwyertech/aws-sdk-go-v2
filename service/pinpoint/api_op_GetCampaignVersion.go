// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignVersionRequest
type GetCampaignVersionInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// CampaignId is a required field
	CampaignId *string `location:"uri" locationName:"campaign-id" type:"string" required:"true"`

	// Version is a required field
	Version *string `location:"uri" locationName:"version" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCampaignVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCampaignVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCampaignVersionInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.CampaignId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CampaignId"))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCampaignVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CampaignId != nil {
		v := *s.CampaignId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "campaign-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignVersionResponse
type GetCampaignVersionOutput struct {
	_ struct{} `type:"structure" payload:"CampaignResponse"`

	// Provides information about the status, configuration, and other settings
	// for a campaign.
	//
	// CampaignResponse is a required field
	CampaignResponse *CampaignResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetCampaignVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCampaignVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CampaignResponse != nil {
		v := s.CampaignResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CampaignResponse", v, metadata)
	}
	return nil
}

const opGetCampaignVersion = "GetCampaignVersion"

// GetCampaignVersionRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status, configuration, and other settings
// for a specific version of a campaign.
//
//    // Example sending a request using GetCampaignVersionRequest.
//    req := client.GetCampaignVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignVersion
func (c *Client) GetCampaignVersionRequest(input *GetCampaignVersionInput) GetCampaignVersionRequest {
	op := &aws.Operation{
		Name:       opGetCampaignVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/campaigns/{campaign-id}/versions/{version}",
	}

	if input == nil {
		input = &GetCampaignVersionInput{}
	}

	req := c.newRequest(op, input, &GetCampaignVersionOutput{})
	return GetCampaignVersionRequest{Request: req, Input: input, Copy: c.GetCampaignVersionRequest}
}

// GetCampaignVersionRequest is the request type for the
// GetCampaignVersion API operation.
type GetCampaignVersionRequest struct {
	*aws.Request
	Input *GetCampaignVersionInput
	Copy  func(*GetCampaignVersionInput) GetCampaignVersionRequest
}

// Send marshals and sends the GetCampaignVersion API request.
func (r GetCampaignVersionRequest) Send(ctx context.Context) (*GetCampaignVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCampaignVersionResponse{
		GetCampaignVersionOutput: r.Request.Data.(*GetCampaignVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCampaignVersionResponse is the response type for the
// GetCampaignVersion API operation.
type GetCampaignVersionResponse struct {
	*GetCampaignVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCampaignVersion request.
func (r *GetCampaignVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
