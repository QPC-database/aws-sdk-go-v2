// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the problems with your application.
func (c *Client) ListProblems(ctx context.Context, params *ListProblemsInput, optFns ...func(*Options)) (*ListProblemsOutput, error) {
	if params == nil {
		params = &ListProblemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProblems", params, optFns, addOperationListProblemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProblemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProblemsInput struct {

	// The time when the problem ended, in epoch seconds. If not specified, problems
	// within the past seven days are returned.
	EndTime *time.Time

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The name of the resource group.
	ResourceGroupName *string

	// The time when the problem was detected, in epoch seconds. If you don't specify a
	// time frame for the request, problems within the past seven days are returned.
	StartTime *time.Time
}

type ListProblemsOutput struct {

	// The token used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The list of problems.
	ProblemList []*types.Problem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProblemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProblems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProblems{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProblems(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListProblems(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "ListProblems",
	}
}
