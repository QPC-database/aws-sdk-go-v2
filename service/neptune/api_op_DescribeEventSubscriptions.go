// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the subscription descriptions for a customer account. The description
// for a subscription includes SubscriptionName, SNSTopicARN, CustomerID,
// SourceType, SourceID, CreationTime, and Status. If you specify a
// SubscriptionName, lists the description for that subscription.
func (c *Client) DescribeEventSubscriptions(ctx context.Context, params *DescribeEventSubscriptionsInput, optFns ...func(*Options)) (*DescribeEventSubscriptionsOutput, error) {
	if params == nil {
		params = &DescribeEventSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventSubscriptions", params, optFns, addOperationDescribeEventSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventSubscriptionsInput struct {

	// This parameter is not currently supported.
	Filters []*types.Filter

	// An optional pagination token provided by a previous
	// DescribeOrderableDBInstanceOptions request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The name of the event notification subscription you want to describe.
	SubscriptionName *string
}

type DescribeEventSubscriptionsOutput struct {

	// A list of EventSubscriptions data types.
	EventSubscriptionsList []*types.EventSubscription

	// An optional pagination token provided by a previous
	// DescribeOrderableDBInstanceOptions request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEventSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEventSubscriptions{}, middleware.After)
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
	addOpDescribeEventSubscriptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventSubscriptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEventSubscriptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeEventSubscriptions",
	}
}
