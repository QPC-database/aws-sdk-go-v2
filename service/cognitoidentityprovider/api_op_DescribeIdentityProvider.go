// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a specific identity provider.
func (c *Client) DescribeIdentityProvider(ctx context.Context, params *DescribeIdentityProviderInput, optFns ...func(*Options)) (*DescribeIdentityProviderOutput, error) {
	if params == nil {
		params = &DescribeIdentityProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIdentityProvider", params, optFns, addOperationDescribeIdentityProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIdentityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIdentityProviderInput struct {

	// The identity provider name.
	//
	// This member is required.
	ProviderName *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string
}

type DescribeIdentityProviderOutput struct {

	// The identity provider that was deleted.
	//
	// This member is required.
	IdentityProvider *types.IdentityProviderType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeIdentityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeIdentityProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeIdentityProvider{}, middleware.After)
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
	addOpDescribeIdentityProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIdentityProvider(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeIdentityProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "DescribeIdentityProvider",
	}
}
