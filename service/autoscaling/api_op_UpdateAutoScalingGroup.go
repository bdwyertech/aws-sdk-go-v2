// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/UpdateAutoScalingGroupType
type UpdateAutoScalingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// One or more Availability Zones for the group.
	AvailabilityZones []string `min:"1" type:"list"`

	// The amount of time, in seconds, after a scaling activity completes before
	// another scaling activity can start. The default value is 300. This cooldown
	// period is not used when a scaling-specific cooldown is specified.
	//
	// Cooldown periods are not supported for target tracking scaling policies,
	// step scaling policies, or scheduled scaling. For more information, see Scaling
	// Cooldowns (https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	DefaultCooldown *int64 `type:"integer"`

	// The number of EC2 instances that should be running in the Auto Scaling group.
	// This number must be greater than or equal to the minimum size of the group
	// and less than or equal to the maximum size of the group.
	DesiredCapacity *int64 `type:"integer"`

	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before
	// checking the health status of an EC2 instance that has come into service.
	// The default value is 0.
	//
	// For more information, see Health Checks for Auto Scaling Instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// Conditional: This parameter is required if you are adding an ELB health check.
	HealthCheckGracePeriod *int64 `type:"integer"`

	// The service to use for the health checks. The valid values are EC2 and ELB.
	// If you configure an Auto Scaling group to use ELB health checks, it considers
	// the instance unhealthy if it fails either the EC2 status checks or the load
	// balancer health checks.
	HealthCheckType *string `min:"1" type:"string"`

	// The name of the launch configuration. If you specify LaunchConfigurationName
	// in your update request, you can't specify LaunchTemplate or MixedInstancesPolicy.
	//
	// To update an Auto Scaling group with a launch configuration with InstanceMonitoring
	// set to false, you must first disable the collection of group metrics. Otherwise,
	// you get an error. If you have previously enabled the collection of group
	// metrics, you can disable it using DisableMetricsCollection.
	LaunchConfigurationName *string `min:"1" type:"string"`

	// The launch template and version to use to specify the updates. If you specify
	// LaunchTemplate in your update request, you can't specify LaunchConfigurationName
	// or MixedInstancesPolicy.
	LaunchTemplate *LaunchTemplateSpecification `type:"structure"`

	// The maximum size of the Auto Scaling group.
	MaxSize *int64 `type:"integer"`

	// The minimum size of the Auto Scaling group.
	MinSize *int64 `type:"integer"`

	// An embedded object that specifies a mixed instances policy.
	//
	// In your call to UpdateAutoScalingGroup, you can make changes to the policy
	// that is specified. All optional parameters are left unchanged if not specified.
	//
	// For more information, see Auto Scaling Groups with Multiple Instance Types
	// and Purchase Options (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-purchase-options.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	MixedInstancesPolicy *MixedInstancesPolicy `type:"structure"`

	// Indicates whether newly launched instances are protected from termination
	// by Amazon EC2 Auto Scaling when scaling in.
	//
	// For more information about preventing instances from terminating on scale
	// in, see Instance Protection (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection)
	// in the Amazon EC2 Auto Scaling User Guide.
	NewInstancesProtectedFromScaleIn *bool `type:"boolean"`

	// The name of the placement group into which to launch your instances, if any.
	// A placement group is a logical grouping of instances within a single Availability
	// Zone. You cannot specify multiple Availability Zones and a placement group.
	// For more information, see Placement Groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	PlacementGroup *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling
	// group uses to call other AWS services on your behalf. For more information,
	// see Service-Linked Roles (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	ServiceLinkedRoleARN *string `min:"1" type:"string"`

	// A standalone termination policy or a list of termination policies used to
	// select the instance to terminate. The policies are executed in the order
	// that they are listed.
	//
	// For more information, see Controlling Which Instances Auto Scaling Terminates
	// During Scale In (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	TerminationPolicies []string `type:"list"`

	// A comma-separated list of subnet IDs for virtual private cloud (VPC).
	//
	// If you specify VPCZoneIdentifier with AvailabilityZones, the subnets that
	// you specify for this parameter must reside in those Availability Zones.
	VPCZoneIdentifier *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateAutoScalingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAutoScalingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAutoScalingGroupInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}
	if s.AvailabilityZones != nil && len(s.AvailabilityZones) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AvailabilityZones", 1))
	}
	if s.HealthCheckType != nil && len(*s.HealthCheckType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HealthCheckType", 1))
	}
	if s.LaunchConfigurationName != nil && len(*s.LaunchConfigurationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LaunchConfigurationName", 1))
	}
	if s.PlacementGroup != nil && len(*s.PlacementGroup) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementGroup", 1))
	}
	if s.ServiceLinkedRoleARN != nil && len(*s.ServiceLinkedRoleARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceLinkedRoleARN", 1))
	}
	if s.VPCZoneIdentifier != nil && len(*s.VPCZoneIdentifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VPCZoneIdentifier", 1))
	}
	if s.LaunchTemplate != nil {
		if err := s.LaunchTemplate.Validate(); err != nil {
			invalidParams.AddNested("LaunchTemplate", err.(aws.ErrInvalidParams))
		}
	}
	if s.MixedInstancesPolicy != nil {
		if err := s.MixedInstancesPolicy.Validate(); err != nil {
			invalidParams.AddNested("MixedInstancesPolicy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/UpdateAutoScalingGroupOutput
type UpdateAutoScalingGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAutoScalingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAutoScalingGroup = "UpdateAutoScalingGroup"

// UpdateAutoScalingGroupRequest returns a request value for making API operation for
// Auto Scaling.
//
// Updates the configuration for the specified Auto Scaling group.
//
// To update an Auto Scaling group, specify the name of the group and the parameter
// that you want to change. Any parameters that you don't specify are not changed
// by this update request. The new settings take effect on any scaling activities
// after this call returns. Scaling activities that are currently in progress
// aren't affected.
//
// If you associate a new launch configuration or template with an Auto Scaling
// group, all new instances will get the updated configuration, but existing
// instances continue to run with the configuration that they were originally
// launched with. When you update a group to specify a mixed instances policy
// instead of a launch configuration or template, existing instances may be
// replaced to match the new purchasing options that you specified in the policy.
// For example, if the group currently has 100% On-Demand capacity and the policy
// specifies 50% Spot capacity, this means that half of your instances will
// be gradually terminated and relaunched as Spot Instances. When replacing
// instances, Amazon EC2 Auto Scaling launches new instances before terminating
// the old ones, so that updating your group does not compromise the performance
// or availability of your application.
//
// Note the following about changing DesiredCapacity, MaxSize, or MinSize:
//
//    * If a scale-in event occurs as a result of a new DesiredCapacity value
//    that is lower than the current size of the group, the Auto Scaling group
//    uses its termination policy to determine which instances to terminate.
//
//    * If you specify a new value for MinSize without specifying a value for
//    DesiredCapacity, and the new MinSize is larger than the current size of
//    the group, this sets the group's DesiredCapacity to the new MinSize value.
//
//    * If you specify a new value for MaxSize without specifying a value for
//    DesiredCapacity, and the new MaxSize is smaller than the current size
//    of the group, this sets the group's DesiredCapacity to the new MaxSize
//    value.
//
// To see which parameters have been set, use DescribeAutoScalingGroups. You
// can also view the scaling policies for an Auto Scaling group using DescribePolicies.
// If the group has scaling policies, you can update them using PutScalingPolicy.
//
//    // Example sending a request using UpdateAutoScalingGroupRequest.
//    req := client.UpdateAutoScalingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/UpdateAutoScalingGroup
func (c *Client) UpdateAutoScalingGroupRequest(input *UpdateAutoScalingGroupInput) UpdateAutoScalingGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateAutoScalingGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAutoScalingGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateAutoScalingGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateAutoScalingGroupRequest{Request: req, Input: input, Copy: c.UpdateAutoScalingGroupRequest}
}

// UpdateAutoScalingGroupRequest is the request type for the
// UpdateAutoScalingGroup API operation.
type UpdateAutoScalingGroupRequest struct {
	*aws.Request
	Input *UpdateAutoScalingGroupInput
	Copy  func(*UpdateAutoScalingGroupInput) UpdateAutoScalingGroupRequest
}

// Send marshals and sends the UpdateAutoScalingGroup API request.
func (r UpdateAutoScalingGroupRequest) Send(ctx context.Context) (*UpdateAutoScalingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAutoScalingGroupResponse{
		UpdateAutoScalingGroupOutput: r.Request.Data.(*UpdateAutoScalingGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAutoScalingGroupResponse is the response type for the
// UpdateAutoScalingGroup API operation.
type UpdateAutoScalingGroupResponse struct {
	*UpdateAutoScalingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAutoScalingGroup request.
func (r *UpdateAutoScalingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
