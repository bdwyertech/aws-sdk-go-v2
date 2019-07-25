// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A POST request to import an API to API Gateway using an input of an API definition
// file.
type ImportRestApiInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// [Required] The POST request body containing external API definitions. Currently,
	// only OpenAPI definition JSON/YAML files are supported. The maximum size of
	// the API definition file is 2MB.
	//
	// Body is a required field
	Body []byte `locationName:"body" type:"blob" required:"true"`

	// A query parameter to indicate whether to rollback the API creation (true)
	// or not (false) when a warning is encountered. The default value is false.
	FailOnWarnings *bool `location:"querystring" locationName:"failonwarnings" type:"boolean"`

	// A key-value map of context-specific query string parameters specifying the
	// behavior of different API importing operations. The following shows operation-specific
	// parameters and their supported values.
	//
	// To exclude DocumentationParts from the import, set parameters as ignore=documentation.
	//
	// To configure the endpoint type, set parameters as endpointConfigurationTypes=EDGE,
	// endpointConfigurationTypes=REGIONAL, or endpointConfigurationTypes=PRIVATE.
	// The default endpoint type is EDGE.
	//
	// To handle imported basepath, set parameters as basepath=ignore, basepath=prepend
	// or basepath=split.
	//
	// For example, the AWS CLI command to exclude documentation from the imported
	// API is:
	//
	//    aws apigateway import-rest-api --parameters ignore=documentation --body
	//    'file:///path/to/imported-api-body.json'
	//
	// The AWS CLI command to set the regional endpoint on the imported API is:
	//
	//    aws apigateway import-rest-api --parameters endpointConfigurationTypes=REGIONAL
	//    --body 'file:///path/to/imported-api-body.json'
	Parameters map[string]string `location:"querystring" locationName:"parameters" type:"map"`
}

// String returns the string representation
func (s ImportRestApiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportRestApiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportRestApiInput"}

	if s.Body == nil {
		invalidParams.Add(aws.NewErrParamRequired("Body"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ImportRestApiInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "body", protocol.BytesStream(v), metadata)
	}
	if s.FailOnWarnings != nil {
		v := *s.FailOnWarnings

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "failonwarnings", protocol.BoolValue(v), metadata)
	}
	if s.Parameters != nil {
		v := s.Parameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.QueryTarget, "parameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Represents a REST API.
//
// Create an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type ImportRestApiOutput struct {
	_ struct{} `type:"structure"`

	// The source of the API key for metering requests according to a usage plan.
	// Valid values are:
	//    * HEADER to read the API key from the X-API-Key header of a request.
	//
	//    * AUTHORIZER to read the API key from the UsageIdentifierKey from a custom
	//    authorizer.
	ApiKeySource ApiKeySourceType `locationName:"apiKeySource" type:"string" enum:"true"`

	// The list of binary media types supported by the RestApi. By default, the
	// RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []string `locationName:"binaryMediaTypes" type:"list"`

	// The timestamp when the API was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// The API's description.
	Description *string `locationName:"description" type:"string"`

	// The endpoint configuration of this RestApi showing the endpoint types of
	// the API.
	EndpointConfiguration *EndpointConfiguration `locationName:"endpointConfiguration" type:"structure"`

	// The API's identifier. This identifier is unique across all of your APIs in
	// API Gateway.
	Id *string `locationName:"id" type:"string"`

	// A nullable integer that is used to enable compression (with non-negative
	// between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with
	// a null value) on an API. When compression is enabled, compression or decompression
	// is not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int64 `locationName:"minimumCompressionSize" type:"integer"`

	// The API's name.
	Name *string `locationName:"name" type:"string"`

	// A stringified JSON policy document that applies to this RestApi regardless
	// of the caller and Method configuration.
	Policy *string `locationName:"policy" type:"string"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// A version identifier for the API.
	Version *string `locationName:"version" type:"string"`

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s ImportRestApiOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ImportRestApiOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ApiKeySource) > 0 {
		v := s.ApiKeySource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeySource", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.BinaryMediaTypes != nil {
		v := s.BinaryMediaTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "binaryMediaTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointConfiguration != nil {
		v := s.EndpointConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "endpointConfiguration", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MinimumCompressionSize != nil {
		v := *s.MinimumCompressionSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "minimumCompressionSize", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Warnings != nil {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opImportRestApi = "ImportRestApi"

// ImportRestApiRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// A feature of the API Gateway control service for creating a new API from
// an external API definition file.
//
//    // Example sending a request using ImportRestApiRequest.
//    req := client.ImportRestApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ImportRestApiRequest(input *ImportRestApiInput) ImportRestApiRequest {
	op := &aws.Operation{
		Name:       opImportRestApi,
		HTTPMethod: "POST",
		HTTPPath:   "/restapis?mode=import",
	}

	if input == nil {
		input = &ImportRestApiInput{}
	}

	req := c.newRequest(op, input, &ImportRestApiOutput{})
	return ImportRestApiRequest{Request: req, Input: input, Copy: c.ImportRestApiRequest}
}

// ImportRestApiRequest is the request type for the
// ImportRestApi API operation.
type ImportRestApiRequest struct {
	*aws.Request
	Input *ImportRestApiInput
	Copy  func(*ImportRestApiInput) ImportRestApiRequest
}

// Send marshals and sends the ImportRestApi API request.
func (r ImportRestApiRequest) Send(ctx context.Context) (*ImportRestApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportRestApiResponse{
		ImportRestApiOutput: r.Request.Data.(*ImportRestApiOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportRestApiResponse is the response type for the
// ImportRestApi API operation.
type ImportRestApiResponse struct {
	*ImportRestApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportRestApi request.
func (r *ImportRestApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
