// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// This examples serializes a streaming blob shape in the request body. In this
// example, no JSON document is synthesized because the payload is not a structure
// or a union type.
func (c *Client) StreamingTraits(ctx context.Context, params *StreamingTraitsInput, optFns ...func(*Options)) (*StreamingTraitsOutput, error) {
	if params == nil {
		params = &StreamingTraitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StreamingTraits", params, optFns, addOperationStreamingTraitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StreamingTraitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StreamingTraitsInput struct {
	Blob io.Reader

	Foo *string
}

type StreamingTraitsOutput struct {
	Blob io.ReadCloser

	Foo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStreamingTraitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStreamingTraits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStreamingTraits{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opStreamingTraits(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStreamingTraits(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StreamingTraits",
	}
}
