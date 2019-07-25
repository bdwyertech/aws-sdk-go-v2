// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ioteventsdata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/DescribeDetectorRequest
type DescribeDetectorInput struct {
	_ struct{} `type:"structure"`

	// The name of the detector model whose detectors (instances) you want information
	// about.
	//
	// DetectorModelName is a required field
	DetectorModelName *string `location:"uri" locationName:"detectorModelName" min:"1" type:"string" required:"true"`

	// A filter used to limit results to detectors (instances) created because of
	// the given key ID.
	KeyValue *string `location:"querystring" locationName:"keyValue" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDetectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDetectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDetectorInput"}

	if s.DetectorModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorModelName"))
	}
	if s.DetectorModelName != nil && len(*s.DetectorModelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorModelName", 1))
	}
	if s.KeyValue != nil && len(*s.KeyValue) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyValue", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDetectorInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DetectorModelName != nil {
		v := *s.DetectorModelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorModelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KeyValue != nil {
		v := *s.KeyValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "keyValue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/DescribeDetectorResponse
type DescribeDetectorOutput struct {
	_ struct{} `type:"structure"`

	// Information about the detector (instance).
	Detector *Detector `locationName:"detector" type:"structure"`
}

// String returns the string representation
func (s DescribeDetectorOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDetectorOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Detector != nil {
		v := s.Detector

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "detector", v, metadata)
	}
	return nil
}

const opDescribeDetector = "DescribeDetector"

// DescribeDetectorRequest returns a request value for making API operation for
// AWS IoT Events Data.
//
// Returns information about the specified detector (instance).
//
//    // Example sending a request using DescribeDetectorRequest.
//    req := client.DescribeDetectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/DescribeDetector
func (c *Client) DescribeDetectorRequest(input *DescribeDetectorInput) DescribeDetectorRequest {
	op := &aws.Operation{
		Name:       opDescribeDetector,
		HTTPMethod: "GET",
		HTTPPath:   "/detectors/{detectorModelName}/keyValues/",
	}

	if input == nil {
		input = &DescribeDetectorInput{}
	}

	req := c.newRequest(op, input, &DescribeDetectorOutput{})
	return DescribeDetectorRequest{Request: req, Input: input, Copy: c.DescribeDetectorRequest}
}

// DescribeDetectorRequest is the request type for the
// DescribeDetector API operation.
type DescribeDetectorRequest struct {
	*aws.Request
	Input *DescribeDetectorInput
	Copy  func(*DescribeDetectorInput) DescribeDetectorRequest
}

// Send marshals and sends the DescribeDetector API request.
func (r DescribeDetectorRequest) Send(ctx context.Context) (*DescribeDetectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDetectorResponse{
		DescribeDetectorOutput: r.Request.Data.(*DescribeDetectorOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDetectorResponse is the response type for the
// DescribeDetector API operation.
type DescribeDetectorResponse struct {
	*DescribeDetectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDetector request.
func (r *DescribeDetectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
