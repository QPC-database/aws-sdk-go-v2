// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example ensures that mediaType strings are base64 encoded in headers.
func (c *Client) MediaTypeHeader(ctx context.Context, params *MediaTypeHeaderInput, optFns ...func(*Options)) (*MediaTypeHeaderOutput, error) {
	if params == nil {
		params = &MediaTypeHeaderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MediaTypeHeader", params, optFns, addOperationMediaTypeHeaderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MediaTypeHeaderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MediaTypeHeaderInput struct {
	Json *string
}

type MediaTypeHeaderOutput struct {
	Json *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationMediaTypeHeaderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpMediaTypeHeader{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpMediaTypeHeader{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opMediaTypeHeader(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opMediaTypeHeader(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "MediaTypeHeader",
	}
}
