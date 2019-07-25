// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ioteventsdata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/ListDetectorsRequest
type ListDetectorsInput struct {
	_ struct{} `type:"structure"`

	// The name of the detector model whose detectors (instances) are listed.
	//
	// DetectorModelName is a required field
	DetectorModelName *string `location:"uri" locationName:"detectorModelName" min:"1" type:"string" required:"true"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// A filter that limits results to those detectors (instances) in the given
	// state.
	StateName *string `location:"querystring" locationName:"stateName" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDetectorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDetectorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDetectorsInput"}

	if s.DetectorModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorModelName"))
	}
	if s.DetectorModelName != nil && len(*s.DetectorModelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorModelName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.StateName != nil && len(*s.StateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDetectorsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DetectorModelName != nil {
		v := *s.DetectorModelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorModelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StateName != nil {
		v := *s.StateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "stateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/ListDetectorsResponse
type ListDetectorsOutput struct {
	_ struct{} `type:"structure"`

	// A list of summary information about the detectors (instances).
	DetectorSummaries []DetectorSummary `locationName:"detectorSummaries" type:"list"`

	// A token to retrieve the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDetectorsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDetectorsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DetectorSummaries != nil {
		v := s.DetectorSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "detectorSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDetectors = "ListDetectors"

// ListDetectorsRequest returns a request value for making API operation for
// AWS IoT Events Data.
//
// Lists detectors (the instances of a detector model).
//
//    // Example sending a request using ListDetectorsRequest.
//    req := client.ListDetectorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/ListDetectors
func (c *Client) ListDetectorsRequest(input *ListDetectorsInput) ListDetectorsRequest {
	op := &aws.Operation{
		Name:       opListDetectors,
		HTTPMethod: "GET",
		HTTPPath:   "/detectors/{detectorModelName}",
	}

	if input == nil {
		input = &ListDetectorsInput{}
	}

	req := c.newRequest(op, input, &ListDetectorsOutput{})
	return ListDetectorsRequest{Request: req, Input: input, Copy: c.ListDetectorsRequest}
}

// ListDetectorsRequest is the request type for the
// ListDetectors API operation.
type ListDetectorsRequest struct {
	*aws.Request
	Input *ListDetectorsInput
	Copy  func(*ListDetectorsInput) ListDetectorsRequest
}

// Send marshals and sends the ListDetectors API request.
func (r ListDetectorsRequest) Send(ctx context.Context) (*ListDetectorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDetectorsResponse{
		ListDetectorsOutput: r.Request.Data.(*ListDetectorsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDetectorsResponse is the response type for the
// ListDetectors API operation.
type ListDetectorsResponse struct {
	*ListDetectorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDetectors request.
func (r *ListDetectorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
