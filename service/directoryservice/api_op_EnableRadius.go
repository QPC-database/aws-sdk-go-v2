// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables multi-factor authentication (MFA) with the Remote Authentication Dial In
// User Service (RADIUS) server for an AD Connector or Microsoft AD directory.
func (c *Client) EnableRadius(ctx context.Context, params *EnableRadiusInput, optFns ...func(*Options)) (*EnableRadiusOutput, error) {
	if params == nil {
		params = &EnableRadiusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableRadius", params, optFns, addOperationEnableRadiusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableRadiusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the EnableRadius () operation.
type EnableRadiusInput struct {

	// The identifier of the directory for which to enable MFA.
	//
	// This member is required.
	DirectoryId *string

	// A RadiusSettings () object that contains information about the RADIUS server.
	//
	// This member is required.
	RadiusSettings *types.RadiusSettings
}

// Contains the results of the EnableRadius () operation.
type EnableRadiusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableRadiusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableRadius{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableRadius{}, middleware.After)
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
	addOpEnableRadiusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableRadius(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEnableRadius(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "EnableRadius",
	}
}
