// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetExportJobRequest
type GetExportJobInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// JobId is a required field
	JobId *string `location:"uri" locationName:"job-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetExportJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetExportJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetExportJobInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetExportJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "job-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetExportJobResponse
type GetExportJobOutput struct {
	_ struct{} `type:"structure" payload:"ExportJobResponse"`

	// Provides information about the status and settings of a job that exports
	// endpoint definitions to a file. The file can be added directly to an Amazon
	// Simple Storage Service (Amazon S3) bucket by using the Amazon Pinpoint API
	// or downloaded directly to a computer by using the Amazon Pinpoint console.
	//
	// ExportJobResponse is a required field
	ExportJobResponse *ExportJobResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetExportJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetExportJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExportJobResponse != nil {
		v := s.ExportJobResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ExportJobResponse", v, metadata)
	}
	return nil
}

const opGetExportJob = "GetExportJob"

// GetExportJobRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of a specific export
// job for an application.
//
//    // Example sending a request using GetExportJobRequest.
//    req := client.GetExportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetExportJob
func (c *Client) GetExportJobRequest(input *GetExportJobInput) GetExportJobRequest {
	op := &aws.Operation{
		Name:       opGetExportJob,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/jobs/export/{job-id}",
	}

	if input == nil {
		input = &GetExportJobInput{}
	}

	req := c.newRequest(op, input, &GetExportJobOutput{})
	return GetExportJobRequest{Request: req, Input: input, Copy: c.GetExportJobRequest}
}

// GetExportJobRequest is the request type for the
// GetExportJob API operation.
type GetExportJobRequest struct {
	*aws.Request
	Input *GetExportJobInput
	Copy  func(*GetExportJobInput) GetExportJobRequest
}

// Send marshals and sends the GetExportJob API request.
func (r GetExportJobRequest) Send(ctx context.Context) (*GetExportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetExportJobResponse{
		GetExportJobOutput: r.Request.Data.(*GetExportJobOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetExportJobResponse is the response type for the
// GetExportJob API operation.
type GetExportJobResponse struct {
	*GetExportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetExportJob request.
func (r *GetExportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
