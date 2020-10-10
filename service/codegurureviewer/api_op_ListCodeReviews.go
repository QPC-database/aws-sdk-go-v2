// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the code reviews that the customer has created in the past 90 days.
func (c *Client) ListCodeReviews(ctx context.Context, params *ListCodeReviewsInput, optFns ...func(*Options)) (*ListCodeReviewsOutput, error) {
	if params == nil {
		params = &ListCodeReviewsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCodeReviews", params, optFns, addOperationListCodeReviewsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCodeReviewsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCodeReviewsInput struct {

	// The type of code reviews to list in the response.
	//
	// This member is required.
	Type types.Type

	// The maximum number of results that are returned per call. The default is 100.
	MaxResults *int32

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged.
	NextToken *string

	// List of provider types for filtering that needs to be applied before displaying
	// the result. For example, providerTypes=[GitHub] lists code reviews from GitHub.
	ProviderTypes []types.ProviderType

	// List of repository names for filtering that needs to be applied before
	// displaying the result.
	RepositoryNames []*string

	// List of states for filtering that needs to be applied before displaying the
	// result. For example, states=[Pending] lists code reviews in the Pending state.
	// The valid code review states are:
	//
	//     * Completed: The code review is
	// complete.
	//
	//     * Pending: The code review started and has not completed or
	// failed.
	//
	//     * Failed: The code review failed.
	//
	//     * Deleting: The code review
	// is being deleted.
	States []types.JobState
}

type ListCodeReviewsOutput struct {

	// A list of code reviews that meet the criteria of the request.
	CodeReviewSummaries []*types.CodeReviewSummary

	// Pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCodeReviewsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCodeReviews{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCodeReviews{}, middleware.After)
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
	addOpListCodeReviewsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCodeReviews(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCodeReviews(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "ListCodeReviews",
	}
}
