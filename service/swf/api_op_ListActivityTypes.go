// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListActivityTypesInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain in which the activity types have been registered.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// The maximum number of results that are returned per call. Use nextPageToken
	// to obtain further pages of results.
	MaximumPageSize *int64 `locationName:"maximumPageSize" type:"integer"`

	// If specified, only lists the activity types that have this name.
	Name *string `locationName:"name" min:"1" type:"string"`

	// If NextPageToken is returned there are more results available. The value
	// of NextPageToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged. Each pagination token expires after 60 seconds. Using
	// an expired pagination token will return a 400 error: "Specified token has
	// exceeded its maximum lifetime".
	//
	// The configured maximumPageSize determines how many results can be returned
	// in a single call.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`

	// Specifies the registration status of the activity types to list.
	//
	// RegistrationStatus is a required field
	RegistrationStatus RegistrationStatus `locationName:"registrationStatus" type:"string" required:"true" enum:"true"`

	// When set to true, returns the results in reverse order. By default, the results
	// are returned in ascending alphabetical order by name of the activity types.
	ReverseOrder *bool `locationName:"reverseOrder" type:"boolean"`
}

// String returns the string representation
func (s ListActivityTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListActivityTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListActivityTypesInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.RegistrationStatus) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RegistrationStatus"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains a paginated list of activity type information structures.
type ListActivityTypesOutput struct {
	_ struct{} `type:"structure"`

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using
	// the returned token in nextPageToken. Keep all other arguments unchanged.
	//
	// The configured maximumPageSize determines how many results can be returned
	// in a single call.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`

	// List of activity type information.
	//
	// TypeInfos is a required field
	TypeInfos []ActivityTypeInfo `locationName:"typeInfos" type:"list" required:"true"`
}

// String returns the string representation
func (s ListActivityTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListActivityTypes = "ListActivityTypes"

// ListActivityTypesRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns information about all activities registered in the specified domain
// that match the specified name and registration status. The result includes
// information like creation date, current status of the activity, etc. The
// results may be split into multiple pages. To retrieve subsequent pages, make
// the call again using the nextPageToken returned by the initial call.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using ListActivityTypesRequest.
//    req := client.ListActivityTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListActivityTypesRequest(input *ListActivityTypesInput) ListActivityTypesRequest {
	op := &aws.Operation{
		Name:       opListActivityTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextPageToken"},
			OutputTokens:    []string{"nextPageToken"},
			LimitToken:      "maximumPageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListActivityTypesInput{}
	}

	req := c.newRequest(op, input, &ListActivityTypesOutput{})
	return ListActivityTypesRequest{Request: req, Input: input, Copy: c.ListActivityTypesRequest}
}

// ListActivityTypesRequest is the request type for the
// ListActivityTypes API operation.
type ListActivityTypesRequest struct {
	*aws.Request
	Input *ListActivityTypesInput
	Copy  func(*ListActivityTypesInput) ListActivityTypesRequest
}

// Send marshals and sends the ListActivityTypes API request.
func (r ListActivityTypesRequest) Send(ctx context.Context) (*ListActivityTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListActivityTypesResponse{
		ListActivityTypesOutput: r.Request.Data.(*ListActivityTypesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListActivityTypesRequestPaginator returns a paginator for ListActivityTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListActivityTypesRequest(input)
//   p := swf.NewListActivityTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListActivityTypesPaginator(req ListActivityTypesRequest) ListActivityTypesPaginator {
	return ListActivityTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListActivityTypesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListActivityTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListActivityTypesPaginator struct {
	aws.Pager
}

func (p *ListActivityTypesPaginator) CurrentPage() *ListActivityTypesOutput {
	return p.Pager.CurrentPage().(*ListActivityTypesOutput)
}

// ListActivityTypesResponse is the response type for the
// ListActivityTypes API operation.
type ListActivityTypesResponse struct {
	*ListActivityTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListActivityTypes request.
func (r *ListActivityTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
