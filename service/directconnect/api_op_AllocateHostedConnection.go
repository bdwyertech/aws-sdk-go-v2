// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocateHostedConnectionRequest
type AllocateHostedConnectionInput struct {
	_ struct{} `type:"structure"`

	// The bandwidth of the connection. The possible values are 50Mbps, 100Mbps,
	// 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, and 10Gbps. Note
	// that only those AWS Direct Connect Partners who have met specific requirements
	// are allowed to create a 1Gbps, 2Gbps, 5Gbps or 10Gbps hosted connection.
	//
	// Bandwidth is a required field
	Bandwidth *string `locationName:"bandwidth" type:"string" required:"true"`

	// The ID of the interconnect or LAG.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`

	// The name of the hosted connection.
	//
	// ConnectionName is a required field
	ConnectionName *string `locationName:"connectionName" type:"string" required:"true"`

	// The ID of the AWS account ID of the customer for the connection.
	//
	// OwnerAccount is a required field
	OwnerAccount *string `locationName:"ownerAccount" type:"string" required:"true"`

	// The tags to assign to the hosted connection.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`

	// The dedicated VLAN provisioned to the hosted connection.
	//
	// Vlan is a required field
	Vlan *int64 `locationName:"vlan" type:"integer" required:"true"`
}

// String returns the string representation
func (s AllocateHostedConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AllocateHostedConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AllocateHostedConnectionInput"}

	if s.Bandwidth == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bandwidth"))
	}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if s.ConnectionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionName"))
	}

	if s.OwnerAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("OwnerAccount"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if s.Vlan == nil {
		invalidParams.Add(aws.NewErrParamRequired("Vlan"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about an AWS Direct Connect connection.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/Connection
type AllocateHostedConnectionOutput struct {
	_ struct{} `type:"structure"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string `locationName:"awsDevice" deprecated:"true" type:"string"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDeviceV2 *string `locationName:"awsDeviceV2" type:"string"`

	// The bandwidth of the connection.
	Bandwidth *string `locationName:"bandwidth" type:"string"`

	// The ID of the connection.
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// The name of the connection.
	ConnectionName *string `locationName:"connectionName" type:"string"`

	// The state of the connection. The following are the possible values:
	//
	//    * ordering: The initial state of a hosted connection provisioned on an
	//    interconnect. The connection stays in the ordering state until the owner
	//    of the hosted connection confirms or declines the connection order.
	//
	//    * requested: The initial state of a standard connection. The connection
	//    stays in the requested state until the Letter of Authorization (LOA) is
	//    sent to the customer.
	//
	//    * pending: The connection has been approved and is being initialized.
	//
	//    * available: The network link is up and the connection is ready for use.
	//
	//    * down: The network link is down.
	//
	//    * deleting: The connection is being deleted.
	//
	//    * deleted: The connection has been deleted.
	//
	//    * rejected: A hosted connection in the ordering state enters the rejected
	//    state if it is deleted by the customer.
	//
	//    * unknown: The state of the connection is not available.
	ConnectionState ConnectionState `locationName:"connectionState" type:"string" enum:"true"`

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy HasLogicalRedundancy `locationName:"hasLogicalRedundancy" type:"string" enum:"true"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `locationName:"jumboFrameCapable" type:"boolean"`

	// The ID of the LAG.
	LagId *string `locationName:"lagId" type:"string"`

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time `locationName:"loaIssueTime" type:"timestamp" timestampFormat:"unix"`

	// The location of the connection.
	Location *string `locationName:"location" type:"string"`

	// The ID of the AWS account that owns the connection.
	OwnerAccount *string `locationName:"ownerAccount" type:"string"`

	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName *string `locationName:"partnerName" type:"string"`

	// The AWS Region where the connection is located.
	Region *string `locationName:"region" type:"string"`

	// Any tags assigned to the connection.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`

	// The ID of the VLAN.
	Vlan *int64 `locationName:"vlan" type:"integer"`
}

// String returns the string representation
func (s AllocateHostedConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opAllocateHostedConnection = "AllocateHostedConnection"

// AllocateHostedConnectionRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Creates a hosted connection on the specified interconnect or a link aggregation
// group (LAG) of interconnects.
//
// Allocates a VLAN number and a specified amount of capacity (bandwidth) for
// use by a hosted connection on the specified interconnect or LAG of interconnects.
// AWS polices the hosted connection for the specified capacity and the AWS
// Direct Connect Partner must also police the hosted connection for the specified
// capacity.
//
// Intended for use by AWS Direct Connect Partners only.
//
//    // Example sending a request using AllocateHostedConnectionRequest.
//    req := client.AllocateHostedConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocateHostedConnection
func (c *Client) AllocateHostedConnectionRequest(input *AllocateHostedConnectionInput) AllocateHostedConnectionRequest {
	op := &aws.Operation{
		Name:       opAllocateHostedConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AllocateHostedConnectionInput{}
	}

	req := c.newRequest(op, input, &AllocateHostedConnectionOutput{})
	return AllocateHostedConnectionRequest{Request: req, Input: input, Copy: c.AllocateHostedConnectionRequest}
}

// AllocateHostedConnectionRequest is the request type for the
// AllocateHostedConnection API operation.
type AllocateHostedConnectionRequest struct {
	*aws.Request
	Input *AllocateHostedConnectionInput
	Copy  func(*AllocateHostedConnectionInput) AllocateHostedConnectionRequest
}

// Send marshals and sends the AllocateHostedConnection API request.
func (r AllocateHostedConnectionRequest) Send(ctx context.Context) (*AllocateHostedConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AllocateHostedConnectionResponse{
		AllocateHostedConnectionOutput: r.Request.Data.(*AllocateHostedConnectionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AllocateHostedConnectionResponse is the response type for the
// AllocateHostedConnection API operation.
type AllocateHostedConnectionResponse struct {
	*AllocateHostedConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AllocateHostedConnection request.
func (r *AllocateHostedConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
