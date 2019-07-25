// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/UpdateTemplateRequest
type UpdateTemplateInput struct {
	_ struct{} `type:"structure"`

	// The content of the email, composed of a subject line, an HTML part, and a
	// text-only part.
	//
	// Template is a required field
	Template *Template `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTemplateInput"}

	if s.Template == nil {
		invalidParams.Add(aws.NewErrParamRequired("Template"))
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			invalidParams.AddNested("Template", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/UpdateTemplateResponse
type UpdateTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTemplate = "UpdateTemplate"

// UpdateTemplateRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Updates an email template. Email templates enable you to send personalized
// email to one or more destinations in a single API operation. For more information,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using UpdateTemplateRequest.
//    req := client.UpdateTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/UpdateTemplate
func (c *Client) UpdateTemplateRequest(input *UpdateTemplateInput) UpdateTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdateTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTemplateInput{}
	}

	req := c.newRequest(op, input, &UpdateTemplateOutput{})
	return UpdateTemplateRequest{Request: req, Input: input, Copy: c.UpdateTemplateRequest}
}

// UpdateTemplateRequest is the request type for the
// UpdateTemplate API operation.
type UpdateTemplateRequest struct {
	*aws.Request
	Input *UpdateTemplateInput
	Copy  func(*UpdateTemplateInput) UpdateTemplateRequest
}

// Send marshals and sends the UpdateTemplate API request.
func (r UpdateTemplateRequest) Send(ctx context.Context) (*UpdateTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTemplateResponse{
		UpdateTemplateOutput: r.Request.Data.(*UpdateTemplateOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTemplateResponse is the response type for the
// UpdateTemplate API operation.
type UpdateTemplateResponse struct {
	*UpdateTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTemplate request.
func (r *UpdateTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
