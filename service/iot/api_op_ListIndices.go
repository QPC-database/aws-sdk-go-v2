// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the search indices.
func (c *Client) ListIndices(ctx context.Context, params *ListIndicesInput, optFns ...func(*Options)) (*ListIndicesOutput, error) {
	if params == nil {
		params = &ListIndicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIndices", params, optFns, addOperationListIndicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIndicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIndicesInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string
}

type ListIndicesOutput struct {

	// The index names.
	IndexNames []*string

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListIndicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIndices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIndices{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIndices(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListIndices(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListIndices",
	}
}
