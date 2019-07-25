// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateWorkflowRequest
type CreateWorkflowInput struct {
	_ struct{} `type:"structure"`

	// A collection of properties to be used as part of each execution of the workflow.
	DefaultRunProperties map[string]string `type:"map"`

	// A description of the workflow.
	Description *string `type:"string"`

	// The name to be assigned to the workflow. It should be unique within your
	// account.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The tags to be used with this workflow.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s CreateWorkflowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateWorkflowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateWorkflowInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateWorkflowResponse
type CreateWorkflowOutput struct {
	_ struct{} `type:"structure"`

	// The name of the workflow which was provided as part of the request.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateWorkflowOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateWorkflow = "CreateWorkflow"

// CreateWorkflowRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a new workflow.
//
//    // Example sending a request using CreateWorkflowRequest.
//    req := client.CreateWorkflowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateWorkflow
func (c *Client) CreateWorkflowRequest(input *CreateWorkflowInput) CreateWorkflowRequest {
	op := &aws.Operation{
		Name:       opCreateWorkflow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateWorkflowInput{}
	}

	req := c.newRequest(op, input, &CreateWorkflowOutput{})
	return CreateWorkflowRequest{Request: req, Input: input, Copy: c.CreateWorkflowRequest}
}

// CreateWorkflowRequest is the request type for the
// CreateWorkflow API operation.
type CreateWorkflowRequest struct {
	*aws.Request
	Input *CreateWorkflowInput
	Copy  func(*CreateWorkflowInput) CreateWorkflowRequest
}

// Send marshals and sends the CreateWorkflow API request.
func (r CreateWorkflowRequest) Send(ctx context.Context) (*CreateWorkflowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateWorkflowResponse{
		CreateWorkflowOutput: r.Request.Data.(*CreateWorkflowOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateWorkflowResponse is the response type for the
// CreateWorkflow API operation.
type CreateWorkflowResponse struct {
	*CreateWorkflowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateWorkflow request.
func (r *CreateWorkflowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
