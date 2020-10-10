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

// Searches contacts and lists the ones that meet a set of filter and sort
// criteria.
func (c *Client) SearchContacts(ctx context.Context, params *SearchContactsInput, optFns ...func(*Options)) (*SearchContactsOutput, error) {
	if params == nil {
		params = &SearchContactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchContacts", params, optFns, addOperationSearchContactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchContactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchContactsInput struct {

	// The filters to use to list a specified set of address books. The supported
	// filter keys are DisplayName, FirstName, LastName, and AddressBookArns.
	Filters []*types.Filter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response only
	// includes results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The sort order to use in listing the specified set of contacts. The supported
	// sort keys are DisplayName, FirstName, and LastName.
	SortCriteria []*types.Sort
}

type SearchContactsOutput struct {

	// The contacts that meet the specified set of filter criteria, in sort order.
	Contacts []*types.ContactData

	// The token returned to indicate that there is more data available.
	NextToken *string

	// The total number of contacts returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchContactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchContacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchContacts{}, middleware.After)
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
	addOpSearchContactsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchContacts(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSearchContacts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchContacts",
	}
}
