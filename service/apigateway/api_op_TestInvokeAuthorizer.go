// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Make a request to simulate the execution of an Authorizer.
type TestInvokeAuthorizerInput struct {
	_ struct{} `type:"structure"`

	// [Optional] A key-value map of additional context variables.
	AdditionalContext map[string]string `locationName:"additionalContext" type:"map"`

	// [Required] Specifies a test invoke authorizer request's Authorizer ID.
	//
	// AuthorizerId is a required field
	AuthorizerId *string `location:"uri" locationName:"authorizer_id" type:"string" required:"true"`

	// [Optional] The simulated request body of an incoming invocation request.
	Body *string `locationName:"body" type:"string"`

	// [Required] A key-value map of headers to simulate an incoming invocation
	// request. This is where the incoming authorization token, or identity source,
	// should be specified.
	Headers map[string]string `locationName:"headers" type:"map"`

	// [Optional] The headers as a map from string to list of values to simulate
	// an incoming invocation request. This is where the incoming authorization
	// token, or identity source, may be specified.
	MultiValueHeaders map[string][]string `locationName:"multiValueHeaders" type:"map"`

	// [Optional] The URI path, including query string, of the simulated invocation
	// request. Use this to specify path parameters and query string parameters.
	PathWithQueryString *string `locationName:"pathWithQueryString" type:"string"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// A key-value map of stage variables to simulate an invocation on a deployed
	// Stage.
	StageVariables map[string]string `locationName:"stageVariables" type:"map"`
}

// String returns the string representation
func (s TestInvokeAuthorizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestInvokeAuthorizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestInvokeAuthorizerInput"}

	if s.AuthorizerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizerId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TestInvokeAuthorizerInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AdditionalContext != nil {
		v := s.AdditionalContext

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "additionalContext", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Body != nil {
		v := *s.Body

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "body", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Headers != nil {
		v := s.Headers

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "headers", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.MultiValueHeaders != nil {
		v := s.MultiValueHeaders

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "multiValueHeaders", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ls1 := ms0.List(k1)
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v2)})
			}
			ls1.End()
		}
		ms0.End()

	}
	if s.PathWithQueryString != nil {
		v := *s.PathWithQueryString

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pathWithQueryString", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageVariables != nil {
		v := s.StageVariables

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "stageVariables", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.AuthorizerId != nil {
		v := *s.AuthorizerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "authorizer_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents the response of the test invoke request for a custom Authorizer
type TestInvokeAuthorizerOutput struct {
	_ struct{} `type:"structure"`

	Authorization map[string][]string `locationName:"authorization" type:"map"`

	// The open identity claims (https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims),
	// with any supported custom attributes, returned from the Cognito Your User
	// Pool configured for the API.
	Claims map[string]string `locationName:"claims" type:"map"`

	// The HTTP status code that the client would have received. Value is 0 if the
	// authorizer succeeded.
	ClientStatus *int64 `locationName:"clientStatus" type:"integer"`

	// The execution latency of the test authorizer request.
	Latency *int64 `locationName:"latency" type:"long"`

	// The API Gateway execution log for the test authorizer request.
	Log *string `locationName:"log" type:"string"`

	// The JSON policy document returned by the Authorizer
	Policy *string `locationName:"policy" type:"string"`

	// The principal identity returned by the Authorizer
	PrincipalId *string `locationName:"principalId" type:"string"`
}

// String returns the string representation
func (s TestInvokeAuthorizerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TestInvokeAuthorizerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Authorization != nil {
		v := s.Authorization

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "authorization", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ls1 := ms0.List(k1)
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v2)})
			}
			ls1.End()
		}
		ms0.End()

	}
	if s.Claims != nil {
		v := s.Claims

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "claims", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ClientStatus != nil {
		v := *s.ClientStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientStatus", protocol.Int64Value(v), metadata)
	}
	if s.Latency != nil {
		v := *s.Latency

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "latency", protocol.Int64Value(v), metadata)
	}
	if s.Log != nil {
		v := *s.Log

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "log", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PrincipalId != nil {
		v := *s.PrincipalId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "principalId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opTestInvokeAuthorizer = "TestInvokeAuthorizer"

// TestInvokeAuthorizerRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Simulate the execution of an Authorizer in your RestApi with headers, parameters,
// and an incoming request body.
//
// Use Lambda Function as Authorizer (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html)
// Use Cognito User Pool as Authorizer (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html)
//
//    // Example sending a request using TestInvokeAuthorizerRequest.
//    req := client.TestInvokeAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) TestInvokeAuthorizerRequest(input *TestInvokeAuthorizerInput) TestInvokeAuthorizerRequest {
	op := &aws.Operation{
		Name:       opTestInvokeAuthorizer,
		HTTPMethod: "POST",
		HTTPPath:   "/restapis/{restapi_id}/authorizers/{authorizer_id}",
	}

	if input == nil {
		input = &TestInvokeAuthorizerInput{}
	}

	req := c.newRequest(op, input, &TestInvokeAuthorizerOutput{})
	return TestInvokeAuthorizerRequest{Request: req, Input: input, Copy: c.TestInvokeAuthorizerRequest}
}

// TestInvokeAuthorizerRequest is the request type for the
// TestInvokeAuthorizer API operation.
type TestInvokeAuthorizerRequest struct {
	*aws.Request
	Input *TestInvokeAuthorizerInput
	Copy  func(*TestInvokeAuthorizerInput) TestInvokeAuthorizerRequest
}

// Send marshals and sends the TestInvokeAuthorizer API request.
func (r TestInvokeAuthorizerRequest) Send(ctx context.Context) (*TestInvokeAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestInvokeAuthorizerResponse{
		TestInvokeAuthorizerOutput: r.Request.Data.(*TestInvokeAuthorizerOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestInvokeAuthorizerResponse is the response type for the
// TestInvokeAuthorizer API operation.
type TestInvokeAuthorizerResponse struct {
	*TestInvokeAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestInvokeAuthorizer request.
func (r *TestInvokeAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
