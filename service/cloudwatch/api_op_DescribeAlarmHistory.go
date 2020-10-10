// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves the history for the specified alarm. You can filter the results by
// date range or item type. If an alarm name is not specified, the histories for
// either all metric alarms or all composite alarms are returned. CloudWatch
// retains the history of an alarm even if you delete the alarm.
func (c *Client) DescribeAlarmHistory(ctx context.Context, params *DescribeAlarmHistoryInput, optFns ...func(*Options)) (*DescribeAlarmHistoryOutput, error) {
	if params == nil {
		params = &DescribeAlarmHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAlarmHistory", params, optFns, addOperationDescribeAlarmHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAlarmHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAlarmHistoryInput struct {

	// The name of the alarm.
	AlarmName *string

	// Use this parameter to specify whether you want the operation to return metric
	// alarms or composite alarms. If you omit this parameter, only metric alarms are
	// returned.
	AlarmTypes []types.AlarmType

	// The ending date to retrieve alarm history.
	EndDate *time.Time

	// The type of alarm histories to retrieve.
	HistoryItemType types.HistoryItemType

	// The maximum number of alarm history records to retrieve.
	MaxRecords *int32

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string

	// Specified whether to return the newest or oldest alarm history first. Specify
	// TimestampDescending to have the newest event history returned first, and specify
	// TimestampAscending to have the oldest history returned first.
	ScanBy types.ScanBy

	// The starting date to retrieve alarm history.
	StartDate *time.Time
}

type DescribeAlarmHistoryOutput struct {

	// The alarm histories, in JSON format.
	AlarmHistoryItems []*types.AlarmHistoryItem

	// The token that marks the start of the next batch of returned results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAlarmHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAlarmHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAlarmHistory{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAlarmHistory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAlarmHistory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DescribeAlarmHistory",
	}
}
