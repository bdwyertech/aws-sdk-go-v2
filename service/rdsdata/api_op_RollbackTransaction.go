// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsdata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request parameters represent the input of a request to perform a rollback
// of a transaction.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/RollbackTransactionRequest
type RollbackTransactionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.
	//
	// ResourceArn is a required field
	ResourceArn *string `locationName:"resourceArn" type:"string" required:"true"`

	// The name or ARN of the secret that enables access to the DB cluster.
	//
	// SecretArn is a required field
	SecretArn *string `locationName:"secretArn" type:"string" required:"true"`

	// The identifier of the transaction to roll back.
	//
	// TransactionId is a required field
	TransactionId *string `locationName:"transactionId" type:"string" required:"true"`
}

// String returns the string representation
func (s RollbackTransactionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RollbackTransactionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RollbackTransactionInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if s.SecretArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretArn"))
	}

	if s.TransactionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransactionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RollbackTransactionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecretArn != nil {
		v := *s.SecretArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "secretArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TransactionId != nil {
		v := *s.TransactionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "transactionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The response elements represent the output of a request to perform a rollback
// of a transaction.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/RollbackTransactionResponse
type RollbackTransactionOutput struct {
	_ struct{} `type:"structure"`

	// The status of the rollback operation.
	TransactionStatus *string `locationName:"transactionStatus" type:"string"`
}

// String returns the string representation
func (s RollbackTransactionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RollbackTransactionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TransactionStatus != nil {
		v := *s.TransactionStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "transactionStatus", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opRollbackTransaction = "RollbackTransaction"

// RollbackTransactionRequest returns a request value for making API operation for
// AWS RDS DataService.
//
// Performs a rollback of a transaction. Rolling back a transaction cancels
// its changes.
//
//    // Example sending a request using RollbackTransactionRequest.
//    req := client.RollbackTransactionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/RollbackTransaction
func (c *Client) RollbackTransactionRequest(input *RollbackTransactionInput) RollbackTransactionRequest {
	op := &aws.Operation{
		Name:       opRollbackTransaction,
		HTTPMethod: "POST",
		HTTPPath:   "/RollbackTransaction",
	}

	if input == nil {
		input = &RollbackTransactionInput{}
	}

	req := c.newRequest(op, input, &RollbackTransactionOutput{})
	return RollbackTransactionRequest{Request: req, Input: input, Copy: c.RollbackTransactionRequest}
}

// RollbackTransactionRequest is the request type for the
// RollbackTransaction API operation.
type RollbackTransactionRequest struct {
	*aws.Request
	Input *RollbackTransactionInput
	Copy  func(*RollbackTransactionInput) RollbackTransactionRequest
}

// Send marshals and sends the RollbackTransaction API request.
func (r RollbackTransactionRequest) Send(ctx context.Context) (*RollbackTransactionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RollbackTransactionResponse{
		RollbackTransactionOutput: r.Request.Data.(*RollbackTransactionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RollbackTransactionResponse is the response type for the
// RollbackTransaction API operation.
type RollbackTransactionResponse struct {
	*RollbackTransactionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RollbackTransaction request.
func (r *RollbackTransactionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
