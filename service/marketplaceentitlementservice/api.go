// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplaceentitlementservice

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

const opGetEntitlements = "GetEntitlements"

// GetEntitlementsRequest generates a "aws.Request" representing the
// client's request for the GetEntitlements operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetEntitlements for more information on using the GetEntitlements
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetEntitlementsRequest method.
//    req, resp := client.GetEntitlementsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/entitlement.marketplace-2017-01-11/GetEntitlements
func (c *MarketplaceEntitlementService) GetEntitlementsRequest(input *GetEntitlementsInput) (req *aws.Request, output *GetEntitlementsOutput) {
	op := &aws.Operation{
		Name:       opGetEntitlements,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetEntitlementsInput{}
	}

	output = &GetEntitlementsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetEntitlements API operation for AWS Marketplace Entitlement Service.
//
// GetEntitlements retrieves entitlement values for a given product. The results
// can be filtered based on customer identifier or product dimensions.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Marketplace Entitlement Service's
// API operation GetEntitlements for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   One or more parameters in your request was invalid.
//
//   * ErrCodeThrottlingException "ThrottlingException"
//   The calls to the GetEntitlements API are throttled.
//
//   * ErrCodeInternalServiceErrorException "InternalServiceErrorException"
//   An internal error has occurred. Retry your request. If the problem persists,
//   post a message with details on the AWS forums.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/entitlement.marketplace-2017-01-11/GetEntitlements
func (c *MarketplaceEntitlementService) GetEntitlements(input *GetEntitlementsInput) (*GetEntitlementsOutput, error) {
	req, out := c.GetEntitlementsRequest(input)
	return out, req.Send()
}

// GetEntitlementsWithContext is the same as GetEntitlements with the addition of
// the ability to pass a context and additional request options.
//
// See GetEntitlements for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MarketplaceEntitlementService) GetEntitlementsWithContext(ctx aws.Context, input *GetEntitlementsInput, opts ...aws.Option) (*GetEntitlementsOutput, error) {
	req, out := c.GetEntitlementsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// An entitlement represents capacity in a product owned by the customer. For
// example, a customer might own some number of users or seats in an SaaS application
// or some amount of data capacity in a multi-tenant database.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/entitlement.marketplace-2017-01-11/Entitlement
type Entitlement struct {
	_ struct{} `type:"structure"`

	// The customer identifier is a handle to each unique customer in an application.
	// Customer identifiers are obtained through the ResolveCustomer operation in
	// AWS Marketplace Metering Service.
	CustomerIdentifier *string `type:"string"`

	// The dimension for which the given entitlement applies. Dimensions represent
	// categories of capacity in a product and are specified when the product is
	// listed in AWS Marketplace.
	Dimension *string `type:"string"`

	// The expiration date represents the minimum date through which this entitlement
	// is expected to remain valid. For contractual products listed on AWS Marketplace,
	// the expiration date is the date at which the customer will renew or cancel
	// their contract. Customers who are opting to renew their contract will still
	// have entitlements with an expiration date.
	ExpirationDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The product code for which the given entitlement applies. Product codes are
	// provided by AWS Marketplace when the product listing is created.
	ProductCode *string `min:"1" type:"string"`

	// The EntitlementValue represents the amount of capacity that the customer
	// is entitled to for the product.
	Value *EntitlementValue `type:"structure"`
}

// String returns the string representation
func (s Entitlement) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Entitlement) GoString() string {
	return s.String()
}

// SetCustomerIdentifier sets the CustomerIdentifier field's value.
func (s *Entitlement) SetCustomerIdentifier(v string) *Entitlement {
	s.CustomerIdentifier = &v
	return s
}

// SetDimension sets the Dimension field's value.
func (s *Entitlement) SetDimension(v string) *Entitlement {
	s.Dimension = &v
	return s
}

// SetExpirationDate sets the ExpirationDate field's value.
func (s *Entitlement) SetExpirationDate(v time.Time) *Entitlement {
	s.ExpirationDate = &v
	return s
}

// SetProductCode sets the ProductCode field's value.
func (s *Entitlement) SetProductCode(v string) *Entitlement {
	s.ProductCode = &v
	return s
}

// SetValue sets the Value field's value.
func (s *Entitlement) SetValue(v *EntitlementValue) *Entitlement {
	s.Value = v
	return s
}

// The EntitlementValue represents the amount of capacity that the customer
// is entitled to for the product.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/entitlement.marketplace-2017-01-11/EntitlementValue
type EntitlementValue struct {
	_ struct{} `type:"structure"`

	// The BooleanValue field will be populated with a boolean value when the entitlement
	// is a boolean type. Otherwise, the field will not be set.
	BooleanValue *bool `type:"boolean"`

	// The DoubleValue field will be populated with a double value when the entitlement
	// is a double type. Otherwise, the field will not be set.
	DoubleValue *float64 `type:"double"`

	// The IntegerValue field will be populated with an integer value when the entitlement
	// is an integer type. Otherwise, the field will not be set.
	IntegerValue *int64 `type:"integer"`

	// The StringValue field will be populated with a string value when the entitlement
	// is a string type. Otherwise, the field will not be set.
	StringValue *string `type:"string"`
}

// String returns the string representation
func (s EntitlementValue) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EntitlementValue) GoString() string {
	return s.String()
}

// SetBooleanValue sets the BooleanValue field's value.
func (s *EntitlementValue) SetBooleanValue(v bool) *EntitlementValue {
	s.BooleanValue = &v
	return s
}

// SetDoubleValue sets the DoubleValue field's value.
func (s *EntitlementValue) SetDoubleValue(v float64) *EntitlementValue {
	s.DoubleValue = &v
	return s
}

// SetIntegerValue sets the IntegerValue field's value.
func (s *EntitlementValue) SetIntegerValue(v int64) *EntitlementValue {
	s.IntegerValue = &v
	return s
}

// SetStringValue sets the StringValue field's value.
func (s *EntitlementValue) SetStringValue(v string) *EntitlementValue {
	s.StringValue = &v
	return s
}

// The GetEntitlementsRequest contains parameters for the GetEntitlements operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/entitlement.marketplace-2017-01-11/GetEntitlementsRequest
type GetEntitlementsInput struct {
	_ struct{} `type:"structure"`

	// Filter is used to return entitlements for a specific customer or for a specific
	// dimension. Filters are described as keys mapped to a lists of values. Filtered
	// requests are unioned for each value in the value list, and then intersected
	// for each filter key.
	Filter map[string][]*string `type:"map"`

	// The maximum number of items to retrieve from the GetEntitlements operation.
	// For pagination, use the NextToken field in subsequent calls to GetEntitlements.
	MaxResults *int64 `type:"integer"`

	// For paginated calls to GetEntitlements, pass the NextToken from the previous
	// GetEntitlementsResult.
	NextToken *string `type:"string"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code will be provided by AWS Marketplace when the product listing
	// is created.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetEntitlementsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetEntitlementsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEntitlementsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEntitlementsInput"}
	if s.ProductCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilter sets the Filter field's value.
func (s *GetEntitlementsInput) SetFilter(v map[string][]*string) *GetEntitlementsInput {
	s.Filter = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *GetEntitlementsInput) SetMaxResults(v int64) *GetEntitlementsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetEntitlementsInput) SetNextToken(v string) *GetEntitlementsInput {
	s.NextToken = &v
	return s
}

// SetProductCode sets the ProductCode field's value.
func (s *GetEntitlementsInput) SetProductCode(v string) *GetEntitlementsInput {
	s.ProductCode = &v
	return s
}

// The GetEntitlementsRequest contains results from the GetEntitlements operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/entitlement.marketplace-2017-01-11/GetEntitlementsResult
type GetEntitlementsOutput struct {
	_ struct{} `type:"structure"`

	// The set of entitlements found through the GetEntitlements operation. If the
	// result contains an empty set of entitlements, NextToken might still be present
	// and should be used.
	Entitlements []*Entitlement `type:"list"`

	// For paginated results, use NextToken in subsequent calls to GetEntitlements.
	// If the result contains an empty set of entitlements, NextToken might still
	// be present and should be used.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetEntitlementsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetEntitlementsOutput) GoString() string {
	return s.String()
}

// SetEntitlements sets the Entitlements field's value.
func (s *GetEntitlementsOutput) SetEntitlements(v []*Entitlement) *GetEntitlementsOutput {
	s.Entitlements = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetEntitlementsOutput) SetNextToken(v string) *GetEntitlementsOutput {
	s.NextToken = &v
	return s
}

const (
	// GetEntitlementFilterNameCustomerIdentifier is a GetEntitlementFilterName enum value
	GetEntitlementFilterNameCustomerIdentifier = "CUSTOMER_IDENTIFIER"

	// GetEntitlementFilterNameDimension is a GetEntitlementFilterName enum value
	GetEntitlementFilterNameDimension = "DIMENSION"
)
