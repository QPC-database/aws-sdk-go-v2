// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the specified directory. Disabled directories cannot be read or written
// to. Only enabled directories can be disabled. Disabled directories may be
// reenabled.
func (c *Client) DisableDirectory(ctx context.Context, params *DisableDirectoryInput, optFns ...func(*Options)) (*DisableDirectoryOutput, error) {
	if params == nil {
		params = &DisableDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableDirectory", params, optFns, addOperationDisableDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableDirectoryInput struct {

	// The ARN of the directory to disable.
	//
	// This member is required.
	DirectoryArn *string
}

type DisableDirectoryOutput struct {

	// The ARN of the directory that has been disabled.
	//
	// This member is required.
	DirectoryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisableDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisableDirectory{}, middleware.After)
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
	addOpDisableDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableDirectory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisableDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "DisableDirectory",
	}
}
