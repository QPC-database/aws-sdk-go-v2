// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) SimpleScalarProperties(ctx context.Context, params *SimpleScalarPropertiesInput, optFns ...func(*Options)) (*SimpleScalarPropertiesOutput, error) {
	if params == nil {
		params = &SimpleScalarPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SimpleScalarProperties", params, optFns, addOperationSimpleScalarPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SimpleScalarPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SimpleScalarPropertiesInput struct {
	ByteValue *int8

	DoubleValue *float64

	FalseBooleanValue *bool

	FloatValue *float32

	Foo *string

	IntegerValue *int32

	LongValue *int64

	ShortValue *int16

	StringValue *string

	TrueBooleanValue *bool
}

type SimpleScalarPropertiesOutput struct {
	ByteValue *int8

	DoubleValue *float64

	FalseBooleanValue *bool

	FloatValue *float32

	Foo *string

	IntegerValue *int32

	LongValue *int64

	ShortValue *int16

	StringValue *string

	TrueBooleanValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSimpleScalarPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpSimpleScalarProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpSimpleScalarProperties{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opSimpleScalarProperties(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSimpleScalarProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SimpleScalarProperties",
	}
}
