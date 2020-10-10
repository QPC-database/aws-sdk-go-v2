// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of pull requests for a specified repository. The return list can
// be refined by pull request status or pull request author ARN.
func (c *Client) ListPullRequests(ctx context.Context, params *ListPullRequestsInput, optFns ...func(*Options)) (*ListPullRequestsOutput, error) {
	if params == nil {
		params = &ListPullRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPullRequests", params, optFns, addOperationListPullRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPullRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPullRequestsInput struct {

	// The name of the repository for which you want to list pull requests.
	//
	// This member is required.
	RepositoryName *string

	// Optional. The Amazon Resource Name (ARN) of the user who created the pull
	// request. If used, this filters the results to pull requests created by that
	// user.
	AuthorArn *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string

	// Optional. The status of the pull request. If used, this refines the results to
	// the pull requests that match the specified status.
	PullRequestStatus types.PullRequestStatusEnum
}

type ListPullRequestsOutput struct {

	// The system-generated IDs of the pull requests.
	//
	// This member is required.
	PullRequestIds []*string

	// An enumeration token that allows the operation to batch the next results of the
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPullRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPullRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPullRequests{}, middleware.After)
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
	addOpListPullRequestsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPullRequests(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListPullRequests(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "ListPullRequests",
	}
}
