// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// The error that has occurred when attempting to retrieve a batch of Records.
type BatchGetRecordError struct {

	// The error code of an error that has occured when attempting to retrieve a batch
	// of Records. For more information on errors, see  Errors
	// (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_feature_store_GetRecord.html#API_feature_store_GetRecord_Errors).
	//
	// This member is required.
	ErrorCode *string

	// The error message of an error that has occured when attempting to retrieve a
	// record in the batch.
	//
	// This member is required.
	ErrorMessage *string

	// The name of the feature group that the record belongs to.
	//
	// This member is required.
	FeatureGroupName *string

	// The value for the RecordIdentifier in string format of a Record from a
	// FeatureGroup that is causing an error when attempting to be retrieved.
	//
	// This member is required.
	RecordIdentifierValueAsString *string
}

// The identifier that identifies the batch of Records you are retrieving in a
// batch.
type BatchGetRecordIdentifier struct {

	// A FeatureGroupName containing Records you are retrieving in a batch.
	//
	// This member is required.
	FeatureGroupName *string

	// The value for a list of record identifiers in string format.
	//
	// This member is required.
	RecordIdentifiersValueAsString []string

	// List of names of Features to be retrieved. If not specified, the latest value
	// for all the Features are returned.
	FeatureNames []string
}

// The output of Records that have been retrieved in a batch.
type BatchGetRecordResultDetail struct {

	// The FeatureGroupName containing Records you retrieved in a batch.
	//
	// This member is required.
	FeatureGroupName *string

	// The Record retrieved.
	//
	// This member is required.
	Record []FeatureValue

	// The value of the record identifer in string format.
	//
	// This member is required.
	RecordIdentifierValueAsString *string
}

// The value associated with a feature.
type FeatureValue struct {

	// The name of a feature that a feature value corresponds to.
	//
	// This member is required.
	FeatureName *string

	// The value associated with a feature, in string format. Note that features types
	// can be String, Integral, or Fractional. This value represents all three types as
	// a string.
	//
	// This member is required.
	ValueAsString *string
}
