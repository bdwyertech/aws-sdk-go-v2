// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignActivitiesRequest
type GetCampaignActivitiesInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// CampaignId is a required field
	CampaignId *string `location:"uri" locationName:"campaign-id" type:"string" required:"true"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s GetCampaignActivitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCampaignActivitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCampaignActivitiesInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.CampaignId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CampaignId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCampaignActivitiesInput) MarshalFields(e protocol.FieldEncoder) error {
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignActivitiesResponse
type GetCampaignActivitiesOutput struct {
	_ struct{} `type:"structure" payload:"ActivitiesResponse"`

	// Provides information about the activities that were performed by a campaign.
	//
	// ActivitiesResponse is a required field
	ActivitiesResponse *ActivitiesResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetCampaignActivitiesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCampaignActivitiesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ActivitiesResponse != nil {
		v := s.ActivitiesResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ActivitiesResponse", v, metadata)
	}
	return nil
}

const opGetCampaignActivities = "GetCampaignActivities"

// GetCampaignActivitiesRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the activity performed by a campaign.
//
//    // Example sending a request using GetCampaignActivitiesRequest.
//    req := client.GetCampaignActivitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignActivities
func (c *Client) GetCampaignActivitiesRequest(input *GetCampaignActivitiesInput) GetCampaignActivitiesRequest {
	op := &aws.Operation{
		Name:       opGetCampaignActivities,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/campaigns/{campaign-id}/activities",
	}

	if input == nil {
		input = &GetCampaignActivitiesInput{}
	}

	req := c.newRequest(op, input, &GetCampaignActivitiesOutput{})
	return GetCampaignActivitiesRequest{Request: req, Input: input, Copy: c.GetCampaignActivitiesRequest}
}

// GetCampaignActivitiesRequest is the request type for the
// GetCampaignActivities API operation.
type GetCampaignActivitiesRequest struct {
	*aws.Request
	Input *GetCampaignActivitiesInput
	Copy  func(*GetCampaignActivitiesInput) GetCampaignActivitiesRequest
}

// Send marshals and sends the GetCampaignActivities API request.
func (r GetCampaignActivitiesRequest) Send(ctx context.Context) (*GetCampaignActivitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCampaignActivitiesResponse{
		GetCampaignActivitiesOutput: r.Request.Data.(*GetCampaignActivitiesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCampaignActivitiesResponse is the response type for the
// GetCampaignActivities API operation.
type GetCampaignActivitiesResponse struct {
	*GetCampaignActivitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCampaignActivities request.
func (r *GetCampaignActivitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
