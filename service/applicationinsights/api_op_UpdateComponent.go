// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/UpdateComponentRequest
type UpdateComponentInput struct {
	_ struct{} `type:"structure"`

	// The name of the component.
	//
	// ComponentName is a required field
	ComponentName *string `type:"string" required:"true"`

	// The new name of the component.
	NewComponentName *string `type:"string"`

	// The name of the resource group.
	//
	// ResourceGroupName is a required field
	ResourceGroupName *string `type:"string" required:"true"`

	// The list of resource ARNs that belong to the component.
	ResourceList []string `type:"list"`
}

// String returns the string representation
func (s UpdateComponentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateComponentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateComponentInput"}

	if s.ComponentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComponentName"))
	}

	if s.ResourceGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/UpdateComponentResponse
type UpdateComponentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateComponentOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateComponent = "UpdateComponent"

// UpdateComponentRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Updates the custom component name and/or the list of resources that make
// up the component.
//
//    // Example sending a request using UpdateComponentRequest.
//    req := client.UpdateComponentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/UpdateComponent
func (c *Client) UpdateComponentRequest(input *UpdateComponentInput) UpdateComponentRequest {
	op := &aws.Operation{
		Name:       opUpdateComponent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateComponentInput{}
	}

	req := c.newRequest(op, input, &UpdateComponentOutput{})
	return UpdateComponentRequest{Request: req, Input: input, Copy: c.UpdateComponentRequest}
}

// UpdateComponentRequest is the request type for the
// UpdateComponent API operation.
type UpdateComponentRequest struct {
	*aws.Request
	Input *UpdateComponentInput
	Copy  func(*UpdateComponentInput) UpdateComponentRequest
}

// Send marshals and sends the UpdateComponent API request.
func (r UpdateComponentRequest) Send(ctx context.Context) (*UpdateComponentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateComponentResponse{
		UpdateComponentOutput: r.Request.Data.(*UpdateComponentOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateComponentResponse is the response type for the
// UpdateComponent API operation.
type UpdateComponentResponse struct {
	*UpdateComponentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateComponent request.
func (r *UpdateComponentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
