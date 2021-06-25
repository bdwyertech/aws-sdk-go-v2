// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a friendly name for a customer master key (CMK). Adding, deleting, or
// updating an alias can allow or deny permission to the CMK. For details, see
// Using ABAC in AWS KMS
// (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the AWS Key
// Management Service Developer Guide. You can use an alias to identify a CMK in
// the AWS KMS console, in the DescribeKey operation and in cryptographic
// operations
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations),
// such as Encrypt and GenerateDataKey. You can also change the CMK that's
// associated with the alias (UpdateAlias) or delete the alias (DeleteAlias) at any
// time. These operations don't affect the underlying CMK. You can associate the
// alias with any customer managed CMK in the same AWS Region. Each alias is
// associated with only one CMK at a time, but a CMK can have multiple aliases. A
// valid CMK is required. You can't create an alias without a CMK. The alias must
// be unique in the account and Region, but you can have aliases with the same name
// in different Regions. For detailed information about aliases, see Using aliases
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html) in the
// AWS Key Management Service Developer Guide. This operation does not return a
// response. To get the alias that you created, use the ListAliases operation. The
// CMK that you use for this operation must be in a compatible key state. For
// details, see Key state: Effect on your CMK
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on an alias in a different AWS account. Required
// permissions
//
// * kms:CreateAlias
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// on the alias (IAM policy).
//
// * kms:CreateAlias
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// on the CMK (key policy).
//
// For details, see Controlling access to aliases
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html#alias-access)
// in the AWS Key Management Service Developer Guide. Related operations:
//
// *
// DeleteAlias
//
// * ListAliases
//
// * UpdateAlias
func (c *Client) CreateAlias(ctx context.Context, params *CreateAliasInput, optFns ...func(*Options)) (*CreateAliasOutput, error) {
	if params == nil {
		params = &CreateAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAlias", params, optFns, c.addOperationCreateAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAliasInput struct {

	// Specifies the alias name. This value must begin with alias/ followed by a name,
	// such as alias/ExampleAlias. The AliasName value must be string of 1-256
	// characters. It can contain only alphanumeric characters, forward slashes (/),
	// underscores (_), and dashes (-). The alias name cannot begin with alias/aws/.
	// The alias/aws/ prefix is reserved for AWS managed CMKs
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
	//
	// This member is required.
	AliasName *string

	// Associates the alias with the specified customer managed CMK
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk).
	// The CMK must be in the same AWS Region. A valid CMK ID is required. If you
	// supply a null or empty string value, this operation returns an error. For help
	// finding the key ID and ARN, see Finding the Key ID and ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn)
	// in the AWS Key Management Service Developer Guide. Specify the key ID or key ARN
	// of the CMK. For example:
	//
	// * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key
	// ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// This member is required.
	TargetKeyId *string
}

type CreateAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAlias(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "CreateAlias",
	}
}
