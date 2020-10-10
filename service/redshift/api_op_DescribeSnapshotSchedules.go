// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of snapshot schedules.
func (c *Client) DescribeSnapshotSchedules(ctx context.Context, params *DescribeSnapshotSchedulesInput, optFns ...func(*Options)) (*DescribeSnapshotSchedulesOutput, error) {
	if params == nil {
		params = &DescribeSnapshotSchedulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSnapshotSchedules", params, optFns, addOperationDescribeSnapshotSchedulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSnapshotSchedulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSnapshotSchedulesInput struct {

	// The unique identifier for the cluster whose snapshot schedules you want to view.
	ClusterIdentifier *string

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// The maximum number or response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	MaxRecords *int32

	// A unique identifier for a snapshot schedule.
	ScheduleIdentifier *string

	// The key value for a snapshot schedule tag.
	TagKeys []*string

	// The value corresponding to the key of the snapshot schedule tag.
	TagValues []*string
}

type DescribeSnapshotSchedulesOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// A list of SnapshotSchedules.
	SnapshotSchedules []*types.SnapshotSchedule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSnapshotSchedulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeSnapshotSchedules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeSnapshotSchedules{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshotSchedules(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSnapshotSchedules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeSnapshotSchedules",
	}
}
