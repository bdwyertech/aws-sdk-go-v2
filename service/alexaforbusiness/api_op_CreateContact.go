// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateContactRequest
type CreateContactInput struct {
	_ struct{} `type:"structure"`

	// A unique, user-specified identifier for this request that ensures idempotency.
	ClientRequestToken *string `min:"10" type:"string" idempotencyToken:"true"`

	// The name of the contact to display on the console.
	DisplayName *string `min:"1" type:"string"`

	// The first name of the contact that is used to call the contact on the device.
	//
	// FirstName is a required field
	FirstName *string `min:"1" type:"string" required:"true"`

	// The last name of the contact that is used to call the contact on the device.
	LastName *string `min:"1" type:"string"`

	// The phone number of the contact in E.164 format. The phone number type defaults
	// to WORK. You can specify PhoneNumber or PhoneNumbers. We recommend that you
	// use PhoneNumbers, which lets you specify the phone number type and multiple
	// numbers.
	PhoneNumber *string `type:"string"`

	// The list of phone numbers for the contact.
	PhoneNumbers []PhoneNumber `type:"list"`

	// The list of SIP addresses for the contact.
	SipAddresses []SipAddress `type:"list"`
}

// String returns the string representation
func (s CreateContactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateContactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateContactInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 10))
	}
	if s.DisplayName != nil && len(*s.DisplayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayName", 1))
	}

	if s.FirstName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FirstName"))
	}
	if s.FirstName != nil && len(*s.FirstName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FirstName", 1))
	}
	if s.LastName != nil && len(*s.LastName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LastName", 1))
	}
	if s.PhoneNumbers != nil {
		for i, v := range s.PhoneNumbers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PhoneNumbers", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SipAddresses != nil {
		for i, v := range s.SipAddresses {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SipAddresses", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateContactResponse
type CreateContactOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the newly created address book.
	ContactArn *string `type:"string"`
}

// String returns the string representation
func (s CreateContactOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateContact = "CreateContact"

// CreateContactRequest returns a request value for making API operation for
// Alexa For Business.
//
// Creates a contact with the specified details.
//
//    // Example sending a request using CreateContactRequest.
//    req := client.CreateContactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateContact
func (c *Client) CreateContactRequest(input *CreateContactInput) CreateContactRequest {
	op := &aws.Operation{
		Name:       opCreateContact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateContactInput{}
	}

	req := c.newRequest(op, input, &CreateContactOutput{})
	return CreateContactRequest{Request: req, Input: input, Copy: c.CreateContactRequest}
}

// CreateContactRequest is the request type for the
// CreateContact API operation.
type CreateContactRequest struct {
	*aws.Request
	Input *CreateContactInput
	Copy  func(*CreateContactInput) CreateContactRequest
}

// Send marshals and sends the CreateContact API request.
func (r CreateContactRequest) Send(ctx context.Context) (*CreateContactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateContactResponse{
		CreateContactOutput: r.Request.Data.(*CreateContactOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateContactResponse is the response type for the
// CreateContact API operation.
type CreateContactResponse struct {
	*CreateContactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateContact request.
func (r *CreateContactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
