// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Revokes any permissions in the queue policy that matches the specified Label
// parameter.
//
//     * Only the owner of a queue can remove permissions from it.
//
//
// * Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
//     * To remove the ability
// to change queue permissions, you must deny permission to the AddPermission,
// RemovePermission, and SetQueueAttributes actions in your IAM policy.
func (c *Client) RemovePermission(ctx context.Context, params *RemovePermissionInput, optFns ...func(*Options)) (*RemovePermissionOutput, error) {
	if params == nil {
		params = &RemovePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemovePermission", params, optFns, addOperationRemovePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemovePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RemovePermissionInput struct {

	// The identification of the permission to remove. This is the label added using
	// the AddPermission () action.
	//
	// This member is required.
	Label *string

	// The URL of the Amazon SQS queue from which permissions are removed. Queue URLs
	// and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string
}

type RemovePermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemovePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRemovePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRemovePermission{}, middleware.After)
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
	addOpRemovePermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemovePermission(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemovePermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "RemovePermission",
	}
}
