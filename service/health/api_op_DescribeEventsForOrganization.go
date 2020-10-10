// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about events across your organization in AWS Organizations,
// meeting the specified filter criteria. Events are returned in a summary form and
// do not include the accounts impacted, detailed description, any additional
// metadata that depends on the event type, or any affected resources. To retrieve
// that information, use the DescribeAffectedAccountsForOrganization (),
// DescribeEventDetailsForOrganization (), and
// DescribeAffectedEntitiesForOrganization () operations. If no filter criteria are
// specified, all events across your organization are returned. Results are sorted
// by lastModifiedTime, starting with the most recent. Before you can call this
// operation, you must first enable Health to work with AWS Organizations. To do
// this, call the EnableHealthServiceAccessForOrganization () operation from your
// organization's master account.
func (c *Client) DescribeEventsForOrganization(ctx context.Context, params *DescribeEventsForOrganizationInput, optFns ...func(*Options)) (*DescribeEventsForOrganizationOutput, error) {
	if params == nil {
		params = &DescribeEventsForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventsForOrganization", params, optFns, addOperationDescribeEventsForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventsForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventsForOrganizationInput struct {

	// Values to narrow the results returned.
	Filter *types.OrganizationEventFilter

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string

	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	MaxResults *int32

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string
}

type DescribeEventsForOrganizationOutput struct {

	// The events that match the specified filter criteria.
	Events []*types.OrganizationEvent

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventsForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventsForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventsForOrganization{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventsForOrganization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEventsForOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeEventsForOrganization",
	}
}
