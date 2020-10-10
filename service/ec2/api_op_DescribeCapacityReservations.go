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

// Describes one or more of your Capacity Reservations. The results describe only
// the Capacity Reservations in the AWS Region that you're currently using.
func (c *Client) DescribeCapacityReservations(ctx context.Context, params *DescribeCapacityReservationsInput, optFns ...func(*Options)) (*DescribeCapacityReservationsOutput, error) {
	if params == nil {
		params = &DescribeCapacityReservationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCapacityReservations", params, optFns, addOperationDescribeCapacityReservationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCapacityReservationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCapacityReservationsInput struct {

	// The ID of the Capacity Reservation.
	CapacityReservationIds []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * instance-type - The type of instance for which the
	// Capacity Reservation reserves capacity.
	//
	//     * owner-id - The ID of the AWS
	// account that owns the Capacity Reservation.
	//
	//     * availability-zone-id - The
	// Availability Zone ID of the Capacity Reservation.
	//
	//     * instance-platform - The
	// type of operating system for which the Capacity Reservation reserves capacity.
	//
	//
	// * availability-zone - The Availability Zone ID of the Capacity Reservation.
	//
	//
	// * tenancy - Indicates the tenancy of the Capacity Reservation. A Capacity
	// Reservation can have one of the following tenancy settings:
	//
	//         * default -
	// The Capacity Reservation is created on hardware that is shared with other AWS
	// accounts.
	//
	//         * dedicated - The Capacity Reservation is created on
	// single-tenant hardware that is dedicated to a single AWS account.
	//
	//     * state -
	// The current state of the Capacity Reservation. A Capacity Reservation can be in
	// one of the following states:
	//
	//         * active- The Capacity Reservation is
	// active and the capacity is available for your use.
	//
	//         * expired - The
	// Capacity Reservation expired automatically at the date and time specified in
	// your request. The reserved capacity is no longer available for your use.
	//
	//
	// * cancelled - The Capacity Reservation was manually cancelled. The reserved
	// capacity is no longer available for your use.
	//
	//         * pending - The Capacity
	// Reservation request was successful but the capacity provisioning is still
	// pending.
	//
	//         * failed - The Capacity Reservation request has failed. A
	// request might fail due to invalid request parameters, capacity constraints, or
	// instance limit constraints. Failed requests are retained for 60 minutes.
	//
	//     *
	// end-date - The date and time at which the Capacity Reservation expires. When a
	// Capacity Reservation expires, the reserved capacity is released and you can no
	// longer launch instances into it. The Capacity Reservation's state changes to
	// expired when it reaches its end date and time.
	//
	//     * end-date-type - Indicates
	// the way in which the Capacity Reservation ends. A Capacity Reservation can have
	// one of the following end types:
	//
	//         * unlimited - The Capacity Reservation
	// remains active until you explicitly cancel it.
	//
	//         * limited - The Capacity
	// Reservation expires automatically at a specified date and time.
	//
	//     *
	// instance-match-criteria - Indicates the type of instance launches that the
	// Capacity Reservation accepts. The options include:
	//
	//         * open - The
	// Capacity Reservation accepts all instances that have matching attributes
	// (instance type, platform, and Availability Zone). Instances that have matching
	// attributes launch into the Capacity Reservation automatically without specifying
	// any additional parameters.
	//
	//         * targeted - The Capacity Reservation only
	// accepts instances that have matching attributes (instance type, platform, and
	// Availability Zone), and explicitly target the Capacity Reservation. This ensures
	// that only permitted instances can use the reserved capacity.
	Filters []*types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given a
	// larger value than 500, you receive an error.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string
}

type DescribeCapacityReservationsOutput struct {

	// Information about the Capacity Reservations.
	CapacityReservations []*types.CapacityReservation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCapacityReservationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeCapacityReservations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeCapacityReservations{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCapacityReservations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeCapacityReservations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeCapacityReservations",
	}
}
