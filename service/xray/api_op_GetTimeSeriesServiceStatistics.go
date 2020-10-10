// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Get an aggregation of service statistics defined by a specific time range.
func (c *Client) GetTimeSeriesServiceStatistics(ctx context.Context, params *GetTimeSeriesServiceStatisticsInput, optFns ...func(*Options)) (*GetTimeSeriesServiceStatisticsOutput, error) {
	if params == nil {
		params = &GetTimeSeriesServiceStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTimeSeriesServiceStatistics", params, optFns, addOperationGetTimeSeriesServiceStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTimeSeriesServiceStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTimeSeriesServiceStatisticsInput struct {

	// The end of the time frame for which to aggregate statistics.
	//
	// This member is required.
	EndTime *time.Time

	// The start of the time frame for which to aggregate statistics.
	//
	// This member is required.
	StartTime *time.Time

	// A filter expression defining entities that will be aggregated for statistics.
	// Supports ID, service, and edge functions. If no selector expression is
	// specified, edge statistics are returned.
	EntitySelectorExpression *string

	// The ARN of the group for which to pull statistics from.
	GroupARN *string

	// The case-sensitive name of the group for which to pull statistics from.
	GroupName *string

	// Pagination token.
	NextToken *string

	// Aggregation period in seconds.
	Period *int32
}

type GetTimeSeriesServiceStatisticsOutput struct {

	// A flag indicating whether or not a group's filter expression has been
	// consistent, or if a returned aggregation may show statistics from an older
	// version of the group's filter expression.
	ContainsOldGroupVersions *bool

	// Pagination token.
	NextToken *string

	// The collection of statistics.
	TimeSeriesServiceStatistics []*types.TimeSeriesServiceStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTimeSeriesServiceStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTimeSeriesServiceStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTimeSeriesServiceStatistics{}, middleware.After)
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
	addOpGetTimeSeriesServiceStatisticsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTimeSeriesServiceStatistics(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetTimeSeriesServiceStatistics(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetTimeSeriesServiceStatistics",
	}
}
