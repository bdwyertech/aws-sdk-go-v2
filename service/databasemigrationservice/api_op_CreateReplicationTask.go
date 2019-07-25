// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/CreateReplicationTaskMessage
type CreateReplicationTaskInput struct {
	_ struct{} `type:"structure"`

	// Indicates when you want a change data capture (CDC) operation to start. Use
	// either CdcStartPosition or CdcStartTime to specify when you want a CDC operation
	// to start. Specifying both values results in an error.
	//
	// The value can be in date, checkpoint, or LSN/SCN format.
	//
	// Date Example: --cdc-start-position “2018-03-08T12:12:12”
	//
	// Checkpoint Example: --cdc-start-position "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"
	//
	// LSN Example: --cdc-start-position “mysql-bin-changelog.000024:373”
	CdcStartPosition *string `type:"string"`

	// Indicates the start time for a change data capture (CDC) operation. Use either
	// CdcStartTime or CdcStartPosition to specify when you want a CDC operation
	// to start. Specifying both values results in an error.
	//
	// Timestamp Example: --cdc-start-time “2018-03-08T12:12:12”
	CdcStartTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Indicates when you want a change data capture (CDC) operation to stop. The
	// value can be either server time or commit time.
	//
	// Server time example: --cdc-stop-position “server_time:3018-02-09T12:12:12”
	//
	// Commit time example: --cdc-stop-position “commit_time: 3018-02-09T12:12:12
	// “
	CdcStopPosition *string `type:"string"`

	// The migration type. Valid values: full-load | cdc | full-load-and-cdc
	//
	// MigrationType is a required field
	MigrationType MigrationTypeValue `type:"string" required:"true" enum:"true"`

	// The Amazon Resource Name (ARN) of a replication instance.
	//
	// ReplicationInstanceArn is a required field
	ReplicationInstanceArn *string `type:"string" required:"true"`

	// An identifier for the replication task.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 255 alphanumeric characters or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// ReplicationTaskIdentifier is a required field
	ReplicationTaskIdentifier *string `type:"string" required:"true"`

	// Overall settings for the task, in JSON format. For more information, see
	// Task Settings (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html)
	// in the AWS Database Migration User Guide.
	ReplicationTaskSettings *string `type:"string"`

	// An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.
	//
	// SourceEndpointArn is a required field
	SourceEndpointArn *string `type:"string" required:"true"`

	// The table mappings for the task, in JSON format. For more information, see
	// Table Mapping (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	// in the AWS Database Migration User Guide.
	//
	// TableMappings is a required field
	TableMappings *string `type:"string" required:"true"`

	// One or more tags to be assigned to the replication task.
	Tags []Tag `type:"list"`

	// An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.
	//
	// TargetEndpointArn is a required field
	TargetEndpointArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateReplicationTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateReplicationTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateReplicationTaskInput"}
	if len(s.MigrationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("MigrationType"))
	}

	if s.ReplicationInstanceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationInstanceArn"))
	}

	if s.ReplicationTaskIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationTaskIdentifier"))
	}

	if s.SourceEndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceEndpointArn"))
	}

	if s.TableMappings == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableMappings"))
	}

	if s.TargetEndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetEndpointArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/CreateReplicationTaskResponse
type CreateReplicationTaskOutput struct {
	_ struct{} `type:"structure"`

	// The replication task that was created.
	ReplicationTask *ReplicationTask `type:"structure"`
}

// String returns the string representation
func (s CreateReplicationTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateReplicationTask = "CreateReplicationTask"

// CreateReplicationTaskRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Creates a replication task using the specified parameters.
//
//    // Example sending a request using CreateReplicationTaskRequest.
//    req := client.CreateReplicationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/CreateReplicationTask
func (c *Client) CreateReplicationTaskRequest(input *CreateReplicationTaskInput) CreateReplicationTaskRequest {
	op := &aws.Operation{
		Name:       opCreateReplicationTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateReplicationTaskInput{}
	}

	req := c.newRequest(op, input, &CreateReplicationTaskOutput{})
	return CreateReplicationTaskRequest{Request: req, Input: input, Copy: c.CreateReplicationTaskRequest}
}

// CreateReplicationTaskRequest is the request type for the
// CreateReplicationTask API operation.
type CreateReplicationTaskRequest struct {
	*aws.Request
	Input *CreateReplicationTaskInput
	Copy  func(*CreateReplicationTaskInput) CreateReplicationTaskRequest
}

// Send marshals and sends the CreateReplicationTask API request.
func (r CreateReplicationTaskRequest) Send(ctx context.Context) (*CreateReplicationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReplicationTaskResponse{
		CreateReplicationTaskOutput: r.Request.Data.(*CreateReplicationTaskOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReplicationTaskResponse is the response type for the
// CreateReplicationTask API operation.
type CreateReplicationTaskResponse struct {
	*CreateReplicationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReplicationTask request.
func (r *CreateReplicationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
