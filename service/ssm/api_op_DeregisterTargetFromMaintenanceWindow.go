// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeregisterTargetFromMaintenanceWindowRequest
type DeregisterTargetFromMaintenanceWindowInput struct {
	_ struct{} `type:"structure"`

	// The system checks if the target is being referenced by a task. If the target
	// is being referenced, the system returns an error and does not deregister
	// the target from the maintenance window.
	Safe *bool `type:"boolean"`

	// The ID of the maintenance window the target should be removed from.
	//
	// WindowId is a required field
	WindowId *string `min:"20" type:"string" required:"true"`

	// The ID of the target definition to remove.
	//
	// WindowTargetId is a required field
	WindowTargetId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterTargetFromMaintenanceWindowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterTargetFromMaintenanceWindowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterTargetFromMaintenanceWindowInput"}

	if s.WindowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowId"))
	}
	if s.WindowId != nil && len(*s.WindowId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowId", 20))
	}

	if s.WindowTargetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowTargetId"))
	}
	if s.WindowTargetId != nil && len(*s.WindowTargetId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowTargetId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeregisterTargetFromMaintenanceWindowResult
type DeregisterTargetFromMaintenanceWindowOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the maintenance window the target was removed from.
	WindowId *string `min:"20" type:"string"`

	// The ID of the removed target definition.
	WindowTargetId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s DeregisterTargetFromMaintenanceWindowOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterTargetFromMaintenanceWindow = "DeregisterTargetFromMaintenanceWindow"

// DeregisterTargetFromMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Removes a target from a maintenance window.
//
//    // Example sending a request using DeregisterTargetFromMaintenanceWindowRequest.
//    req := client.DeregisterTargetFromMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeregisterTargetFromMaintenanceWindow
func (c *Client) DeregisterTargetFromMaintenanceWindowRequest(input *DeregisterTargetFromMaintenanceWindowInput) DeregisterTargetFromMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opDeregisterTargetFromMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterTargetFromMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &DeregisterTargetFromMaintenanceWindowOutput{})
	return DeregisterTargetFromMaintenanceWindowRequest{Request: req, Input: input, Copy: c.DeregisterTargetFromMaintenanceWindowRequest}
}

// DeregisterTargetFromMaintenanceWindowRequest is the request type for the
// DeregisterTargetFromMaintenanceWindow API operation.
type DeregisterTargetFromMaintenanceWindowRequest struct {
	*aws.Request
	Input *DeregisterTargetFromMaintenanceWindowInput
	Copy  func(*DeregisterTargetFromMaintenanceWindowInput) DeregisterTargetFromMaintenanceWindowRequest
}

// Send marshals and sends the DeregisterTargetFromMaintenanceWindow API request.
func (r DeregisterTargetFromMaintenanceWindowRequest) Send(ctx context.Context) (*DeregisterTargetFromMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterTargetFromMaintenanceWindowResponse{
		DeregisterTargetFromMaintenanceWindowOutput: r.Request.Data.(*DeregisterTargetFromMaintenanceWindowOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterTargetFromMaintenanceWindowResponse is the response type for the
// DeregisterTargetFromMaintenanceWindow API operation.
type DeregisterTargetFromMaintenanceWindowResponse struct {
	*DeregisterTargetFromMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterTargetFromMaintenanceWindow request.
func (r *DeregisterTargetFromMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
