// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a registry.
func (c *Client) UpdateRegistry(ctx context.Context, params *UpdateRegistryInput, optFns ...func(*Options)) (*UpdateRegistryOutput, error) {
	if params == nil {
		params = &UpdateRegistryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRegistry", params, optFns, addOperationUpdateRegistryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRegistryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates the registry.
type UpdateRegistryInput struct {

	// The name of the registry.
	//
	// This member is required.
	RegistryName *string

	// The description of the registry to update.
	Description *string
}

type UpdateRegistryOutput struct {

	// The description of the registry.
	Description *string

	// The ARN of the registry.
	RegistryArn *string

	// The name of the registry.
	RegistryName *string

	// Tags associated with the registry.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRegistryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRegistry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRegistry{}, middleware.After)
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
	addOpUpdateRegistryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRegistry(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateRegistry(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "UpdateRegistry",
	}
}
