// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists your policies.
func (c *Client) ListPolicies(ctx context.Context, params *ListPoliciesInput, optFns ...func(*Options)) (*ListPoliciesOutput, error) {
	if params == nil {
		params = &ListPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicies", params, optFns, addOperationListPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListPolicies operation.
type ListPoliciesInput struct {

	// Specifies the order for results. If true, the results are returned in ascending
	// creation order.
	AscendingOrder *bool

	// The marker for the next set of results.
	Marker *string

	// The result page size.
	PageSize *int32
}

// The output from the ListPolicies operation.
type ListPoliciesOutput struct {

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string

	// The descriptions of the policies.
	Policies []*types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPolicies{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicies(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListPolicies",
	}
}
