// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of RepositorySummary
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositorySummary.html)
// objects. Each RepositorySummary contains information about a repository in the
// specified AWS account and that matches the input parameters.
func (c *Client) ListRepositories(ctx context.Context, params *ListRepositoriesInput, optFns ...func(*Options)) (*ListRepositoriesOutput, error) {
	if params == nil {
		params = &ListRepositoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRepositories", params, optFns, addOperationListRepositoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRepositoriesInput struct {

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// A prefix used to filter returned repositories. Only repositories with names that
	// start with repositoryPrefix are returned.
	RepositoryPrefix *string
}

type ListRepositoriesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The returned list of RepositorySummary
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositorySummary.html)
	// objects.
	Repositories []*types.RepositorySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRepositories{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositories(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListRepositories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListRepositories",
	}
}
