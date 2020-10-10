// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a subset of information about all the custom data identifiers for an
// account.
func (c *Client) ListCustomDataIdentifiers(ctx context.Context, params *ListCustomDataIdentifiersInput, optFns ...func(*Options)) (*ListCustomDataIdentifiersOutput, error) {
	if params == nil {
		params = &ListCustomDataIdentifiersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomDataIdentifiers", params, optFns, addOperationListCustomDataIdentifiersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomDataIdentifiersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomDataIdentifiersInput struct {

	// The maximum number of items to include in each page of the response.
	MaxResults *int32

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string
}

type ListCustomDataIdentifiersOutput struct {

	// An array of objects, one for each custom data identifier.
	Items []*types.CustomDataIdentifierSummary

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCustomDataIdentifiersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCustomDataIdentifiers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCustomDataIdentifiers{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomDataIdentifiers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCustomDataIdentifiers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "ListCustomDataIdentifiers",
	}
}
