// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an identity provider for a user pool.
func (c *Client) DeleteIdentityProvider(ctx context.Context, params *DeleteIdentityProviderInput, optFns ...func(*Options)) (*DeleteIdentityProviderOutput, error) {
	if params == nil {
		params = &DeleteIdentityProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIdentityProvider", params, optFns, addOperationDeleteIdentityProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIdentityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIdentityProviderInput struct {

	// The identity provider name.
	//
	// This member is required.
	ProviderName *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string
}

type DeleteIdentityProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteIdentityProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteIdentityProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteIdentityProvider{}, middleware.After)
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
	addOpDeleteIdentityProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIdentityProvider(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteIdentityProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "DeleteIdentityProvider",
	}
}
