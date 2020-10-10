// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a subscription. If the subscription requires authentication for
// deletion, only the owner of the subscription or the topic's owner can
// unsubscribe, and an AWS signature is required. If the Unsubscribe call does not
// require authentication and the requester is not the subscription owner, a final
// cancellation message is delivered to the endpoint, so that the endpoint owner
// can easily resubscribe to the topic if the Unsubscribe request was unintended.
// This action is throttled at 100 transactions per second (TPS).
func (c *Client) Unsubscribe(ctx context.Context, params *UnsubscribeInput, optFns ...func(*Options)) (*UnsubscribeOutput, error) {
	if params == nil {
		params = &UnsubscribeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Unsubscribe", params, optFns, addOperationUnsubscribeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnsubscribeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for Unsubscribe action.
type UnsubscribeInput struct {

	// The ARN of the subscription to be deleted.
	//
	// This member is required.
	SubscriptionArn *string
}

type UnsubscribeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUnsubscribeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUnsubscribe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUnsubscribe{}, middleware.After)
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
	addOpUnsubscribeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUnsubscribe(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUnsubscribe(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "Unsubscribe",
	}
}
