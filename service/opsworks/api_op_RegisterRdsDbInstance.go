// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterRdsDbInstanceRequest
type RegisterRdsDbInstanceInput struct {
	_ struct{} `type:"structure"`

	// The database password.
	//
	// DbPassword is a required field
	DbPassword *string `type:"string" required:"true"`

	// The database's master user name.
	//
	// DbUser is a required field
	DbUser *string `type:"string" required:"true"`

	// The Amazon RDS instance's ARN.
	//
	// RdsDbInstanceArn is a required field
	RdsDbInstanceArn *string `type:"string" required:"true"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterRdsDbInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterRdsDbInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterRdsDbInstanceInput"}

	if s.DbPassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("DbPassword"))
	}

	if s.DbUser == nil {
		invalidParams.Add(aws.NewErrParamRequired("DbUser"))
	}

	if s.RdsDbInstanceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RdsDbInstanceArn"))
	}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterRdsDbInstanceOutput
type RegisterRdsDbInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterRdsDbInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterRdsDbInstance = "RegisterRdsDbInstance"

// RegisterRdsDbInstanceRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Registers an Amazon RDS instance with a stack.
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using RegisterRdsDbInstanceRequest.
//    req := client.RegisterRdsDbInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterRdsDbInstance
func (c *Client) RegisterRdsDbInstanceRequest(input *RegisterRdsDbInstanceInput) RegisterRdsDbInstanceRequest {
	op := &aws.Operation{
		Name:       opRegisterRdsDbInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterRdsDbInstanceInput{}
	}

	req := c.newRequest(op, input, &RegisterRdsDbInstanceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RegisterRdsDbInstanceRequest{Request: req, Input: input, Copy: c.RegisterRdsDbInstanceRequest}
}

// RegisterRdsDbInstanceRequest is the request type for the
// RegisterRdsDbInstance API operation.
type RegisterRdsDbInstanceRequest struct {
	*aws.Request
	Input *RegisterRdsDbInstanceInput
	Copy  func(*RegisterRdsDbInstanceInput) RegisterRdsDbInstanceRequest
}

// Send marshals and sends the RegisterRdsDbInstance API request.
func (r RegisterRdsDbInstanceRequest) Send(ctx context.Context) (*RegisterRdsDbInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterRdsDbInstanceResponse{
		RegisterRdsDbInstanceOutput: r.Request.Data.(*RegisterRdsDbInstanceOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterRdsDbInstanceResponse is the response type for the
// RegisterRdsDbInstance API operation.
type RegisterRdsDbInstanceResponse struct {
	*RegisterRdsDbInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterRdsDbInstance request.
func (r *RegisterRdsDbInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
