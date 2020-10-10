// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information on a specific signing platform.
func (c *Client) GetSigningPlatform(ctx context.Context, params *GetSigningPlatformInput, optFns ...func(*Options)) (*GetSigningPlatformOutput, error) {
	if params == nil {
		params = &GetSigningPlatformInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSigningPlatform", params, optFns, addOperationGetSigningPlatformMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSigningPlatformOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSigningPlatformInput struct {

	// The ID of the target signing platform.
	//
	// This member is required.
	PlatformId *string
}

type GetSigningPlatformOutput struct {

	// The category type of the target signing platform.
	Category types.Category

	// The display name of the target signing platform.
	DisplayName *string

	// The maximum size (in MB) of the payload that can be signed by the target
	// platform.
	MaxSizeInMB *int32

	// A list of partner entities that use the target signing platform.
	Partner *string

	// The ID of the target signing platform.
	PlatformId *string

	// A list of configurations applied to the target platform at signing.
	SigningConfiguration *types.SigningConfiguration

	// The format of the target platform's signing image.
	SigningImageFormat *types.SigningImageFormat

	// The validation template that is used by the target signing platform.
	Target *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSigningPlatformMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSigningPlatform{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSigningPlatform{}, middleware.After)
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
	addOpGetSigningPlatformValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSigningPlatform(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSigningPlatform(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "signer",
		OperationName: "GetSigningPlatform",
	}
}
