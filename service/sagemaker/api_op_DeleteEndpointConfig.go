// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an endpoint configuration. The DeleteEndpointConfig API deletes only the
// specified configuration. It does not delete endpoints created using the
// configuration. You must not delete an EndpointConfig in use by an endpoint that
// is live or while the UpdateEndpoint or CreateEndpoint operations are being
// performed on the endpoint. If you delete the EndpointConfig of an endpoint that
// is active or being created or updated you may lose visibility into the instance
// type the endpoint is using. The endpoint must be deleted in order to stop
// incurring charges.
func (c *Client) DeleteEndpointConfig(ctx context.Context, params *DeleteEndpointConfigInput, optFns ...func(*Options)) (*DeleteEndpointConfigOutput, error) {
	if params == nil {
		params = &DeleteEndpointConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEndpointConfig", params, optFns, addOperationDeleteEndpointConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEndpointConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEndpointConfigInput struct {

	// The name of the endpoint configuration that you want to delete.
	//
	// This member is required.
	EndpointConfigName *string
}

type DeleteEndpointConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteEndpointConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEndpointConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEndpointConfig{}, middleware.After)
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
	addOpDeleteEndpointConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEndpointConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteEndpointConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DeleteEndpointConfig",
	}
}
