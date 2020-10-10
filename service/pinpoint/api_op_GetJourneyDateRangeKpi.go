// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves (queries) pre-aggregated data for a standard engagement metric that
// applies to a journey.
func (c *Client) GetJourneyDateRangeKpi(ctx context.Context, params *GetJourneyDateRangeKpiInput, optFns ...func(*Options)) (*GetJourneyDateRangeKpiOutput, error) {
	if params == nil {
		params = &GetJourneyDateRangeKpiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetJourneyDateRangeKpi", params, optFns, addOperationGetJourneyDateRangeKpiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetJourneyDateRangeKpiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJourneyDateRangeKpiInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier for the journey.
	//
	// This member is required.
	JourneyId *string

	// The name of the metric, also referred to as a key performance indicator (KPI),
	// to retrieve data for. This value describes the associated metric and consists of
	// two or more terms, which are comprised of lowercase alphanumeric characters,
	// separated by a hyphen. Examples are email-open-rate and
	// successful-delivery-rate. For a list of valid values, see the Amazon Pinpoint
	// Developer Guide
	// (https://docs.aws.amazon.com/pinpoint/latest/developerguide/analytics-standard-metrics.html).
	//
	// This member is required.
	KpiName *string

	// The last date and time to retrieve data for, as part of an inclusive date range
	// that filters the query results. This value should be in extended ISO 8601 format
	// and use Coordinated Universal Time (UTC), for example: 2019-07-26T20:00:00Z for
	// 8:00 PM UTC July 26, 2019.
	EndTime *time.Time

	// The string that specifies which page of results to return in a paginated
	// response. This parameter is not supported for application, campaign, and journey
	// metrics.
	NextToken *string

	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string

	// The first date and time to retrieve data for, as part of an inclusive date range
	// that filters the query results. This value should be in extended ISO 8601 format
	// and use Coordinated Universal Time (UTC), for example: 2019-07-19T20:00:00Z for
	// 8:00 PM UTC July 19, 2019. This value should also be fewer than 90 days from the
	// current day.
	StartTime *time.Time
}

type GetJourneyDateRangeKpiOutput struct {

	// Provides the results of a query that retrieved the data for a standard
	// engagement metric that applies to a journey, and provides information about that
	// query.
	//
	// This member is required.
	JourneyDateRangeKpiResponse *types.JourneyDateRangeKpiResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetJourneyDateRangeKpiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetJourneyDateRangeKpi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetJourneyDateRangeKpi{}, middleware.After)
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
	addOpGetJourneyDateRangeKpiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetJourneyDateRangeKpi(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetJourneyDateRangeKpi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetJourneyDateRangeKpi",
	}
}
