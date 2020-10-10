// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets an OpenID token, using a known Cognito ID. This known Cognito ID is
// returned by GetId (). You can optionally add additional logins for the identity.
// Supplying multiple logins creates an implicit link. The OpenId token is valid
// for 10 minutes. This is a public API. You do not need any credentials to call
// this API.
func (c *Client) GetOpenIdToken(ctx context.Context, params *GetOpenIdTokenInput, optFns ...func(*Options)) (*GetOpenIdTokenOutput, error) {
	if params == nil {
		params = &GetOpenIdTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOpenIdToken", params, optFns, addOperationGetOpenIdTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOpenIdTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the GetOpenIdToken action.
type GetOpenIdTokenInput struct {

	// A unique identifier in the format REGION:GUID.
	//
	// This member is required.
	IdentityId *string

	// A set of optional name-value pairs that map provider names to provider tokens.
	// When using graph.facebook.com and www.amazon.com, supply the access_token
	// returned from the provider's authflow. For accounts.google.com, an Amazon
	// Cognito user pool provider, or any other OpenId Connect provider, always include
	// the id_token.
	Logins map[string]*string
}

// Returned in response to a successful GetOpenIdToken request.
type GetOpenIdTokenOutput struct {

	// A unique identifier in the format REGION:GUID. Note that the IdentityId returned
	// may not match the one passed on input.
	IdentityId *string

	// An OpenID token, valid for 10 minutes.
	Token *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetOpenIdTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOpenIdToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOpenIdToken{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetOpenIdTokenValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpenIdToken(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetOpenIdToken(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "GetOpenIdToken",
	}
}
