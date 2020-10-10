// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the SAML provider resource objects defined in IAM in the account. This
// operation requires Signature Version 4
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
func (c *Client) ListSAMLProviders(ctx context.Context, params *ListSAMLProvidersInput, optFns ...func(*Options)) (*ListSAMLProvidersOutput, error) {
	if params == nil {
		params = &ListSAMLProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSAMLProviders", params, optFns, addOperationListSAMLProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSAMLProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSAMLProvidersInput struct {
}

// Contains the response to a successful ListSAMLProviders () request.
type ListSAMLProvidersOutput struct {

	// The list of SAML provider resource objects defined in IAM for this AWS account.
	SAMLProviderList []*types.SAMLProviderListEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSAMLProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListSAMLProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListSAMLProviders{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSAMLProviders(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListSAMLProviders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListSAMLProviders",
	}
}
