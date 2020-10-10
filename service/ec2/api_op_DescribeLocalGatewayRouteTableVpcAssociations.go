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

// Describes the specified associations between VPCs and local gateway route
// tables.
func (c *Client) DescribeLocalGatewayRouteTableVpcAssociations(ctx context.Context, params *DescribeLocalGatewayRouteTableVpcAssociationsInput, optFns ...func(*Options)) (*DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	if params == nil {
		params = &DescribeLocalGatewayRouteTableVpcAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocalGatewayRouteTableVpcAssociations", params, optFns, addOperationDescribeLocalGatewayRouteTableVpcAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocalGatewayRouteTableVpcAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocalGatewayRouteTableVpcAssociationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * local-gateway-id - The ID of a local gateway.
	//
	//     *
	// local-gateway-route-table-id - The ID of the local gateway route table.
	//
	//     *
	// local-gateway-route-table-vpc-association-id - The ID of the association.
	//
	//     *
	// state - The state of the association.
	//
	//     * vpc-id - The ID of the VPC.
	Filters []*types.Filter

	// The IDs of the associations.
	LocalGatewayRouteTableVpcAssociationIds []*string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeLocalGatewayRouteTableVpcAssociationsOutput struct {

	// Information about the associations.
	LocalGatewayRouteTableVpcAssociations []*types.LocalGatewayRouteTableVpcAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLocalGatewayRouteTableVpcAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeLocalGatewayRouteTableVpcAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLocalGatewayRouteTableVpcAssociations{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVpcAssociations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVpcAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeLocalGatewayRouteTableVpcAssociations",
	}
}
