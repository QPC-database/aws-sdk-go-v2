// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a pull request in the specified repository.
func (c *Client) CreatePullRequest(ctx context.Context, params *CreatePullRequestInput, optFns ...func(*Options)) (*CreatePullRequestOutput, error) {
	if params == nil {
		params = &CreatePullRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePullRequest", params, optFns, addOperationCreatePullRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePullRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePullRequestInput struct {

	// The targets for the pull request, including the source of the code to be
	// reviewed (the source branch) and the destination where the creator of the pull
	// request intends the code to be merged after the pull request is closed (the
	// destination branch).
	//
	// This member is required.
	Targets []*types.Target

	// The title of the pull request. This title is used to identify the pull request
	// to other users in the repository.
	//
	// This member is required.
	Title *string

	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request is
	// received with the same parameters and a token is included, the request returns
	// information about the initial request that used that token. The AWS SDKs
	// prepopulate client request tokens. If you are using an AWS SDK, an idempotency
	// token is created for you.
	ClientRequestToken *string

	// A description of the pull request.
	Description *string
}

type CreatePullRequestOutput struct {

	// Information about the newly created pull request.
	//
	// This member is required.
	PullRequest *types.PullRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePullRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePullRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePullRequest{}, middleware.After)
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
	addIdempotencyToken_opCreatePullRequestMiddleware(stack, options)
	addOpCreatePullRequestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePullRequest(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreatePullRequest struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePullRequest) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePullRequest) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePullRequestInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePullRequestInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePullRequestMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreatePullRequest{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePullRequest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "CreatePullRequest",
	}
}
