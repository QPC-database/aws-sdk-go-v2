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

// Deletes a route from a Client VPN endpoint. You can only delete routes that you
// manually added using the CreateClientVpnRoute action. You cannot delete routes
// that were automatically added when associating a subnet. To remove routes that
// have been automatically added, disassociate the target subnet from the Client
// VPN endpoint.
func (c *Client) DeleteClientVpnRoute(ctx context.Context, params *DeleteClientVpnRouteInput, optFns ...func(*Options)) (*DeleteClientVpnRouteOutput, error) {
	if params == nil {
		params = &DeleteClientVpnRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteClientVpnRoute", params, optFns, addOperationDeleteClientVpnRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteClientVpnRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteClientVpnRouteInput struct {

	// The ID of the Client VPN endpoint from which the route is to be deleted.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The IPv4 address range, in CIDR notation, of the route to be deleted.
	//
	// This member is required.
	DestinationCidrBlock *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the target subnet used by the route.
	TargetVpcSubnetId *string
}

type DeleteClientVpnRouteOutput struct {

	// The current state of the route.
	Status *types.ClientVpnRouteStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteClientVpnRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteClientVpnRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteClientVpnRoute{}, middleware.After)
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
	addOpDeleteClientVpnRouteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteClientVpnRoute(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteClientVpnRoute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteClientVpnRoute",
	}
}
