// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListMembersRequest
type ListMembersInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector the member is associated with.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 50. The maximum value is 50.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls
	// to the action fill nextToken in the request with the value of NextToken from
	// the previous response to continue listing data.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// Specifies whether to only return associated members or to return all members
	// (including members which haven't been invited yet or have been disassociated).
	OnlyAssociated *string `location:"querystring" locationName:"onlyAssociated" type:"string"`
}

// String returns the string representation
func (s ListMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMembersInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OnlyAssociated != nil {
		v := *s.OnlyAssociated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "onlyAssociated", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListMembersResponse
type ListMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of members.
	Members []Member `locationName:"members" type:"list"`

	// Pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Members != nil {
		v := s.Members

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "members", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListMembers = "ListMembers"

// ListMembersRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Lists details about all member accounts for the current GuardDuty master
// account.
//
//    // Example sending a request using ListMembersRequest.
//    req := client.ListMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListMembers
func (c *Client) ListMembersRequest(input *ListMembersInput) ListMembersRequest {
	op := &aws.Operation{
		Name:       opListMembers,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/member",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListMembersInput{}
	}

	req := c.newRequest(op, input, &ListMembersOutput{})
	return ListMembersRequest{Request: req, Input: input, Copy: c.ListMembersRequest}
}

// ListMembersRequest is the request type for the
// ListMembers API operation.
type ListMembersRequest struct {
	*aws.Request
	Input *ListMembersInput
	Copy  func(*ListMembersInput) ListMembersRequest
}

// Send marshals and sends the ListMembers API request.
func (r ListMembersRequest) Send(ctx context.Context) (*ListMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMembersResponse{
		ListMembersOutput: r.Request.Data.(*ListMembersOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMembersRequestPaginator returns a paginator for ListMembers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMembersRequest(input)
//   p := guardduty.NewListMembersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMembersPaginator(req ListMembersRequest) ListMembersPaginator {
	return ListMembersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMembersInput
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

// ListMembersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMembersPaginator struct {
	aws.Pager
}

func (p *ListMembersPaginator) CurrentPage() *ListMembersOutput {
	return p.Pager.CurrentPage().(*ListMembersOutput)
}

// ListMembersResponse is the response type for the
// ListMembers API operation.
type ListMembersResponse struct {
	*ListMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMembers request.
func (r *ListMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
