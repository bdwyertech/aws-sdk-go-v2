// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PurchaseProvisionedCapacityInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID of the account that owns the vault. You can either specify
	// an AWS account ID or optionally a single '-' (hyphen), in which case Amazon
	// S3 Glacier uses the AWS account ID associated with the credentials used to
	// sign the request. If you use an account ID, don't include any hyphens ('-')
	// in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`
}

// String returns the string representation
func (s PurchaseProvisionedCapacityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurchaseProvisionedCapacityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PurchaseProvisionedCapacityInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PurchaseProvisionedCapacityInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PurchaseProvisionedCapacityOutput struct {
	_ struct{} `type:"structure"`

	// The ID that identifies the provisioned capacity unit.
	CapacityId *string `location:"header" locationName:"x-amz-capacity-id" type:"string"`
}

// String returns the string representation
func (s PurchaseProvisionedCapacityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PurchaseProvisionedCapacityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CapacityId != nil {
		v := *s.CapacityId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-capacity-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opPurchaseProvisionedCapacity = "PurchaseProvisionedCapacity"

// PurchaseProvisionedCapacityRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation purchases a provisioned capacity unit for an AWS account.
//
//    // Example sending a request using PurchaseProvisionedCapacityRequest.
//    req := client.PurchaseProvisionedCapacityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) PurchaseProvisionedCapacityRequest(input *PurchaseProvisionedCapacityInput) PurchaseProvisionedCapacityRequest {
	op := &aws.Operation{
		Name:       opPurchaseProvisionedCapacity,
		HTTPMethod: "POST",
		HTTPPath:   "/{accountId}/provisioned-capacity",
	}

	if input == nil {
		input = &PurchaseProvisionedCapacityInput{}
	}

	req := c.newRequest(op, input, &PurchaseProvisionedCapacityOutput{})
	return PurchaseProvisionedCapacityRequest{Request: req, Input: input, Copy: c.PurchaseProvisionedCapacityRequest}
}

// PurchaseProvisionedCapacityRequest is the request type for the
// PurchaseProvisionedCapacity API operation.
type PurchaseProvisionedCapacityRequest struct {
	*aws.Request
	Input *PurchaseProvisionedCapacityInput
	Copy  func(*PurchaseProvisionedCapacityInput) PurchaseProvisionedCapacityRequest
}

// Send marshals and sends the PurchaseProvisionedCapacity API request.
func (r PurchaseProvisionedCapacityRequest) Send(ctx context.Context) (*PurchaseProvisionedCapacityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseProvisionedCapacityResponse{
		PurchaseProvisionedCapacityOutput: r.Request.Data.(*PurchaseProvisionedCapacityOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseProvisionedCapacityResponse is the response type for the
// PurchaseProvisionedCapacity API operation.
type PurchaseProvisionedCapacityResponse struct {
	*PurchaseProvisionedCapacityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseProvisionedCapacity request.
func (r *PurchaseProvisionedCapacityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
