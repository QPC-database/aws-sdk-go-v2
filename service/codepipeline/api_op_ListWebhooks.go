// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a listing of all the webhooks in this AWS Region for this account. The
// output lists all webhooks and includes the webhook URL and ARN and the
// configuration for each webhook.
func (c *Client) ListWebhooks(ctx context.Context, params *ListWebhooksInput, optFns ...func(*Options)) (*ListWebhooksOutput, error) {
	if params == nil {
		params = &ListWebhooksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWebhooks", params, optFns, addOperationListWebhooksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWebhooksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWebhooksInput struct {

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token that was returned from the previous ListWebhooks call, which can be
	// used to return the next set of webhooks in the list.
	NextToken *string
}

type ListWebhooksOutput struct {

	// If the amount of returned information is significantly large, an identifier is
	// also returned and can be used in a subsequent ListWebhooks call to return the
	// next set of webhooks in the list.
	NextToken *string

	// The JSON detail returned for each webhook in the list output for the
	// ListWebhooks call.
	Webhooks []*types.ListWebhookItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWebhooksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWebhooks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWebhooks{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWebhooks(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListWebhooks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "ListWebhooks",
	}
}
