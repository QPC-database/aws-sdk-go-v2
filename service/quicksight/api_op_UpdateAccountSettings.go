// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the settings for the Amazon QuickSight subscription in your AWS Account.
func (c *Client) UpdateAccountSettings(ctx context.Context, params *UpdateAccountSettingsInput, optFns ...func(*Options)) (*UpdateAccountSettingsOutput, error) {
	if params == nil {
		params = &UpdateAccountSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccountSettings", params, optFns, addOperationUpdateAccountSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccountSettingsInput struct {

	// The ID for the AWS account that contains the QuickSight namespaces that you want
	// to list.
	//
	// This member is required.
	AwsAccountId *string

	// The default namespace for this AWS Account. Currently, the default is default.
	// IAM users who register for the first time with QuickSight provide an email that
	// becomes associated with the default namespace.
	//
	// This member is required.
	DefaultNamespace *string

	// Email address used to send notifications regarding administration of QuickSight.
	NotificationEmail *string
}

type UpdateAccountSettingsOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAccountSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAccountSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAccountSettings{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateAccountSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccountSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateAccountSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateAccountSettings",
	}
}
