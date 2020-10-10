// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the password for your VM local console. When you log in to the local
// console for the first time, you log in to the VM with the default credentials.
// We recommend that you set a new password. You don't need to know the default
// password to set a new password.
func (c *Client) SetLocalConsolePassword(ctx context.Context, params *SetLocalConsolePasswordInput, optFns ...func(*Options)) (*SetLocalConsolePasswordOutput, error) {
	if params == nil {
		params = &SetLocalConsolePasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetLocalConsolePassword", params, optFns, addOperationSetLocalConsolePasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetLocalConsolePasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// SetLocalConsolePasswordInput
type SetLocalConsolePasswordInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string

	// The password you want to set for your VM local console.
	//
	// This member is required.
	LocalConsolePassword *string
}

type SetLocalConsolePasswordOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetLocalConsolePasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetLocalConsolePassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetLocalConsolePassword{}, middleware.After)
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
	addOpSetLocalConsolePasswordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetLocalConsolePassword(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetLocalConsolePassword(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "SetLocalConsolePassword",
	}
}
