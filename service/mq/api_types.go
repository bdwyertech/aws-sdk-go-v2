// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Name of the availability zone.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/AvailabilityZone
type AvailabilityZone struct {
	_ struct{} `type:"structure"`

	// Id for the availability zone.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s AvailabilityZone) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AvailabilityZone) MarshalFields(e protocol.FieldEncoder) error {
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Types of broker engines.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/BrokerEngineType
type BrokerEngineType struct {
	_ struct{} `type:"structure"`

	// The type of broker engine.
	EngineType EngineType `locationName:"engineType" type:"string" enum:"true"`

	// The list of engine versions.
	EngineVersions []EngineVersion `locationName:"engineVersions" type:"list"`
}

// String returns the string representation
func (s BrokerEngineType) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BrokerEngineType) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.EngineType) > 0 {
		v := s.EngineType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "engineType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.EngineVersions != nil {
		v := s.EngineVersions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "engineVersions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Returns information about all brokers.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/BrokerInstance
type BrokerInstance struct {
	_ struct{} `type:"structure"`

	// The URL of the broker's ActiveMQ Web Console.
	ConsoleURL *string `locationName:"consoleURL" type:"string"`

	// The broker's wire-level protocol endpoints.
	Endpoints []string `locationName:"endpoints" type:"list"`

	// The IP address of the Elastic Network Interface (ENI) attached to the broker.
	IpAddress *string `locationName:"ipAddress" type:"string"`
}

// String returns the string representation
func (s BrokerInstance) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BrokerInstance) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConsoleURL != nil {
		v := *s.ConsoleURL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "consoleURL", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Endpoints != nil {
		v := s.Endpoints

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "endpoints", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.IpAddress != nil {
		v := *s.IpAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ipAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Option for host instance type.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/BrokerInstanceOption
type BrokerInstanceOption struct {
	_ struct{} `type:"structure"`

	// The list of available az.
	AvailabilityZones []AvailabilityZone `locationName:"availabilityZones" type:"list"`

	// The type of broker engine.
	EngineType EngineType `locationName:"engineType" type:"string" enum:"true"`

	// The type of broker instance.
	HostInstanceType *string `locationName:"hostInstanceType" type:"string"`

	// The list of supported engine versions.
	SupportedEngineVersions []string `locationName:"supportedEngineVersions" type:"list"`
}

// String returns the string representation
func (s BrokerInstanceOption) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BrokerInstanceOption) MarshalFields(e protocol.FieldEncoder) error {
	if s.AvailabilityZones != nil {
		v := s.AvailabilityZones

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "availabilityZones", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.EngineType) > 0 {
		v := s.EngineType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "engineType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.HostInstanceType != nil {
		v := *s.HostInstanceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "hostInstanceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SupportedEngineVersions != nil {
		v := s.SupportedEngineVersions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "supportedEngineVersions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// The Amazon Resource Name (ARN) of the broker.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/BrokerSummary
type BrokerSummary struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the broker.
	BrokerArn *string `locationName:"brokerArn" type:"string"`

	// The unique ID that Amazon MQ generates for the broker.
	BrokerId *string `locationName:"brokerId" type:"string"`

	// The name of the broker. This value must be unique in your AWS account, 1-50
	// characters long, must contain only letters, numbers, dashes, and underscores,
	// and must not contain whitespaces, brackets, wildcard characters, or special
	// characters.
	BrokerName *string `locationName:"brokerName" type:"string"`

	// The status of the broker.
	BrokerState BrokerState `locationName:"brokerState" type:"string" enum:"true"`

	// The time when the broker was created.
	Created *time.Time `locationName:"created" type:"timestamp" timestampFormat:"unix"`

	// Required. The deployment mode of the broker.
	DeploymentMode DeploymentMode `locationName:"deploymentMode" type:"string" enum:"true"`

	// The broker's instance type.
	HostInstanceType *string `locationName:"hostInstanceType" type:"string"`
}

// String returns the string representation
func (s BrokerSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BrokerSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.BrokerArn != nil {
		v := *s.BrokerArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BrokerName != nil {
		v := *s.BrokerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.BrokerState) > 0 {
		v := s.BrokerState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Created != nil {
		v := *s.Created

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "created", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if len(s.DeploymentMode) > 0 {
		v := s.DeploymentMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentMode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.HostInstanceType != nil {
		v := *s.HostInstanceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "hostInstanceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Returns information about all configurations.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/Configuration
type Configuration struct {
	_ struct{} `type:"structure"`

	// Required. The ARN of the configuration.
	Arn *string `locationName:"arn" type:"string"`

	// Required. The date and time of the configuration revision.
	Created *time.Time `locationName:"created" type:"timestamp" timestampFormat:"unix"`

	// Required. The description of the configuration.
	Description *string `locationName:"description" type:"string"`

	// Required. The type of broker engine. Note: Currently, Amazon MQ supports
	// only ACTIVEMQ.
	EngineType EngineType `locationName:"engineType" type:"string" enum:"true"`

	// Required. The version of the broker engine. For a list of supported engine
	// versions, see https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html
	EngineVersion *string `locationName:"engineVersion" type:"string"`

	// Required. The unique ID that Amazon MQ generates for the configuration.
	Id *string `locationName:"id" type:"string"`

	// Required. The latest revision of the configuration.
	LatestRevision *ConfigurationRevision `locationName:"latestRevision" type:"structure"`

	// Required. The name of the configuration. This value can contain only alphanumeric
	// characters, dashes, periods, underscores, and tildes (- . _ ~). This value
	// must be 1-150 characters long.
	Name *string `locationName:"name" type:"string"`

	// The list of all tags associated with this configuration.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s Configuration) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Configuration) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Created != nil {
		v := *s.Created

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "created", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.EngineType) > 0 {
		v := s.EngineType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "engineType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.EngineVersion != nil {
		v := *s.EngineVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "engineVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestRevision != nil {
		v := s.LatestRevision

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "latestRevision", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// A list of information about the configuration.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/ConfigurationId
type ConfigurationId struct {
	_ struct{} `type:"structure"`

	// Required. The unique ID that Amazon MQ generates for the configuration.
	Id *string `locationName:"id" type:"string"`

	// The revision number of the configuration.
	Revision *int64 `locationName:"revision" type:"integer"`
}

// String returns the string representation
func (s ConfigurationId) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ConfigurationId) MarshalFields(e protocol.FieldEncoder) error {
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revision", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Returns information about the specified configuration revision.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/ConfigurationRevision
type ConfigurationRevision struct {
	_ struct{} `type:"structure"`

	// Required. The date and time of the configuration revision.
	Created *time.Time `locationName:"created" type:"timestamp" timestampFormat:"unix"`

	// The description of the configuration revision.
	Description *string `locationName:"description" type:"string"`

	// Required. The revision number of the configuration.
	Revision *int64 `locationName:"revision" type:"integer"`
}

// String returns the string representation
func (s ConfigurationRevision) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ConfigurationRevision) MarshalFields(e protocol.FieldEncoder) error {
	if s.Created != nil {
		v := *s.Created

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "created", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revision", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Broker configuration information
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/Configurations
type Configurations struct {
	_ struct{} `type:"structure"`

	// The current configuration of the broker.
	Current *ConfigurationId `locationName:"current" type:"structure"`

	// The history of configurations applied to the broker.
	History []ConfigurationId `locationName:"history" type:"list"`

	// The pending configuration of the broker.
	Pending *ConfigurationId `locationName:"pending" type:"structure"`
}

// String returns the string representation
func (s Configurations) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Configurations) MarshalFields(e protocol.FieldEncoder) error {
	if s.Current != nil {
		v := s.Current

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "current", v, metadata)
	}
	if s.History != nil {
		v := s.History

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "history", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Pending != nil {
		v := s.Pending

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "pending", v, metadata)
	}
	return nil
}

// Encryption options for the broker.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/EncryptionOptions
type EncryptionOptions struct {
	_ struct{} `type:"structure"`

	// The customer master key (CMK) to use for the AWS Key Management Service (KMS).
	// This key is used to encrypt your data at rest. If not provided, Amazon MQ
	// will use a default CMK to encrypt your data.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// Enables the use of an AWS owned CMK using AWS Key Management Service (KMS).
	//
	// UseAwsOwnedKey is a required field
	UseAwsOwnedKey *bool `locationName:"useAwsOwnedKey" type:"boolean" required:"true"`
}

// String returns the string representation
func (s EncryptionOptions) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EncryptionOptions) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EncryptionOptions"}

	if s.UseAwsOwnedKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("UseAwsOwnedKey"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EncryptionOptions) MarshalFields(e protocol.FieldEncoder) error {
	if s.KmsKeyId != nil {
		v := *s.KmsKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "kmsKeyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UseAwsOwnedKey != nil {
		v := *s.UseAwsOwnedKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "useAwsOwnedKey", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Id of the engine version.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/EngineVersion
type EngineVersion struct {
	_ struct{} `type:"structure"`

	// Id for the version.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s EngineVersion) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EngineVersion) MarshalFields(e protocol.FieldEncoder) error {
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The list of information about logs to be enabled for the specified broker.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/Logs
type Logs struct {
	_ struct{} `type:"structure"`

	// Enables audit logging. Every user management action made using JMX or the
	// ActiveMQ Web Console is logged.
	Audit *bool `locationName:"audit" type:"boolean"`

	// Enables general logging.
	General *bool `locationName:"general" type:"boolean"`
}

// String returns the string representation
func (s Logs) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Logs) MarshalFields(e protocol.FieldEncoder) error {
	if s.Audit != nil {
		v := *s.Audit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "audit", protocol.BoolValue(v), metadata)
	}
	if s.General != nil {
		v := *s.General

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "general", protocol.BoolValue(v), metadata)
	}
	return nil
}

// The list of information about logs currently enabled and pending to be deployed
// for the specified broker.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/LogsSummary
type LogsSummary struct {
	_ struct{} `type:"structure"`

	// Enables audit logging. Every user management action made using JMX or the
	// ActiveMQ Web Console is logged.
	Audit *bool `locationName:"audit" type:"boolean"`

	// The location of the CloudWatch Logs log group where audit logs are sent.
	AuditLogGroup *string `locationName:"auditLogGroup" type:"string"`

	// Enables general logging.
	General *bool `locationName:"general" type:"boolean"`

	// The location of the CloudWatch Logs log group where general logs are sent.
	GeneralLogGroup *string `locationName:"generalLogGroup" type:"string"`

	// The list of information about logs pending to be deployed for the specified
	// broker.
	Pending *PendingLogs `locationName:"pending" type:"structure"`
}

// String returns the string representation
func (s LogsSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LogsSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.Audit != nil {
		v := *s.Audit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "audit", protocol.BoolValue(v), metadata)
	}
	if s.AuditLogGroup != nil {
		v := *s.AuditLogGroup

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "auditLogGroup", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.General != nil {
		v := *s.General

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "general", protocol.BoolValue(v), metadata)
	}
	if s.GeneralLogGroup != nil {
		v := *s.GeneralLogGroup

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "generalLogGroup", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Pending != nil {
		v := s.Pending

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "pending", v, metadata)
	}
	return nil
}

// The list of information about logs to be enabled for the specified broker.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/PendingLogs
type PendingLogs struct {
	_ struct{} `type:"structure"`

	// Enables audit logging. Every user management action made using JMX or the
	// ActiveMQ Web Console is logged.
	Audit *bool `locationName:"audit" type:"boolean"`

	// Enables general logging.
	General *bool `locationName:"general" type:"boolean"`
}

// String returns the string representation
func (s PendingLogs) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PendingLogs) MarshalFields(e protocol.FieldEncoder) error {
	if s.Audit != nil {
		v := *s.Audit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "audit", protocol.BoolValue(v), metadata)
	}
	if s.General != nil {
		v := *s.General

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "general", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Returns information about the XML element or attribute that was sanitized
// in the configuration.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/SanitizationWarning
type SanitizationWarning struct {
	_ struct{} `type:"structure"`

	// The name of the XML attribute that has been sanitized.
	AttributeName *string `locationName:"attributeName" type:"string"`

	// The name of the XML element that has been sanitized.
	ElementName *string `locationName:"elementName" type:"string"`

	// Required. The reason for which the XML elements or attributes were sanitized.
	Reason SanitizationWarningReason `locationName:"reason" type:"string" enum:"true"`
}

// String returns the string representation
func (s SanitizationWarning) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SanitizationWarning) MarshalFields(e protocol.FieldEncoder) error {
	if s.AttributeName != nil {
		v := *s.AttributeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "attributeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ElementName != nil {
		v := *s.ElementName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "elementName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Reason) > 0 {
		v := s.Reason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "reason", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// An ActiveMQ user associated with the broker.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/User
type User struct {
	_ struct{} `type:"structure"`

	// Enables access to the the ActiveMQ Web Console for the ActiveMQ user.
	ConsoleAccess *bool `locationName:"consoleAccess" type:"boolean"`

	// The list of groups (20 maximum) to which the ActiveMQ user belongs. This
	// value can contain only alphanumeric characters, dashes, periods, underscores,
	// and tildes (- . _ ~). This value must be 2-100 characters long.
	Groups []string `locationName:"groups" type:"list"`

	// Required. The password of the ActiveMQ user. This value must be at least
	// 12 characters long, must contain at least 4 unique characters, and must not
	// contain commas.
	Password *string `locationName:"password" type:"string"`

	// Required. The username of the ActiveMQ user. This value can contain only
	// alphanumeric characters, dashes, periods, underscores, and tildes (- . _
	// ~). This value must be 2-100 characters long.
	Username *string `locationName:"username" type:"string"`
}

// String returns the string representation
func (s User) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s User) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConsoleAccess != nil {
		v := *s.ConsoleAccess

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "consoleAccess", protocol.BoolValue(v), metadata)
	}
	if s.Groups != nil {
		v := s.Groups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "groups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Password != nil {
		v := *s.Password

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "password", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Username != nil {
		v := *s.Username

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "username", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Returns information about the status of the changes pending for the ActiveMQ
// user.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/UserPendingChanges
type UserPendingChanges struct {
	_ struct{} `type:"structure"`

	// Enables access to the the ActiveMQ Web Console for the ActiveMQ user.
	ConsoleAccess *bool `locationName:"consoleAccess" type:"boolean"`

	// The list of groups (20 maximum) to which the ActiveMQ user belongs. This
	// value can contain only alphanumeric characters, dashes, periods, underscores,
	// and tildes (- . _ ~). This value must be 2-100 characters long.
	Groups []string `locationName:"groups" type:"list"`

	// Required. The type of change pending for the ActiveMQ user.
	PendingChange ChangeType `locationName:"pendingChange" type:"string" enum:"true"`
}

// String returns the string representation
func (s UserPendingChanges) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UserPendingChanges) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConsoleAccess != nil {
		v := *s.ConsoleAccess

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "consoleAccess", protocol.BoolValue(v), metadata)
	}
	if s.Groups != nil {
		v := s.Groups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "groups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.PendingChange) > 0 {
		v := s.PendingChange

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pendingChange", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Returns a list of all ActiveMQ users.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/UserSummary
type UserSummary struct {
	_ struct{} `type:"structure"`

	// The type of change pending for the ActiveMQ user.
	PendingChange ChangeType `locationName:"pendingChange" type:"string" enum:"true"`

	// Required. The username of the ActiveMQ user. This value can contain only
	// alphanumeric characters, dashes, periods, underscores, and tildes (- . _
	// ~). This value must be 2-100 characters long.
	Username *string `locationName:"username" type:"string"`
}

// String returns the string representation
func (s UserSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UserSummary) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.PendingChange) > 0 {
		v := s.PendingChange

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pendingChange", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Username != nil {
		v := *s.Username

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "username", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The scheduled time period relative to UTC during which Amazon MQ begins to
// apply pending updates or patches to the broker.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/WeeklyStartTime
type WeeklyStartTime struct {
	_ struct{} `type:"structure"`

	// Required. The day of the week.
	DayOfWeek DayOfWeek `locationName:"dayOfWeek" type:"string" enum:"true"`

	// Required. The time, in 24-hour format.
	TimeOfDay *string `locationName:"timeOfDay" type:"string"`

	// The time zone, UTC by default, in either the Country/City format, or the
	// UTC offset format.
	TimeZone *string `locationName:"timeZone" type:"string"`
}

// String returns the string representation
func (s WeeklyStartTime) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s WeeklyStartTime) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.DayOfWeek) > 0 {
		v := s.DayOfWeek

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dayOfWeek", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TimeOfDay != nil {
		v := *s.TimeOfDay

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "timeOfDay", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TimeZone != nil {
		v := *s.TimeZone

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "timeZone", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
