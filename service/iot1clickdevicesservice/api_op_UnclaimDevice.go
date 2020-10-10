// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickdevicesservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a device from your AWS account using its device ID.
func (c *Client) UnclaimDevice(ctx context.Context, params *UnclaimDeviceInput, optFns ...func(*Options)) (*UnclaimDeviceOutput, error) {
	if params == nil {
		params = &UnclaimDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnclaimDevice", params, optFns, addOperationUnclaimDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnclaimDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnclaimDeviceInput struct {

	// The unique identifier of the device.
	//
	// This member is required.
	DeviceId *string
}

type UnclaimDeviceOutput struct {

	// The device's final claim state.
	State *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUnclaimDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUnclaimDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUnclaimDevice{}, middleware.After)
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
	addOpUnclaimDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUnclaimDevice(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUnclaimDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot1click",
		OperationName: "UnclaimDevice",
	}
}
