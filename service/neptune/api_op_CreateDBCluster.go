// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateDBClusterMessage
type CreateDBClusterInput struct {
	_ struct{} `type:"structure"`

	// A list of EC2 Availability Zones that instances in the DB cluster can be
	// created in.
	AvailabilityZones []string `locationNameList:"AvailabilityZone" type:"list"`

	// The number of days for which automated backups are retained. You must specify
	// a minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	//    * Must be a value from 1 to 35
	BackupRetentionPeriod *int64 `type:"integer"`

	// A value that indicates that the DB cluster should be associated with the
	// specified CharacterSet.
	CharacterSetName *string `type:"string"`

	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster1
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the DB cluster parameter group to associate with this DB cluster.
	// If this argument is omitted, the default is used.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string `type:"string"`

	// A DB subnet group to associate with this DB cluster.
	//
	// Constraints: Must match the name of an existing DBSubnetGroup. Must not be
	// default.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// The name for your database of up to 64 alpha-numeric characters. If you do
	// not provide a name, Amazon Neptune will not create a database in the DB cluster
	// you are creating.
	DatabaseName *string `type:"string"`

	// The list of log types that need to be enabled for exporting to CloudWatch
	// Logs.
	EnableCloudwatchLogsExports []string `type:"list"`

	// True to enable mapping of AWS Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false.
	//
	// Default: false
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values: neptune
	//
	// Engine is a required field
	Engine *string `type:"string" required:"true"`

	// The version number of the database engine to use.
	//
	// Example: 1.0.1
	EngineVersion *string `type:"string"`

	// The AWS KMS key identifier for an encrypted DB cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are creating a DB cluster with the same AWS account that owns
	// the KMS encryption key used to encrypt the new DB cluster, then you can use
	// the KMS key alias instead of the ARN for the KMS encryption key.
	//
	// If an encryption key is not specified in KmsKeyId:
	//
	//    * If ReplicationSourceIdentifier identifies an encrypted source, then
	//    Amazon Neptune will use the encryption key used to encrypt the source.
	//    Otherwise, Amazon Neptune will use your default encryption key.
	//
	//    * If the StorageEncrypted parameter is true and ReplicationSourceIdentifier
	//    is not specified, then Amazon Neptune will use your default encryption
	//    key.
	//
	// AWS KMS creates the default encryption key for your AWS account. Your AWS
	// account has a different default encryption key for each AWS Region.
	//
	// If you create a Read Replica of an encrypted DB cluster in another AWS Region,
	// you must set KmsKeyId to a KMS key ID that is valid in the destination AWS
	// Region. This key is used to encrypt the Read Replica in that AWS Region.
	KmsKeyId *string `type:"string"`

	// The password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@".
	//
	// Constraints: Must contain from 8 to 41 characters.
	MasterUserPassword *string `type:"string"`

	// The name of the master user for the DB cluster.
	//
	// Constraints:
	//
	//    * Must be 1 to 16 letters or numbers.
	//
	//    * First character must be a letter.
	//
	//    * Cannot be a reserved word for the chosen database engine.
	MasterUsername *string `type:"string"`

	// A value that indicates that the DB cluster should be associated with the
	// specified option group.
	//
	// Permanent options can't be removed from an option group. The option group
	// can't be removed from a DB cluster once it is associated with a DB cluster.
	OptionGroupName *string `type:"string"`

	// The port number on which the instances in the DB cluster accept connections.
	//
	// Default: 8182
	Port *int64 `type:"integer"`

	// This parameter is not currently supported.
	PreSignedUrl *string `type:"string"`

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region. To see the time blocks available, see Adjusting
	// the Preferred Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AdjustingTheMaintenanceWindow.html)
	// in the Amazon Neptune User Guide.
	//
	// Constraints:
	//
	//    * Must be in the format hh24:mi-hh24:mi.
	//
	//    * Must be in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred maintenance window.
	//
	//    * Must be at least 30 minutes.
	PreferredBackupWindow *string `type:"string"`

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region, occurring on a random day of the week. To see
	// the time blocks available, see Adjusting the Preferred Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AdjustingTheMaintenanceWindow.html)
	// in the Amazon Neptune User Guide.
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `type:"string"`

	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if
	// this DB cluster is created as a Read Replica.
	ReplicationSourceIdentifier *string `type:"string"`

	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool `type:"boolean"`

	// The tags to assign to the new DB cluster.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A list of EC2 VPC security groups to associate with this DB cluster.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s CreateDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.Engine == nil {
		invalidParams.Add(aws.NewErrParamRequired("Engine"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateDBClusterResult
type CreateDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters action.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s CreateDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBCluster = "CreateDBCluster"

// CreateDBClusterRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Creates a new Amazon Neptune DB cluster.
//
// You can use the ReplicationSourceIdentifier parameter to create the DB cluster
// as a Read Replica of another DB cluster or Amazon Neptune DB instance.
//
//    // Example sending a request using CreateDBClusterRequest.
//    req := client.CreateDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateDBCluster
func (c *Client) CreateDBClusterRequest(input *CreateDBClusterInput) CreateDBClusterRequest {
	op := &aws.Operation{
		Name:       opCreateDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBClusterInput{}
	}

	req := c.newRequest(op, input, &CreateDBClusterOutput{})
	return CreateDBClusterRequest{Request: req, Input: input, Copy: c.CreateDBClusterRequest}
}

// CreateDBClusterRequest is the request type for the
// CreateDBCluster API operation.
type CreateDBClusterRequest struct {
	*aws.Request
	Input *CreateDBClusterInput
	Copy  func(*CreateDBClusterInput) CreateDBClusterRequest
}

// Send marshals and sends the CreateDBCluster API request.
func (r CreateDBClusterRequest) Send(ctx context.Context) (*CreateDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBClusterResponse{
		CreateDBClusterOutput: r.Request.Data.(*CreateDBClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBClusterResponse is the response type for the
// CreateDBCluster API operation.
type CreateDBClusterResponse struct {
	*CreateDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBCluster request.
func (r *CreateDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
