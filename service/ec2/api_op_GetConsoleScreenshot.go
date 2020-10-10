// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve a JPG-format screenshot of a running instance to help with
// troubleshooting. The returned content is Base64-encoded.
func (c *Client) GetConsoleScreenshot(ctx context.Context, params *GetConsoleScreenshotInput, optFns ...func(*Options)) (*GetConsoleScreenshotOutput, error) {
	if params == nil {
		params = &GetConsoleScreenshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConsoleScreenshot", params, optFns, addOperationGetConsoleScreenshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConsoleScreenshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConsoleScreenshotInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// When set to true, acts as keystroke input and wakes up an instance that's in
	// standby or "sleep" mode.
	WakeUp *bool
}

type GetConsoleScreenshotOutput struct {

	// The data that comprises the image.
	ImageData *string

	// The ID of the instance.
	InstanceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetConsoleScreenshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetConsoleScreenshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetConsoleScreenshot{}, middleware.After)
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
	addOpGetConsoleScreenshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetConsoleScreenshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetConsoleScreenshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetConsoleScreenshot",
	}
}
