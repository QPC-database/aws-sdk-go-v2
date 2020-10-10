// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove the specified thing from the specified group. You must specify either a
// thingGroupArn or a thingGroupName to identify the thing group and either a
// thingArn or a thingName to identify the thing to remove from the thing group.
func (c *Client) RemoveThingFromThingGroup(ctx context.Context, params *RemoveThingFromThingGroupInput, optFns ...func(*Options)) (*RemoveThingFromThingGroupOutput, error) {
	if params == nil {
		params = &RemoveThingFromThingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveThingFromThingGroup", params, optFns, addOperationRemoveThingFromThingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveThingFromThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveThingFromThingGroupInput struct {

	// The ARN of the thing to remove from the group.
	ThingArn *string

	// The group ARN.
	ThingGroupArn *string

	// The group name.
	ThingGroupName *string

	// The name of the thing to remove from the group.
	ThingName *string
}

type RemoveThingFromThingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveThingFromThingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemoveThingFromThingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveThingFromThingGroup{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveThingFromThingGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemoveThingFromThingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RemoveThingFromThingGroup",
	}
}
