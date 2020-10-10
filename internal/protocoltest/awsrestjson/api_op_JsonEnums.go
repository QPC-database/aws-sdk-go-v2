// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes enums as top level properties, in lists, sets, and maps.
func (c *Client) JsonEnums(ctx context.Context, params *JsonEnumsInput, optFns ...func(*Options)) (*JsonEnumsOutput, error) {
	if params == nil {
		params = &JsonEnumsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "JsonEnums", params, optFns, addOperationJsonEnumsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*JsonEnumsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type JsonEnumsInput struct {
	FooEnum1 types.FooEnum

	FooEnum2 types.FooEnum

	FooEnum3 types.FooEnum

	FooEnumList []types.FooEnum

	FooEnumMap map[string]types.FooEnum

	FooEnumSet []types.FooEnum
}

type JsonEnumsOutput struct {
	FooEnum1 types.FooEnum

	FooEnum2 types.FooEnum

	FooEnum3 types.FooEnum

	FooEnumList []types.FooEnum

	FooEnumMap map[string]types.FooEnum

	FooEnumSet []types.FooEnum

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationJsonEnumsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpJsonEnums{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpJsonEnums{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opJsonEnums(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opJsonEnums(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "JsonEnums",
	}
}
