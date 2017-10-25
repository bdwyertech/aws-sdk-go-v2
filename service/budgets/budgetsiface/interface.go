// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package budgetsiface provides an interface to enable mocking the AWS Budgets service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package budgetsiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
)

// BudgetsAPI provides an interface to enable mocking the
// budgets.Budgets service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Budgets.
//    func myFunc(svc budgetsiface.BudgetsAPI) bool {
//        // Make svc.CreateBudget request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := budgets.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockBudgetsClient struct {
//        budgetsiface.BudgetsAPI
//    }
//    func (m *mockBudgetsClient) CreateBudget(input *budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockBudgetsClient{}
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
type BudgetsAPI interface {
	CreateBudget(*budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error)
	CreateBudgetWithContext(aws.Context, *budgets.CreateBudgetInput, ...aws.Option) (*budgets.CreateBudgetOutput, error)
	CreateBudgetRequest(*budgets.CreateBudgetInput) (*aws.Request, *budgets.CreateBudgetOutput)

	CreateNotification(*budgets.CreateNotificationInput) (*budgets.CreateNotificationOutput, error)
	CreateNotificationWithContext(aws.Context, *budgets.CreateNotificationInput, ...aws.Option) (*budgets.CreateNotificationOutput, error)
	CreateNotificationRequest(*budgets.CreateNotificationInput) (*aws.Request, *budgets.CreateNotificationOutput)

	CreateSubscriber(*budgets.CreateSubscriberInput) (*budgets.CreateSubscriberOutput, error)
	CreateSubscriberWithContext(aws.Context, *budgets.CreateSubscriberInput, ...aws.Option) (*budgets.CreateSubscriberOutput, error)
	CreateSubscriberRequest(*budgets.CreateSubscriberInput) (*aws.Request, *budgets.CreateSubscriberOutput)

	DeleteBudget(*budgets.DeleteBudgetInput) (*budgets.DeleteBudgetOutput, error)
	DeleteBudgetWithContext(aws.Context, *budgets.DeleteBudgetInput, ...aws.Option) (*budgets.DeleteBudgetOutput, error)
	DeleteBudgetRequest(*budgets.DeleteBudgetInput) (*aws.Request, *budgets.DeleteBudgetOutput)

	DeleteNotification(*budgets.DeleteNotificationInput) (*budgets.DeleteNotificationOutput, error)
	DeleteNotificationWithContext(aws.Context, *budgets.DeleteNotificationInput, ...aws.Option) (*budgets.DeleteNotificationOutput, error)
	DeleteNotificationRequest(*budgets.DeleteNotificationInput) (*aws.Request, *budgets.DeleteNotificationOutput)

	DeleteSubscriber(*budgets.DeleteSubscriberInput) (*budgets.DeleteSubscriberOutput, error)
	DeleteSubscriberWithContext(aws.Context, *budgets.DeleteSubscriberInput, ...aws.Option) (*budgets.DeleteSubscriberOutput, error)
	DeleteSubscriberRequest(*budgets.DeleteSubscriberInput) (*aws.Request, *budgets.DeleteSubscriberOutput)

	DescribeBudget(*budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error)
	DescribeBudgetWithContext(aws.Context, *budgets.DescribeBudgetInput, ...aws.Option) (*budgets.DescribeBudgetOutput, error)
	DescribeBudgetRequest(*budgets.DescribeBudgetInput) (*aws.Request, *budgets.DescribeBudgetOutput)

	DescribeBudgets(*budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error)
	DescribeBudgetsWithContext(aws.Context, *budgets.DescribeBudgetsInput, ...aws.Option) (*budgets.DescribeBudgetsOutput, error)
	DescribeBudgetsRequest(*budgets.DescribeBudgetsInput) (*aws.Request, *budgets.DescribeBudgetsOutput)

	DescribeNotificationsForBudget(*budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error)
	DescribeNotificationsForBudgetWithContext(aws.Context, *budgets.DescribeNotificationsForBudgetInput, ...aws.Option) (*budgets.DescribeNotificationsForBudgetOutput, error)
	DescribeNotificationsForBudgetRequest(*budgets.DescribeNotificationsForBudgetInput) (*aws.Request, *budgets.DescribeNotificationsForBudgetOutput)

	DescribeSubscribersForNotification(*budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error)
	DescribeSubscribersForNotificationWithContext(aws.Context, *budgets.DescribeSubscribersForNotificationInput, ...aws.Option) (*budgets.DescribeSubscribersForNotificationOutput, error)
	DescribeSubscribersForNotificationRequest(*budgets.DescribeSubscribersForNotificationInput) (*aws.Request, *budgets.DescribeSubscribersForNotificationOutput)

	UpdateBudget(*budgets.UpdateBudgetInput) (*budgets.UpdateBudgetOutput, error)
	UpdateBudgetWithContext(aws.Context, *budgets.UpdateBudgetInput, ...aws.Option) (*budgets.UpdateBudgetOutput, error)
	UpdateBudgetRequest(*budgets.UpdateBudgetInput) (*aws.Request, *budgets.UpdateBudgetOutput)

	UpdateNotification(*budgets.UpdateNotificationInput) (*budgets.UpdateNotificationOutput, error)
	UpdateNotificationWithContext(aws.Context, *budgets.UpdateNotificationInput, ...aws.Option) (*budgets.UpdateNotificationOutput, error)
	UpdateNotificationRequest(*budgets.UpdateNotificationInput) (*aws.Request, *budgets.UpdateNotificationOutput)

	UpdateSubscriber(*budgets.UpdateSubscriberInput) (*budgets.UpdateSubscriberOutput, error)
	UpdateSubscriberWithContext(aws.Context, *budgets.UpdateSubscriberInput, ...aws.Option) (*budgets.UpdateSubscriberOutput, error)
	UpdateSubscriberRequest(*budgets.UpdateSubscriberInput) (*aws.Request, *budgets.UpdateSubscriberOutput)
}

var _ BudgetsAPI = (*budgets.Budgets)(nil)
