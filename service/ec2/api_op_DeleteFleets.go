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

// Deletes the specified EC2 Fleet. After you delete an EC2 Fleet, it launches no
// new instances. You must specify whether an EC2 Fleet should also terminate its
// instances. If you terminate the instances, the EC2 Fleet enters the
// deleted_terminating state. Otherwise, the EC2 Fleet enters the deleted_running
// state, and the instances continue to run until they are interrupted or you
// terminate them manually.
func (c *Client) DeleteFleets(ctx context.Context, params *DeleteFleetsInput, optFns ...func(*Options)) (*DeleteFleetsOutput, error) {
	if params == nil {
		params = &DeleteFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFleets", params, optFns, addOperationDeleteFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFleetsInput struct {

	// The IDs of the EC2 Fleets.
	//
	// This member is required.
	FleetIds []*string

	// Indicates whether to terminate instances for an EC2 Fleet if it is deleted
	// successfully.
	//
	// This member is required.
	TerminateInstances *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteFleetsOutput struct {

	// Information about the EC2 Fleets that are successfully deleted.
	SuccessfulFleetDeletions []*types.DeleteFleetSuccessItem

	// Information about the EC2 Fleets that are not successfully deleted.
	UnsuccessfulFleetDeletions []*types.DeleteFleetErrorItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteFleets{}, middleware.After)
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
	addOpDeleteFleetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFleets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteFleets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteFleets",
	}
}
