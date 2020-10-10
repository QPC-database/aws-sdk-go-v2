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

// Gets information about the IPv6 CIDR block associations for a specified IPv6
// address pool.
func (c *Client) GetAssociatedIpv6PoolCidrs(ctx context.Context, params *GetAssociatedIpv6PoolCidrsInput, optFns ...func(*Options)) (*GetAssociatedIpv6PoolCidrsOutput, error) {
	if params == nil {
		params = &GetAssociatedIpv6PoolCidrsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssociatedIpv6PoolCidrs", params, optFns, addOperationGetAssociatedIpv6PoolCidrsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssociatedIpv6PoolCidrsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssociatedIpv6PoolCidrsInput struct {

	// The ID of the IPv6 address pool.
	//
	// This member is required.
	PoolId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type GetAssociatedIpv6PoolCidrsOutput struct {

	// Information about the IPv6 CIDR block associations.
	Ipv6CidrAssociations []*types.Ipv6CidrAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAssociatedIpv6PoolCidrsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetAssociatedIpv6PoolCidrs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetAssociatedIpv6PoolCidrs{}, middleware.After)
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
	addOpGetAssociatedIpv6PoolCidrsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssociatedIpv6PoolCidrs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAssociatedIpv6PoolCidrs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetAssociatedIpv6PoolCidrs",
	}
}
