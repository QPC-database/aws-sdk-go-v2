// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation uses unions for inputs and outputs.
func (c *Client) JsonUnions(ctx context.Context, params *JsonUnionsInput, optFns ...func(*Options)) (*JsonUnionsOutput, error) {
	if params == nil {
		params = &JsonUnionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "JsonUnions", params, optFns, addOperationJsonUnionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*JsonUnionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A shared structure that contains a single union member.
type JsonUnionsInput struct {

	// A union with a representative set of types for members.
	Contents types.MyUnion
}

// A shared structure that contains a single union member.
type JsonUnionsOutput struct {

	// A union with a representative set of types for members.
	Contents types.MyUnion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationJsonUnionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpJsonUnions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpJsonUnions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opJsonUnions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opJsonUnions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "foo",
		OperationName: "JsonUnions",
	}
}
