// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Null and empty headers are not sent over the wire.
func (c *Client) NullAndEmptyHeadersClient(ctx context.Context, params *NullAndEmptyHeadersClientInput, optFns ...func(*Options)) (*NullAndEmptyHeadersClientOutput, error) {
	if params == nil {
		params = &NullAndEmptyHeadersClientInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NullAndEmptyHeadersClient", params, optFns, addOperationNullAndEmptyHeadersClientMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NullAndEmptyHeadersClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NullAndEmptyHeadersClientInput struct {
	A *string

	B *string

	C []*string
}

type NullAndEmptyHeadersClientOutput struct {
	A *string

	B *string

	C []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationNullAndEmptyHeadersClientMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpNullAndEmptyHeadersClient{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpNullAndEmptyHeadersClient{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opNullAndEmptyHeadersClient(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opNullAndEmptyHeadersClient(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NullAndEmptyHeadersClient",
	}
}
