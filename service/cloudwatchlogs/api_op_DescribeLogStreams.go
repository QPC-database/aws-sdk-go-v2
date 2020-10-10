// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the log streams for the specified log group. You can list all the log
// streams or filter the results by prefix. You can also control how the results
// are ordered. This operation has a limit of five transactions per second, after
// which transactions are throttled.
func (c *Client) DescribeLogStreams(ctx context.Context, params *DescribeLogStreamsInput, optFns ...func(*Options)) (*DescribeLogStreamsOutput, error) {
	if params == nil {
		params = &DescribeLogStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLogStreams", params, optFns, addOperationDescribeLogStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLogStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLogStreamsInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// If the value is true, results are returned in descending order. If the value is
	// to false, results are returned in ascending order. The default value is false.
	Descending *bool

	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit *int32

	// The prefix to match. If orderBy is LastEventTime,you cannot specify this
	// parameter.
	LogStreamNamePrefix *string

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// If the value is LogStreamName, the results are ordered by log stream name. If
	// the value is LastEventTime, the results are ordered by the event time. The
	// default value is LogStreamName. If you order the results by event time, you
	// cannot specify the logStreamNamePrefix parameter. lastEventTimestamp represents
	// the time of the most recent log event in the log stream in CloudWatch Logs. This
	// number is expressed as the number of milliseconds after Jan 1, 1970 00:00:00
	// UTC. lastEventTimeStamp updates on an eventual consistency basis. It typically
	// updates in less than an hour from ingestion, but may take longer in some rare
	// situations.
	OrderBy types.OrderBy
}

type DescribeLogStreamsOutput struct {

	// The log streams.
	LogStreams []*types.LogStream

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLogStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLogStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLogStreams{}, middleware.After)
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
	addOpDescribeLogStreamsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLogStreams(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeLogStreams(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DescribeLogStreams",
	}
}
