// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the given thing from the billing group.
func (c *Client) RemoveThingFromBillingGroup(ctx context.Context, params *RemoveThingFromBillingGroupInput, optFns ...func(*Options)) (*RemoveThingFromBillingGroupOutput, error) {
	if params == nil {
		params = &RemoveThingFromBillingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveThingFromBillingGroup", params, optFns, addOperationRemoveThingFromBillingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveThingFromBillingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveThingFromBillingGroupInput struct {

	// The ARN of the billing group.
	BillingGroupArn *string

	// The name of the billing group.
	BillingGroupName *string

	// The ARN of the thing to be removed from the billing group.
	ThingArn *string

	// The name of the thing to be removed from the billing group.
	ThingName *string
}

type RemoveThingFromBillingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveThingFromBillingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemoveThingFromBillingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveThingFromBillingGroup{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveThingFromBillingGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemoveThingFromBillingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RemoveThingFromBillingGroup",
	}
}
