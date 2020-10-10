// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove cost allocation tags from the specified Amazon SQS queue. For an
// overview, see Tagging Your Amazon SQS Queues
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
// in the Amazon Simple Queue Service Developer Guide. Cross-account permissions
// don't apply to this action. For more information, see Grant Cross-Account
// Permissions to a Role and a User Name
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
func (c *Client) UntagQueue(ctx context.Context, params *UntagQueueInput, optFns ...func(*Options)) (*UntagQueueOutput, error) {
	if params == nil {
		params = &UntagQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UntagQueue", params, optFns, addOperationUntagQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UntagQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagQueueInput struct {

	// The URL of the queue.
	//
	// This member is required.
	QueueUrl *string

	// The list of tags to be removed from the specified queue.
	//
	// This member is required.
	TagKeys []*string
}

type UntagQueueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUntagQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUntagQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUntagQueue{}, middleware.After)
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
	addOpUntagQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUntagQueue(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUntagQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "UntagQueue",
	}
}
