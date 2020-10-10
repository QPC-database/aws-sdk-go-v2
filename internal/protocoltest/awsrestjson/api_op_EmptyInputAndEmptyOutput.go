// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The example tests how requests and responses are serialized when there's no
// request or response payload because the operation has an empty input and empty
// output structure that reuses the same shape. While this should be rare, code
// generators must support this.
func (c *Client) EmptyInputAndEmptyOutput(ctx context.Context, params *EmptyInputAndEmptyOutputInput, optFns ...func(*Options)) (*EmptyInputAndEmptyOutputOutput, error) {
	if params == nil {
		params = &EmptyInputAndEmptyOutputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EmptyInputAndEmptyOutput", params, optFns, addOperationEmptyInputAndEmptyOutputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EmptyInputAndEmptyOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EmptyInputAndEmptyOutputInput struct {
}

type EmptyInputAndEmptyOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEmptyInputAndEmptyOutputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEmptyInputAndEmptyOutput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEmptyInputAndEmptyOutput{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opEmptyInputAndEmptyOutput(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEmptyInputAndEmptyOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EmptyInputAndEmptyOutput",
	}
}
