// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/CreateMembersRequest
type CreateMembersInput struct {
	_ struct{} `type:"structure"`

	// A list of account ID and email address pairs of the accounts to associate
	// with the Security Hub master account.
	AccountDetails []AccountDetails `type:"list"`
}

// String returns the string representation
func (s CreateMembersInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AccountDetails != nil {
		v := s.AccountDetails

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AccountDetails", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/CreateMembersResponse
type CreateMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of account ID and email address pairs of the AWS accounts that weren't
	// processed.
	UnprocessedAccounts []Result `type:"list"`
}

// String returns the string representation
func (s CreateMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UnprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opCreateMembers = "CreateMembers"

// CreateMembersRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Creates a member association in Security Hub between the specified accounts
// and the account used to make the request, which is the master account. To
// successfully create a member, you must use this action from an account that
// already has Security Hub enabled. You can use the EnableSecurityHub to enable
// Security Hub.
//
// After you use CreateMembers to create member account associations in Security
// Hub, you need to use the InviteMembers action, which invites the accounts
// to enable Security Hub and become member accounts in Security Hub. If the
// invitation is accepted by the account owner, the account becomes a member
// account in Security Hub, and a permission policy is added that permits the
// master account to view the findings generated in the member account. When
// Security Hub is enabled in the invited account, findings start being sent
// to both the member and master accounts.
//
// You can remove the association between the master and member accounts by
// using the DisassociateFromMasterAccount or DisassociateMembers operation.
//
//    // Example sending a request using CreateMembersRequest.
//    req := client.CreateMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/CreateMembers
func (c *Client) CreateMembersRequest(input *CreateMembersInput) CreateMembersRequest {
	op := &aws.Operation{
		Name:       opCreateMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/members",
	}

	if input == nil {
		input = &CreateMembersInput{}
	}

	req := c.newRequest(op, input, &CreateMembersOutput{})
	return CreateMembersRequest{Request: req, Input: input, Copy: c.CreateMembersRequest}
}

// CreateMembersRequest is the request type for the
// CreateMembers API operation.
type CreateMembersRequest struct {
	*aws.Request
	Input *CreateMembersInput
	Copy  func(*CreateMembersInput) CreateMembersRequest
}

// Send marshals and sends the CreateMembers API request.
func (r CreateMembersRequest) Send(ctx context.Context) (*CreateMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMembersResponse{
		CreateMembersOutput: r.Request.Data.(*CreateMembersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMembersResponse is the response type for the
// CreateMembers API operation.
type CreateMembersResponse struct {
	*CreateMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMembers request.
func (r *CreateMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
