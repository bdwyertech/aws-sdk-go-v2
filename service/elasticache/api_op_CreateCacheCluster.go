// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateCacheCluster operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateCacheClusterMessage
type CreateCacheClusterInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether the nodes in this Memcached cluster are created in a single
	// Availability Zone or created across multiple Availability Zones in the cluster's
	// region.
	//
	// This parameter is only supported for Memcached clusters.
	//
	// If the AZMode and PreferredAvailabilityZones are not specified, ElastiCache
	// assumes single-az mode.
	AZMode AZMode `type:"string" enum:"true"`

	// Reserved parameter. The password used to access a password protected server.
	//
	// Password constraints:
	//
	//    * Must be only printable ASCII characters.
	//
	//    * Must be at least 16 characters and no more than 128 characters in length.
	//
	//    * Cannot contain any of the following characters: '/', '"', or '@'.
	//
	// For more information, see AUTH password (http://redis.io/commands/AUTH) at
	// http://redis.io/commands/AUTH.
	AuthToken *string `type:"string"`

	// This parameter is currently disabled.
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// The node group (shard) identifier. This parameter is stored as a lowercase
	// string.
	//
	// Constraints:
	//
	//    * A name must contain from 1 to 20 alphanumeric characters or hyphens.
	//
	//    * The first character must be a letter.
	//
	//    * A name cannot end with a hyphen or contain two consecutive hyphens.
	//
	// CacheClusterId is a required field
	CacheClusterId *string `type:"string" required:"true"`

	// The compute and memory capacity of the nodes in the node group (shard).
	//
	// The following node types are supported by ElastiCache. Generally speaking,
	// the current generation types provide more memory and computational power
	// at lower cost when compared to their equivalent previous generation counterparts.
	//
	//    * General purpose: Current generation: M5 node types: cache.m5.large,
	//    cache.m5.xlarge, cache.m5.2xlarge, cache.m5.4xlarge, cache.m5.12xlarge,
	//    cache.m5.24xlarge M4 node types: cache.m4.large, cache.m4.xlarge, cache.m4.2xlarge,
	//    cache.m4.4xlarge, cache.m4.10xlarge T2 node types: cache.t2.micro, cache.t2.small,
	//    cache.t2.medium Previous generation: (not recommended) T1 node types:
	//    cache.t1.micro M1 node types: cache.m1.small, cache.m1.medium, cache.m1.large,
	//    cache.m1.xlarge M3 node types: cache.m3.medium, cache.m3.large, cache.m3.xlarge,
	//    cache.m3.2xlarge
	//
	//    * Compute optimized: Previous generation: (not recommended) C1 node types:
	//    cache.c1.xlarge
	//
	//    * Memory optimized: Current generation: R5 node types: cache.r5.large,
	//    cache.r5.xlarge, cache.r5.2xlarge, cache.r5.4xlarge, cache.r5.12xlarge,
	//    cache.r5.24xlarge R4 node types: cache.r4.large, cache.r4.xlarge, cache.r4.2xlarge,
	//    cache.r4.4xlarge, cache.r4.8xlarge, cache.r4.16xlarge Previous generation:
	//    (not recommended) M2 node types: cache.m2.xlarge, cache.m2.2xlarge, cache.m2.4xlarge
	//    R3 node types: cache.r3.large, cache.r3.xlarge, cache.r3.2xlarge, cache.r3.4xlarge,
	//    cache.r3.8xlarge
	//
	// Additional node type info
	//
	//    * All current generation instance types are created in Amazon VPC by default.
	//
	//    * Redis append-only files (AOF) are not supported for T1 or T2 instances.
	//
	//    * Redis Multi-AZ with automatic failover is not supported on T1 instances.
	//
	//    * Redis configuration variables appendonly and appendfsync are not supported
	//    on Redis version 2.8.22 and later.
	CacheNodeType *string `type:"string"`

	// The name of the parameter group to associate with this cluster. If this argument
	// is omitted, the default parameter group for the specified engine is used.
	// You cannot use any parameter group which has cluster-enabled='yes' when creating
	// a cluster.
	CacheParameterGroupName *string `type:"string"`

	// A list of security group names to associate with this cluster.
	//
	// Use this parameter only when you are creating a cluster outside of an Amazon
	// Virtual Private Cloud (Amazon VPC).
	CacheSecurityGroupNames []string `locationNameList:"CacheSecurityGroupName" type:"list"`

	// The name of the subnet group to be used for the cluster.
	//
	// Use this parameter only when you are creating a cluster in an Amazon Virtual
	// Private Cloud (Amazon VPC).
	//
	// If you're going to launch your cluster in an Amazon VPC, you need to create
	// a subnet group before you start creating a cluster. For more information,
	// see Subnets and Subnet Groups (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html).
	CacheSubnetGroupName *string `type:"string"`

	// The name of the cache engine to be used for this cluster.
	//
	// Valid values for this parameter are: memcached | redis
	Engine *string `type:"string"`

	// The version number of the cache engine to be used for this cluster. To view
	// the supported cache engine versions, use the DescribeCacheEngineVersions
	// operation.
	//
	// Important: You can upgrade to a newer engine version (see Selecting a Cache
	// Engine and Version (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement)),
	// but you cannot downgrade to an earlier engine version. If you want to use
	// an earlier engine version, you must delete the existing cluster or replication
	// group and create it anew with the earlier engine version.
	EngineVersion *string `type:"string"`

	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service
	// (SNS) topic to which notifications are sent.
	//
	// The Amazon SNS topic owner must be the same as the cluster owner.
	NotificationTopicArn *string `type:"string"`

	// The initial number of cache nodes that the cluster has.
	//
	// For clusters running Redis, this value must be 1. For clusters running Memcached,
	// this value must be between 1 and 20.
	//
	// If you need more than 20 nodes for your Memcached cluster, please fill out
	// the ElastiCache Limit Increase Request form at http://aws.amazon.com/contact-us/elasticache-node-limit-request/
	// (http://aws.amazon.com/contact-us/elasticache-node-limit-request/).
	NumCacheNodes *int64 `type:"integer"`

	// The port number on which each of the cache nodes accepts connections.
	Port *int64 `type:"integer"`

	// The EC2 Availability Zone in which the cluster is created.
	//
	// All nodes belonging to this Memcached cluster are placed in the preferred
	// Availability Zone. If you want to create your nodes across multiple Availability
	// Zones, use PreferredAvailabilityZones.
	//
	// Default: System chosen Availability Zone.
	PreferredAvailabilityZone *string `type:"string"`

	// A list of the Availability Zones in which cache nodes are created. The order
	// of the zones in the list is not important.
	//
	// This option is only supported on Memcached.
	//
	// If you are creating your cluster in an Amazon VPC (recommended) you can only
	// locate nodes in Availability Zones that are associated with the subnets in
	// the selected subnet group.
	//
	// The number of Availability Zones listed must equal the value of NumCacheNodes.
	//
	// If you want all the nodes in the same Availability Zone, use PreferredAvailabilityZone
	// instead, or repeat the Availability Zone multiple times in the list.
	//
	// Default: System chosen Availability Zones.
	PreferredAvailabilityZones []string `locationNameList:"PreferredAvailabilityZone" type:"list"`

	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Valid
	// values for ddd are:
	//
	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	//
	// Valid values for ddd are:
	//
	//    * sun
	//
	//    * mon
	//
	//    * tue
	//
	//    * wed
	//
	//    * thu
	//
	//    * fri
	//
	//    * sat
	//
	// Example: sun:23:00-mon:01:30
	PreferredMaintenanceWindow *string `type:"string"`

	// The ID of the replication group to which this cluster should belong. If this
	// parameter is specified, the cluster is added to the specified replication
	// group as a read replica; otherwise, the cluster is a standalone primary that
	// is not part of any replication group.
	//
	// If the specified replication group is Multi-AZ enabled and the Availability
	// Zone is not specified, the cluster is created in Availability Zones that
	// provide the best spread of read replicas across Availability Zones.
	//
	// This parameter is only valid if the Engine parameter is redis.
	ReplicationGroupId *string `type:"string"`

	// One or more VPC security groups associated with the cluster.
	//
	// Use this parameter only when you are creating a cluster in an Amazon Virtual
	// Private Cloud (Amazon VPC).
	SecurityGroupIds []string `locationNameList:"SecurityGroupId" type:"list"`

	// A single-element string list containing an Amazon Resource Name (ARN) that
	// uniquely identifies a Redis RDB snapshot file stored in Amazon S3. The snapshot
	// file is used to populate the node group (shard). The Amazon S3 object name
	// in the ARN cannot contain any commas.
	//
	// This parameter is only valid if the Engine parameter is redis.
	//
	// Example of an Amazon S3 ARN: arn:aws:s3:::my_bucket/snapshot1.rdb
	SnapshotArns []string `locationNameList:"SnapshotArn" type:"list"`

	// The name of a Redis snapshot from which to restore data into the new node
	// group (shard). The snapshot status changes to restoring while the new node
	// group (shard) is being created.
	//
	// This parameter is only valid if the Engine parameter is redis.
	SnapshotName *string `type:"string"`

	// The number of days for which ElastiCache retains automatic snapshots before
	// deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot
	// taken today is retained for 5 days before being deleted.
	//
	// This parameter is only valid if the Engine parameter is redis.
	//
	// Default: 0 (i.e., automatic backups are disabled for this cache cluster).
	SnapshotRetentionLimit *int64 `type:"integer"`

	// The daily time range (in UTC) during which ElastiCache begins taking a daily
	// snapshot of your node group (shard).
	//
	// Example: 05:00-09:00
	//
	// If you do not specify this parameter, ElastiCache automatically chooses an
	// appropriate time range.
	//
	// This parameter is only valid if the Engine parameter is redis.
	SnapshotWindow *string `type:"string"`

	// A list of cost allocation tags to be added to this resource.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateCacheClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCacheClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCacheClusterInput"}

	if s.CacheClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateCacheClusterResult
type CreateCacheClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific cluster.
	CacheCluster *CacheCluster `type:"structure"`
}

// String returns the string representation
func (s CreateCacheClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCacheCluster = "CreateCacheCluster"

// CreateCacheClusterRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Creates a cluster. All nodes in the cluster run the same protocol-compliant
// cache engine software, either Memcached or Redis.
//
// This operation is not supported for Redis (cluster mode enabled) clusters.
//
//    // Example sending a request using CreateCacheClusterRequest.
//    req := client.CreateCacheClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateCacheCluster
func (c *Client) CreateCacheClusterRequest(input *CreateCacheClusterInput) CreateCacheClusterRequest {
	op := &aws.Operation{
		Name:       opCreateCacheCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCacheClusterInput{}
	}

	req := c.newRequest(op, input, &CreateCacheClusterOutput{})
	return CreateCacheClusterRequest{Request: req, Input: input, Copy: c.CreateCacheClusterRequest}
}

// CreateCacheClusterRequest is the request type for the
// CreateCacheCluster API operation.
type CreateCacheClusterRequest struct {
	*aws.Request
	Input *CreateCacheClusterInput
	Copy  func(*CreateCacheClusterInput) CreateCacheClusterRequest
}

// Send marshals and sends the CreateCacheCluster API request.
func (r CreateCacheClusterRequest) Send(ctx context.Context) (*CreateCacheClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCacheClusterResponse{
		CreateCacheClusterOutput: r.Request.Data.(*CreateCacheClusterOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCacheClusterResponse is the response type for the
// CreateCacheCluster API operation.
type CreateCacheClusterResponse struct {
	*CreateCacheClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCacheCluster request.
func (r *CreateCacheClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
