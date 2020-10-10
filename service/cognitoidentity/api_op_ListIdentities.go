// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the identities in an identity pool. You must use AWS Developer credentials
// to call this API.
func (c *Client) ListIdentities(ctx context.Context, params *ListIdentitiesInput, optFns ...func(*Options)) (*ListIdentitiesOutput, error) {
	if params == nil {
		params = &ListIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdentities", params, optFns, addOperationListIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the ListIdentities action.
type ListIdentitiesInput struct {

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	// The maximum number of identities to return.
	//
	// This member is required.
	MaxResults *int32

	// An optional boolean parameter that allows you to hide disabled identities. If
	// omitted, the ListIdentities API will include disabled identities in the
	// response.
	HideDisabled *bool

	// A pagination token.
	NextToken *string
}

// The response to a ListIdentities request.
type ListIdentitiesOutput struct {

	// An object containing a set of identities and associated mappings.
	Identities []*types.IdentityDescription

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string

	// A pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListIdentities{}, middleware.After)
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
	addOpListIdentitiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentities(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListIdentities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "ListIdentities",
	}
}
