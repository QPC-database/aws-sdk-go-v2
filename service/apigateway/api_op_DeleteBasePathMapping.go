// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the BasePathMapping () resource.
func (c *Client) DeleteBasePathMapping(ctx context.Context, params *DeleteBasePathMappingInput, optFns ...func(*Options)) (*DeleteBasePathMappingOutput, error) {
	if params == nil {
		params = &DeleteBasePathMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBasePathMapping", params, optFns, addOperationDeleteBasePathMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBasePathMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete the BasePathMapping () resource.
type DeleteBasePathMappingInput struct {

	// [Required] The base path name of the BasePathMapping () resource to delete. To
	// specify an empty base path, set this parameter to '(none)'.
	//
	// This member is required.
	BasePath *string

	// [Required] The domain name of the BasePathMapping () resource to delete.
	//
	// This member is required.
	DomainName *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

type DeleteBasePathMappingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBasePathMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBasePathMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBasePathMapping{}, middleware.After)
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
	addOpDeleteBasePathMappingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBasePathMapping(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBasePathMapping(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteBasePathMapping",
	}
}
