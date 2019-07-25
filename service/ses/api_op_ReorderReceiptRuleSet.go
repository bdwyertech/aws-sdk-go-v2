// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to reorder the receipt rules within a receipt rule set.
// You use receipt rule sets to receive email with Amazon SES. For more information,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ReorderReceiptRuleSetRequest
type ReorderReceiptRuleSetInput struct {
	_ struct{} `type:"structure"`

	// A list of the specified receipt rule set's receipt rules in the order that
	// you want to put them.
	//
	// RuleNames is a required field
	RuleNames []string `type:"list" required:"true"`

	// The name of the receipt rule set to reorder.
	//
	// RuleSetName is a required field
	RuleSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ReorderReceiptRuleSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReorderReceiptRuleSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReorderReceiptRuleSetInput"}

	if s.RuleNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleNames"))
	}

	if s.RuleSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ReorderReceiptRuleSetResponse
type ReorderReceiptRuleSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ReorderReceiptRuleSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opReorderReceiptRuleSet = "ReorderReceiptRuleSet"

// ReorderReceiptRuleSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Reorders the receipt rules within a receipt rule set.
//
// All of the rules in the rule set must be represented in this request. That
// is, this API will return an error if the reorder request doesn't explicitly
// position all of the rules.
//
// For information about managing receipt rule sets, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rule-sets.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using ReorderReceiptRuleSetRequest.
//    req := client.ReorderReceiptRuleSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ReorderReceiptRuleSet
func (c *Client) ReorderReceiptRuleSetRequest(input *ReorderReceiptRuleSetInput) ReorderReceiptRuleSetRequest {
	op := &aws.Operation{
		Name:       opReorderReceiptRuleSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReorderReceiptRuleSetInput{}
	}

	req := c.newRequest(op, input, &ReorderReceiptRuleSetOutput{})
	return ReorderReceiptRuleSetRequest{Request: req, Input: input, Copy: c.ReorderReceiptRuleSetRequest}
}

// ReorderReceiptRuleSetRequest is the request type for the
// ReorderReceiptRuleSet API operation.
type ReorderReceiptRuleSetRequest struct {
	*aws.Request
	Input *ReorderReceiptRuleSetInput
	Copy  func(*ReorderReceiptRuleSetInput) ReorderReceiptRuleSetRequest
}

// Send marshals and sends the ReorderReceiptRuleSet API request.
func (r ReorderReceiptRuleSetRequest) Send(ctx context.Context) (*ReorderReceiptRuleSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReorderReceiptRuleSetResponse{
		ReorderReceiptRuleSetOutput: r.Request.Data.(*ReorderReceiptRuleSetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReorderReceiptRuleSetResponse is the response type for the
// ReorderReceiptRuleSet API operation.
type ReorderReceiptRuleSetResponse struct {
	*ReorderReceiptRuleSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReorderReceiptRuleSet request.
func (r *ReorderReceiptRuleSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
