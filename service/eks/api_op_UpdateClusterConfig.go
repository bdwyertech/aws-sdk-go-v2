// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/UpdateClusterConfigRequest
type UpdateClusterConfigInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string" idempotencyToken:"true"`

	// Enable or disable exporting the Kubernetes control plane logs for your cluster
	// to CloudWatch Logs. By default, cluster control plane logs aren't exported
	// to CloudWatch Logs. For more information, see Amazon EKS Cluster Control
	// Plane Logs (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
	// in the Amazon EKS User Guide .
	//
	// CloudWatch Logs ingestion, archive storage, and data scanning rates apply
	// to exported control plane logs. For more information, see Amazon CloudWatch
	// Pricing (http://aws.amazon.com/cloudwatch/pricing/).
	Logging *Logging `locationName:"logging" type:"structure"`

	// The name of the Amazon EKS cluster to update.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// An object representing the VPC configuration to use for an Amazon EKS cluster.
	ResourcesVpcConfig *VpcConfigRequest `locationName:"resourcesVpcConfig" type:"structure"`
}

// String returns the string representation
func (s UpdateClusterConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClusterConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateClusterConfigInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClusterConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Logging != nil {
		v := s.Logging

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "logging", v, metadata)
	}
	if s.ResourcesVpcConfig != nil {
		v := s.ResourcesVpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resourcesVpcConfig", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/UpdateClusterConfigResponse
type UpdateClusterConfigOutput struct {
	_ struct{} `type:"structure"`

	// An object representing an asynchronous update.
	Update *Update `locationName:"update" type:"structure"`
}

// String returns the string representation
func (s UpdateClusterConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClusterConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Update != nil {
		v := s.Update

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "update", v, metadata)
	}
	return nil
}

const opUpdateClusterConfig = "UpdateClusterConfig"

// UpdateClusterConfigRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Updates an Amazon EKS cluster configuration. Your cluster continues to function
// during the update. The response output includes an update ID that you can
// use to track the status of your cluster update with the DescribeUpdate API
// operation.
//
// You can use this API operation to enable or disable exporting the Kubernetes
// control plane logs for your cluster to CloudWatch Logs. By default, cluster
// control plane logs aren't exported to CloudWatch Logs. For more information,
// see Amazon EKS Cluster Control Plane Logs (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
// in the Amazon EKS User Guide .
//
// CloudWatch Logs ingestion, archive storage, and data scanning rates apply
// to exported control plane logs. For more information, see Amazon CloudWatch
// Pricing (http://aws.amazon.com/cloudwatch/pricing/).
//
// You can also use this API operation to enable or disable public and private
// access to your cluster's Kubernetes API server endpoint. By default, public
// access is enabled, and private access is disabled. For more information,
// see Amazon EKS Cluster Endpoint Access Control (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html)
// in the Amazon EKS User Guide .
//
// At this time, you can not update the subnets or security group IDs for an
// existing cluster.
//
// Cluster updates are asynchronous, and they should finish within a few minutes.
// During an update, the cluster status moves to UPDATING (this status transition
// is eventually consistent). When the update is complete (either Failed or
// Successful), the cluster status moves to Active.
//
//    // Example sending a request using UpdateClusterConfigRequest.
//    req := client.UpdateClusterConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/UpdateClusterConfig
func (c *Client) UpdateClusterConfigRequest(input *UpdateClusterConfigInput) UpdateClusterConfigRequest {
	op := &aws.Operation{
		Name:       opUpdateClusterConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/clusters/{name}/update-config",
	}

	if input == nil {
		input = &UpdateClusterConfigInput{}
	}

	req := c.newRequest(op, input, &UpdateClusterConfigOutput{})
	return UpdateClusterConfigRequest{Request: req, Input: input, Copy: c.UpdateClusterConfigRequest}
}

// UpdateClusterConfigRequest is the request type for the
// UpdateClusterConfig API operation.
type UpdateClusterConfigRequest struct {
	*aws.Request
	Input *UpdateClusterConfigInput
	Copy  func(*UpdateClusterConfigInput) UpdateClusterConfigRequest
}

// Send marshals and sends the UpdateClusterConfig API request.
func (r UpdateClusterConfigRequest) Send(ctx context.Context) (*UpdateClusterConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClusterConfigResponse{
		UpdateClusterConfigOutput: r.Request.Data.(*UpdateClusterConfigOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClusterConfigResponse is the response type for the
// UpdateClusterConfig API operation.
type UpdateClusterConfigResponse struct {
	*UpdateClusterConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateClusterConfig request.
func (r *UpdateClusterConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
