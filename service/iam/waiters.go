// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilInstanceProfileExists uses the IAM API operation
// GetInstanceProfile to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *IAM) WaitUntilInstanceProfileExists(input *GetInstanceProfileInput) error {
	return c.WaitUntilInstanceProfileExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceProfileExistsWithContext is an extended version of WaitUntilInstanceProfileExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) WaitUntilInstanceProfileExistsWithContext(ctx aws.Context, input *GetInstanceProfileInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInstanceProfileExists",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(1 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 404,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *GetInstanceProfileInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetInstanceProfileRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilUserExists uses the IAM API operation
// GetUser to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *IAM) WaitUntilUserExists(input *GetUserInput) error {
	return c.WaitUntilUserExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilUserExistsWithContext is an extended version of WaitUntilUserExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) WaitUntilUserExistsWithContext(ctx aws.Context, input *GetUserInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilUserExists",
		MaxAttempts: 20,
		Delay:       aws.ConstantWaiterDelay(1 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "NoSuchEntity",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *GetUserInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetUserRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
