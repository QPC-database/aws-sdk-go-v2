// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all signing platforms available in code signing that match the request
// parameters. If additional jobs remain to be listed, code signing returns a
// nextToken value. Use this value in subsequent calls to ListSigningJobs to fetch
// the remaining values. You can continue calling ListSigningJobs with your
// maxResults parameter and with new values that code signing returns in the
// nextToken parameter until all of your signing jobs have been returned.
func (c *Client) ListSigningPlatforms(ctx context.Context, params *ListSigningPlatformsInput, optFns ...func(*Options)) (*ListSigningPlatformsOutput, error) {
	if params == nil {
		params = &ListSigningPlatformsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSigningPlatforms", params, optFns, addOperationListSigningPlatformsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSigningPlatformsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSigningPlatformsInput struct {

	// The category type of a signing platform.
	Category *string

	// The maximum number of results to be returned by this operation.
	MaxResults *int32

	// Value for specifying the next set of paginated results to return. After you
	// receive a response with truncated results, use this parameter in a subsequent
	// request. Set it to the value of nextToken from the response that you just
	// received.
	NextToken *string

	// Any partner entities connected to a signing platform.
	Partner *string

	// The validation template that is used by the target signing platform.
	Target *string
}

type ListSigningPlatformsOutput struct {

	// Value for specifying the next set of paginated results to return.
	NextToken *string

	// A list of all platforms that match the request parameters.
	Platforms []*types.SigningPlatform

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSigningPlatformsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSigningPlatforms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSigningPlatforms{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSigningPlatforms(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListSigningPlatforms(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "signer",
		OperationName: "ListSigningPlatforms",
	}
}
