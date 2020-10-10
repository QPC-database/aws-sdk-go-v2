// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a thing to a billing group.
func (c *Client) AddThingToBillingGroup(ctx context.Context, params *AddThingToBillingGroupInput, optFns ...func(*Options)) (*AddThingToBillingGroupOutput, error) {
	if params == nil {
		params = &AddThingToBillingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddThingToBillingGroup", params, optFns, addOperationAddThingToBillingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddThingToBillingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddThingToBillingGroupInput struct {

	// The ARN of the billing group.
	BillingGroupArn *string

	// The name of the billing group.
	BillingGroupName *string

	// The ARN of the thing to be added to the billing group.
	ThingArn *string

	// The name of the thing to be added to the billing group.
	ThingName *string
}

type AddThingToBillingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddThingToBillingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddThingToBillingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddThingToBillingGroup{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddThingToBillingGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAddThingToBillingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "AddThingToBillingGroup",
	}
}
