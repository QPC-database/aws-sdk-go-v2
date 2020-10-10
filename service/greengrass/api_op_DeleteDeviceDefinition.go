// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a device definition.
func (c *Client) DeleteDeviceDefinition(ctx context.Context, params *DeleteDeviceDefinitionInput, optFns ...func(*Options)) (*DeleteDeviceDefinitionOutput, error) {
	if params == nil {
		params = &DeleteDeviceDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDeviceDefinition", params, optFns, addOperationDeleteDeviceDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDeviceDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDeviceDefinitionInput struct {

	// The ID of the device definition.
	//
	// This member is required.
	DeviceDefinitionId *string
}

type DeleteDeviceDefinitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDeviceDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDeviceDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDeviceDefinition{}, middleware.After)
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
	addOpDeleteDeviceDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDeviceDefinition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDeviceDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "DeleteDeviceDefinition",
	}
}
