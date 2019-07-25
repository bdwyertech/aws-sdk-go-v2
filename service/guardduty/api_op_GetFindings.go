// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetFindingsRequest
type GetFindingsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the detector that specifies the GuardDuty service whose findings
	// you want to retrieve.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// IDs of the findings that you want to retrieve.
	//
	// FindingIds is a required field
	FindingIds []string `locationName:"findingIds" type:"list" required:"true"`

	// Represents the criteria used for sorting findings.
	SortCriteria *SortCriteria `locationName:"sortCriteria" type:"structure"`
}

// String returns the string representation
func (s GetFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFindingsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.FindingIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("FindingIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.FindingIds != nil {
		v := s.FindingIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "findingIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SortCriteria != nil {
		v := s.SortCriteria

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "sortCriteria", v, metadata)
	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetFindingsResponse
type GetFindingsOutput struct {
	_ struct{} `type:"structure"`

	// A list of findings.
	//
	// Findings is a required field
	Findings []Finding `locationName:"findings" type:"list" required:"true"`
}

// String returns the string representation
func (s GetFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Findings != nil {
		v := s.Findings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "findings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetFindings = "GetFindings"

// GetFindingsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Describes Amazon GuardDuty findings specified by finding IDs.
//
//    // Example sending a request using GetFindingsRequest.
//    req := client.GetFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetFindings
func (c *Client) GetFindingsRequest(input *GetFindingsInput) GetFindingsRequest {
	op := &aws.Operation{
		Name:       opGetFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/findings/get",
	}

	if input == nil {
		input = &GetFindingsInput{}
	}

	req := c.newRequest(op, input, &GetFindingsOutput{})
	return GetFindingsRequest{Request: req, Input: input, Copy: c.GetFindingsRequest}
}

// GetFindingsRequest is the request type for the
// GetFindings API operation.
type GetFindingsRequest struct {
	*aws.Request
	Input *GetFindingsInput
	Copy  func(*GetFindingsInput) GetFindingsRequest
}

// Send marshals and sends the GetFindings API request.
func (r GetFindingsRequest) Send(ctx context.Context) (*GetFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFindingsResponse{
		GetFindingsOutput: r.Request.Data.(*GetFindingsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFindingsResponse is the response type for the
// GetFindings API operation.
type GetFindingsResponse struct {
	*GetFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFindings request.
func (r *GetFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
