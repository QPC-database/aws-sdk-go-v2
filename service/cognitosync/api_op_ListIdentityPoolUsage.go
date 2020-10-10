// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of identity pools registered with Cognito. ListIdentityPoolUsage can
// only be called with developer credentials. You cannot make this API call with
// the temporary user credentials provided by Cognito Identity.
func (c *Client) ListIdentityPoolUsage(ctx context.Context, params *ListIdentityPoolUsageInput, optFns ...func(*Options)) (*ListIdentityPoolUsageOutput, error) {
	if params == nil {
		params = &ListIdentityPoolUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdentityPoolUsage", params, optFns, addOperationListIdentityPoolUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdentityPoolUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for usage information on an identity pool.
type ListIdentityPoolUsageInput struct {

	// The maximum number of results to be returned.
	MaxResults *int32

	// A pagination token for obtaining the next page of results.
	NextToken *string
}

// Returned for a successful ListIdentityPoolUsage request.
type ListIdentityPoolUsageOutput struct {

	// Total number of identities for the identity pool.
	Count *int32

	// Usage information for the identity pools.
	IdentityPoolUsages []*types.IdentityPoolUsage

	// The maximum number of results to be returned.
	MaxResults *int32

	// A pagination token for obtaining the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListIdentityPoolUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIdentityPoolUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIdentityPoolUsage{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentityPoolUsage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListIdentityPoolUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "ListIdentityPoolUsage",
	}
}
