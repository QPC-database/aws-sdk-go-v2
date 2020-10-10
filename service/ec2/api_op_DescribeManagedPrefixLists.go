// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes your managed prefix lists and any AWS-managed prefix lists. To view
// the entries for your prefix list, use GetManagedPrefixListEntries ().
func (c *Client) DescribeManagedPrefixLists(ctx context.Context, params *DescribeManagedPrefixListsInput, optFns ...func(*Options)) (*DescribeManagedPrefixListsOutput, error) {
	if params == nil {
		params = &DescribeManagedPrefixListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeManagedPrefixLists", params, optFns, addOperationDescribeManagedPrefixListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeManagedPrefixListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeManagedPrefixListsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * owner-id - The ID of the prefix list owner.
	//
	//     *
	// prefix-list-id - The ID of the prefix list.
	//
	//     * prefix-list-name - The name
	// of the prefix list.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more prefix list IDs.
	PrefixListIds []*string
}

type DescribeManagedPrefixListsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the prefix lists.
	PrefixLists []*types.ManagedPrefixList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeManagedPrefixListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeManagedPrefixLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeManagedPrefixLists{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeManagedPrefixLists(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeManagedPrefixLists(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeManagedPrefixLists",
	}
}
