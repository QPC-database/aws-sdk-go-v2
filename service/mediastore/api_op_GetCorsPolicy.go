// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastore

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the cross-origin resource sharing (CORS) configuration information that
// is set for the container. To use this operation, you must have permission to
// perform the MediaStore:GetCorsPolicy action. By default, the container owner has
// this permission and can grant it to others.
func (c *Client) GetCorsPolicy(ctx context.Context, params *GetCorsPolicyInput, optFns ...func(*Options)) (*GetCorsPolicyOutput, error) {
	if params == nil {
		params = &GetCorsPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCorsPolicy", params, optFns, addOperationGetCorsPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCorsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCorsPolicyInput struct {

	// The name of the container that the policy is assigned to.
	//
	// This member is required.
	ContainerName *string
}

type GetCorsPolicyOutput struct {

	// The CORS policy assigned to the container.
	//
	// This member is required.
	CorsPolicy []*types.CorsRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCorsPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCorsPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCorsPolicy{}, middleware.After)
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
	addOpGetCorsPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCorsPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCorsPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediastore",
		OperationName: "GetCorsPolicy",
	}
}
