// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Omits null, but serializes empty string value.
func (c *Client) OmitsNullSerializesEmptyString(ctx context.Context, params *OmitsNullSerializesEmptyStringInput, optFns ...func(*Options)) (*OmitsNullSerializesEmptyStringOutput, error) {
	if params == nil {
		params = &OmitsNullSerializesEmptyStringInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "OmitsNullSerializesEmptyString", params, optFns, addOperationOmitsNullSerializesEmptyStringMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*OmitsNullSerializesEmptyStringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type OmitsNullSerializesEmptyStringInput struct {
	EmptyString *string

	NullValue *string
}

type OmitsNullSerializesEmptyStringOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationOmitsNullSerializesEmptyStringMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpOmitsNullSerializesEmptyString{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpOmitsNullSerializesEmptyString{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opOmitsNullSerializesEmptyString(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opOmitsNullSerializesEmptyString(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "OmitsNullSerializesEmptyString",
	}
}
