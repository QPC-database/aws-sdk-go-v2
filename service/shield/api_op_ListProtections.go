// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all Protection () objects for the account.
func (c *Client) ListProtections(ctx context.Context, params *ListProtectionsInput, optFns ...func(*Options)) (*ListProtectionsOutput, error) {
	if params == nil {
		params = &ListProtectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProtections", params, optFns, addOperationListProtectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProtectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProtectionsInput struct {

	// The maximum number of Protection () objects to be returned. If this is left
	// blank the first 20 results will be returned. This is a maximum value; it is
	// possible that AWS WAF will return the results in smaller batches. That is, the
	// number of Protection () objects returned could be less than MaxResults, even if
	// there are still more Protection () objects yet to return. If there are more
	// Protection () objects to return, AWS WAF will always also return a NextToken.
	MaxResults *int32

	// The ListProtectionsRequest.NextToken value from a previous call to
	// ListProtections. Pass null if this is the first call.
	NextToken *string
}

type ListProtectionsOutput struct {

	// If you specify a value for MaxResults and you have more Protections than the
	// value of MaxResults, AWS Shield Advanced returns a NextToken value in the
	// response that allows you to list another group of Protections. For the second
	// and subsequent ListProtections requests, specify the value of NextToken from the
	// previous response to get information about another batch of Protections. AWS WAF
	// might return the list of Protection () objects in batches smaller than the
	// number specified by MaxResults. If there are more Protection () objects to
	// return, AWS WAF will always also return a NextToken.
	NextToken *string

	// The array of enabled Protection () objects.
	Protections []*types.Protection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProtectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProtections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProtections{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProtections(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListProtections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "ListProtections",
	}
}
