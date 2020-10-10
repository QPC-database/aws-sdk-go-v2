// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a cache policy. You cannot delete a cache policy if it’s attached to a
// cache behavior. First update your distributions to remove the cache policy from
// all cache behaviors, then delete the cache policy. To delete a cache policy, you
// must provide the policy’s identifier and version. To get these values, you can
// use ListCachePolicies or GetCachePolicy.
func (c *Client) DeleteCachePolicy(ctx context.Context, params *DeleteCachePolicyInput, optFns ...func(*Options)) (*DeleteCachePolicyOutput, error) {
	if params == nil {
		params = &DeleteCachePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCachePolicy", params, optFns, addOperationDeleteCachePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCachePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCachePolicyInput struct {

	// The unique identifier for the cache policy that you are deleting. To get the
	// identifier, you can use ListCachePolicies.
	//
	// This member is required.
	Id *string

	// The version of the cache policy that you are deleting. The version is the cache
	// policy’s ETag value, which you can get using ListCachePolicies, GetCachePolicy,
	// or GetCachePolicyConfig.
	IfMatch *string
}

type DeleteCachePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCachePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteCachePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteCachePolicy{}, middleware.After)
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
	addOpDeleteCachePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCachePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCachePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteCachePolicy",
	}
}
