// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListReviewableHITs operation retrieves the HITs with Status equal to
// Reviewable or Status equal to Reviewing that belong to the Requester calling the
// operation.
func (c *Client) ListReviewableHITs(ctx context.Context, params *ListReviewableHITsInput, optFns ...func(*Options)) (*ListReviewableHITsOutput, error) {
	if params == nil {
		params = &ListReviewableHITsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReviewableHITs", params, optFns, addOperationListReviewableHITsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReviewableHITsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReviewableHITsInput struct {

	// The ID of the HIT type of the HITs to consider for the query. If not specified,
	// all HITs for the Reviewer are considered
	HITTypeId *string

	// Limit the number of results returned.
	MaxResults *int32

	// Pagination Token
	NextToken *string

	// Can be either Reviewable or Reviewing. Reviewable is the default value.
	Status types.ReviewableHITStatus
}

type ListReviewableHITsOutput struct {

	// The list of HIT elements returned by the query.
	HITs []*types.HIT

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of HITs on this page in the filtered results list, equivalent to the
	// number of HITs being returned by this call.
	NumResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListReviewableHITsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListReviewableHITs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReviewableHITs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListReviewableHITs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListReviewableHITs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListReviewableHITs",
	}
}
