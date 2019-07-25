// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// SetVaultAccessPolicy input.
type SetVaultAccessPolicyInput struct {
	_ struct{} `type:"structure" payload:"Policy"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The vault access policy as a JSON string.
	Policy *VaultAccessPolicy `locationName:"policy" type:"structure"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s SetVaultAccessPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetVaultAccessPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetVaultAccessPolicyInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetVaultAccessPolicyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policy != nil {
		v := s.Policy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "policy", v, metadata)
	}
	return nil
}

type SetVaultAccessPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetVaultAccessPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetVaultAccessPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opSetVaultAccessPolicy = "SetVaultAccessPolicy"

// SetVaultAccessPolicyRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation configures an access policy for a vault and will overwrite
// an existing policy. To configure a vault access policy, send a PUT request
// to the access-policy subresource of the vault. An access policy is specific
// to a vault and is also called a vault subresource. You can set one access
// policy per vault and the policy can be up to 20 KB in size. For more information
// about vault access policies, see Amazon Glacier Access Control with Vault
// Access Policies (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html).
//
//    // Example sending a request using SetVaultAccessPolicyRequest.
//    req := client.SetVaultAccessPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) SetVaultAccessPolicyRequest(input *SetVaultAccessPolicyInput) SetVaultAccessPolicyRequest {
	op := &aws.Operation{
		Name:       opSetVaultAccessPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/access-policy",
	}

	if input == nil {
		input = &SetVaultAccessPolicyInput{}
	}

	req := c.newRequest(op, input, &SetVaultAccessPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetVaultAccessPolicyRequest{Request: req, Input: input, Copy: c.SetVaultAccessPolicyRequest}
}

// SetVaultAccessPolicyRequest is the request type for the
// SetVaultAccessPolicy API operation.
type SetVaultAccessPolicyRequest struct {
	*aws.Request
	Input *SetVaultAccessPolicyInput
	Copy  func(*SetVaultAccessPolicyInput) SetVaultAccessPolicyRequest
}

// Send marshals and sends the SetVaultAccessPolicy API request.
func (r SetVaultAccessPolicyRequest) Send(ctx context.Context) (*SetVaultAccessPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetVaultAccessPolicyResponse{
		SetVaultAccessPolicyOutput: r.Request.Data.(*SetVaultAccessPolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetVaultAccessPolicyResponse is the response type for the
// SetVaultAccessPolicy API operation.
type SetVaultAccessPolicyResponse struct {
	*SetVaultAccessPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetVaultAccessPolicy request.
func (r *SetVaultAccessPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
