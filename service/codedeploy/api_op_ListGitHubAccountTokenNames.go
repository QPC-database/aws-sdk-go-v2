// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the names of stored connections to GitHub accounts.
func (c *Client) ListGitHubAccountTokenNames(ctx context.Context, params *ListGitHubAccountTokenNamesInput, optFns ...func(*Options)) (*ListGitHubAccountTokenNamesOutput, error) {
	if params == nil {
		params = &ListGitHubAccountTokenNamesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGitHubAccountTokenNames", params, optFns, addOperationListGitHubAccountTokenNamesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGitHubAccountTokenNamesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListGitHubAccountTokenNames operation.
type ListGitHubAccountTokenNamesInput struct {

	// An identifier returned from the previous ListGitHubAccountTokenNames call. It
	// can be used to return the next set of names in the list.
	NextToken *string
}

// Represents the output of a ListGitHubAccountTokenNames operation.
type ListGitHubAccountTokenNamesOutput struct {

	// If a large amount of information is returned, an identifier is also returned. It
	// can be used in a subsequent ListGitHubAccountTokenNames call to return the next
	// set of names in the list.
	NextToken *string

	// A list of names of connections to GitHub accounts.
	TokenNameList []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListGitHubAccountTokenNamesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGitHubAccountTokenNames{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGitHubAccountTokenNames{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListGitHubAccountTokenNames(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListGitHubAccountTokenNames(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListGitHubAccountTokenNames",
	}
}
