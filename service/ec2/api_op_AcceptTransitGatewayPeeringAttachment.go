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

// Accepts a transit gateway peering attachment request. The peering attachment
// must be in the pendingAcceptance state.
func (c *Client) AcceptTransitGatewayPeeringAttachment(ctx context.Context, params *AcceptTransitGatewayPeeringAttachmentInput, optFns ...func(*Options)) (*AcceptTransitGatewayPeeringAttachmentOutput, error) {
	if params == nil {
		params = &AcceptTransitGatewayPeeringAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptTransitGatewayPeeringAttachment", params, optFns, addOperationAcceptTransitGatewayPeeringAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptTransitGatewayPeeringAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptTransitGatewayPeeringAttachmentInput struct {

	// The ID of the transit gateway attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type AcceptTransitGatewayPeeringAttachmentOutput struct {

	// The transit gateway peering attachment.
	TransitGatewayPeeringAttachment *types.TransitGatewayPeeringAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAcceptTransitGatewayPeeringAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAcceptTransitGatewayPeeringAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAcceptTransitGatewayPeeringAttachment{}, middleware.After)
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
	addOpAcceptTransitGatewayPeeringAttachmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptTransitGatewayPeeringAttachment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAcceptTransitGatewayPeeringAttachment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AcceptTransitGatewayPeeringAttachment",
	}
}
