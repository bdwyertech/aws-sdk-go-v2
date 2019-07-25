// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/ListClustersRequest
type ListClustersInput struct {
	_ struct{} `type:"structure"`

	ClusterNameFilter *string `location:"querystring" locationName:"clusterNameFilter" type:"string"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClustersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListClustersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListClustersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListClustersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ClusterNameFilter != nil {
		v := *s.ClusterNameFilter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "clusterNameFilter", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	return nil
}

// The response contains an array containing cluster information and a next
// token if the response is truncated.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/ListClustersResponse
type ListClustersOutput struct {
	_ struct{} `type:"structure"`

	// Information on each of the MSK clusters in the response.
	ClusterInfoList []ClusterInfo `locationName:"clusterInfoList" type:"list"`

	// The paginated results marker. When the result of a ListClusters operation
	// is truncated, the call returns NextToken in the response. To get another
	// batch of clusters, provide this token in your next request.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClustersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListClustersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClusterInfoList != nil {
		v := s.ClusterInfoList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "clusterInfoList", metadata)
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

const opListClusters = "ListClusters"

// ListClustersRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Returns a list of all the MSK clusters in the current Region.
//
//    // Example sending a request using ListClustersRequest.
//    req := client.ListClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/ListClusters
func (c *Client) ListClustersRequest(input *ListClustersInput) ListClustersRequest {
	op := &aws.Operation{
		Name:       opListClusters,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/clusters",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListClustersInput{}
	}

	req := c.newRequest(op, input, &ListClustersOutput{})
	return ListClustersRequest{Request: req, Input: input, Copy: c.ListClustersRequest}
}

// ListClustersRequest is the request type for the
// ListClusters API operation.
type ListClustersRequest struct {
	*aws.Request
	Input *ListClustersInput
	Copy  func(*ListClustersInput) ListClustersRequest
}

// Send marshals and sends the ListClusters API request.
func (r ListClustersRequest) Send(ctx context.Context) (*ListClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListClustersResponse{
		ListClustersOutput: r.Request.Data.(*ListClustersOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListClustersRequestPaginator returns a paginator for ListClusters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListClustersRequest(input)
//   p := kafka.NewListClustersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListClustersPaginator(req ListClustersRequest) ListClustersPaginator {
	return ListClustersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListClustersInput
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

// ListClustersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListClustersPaginator struct {
	aws.Pager
}

func (p *ListClustersPaginator) CurrentPage() *ListClustersOutput {
	return p.Pager.CurrentPage().(*ListClustersOutput)
}

// ListClustersResponse is the response type for the
// ListClusters API operation.
type ListClustersResponse struct {
	*ListClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListClusters request.
func (r *ListClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
