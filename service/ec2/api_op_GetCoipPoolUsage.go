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

// Describes the allocations from the specified customer-owned address pool.
func (c *Client) GetCoipPoolUsage(ctx context.Context, params *GetCoipPoolUsageInput, optFns ...func(*Options)) (*GetCoipPoolUsageOutput, error) {
	if params == nil {
		params = &GetCoipPoolUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCoipPoolUsage", params, optFns, addOperationGetCoipPoolUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCoipPoolUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCoipPoolUsageInput struct {

	// The ID of the address pool.
	//
	// This member is required.
	PoolId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters. The following are the possible values:
	//
	//     *
	// coip-address-usage.allocation-id
	//
	//     * coip-address-usage.aws-account-id
	//
	//     *
	// coip-address-usage.aws-service
	//
	//     * coip-address-usage.co-ip
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type GetCoipPoolUsageOutput struct {

	// Information about the address usage.
	CoipAddressUsages []*types.CoipAddressUsage

	// The ID of the customer-owned address pool.
	CoipPoolId *string

	// The ID of the local gateway route table.
	LocalGatewayRouteTableId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCoipPoolUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetCoipPoolUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetCoipPoolUsage{}, middleware.After)
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
	addOpGetCoipPoolUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCoipPoolUsage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCoipPoolUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetCoipPoolUsage",
	}
}
