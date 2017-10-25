// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package configserviceiface provides an interface to enable mocking the AWS Config service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package configserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

// ConfigServiceAPI provides an interface to enable mocking the
// configservice.ConfigService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Config.
//    func myFunc(svc configserviceiface.ConfigServiceAPI) bool {
//        // Make svc.DeleteConfigRule request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := configservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockConfigServiceClient struct {
//        configserviceiface.ConfigServiceAPI
//    }
//    func (m *mockConfigServiceClient) DeleteConfigRule(input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockConfigServiceClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ConfigServiceAPI interface {
	DeleteConfigRule(*configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error)
	DeleteConfigRuleWithContext(aws.Context, *configservice.DeleteConfigRuleInput, ...aws.Option) (*configservice.DeleteConfigRuleOutput, error)
	DeleteConfigRuleRequest(*configservice.DeleteConfigRuleInput) (*aws.Request, *configservice.DeleteConfigRuleOutput)

	DeleteConfigurationRecorder(*configservice.DeleteConfigurationRecorderInput) (*configservice.DeleteConfigurationRecorderOutput, error)
	DeleteConfigurationRecorderWithContext(aws.Context, *configservice.DeleteConfigurationRecorderInput, ...aws.Option) (*configservice.DeleteConfigurationRecorderOutput, error)
	DeleteConfigurationRecorderRequest(*configservice.DeleteConfigurationRecorderInput) (*aws.Request, *configservice.DeleteConfigurationRecorderOutput)

	DeleteDeliveryChannel(*configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error)
	DeleteDeliveryChannelWithContext(aws.Context, *configservice.DeleteDeliveryChannelInput, ...aws.Option) (*configservice.DeleteDeliveryChannelOutput, error)
	DeleteDeliveryChannelRequest(*configservice.DeleteDeliveryChannelInput) (*aws.Request, *configservice.DeleteDeliveryChannelOutput)

	DeleteEvaluationResults(*configservice.DeleteEvaluationResultsInput) (*configservice.DeleteEvaluationResultsOutput, error)
	DeleteEvaluationResultsWithContext(aws.Context, *configservice.DeleteEvaluationResultsInput, ...aws.Option) (*configservice.DeleteEvaluationResultsOutput, error)
	DeleteEvaluationResultsRequest(*configservice.DeleteEvaluationResultsInput) (*aws.Request, *configservice.DeleteEvaluationResultsOutput)

	DeliverConfigSnapshot(*configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error)
	DeliverConfigSnapshotWithContext(aws.Context, *configservice.DeliverConfigSnapshotInput, ...aws.Option) (*configservice.DeliverConfigSnapshotOutput, error)
	DeliverConfigSnapshotRequest(*configservice.DeliverConfigSnapshotInput) (*aws.Request, *configservice.DeliverConfigSnapshotOutput)

	DescribeComplianceByConfigRule(*configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error)
	DescribeComplianceByConfigRuleWithContext(aws.Context, *configservice.DescribeComplianceByConfigRuleInput, ...aws.Option) (*configservice.DescribeComplianceByConfigRuleOutput, error)
	DescribeComplianceByConfigRuleRequest(*configservice.DescribeComplianceByConfigRuleInput) (*aws.Request, *configservice.DescribeComplianceByConfigRuleOutput)

	DescribeComplianceByResource(*configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error)
	DescribeComplianceByResourceWithContext(aws.Context, *configservice.DescribeComplianceByResourceInput, ...aws.Option) (*configservice.DescribeComplianceByResourceOutput, error)
	DescribeComplianceByResourceRequest(*configservice.DescribeComplianceByResourceInput) (*aws.Request, *configservice.DescribeComplianceByResourceOutput)

	DescribeConfigRuleEvaluationStatus(*configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)
	DescribeConfigRuleEvaluationStatusWithContext(aws.Context, *configservice.DescribeConfigRuleEvaluationStatusInput, ...aws.Option) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)
	DescribeConfigRuleEvaluationStatusRequest(*configservice.DescribeConfigRuleEvaluationStatusInput) (*aws.Request, *configservice.DescribeConfigRuleEvaluationStatusOutput)

	DescribeConfigRules(*configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error)
	DescribeConfigRulesWithContext(aws.Context, *configservice.DescribeConfigRulesInput, ...aws.Option) (*configservice.DescribeConfigRulesOutput, error)
	DescribeConfigRulesRequest(*configservice.DescribeConfigRulesInput) (*aws.Request, *configservice.DescribeConfigRulesOutput)

	DescribeConfigurationRecorderStatus(*configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
	DescribeConfigurationRecorderStatusWithContext(aws.Context, *configservice.DescribeConfigurationRecorderStatusInput, ...aws.Option) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
	DescribeConfigurationRecorderStatusRequest(*configservice.DescribeConfigurationRecorderStatusInput) (*aws.Request, *configservice.DescribeConfigurationRecorderStatusOutput)

	DescribeConfigurationRecorders(*configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error)
	DescribeConfigurationRecordersWithContext(aws.Context, *configservice.DescribeConfigurationRecordersInput, ...aws.Option) (*configservice.DescribeConfigurationRecordersOutput, error)
	DescribeConfigurationRecordersRequest(*configservice.DescribeConfigurationRecordersInput) (*aws.Request, *configservice.DescribeConfigurationRecordersOutput)

	DescribeDeliveryChannelStatus(*configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error)
	DescribeDeliveryChannelStatusWithContext(aws.Context, *configservice.DescribeDeliveryChannelStatusInput, ...aws.Option) (*configservice.DescribeDeliveryChannelStatusOutput, error)
	DescribeDeliveryChannelStatusRequest(*configservice.DescribeDeliveryChannelStatusInput) (*aws.Request, *configservice.DescribeDeliveryChannelStatusOutput)

	DescribeDeliveryChannels(*configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error)
	DescribeDeliveryChannelsWithContext(aws.Context, *configservice.DescribeDeliveryChannelsInput, ...aws.Option) (*configservice.DescribeDeliveryChannelsOutput, error)
	DescribeDeliveryChannelsRequest(*configservice.DescribeDeliveryChannelsInput) (*aws.Request, *configservice.DescribeDeliveryChannelsOutput)

	GetComplianceDetailsByConfigRule(*configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)
	GetComplianceDetailsByConfigRuleWithContext(aws.Context, *configservice.GetComplianceDetailsByConfigRuleInput, ...aws.Option) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)
	GetComplianceDetailsByConfigRuleRequest(*configservice.GetComplianceDetailsByConfigRuleInput) (*aws.Request, *configservice.GetComplianceDetailsByConfigRuleOutput)

	GetComplianceDetailsByResource(*configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error)
	GetComplianceDetailsByResourceWithContext(aws.Context, *configservice.GetComplianceDetailsByResourceInput, ...aws.Option) (*configservice.GetComplianceDetailsByResourceOutput, error)
	GetComplianceDetailsByResourceRequest(*configservice.GetComplianceDetailsByResourceInput) (*aws.Request, *configservice.GetComplianceDetailsByResourceOutput)

	GetComplianceSummaryByConfigRule(*configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)
	GetComplianceSummaryByConfigRuleWithContext(aws.Context, *configservice.GetComplianceSummaryByConfigRuleInput, ...aws.Option) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)
	GetComplianceSummaryByConfigRuleRequest(*configservice.GetComplianceSummaryByConfigRuleInput) (*aws.Request, *configservice.GetComplianceSummaryByConfigRuleOutput)

	GetComplianceSummaryByResourceType(*configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)
	GetComplianceSummaryByResourceTypeWithContext(aws.Context, *configservice.GetComplianceSummaryByResourceTypeInput, ...aws.Option) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)
	GetComplianceSummaryByResourceTypeRequest(*configservice.GetComplianceSummaryByResourceTypeInput) (*aws.Request, *configservice.GetComplianceSummaryByResourceTypeOutput)

	GetDiscoveredResourceCounts(*configservice.GetDiscoveredResourceCountsInput) (*configservice.GetDiscoveredResourceCountsOutput, error)
	GetDiscoveredResourceCountsWithContext(aws.Context, *configservice.GetDiscoveredResourceCountsInput, ...aws.Option) (*configservice.GetDiscoveredResourceCountsOutput, error)
	GetDiscoveredResourceCountsRequest(*configservice.GetDiscoveredResourceCountsInput) (*aws.Request, *configservice.GetDiscoveredResourceCountsOutput)

	GetResourceConfigHistory(*configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error)
	GetResourceConfigHistoryWithContext(aws.Context, *configservice.GetResourceConfigHistoryInput, ...aws.Option) (*configservice.GetResourceConfigHistoryOutput, error)
	GetResourceConfigHistoryRequest(*configservice.GetResourceConfigHistoryInput) (*aws.Request, *configservice.GetResourceConfigHistoryOutput)

	GetResourceConfigHistoryPages(*configservice.GetResourceConfigHistoryInput, func(*configservice.GetResourceConfigHistoryOutput, bool) bool) error
	GetResourceConfigHistoryPagesWithContext(aws.Context, *configservice.GetResourceConfigHistoryInput, func(*configservice.GetResourceConfigHistoryOutput, bool) bool, ...aws.Option) error

	ListDiscoveredResources(*configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesWithContext(aws.Context, *configservice.ListDiscoveredResourcesInput, ...aws.Option) (*configservice.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesRequest(*configservice.ListDiscoveredResourcesInput) (*aws.Request, *configservice.ListDiscoveredResourcesOutput)

	PutConfigRule(*configservice.PutConfigRuleInput) (*configservice.PutConfigRuleOutput, error)
	PutConfigRuleWithContext(aws.Context, *configservice.PutConfigRuleInput, ...aws.Option) (*configservice.PutConfigRuleOutput, error)
	PutConfigRuleRequest(*configservice.PutConfigRuleInput) (*aws.Request, *configservice.PutConfigRuleOutput)

	PutConfigurationRecorder(*configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error)
	PutConfigurationRecorderWithContext(aws.Context, *configservice.PutConfigurationRecorderInput, ...aws.Option) (*configservice.PutConfigurationRecorderOutput, error)
	PutConfigurationRecorderRequest(*configservice.PutConfigurationRecorderInput) (*aws.Request, *configservice.PutConfigurationRecorderOutput)

	PutDeliveryChannel(*configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error)
	PutDeliveryChannelWithContext(aws.Context, *configservice.PutDeliveryChannelInput, ...aws.Option) (*configservice.PutDeliveryChannelOutput, error)
	PutDeliveryChannelRequest(*configservice.PutDeliveryChannelInput) (*aws.Request, *configservice.PutDeliveryChannelOutput)

	PutEvaluations(*configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error)
	PutEvaluationsWithContext(aws.Context, *configservice.PutEvaluationsInput, ...aws.Option) (*configservice.PutEvaluationsOutput, error)
	PutEvaluationsRequest(*configservice.PutEvaluationsInput) (*aws.Request, *configservice.PutEvaluationsOutput)

	StartConfigRulesEvaluation(*configservice.StartConfigRulesEvaluationInput) (*configservice.StartConfigRulesEvaluationOutput, error)
	StartConfigRulesEvaluationWithContext(aws.Context, *configservice.StartConfigRulesEvaluationInput, ...aws.Option) (*configservice.StartConfigRulesEvaluationOutput, error)
	StartConfigRulesEvaluationRequest(*configservice.StartConfigRulesEvaluationInput) (*aws.Request, *configservice.StartConfigRulesEvaluationOutput)

	StartConfigurationRecorder(*configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error)
	StartConfigurationRecorderWithContext(aws.Context, *configservice.StartConfigurationRecorderInput, ...aws.Option) (*configservice.StartConfigurationRecorderOutput, error)
	StartConfigurationRecorderRequest(*configservice.StartConfigurationRecorderInput) (*aws.Request, *configservice.StartConfigurationRecorderOutput)

	StopConfigurationRecorder(*configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error)
	StopConfigurationRecorderWithContext(aws.Context, *configservice.StopConfigurationRecorderInput, ...aws.Option) (*configservice.StopConfigurationRecorderOutput, error)
	StopConfigurationRecorderRequest(*configservice.StopConfigurationRecorderInput) (*aws.Request, *configservice.StopConfigurationRecorderOutput)
}

var _ ConfigServiceAPI = (*configservice.ConfigService)(nil)
