// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchGetWorkflowsRequest
type BatchGetWorkflowsInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether to include a graph when returning the workflow resource
	// metadata.
	IncludeGraph *bool `type:"boolean"`

	// A list of workflow names, which may be the names returned from the ListWorkflows
	// operation.
	//
	// Names is a required field
	Names []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetWorkflowsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetWorkflowsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetWorkflowsInput"}

	if s.Names == nil {
		invalidParams.Add(aws.NewErrParamRequired("Names"))
	}
	if s.Names != nil && len(s.Names) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Names", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchGetWorkflowsResponse
type BatchGetWorkflowsOutput struct {
	_ struct{} `type:"structure"`

	// A list of names of workflows not found.
	MissingWorkflows []string `min:"1" type:"list"`

	// A list of workflow resource metadata.
	Workflows []Workflow `min:"1" type:"list"`
}

// String returns the string representation
func (s BatchGetWorkflowsOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetWorkflows = "BatchGetWorkflows"

// BatchGetWorkflowsRequest returns a request value for making API operation for
// AWS Glue.
//
// Returns a list of resource metadata for a given list of workflow names. After
// calling the ListWorkflows operation, you can call this operation to access
// the data to which you have been granted permissions. This operation supports
// all IAM permissions, including permission conditions that uses tags.
//
//    // Example sending a request using BatchGetWorkflowsRequest.
//    req := client.BatchGetWorkflowsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchGetWorkflows
func (c *Client) BatchGetWorkflowsRequest(input *BatchGetWorkflowsInput) BatchGetWorkflowsRequest {
	op := &aws.Operation{
		Name:       opBatchGetWorkflows,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetWorkflowsInput{}
	}

	req := c.newRequest(op, input, &BatchGetWorkflowsOutput{})
	return BatchGetWorkflowsRequest{Request: req, Input: input, Copy: c.BatchGetWorkflowsRequest}
}

// BatchGetWorkflowsRequest is the request type for the
// BatchGetWorkflows API operation.
type BatchGetWorkflowsRequest struct {
	*aws.Request
	Input *BatchGetWorkflowsInput
	Copy  func(*BatchGetWorkflowsInput) BatchGetWorkflowsRequest
}

// Send marshals and sends the BatchGetWorkflows API request.
func (r BatchGetWorkflowsRequest) Send(ctx context.Context) (*BatchGetWorkflowsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetWorkflowsResponse{
		BatchGetWorkflowsOutput: r.Request.Data.(*BatchGetWorkflowsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetWorkflowsResponse is the response type for the
// BatchGetWorkflows API operation.
type BatchGetWorkflowsResponse struct {
	*BatchGetWorkflowsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetWorkflows request.
func (r *BatchGetWorkflowsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
