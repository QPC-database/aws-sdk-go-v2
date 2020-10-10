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

// Gets information about the associations for the specified transit gateway route
// table.
func (c *Client) GetTransitGatewayRouteTableAssociations(ctx context.Context, params *GetTransitGatewayRouteTableAssociationsInput, optFns ...func(*Options)) (*GetTransitGatewayRouteTableAssociationsOutput, error) {
	if params == nil {
		params = &GetTransitGatewayRouteTableAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayRouteTableAssociations", params, optFns, addOperationGetTransitGatewayRouteTableAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayRouteTableAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayRouteTableAssociationsInput struct {

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//     * resource-id - The ID of the
	// resource.
	//
	//     * resource-type - The resource type (vpc | vpn).
	//
	//     *
	// transit-gateway-attachment-id - The ID of the attachment.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type GetTransitGatewayRouteTableAssociationsOutput struct {

	// Information about the associations.
	Associations []*types.TransitGatewayRouteTableAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTransitGatewayRouteTableAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetTransitGatewayRouteTableAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetTransitGatewayRouteTableAssociations{}, middleware.After)
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
	addOpGetTransitGatewayRouteTableAssociationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayRouteTableAssociations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetTransitGatewayRouteTableAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetTransitGatewayRouteTableAssociations",
	}
}
