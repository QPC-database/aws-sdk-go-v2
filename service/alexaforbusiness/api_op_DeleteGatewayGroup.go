// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a gateway group.
func (c *Client) DeleteGatewayGroup(ctx context.Context, params *DeleteGatewayGroupInput, optFns ...func(*Options)) (*DeleteGatewayGroupOutput, error) {
	if params == nil {
		params = &DeleteGatewayGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGatewayGroup", params, optFns, addOperationDeleteGatewayGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGatewayGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGatewayGroupInput struct {

	// The ARN of the gateway group to delete.
	//
	// This member is required.
	GatewayGroupArn *string
}

type DeleteGatewayGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteGatewayGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteGatewayGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteGatewayGroup{}, middleware.After)
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
	addOpDeleteGatewayGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGatewayGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteGatewayGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "DeleteGatewayGroup",
	}
}
