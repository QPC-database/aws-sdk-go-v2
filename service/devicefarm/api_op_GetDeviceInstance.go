// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a device instance that belongs to a private device
// fleet.
func (c *Client) GetDeviceInstance(ctx context.Context, params *GetDeviceInstanceInput, optFns ...func(*Options)) (*GetDeviceInstanceOutput, error) {
	if params == nil {
		params = &GetDeviceInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeviceInstance", params, optFns, addOperationGetDeviceInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeviceInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeviceInstanceInput struct {

	// The Amazon Resource Name (ARN) of the instance you're requesting information
	// about.
	//
	// This member is required.
	Arn *string
}

type GetDeviceInstanceOutput struct {

	// An object that contains information about your device instance.
	DeviceInstance *types.DeviceInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDeviceInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeviceInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeviceInstance{}, middleware.After)
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
	addOpGetDeviceInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeviceInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDeviceInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "GetDeviceInstance",
	}
}
