// Code generated by smithy-go-codegen DO NOT EDIT.

package commander

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/commander/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists timeline events of the specified incident record.
func (c *Client) ListTimelineEvents(ctx context.Context, params *ListTimelineEventsInput, optFns ...func(*Options)) (*ListTimelineEventsOutput, error) {
	if params == nil {
		params = &ListTimelineEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTimelineEvents", params, optFns, addOperationListTimelineEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTimelineEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTimelineEventsInput struct {

	// The Amazon Resource Name (ARN) of the incident that the event is part of.
	//
	// This member is required.
	IncidentRecordArn *string

	// Filters the timeline events based on the provided conditional values. You can
	// filter timeline events using the following keys:
	//
	// * eventTime
	//
	// * eventType
	Filters []types.Filter

	// The maximum number of results per page.
	MaxResults *int32

	// The pagination token to continue to the next page of results.
	NextToken *string

	// Sort by the specified key value pair.
	SortBy types.TimelineEventSort

	// Sorts the order of timeline events by the value specified in the sortBy field.
	SortOrder types.SortOrder
}

type ListTimelineEventsOutput struct {

	// Details about each event that occurred during the incident.
	//
	// This member is required.
	EventSummaries []types.EventSummary

	// The pagination token to continue to the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTimelineEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTimelineEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTimelineEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListTimelineEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTimelineEvents(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListTimelineEventsAPIClient is a client that implements the ListTimelineEvents
// operation.
type ListTimelineEventsAPIClient interface {
	ListTimelineEvents(context.Context, *ListTimelineEventsInput, ...func(*Options)) (*ListTimelineEventsOutput, error)
}

var _ ListTimelineEventsAPIClient = (*Client)(nil)

// ListTimelineEventsPaginatorOptions is the paginator options for
// ListTimelineEvents
type ListTimelineEventsPaginatorOptions struct {
	// The maximum number of results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTimelineEventsPaginator is a paginator for ListTimelineEvents
type ListTimelineEventsPaginator struct {
	options   ListTimelineEventsPaginatorOptions
	client    ListTimelineEventsAPIClient
	params    *ListTimelineEventsInput
	nextToken *string
	firstPage bool
}

// NewListTimelineEventsPaginator returns a new ListTimelineEventsPaginator
func NewListTimelineEventsPaginator(client ListTimelineEventsAPIClient, params *ListTimelineEventsInput, optFns ...func(*ListTimelineEventsPaginatorOptions)) *ListTimelineEventsPaginator {
	if params == nil {
		params = &ListTimelineEventsInput{}
	}

	options := ListTimelineEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTimelineEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTimelineEventsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTimelineEvents page.
func (p *ListTimelineEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTimelineEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListTimelineEvents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListTimelineEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-incidents",
		OperationName: "ListTimelineEvents",
	}
}
