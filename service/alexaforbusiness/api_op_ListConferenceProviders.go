// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists conference providers under a specific AWS account.
func (c *Client) ListConferenceProviders(ctx context.Context, params *ListConferenceProvidersInput, optFns ...func(*Options)) (*ListConferenceProvidersOutput, error) {
	if params == nil {
		params = &ListConferenceProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConferenceProviders", params, optFns, addOperationListConferenceProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConferenceProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConferenceProvidersInput struct {

	// The maximum number of conference providers to be returned, per paginated calls.
	MaxResults *int32

	// The tokens used for pagination.
	NextToken *string
}

type ListConferenceProvidersOutput struct {

	// The conference providers.
	ConferenceProviders []*types.ConferenceProvider

	// The tokens used for pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListConferenceProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListConferenceProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConferenceProviders{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListConferenceProviders(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListConferenceProviders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ListConferenceProviders",
	}
}
