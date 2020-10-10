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

// Creates the user pool client.
func (c *Client) CreateUserPoolClient(ctx context.Context, params *CreateUserPoolClientInput, optFns ...func(*Options)) (*CreateUserPoolClientOutput, error) {
	if params == nil {
		params = &CreateUserPoolClientInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserPoolClient", params, optFns, addOperationCreateUserPoolClientMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserPoolClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to create a user pool client.
type CreateUserPoolClientInput struct {

	// The client name for the user pool client you would like to create.
	//
	// This member is required.
	ClientName *string

	// The user pool ID for the user pool where you want to create a user pool client.
	//
	// This member is required.
	UserPoolId *string

	// The allowed OAuth flows. Set to code to initiate a code grant flow, which
	// provides an authorization code as the response. This code can be exchanged for
	// access tokens with the token endpoint. Set to implicit to specify that the
	// client should get the access token (and, optionally, ID token, based on scopes)
	// directly. Set to client_credentials to specify that the client should get the
	// access token (and, optionally, ID token, based on scopes) from the token
	// endpoint using a combination of client and client_secret.
	AllowedOAuthFlows []types.OAuthFlowType

	// Set to true if the client is allowed to follow the OAuth protocol when
	// interacting with Cognito user pools.
	AllowedOAuthFlowsUserPoolClient *bool

	// The allowed OAuth scopes. Possible values provided by OAuth are: phone, email,
	// openid, and profile. Possible values provided by AWS are:
	// aws.cognito.signin.user.admin. Custom scopes created in Resource Servers are
	// also supported.
	AllowedOAuthScopes []*string

	// The Amazon Pinpoint analytics configuration for collecting metrics for this user
	// pool. Cognito User Pools only supports sending events to Amazon Pinpoint
	// projects in the US East (N. Virginia) us-east-1 Region, regardless of the region
	// in which the user pool resides.
	AnalyticsConfiguration *types.AnalyticsConfigurationType

	// A list of allowed redirect (callback) URLs for the identity providers. A
	// redirect URI must:
	//
	//     * Be an absolute URI.
	//
	//     * Be registered with the
	// authorization server.
	//
	//     * Not include a fragment component.
	//
	// See OAuth 2.0 -
	// Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2). Amazon
	// Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only. App callback URLs such as myapp://example are also supported.
	CallbackURLs []*string

	// The default redirect URI. Must be in the CallbackURLs list. A redirect URI
	// must:
	//
	//     * Be an absolute URI.
	//
	//     * Be registered with the authorization
	// server.
	//
	//     * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection
	// Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2). Amazon Cognito
	// requires HTTPS over HTTP except for http://localhost for testing purposes only.
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectURI *string

	// The authentication flows that are supported by the user pool clients. Flow names
	// without the ALLOW_ prefix are deprecated in favor of new names with the ALLOW_
	// prefix. Note that values with ALLOW_ prefix cannot be used along with values
	// without ALLOW_ prefix. Valid values include:
	//
	//     *
	// ALLOW_ADMIN_USER_PASSWORD_AUTH: Enable admin based user password authentication
	// flow ADMIN_USER_PASSWORD_AUTH. This setting replaces the ADMIN_NO_SRP_AUTH
	// setting. With this authentication flow, Cognito receives the password in the
	// request instead of using the SRP (Secure Remote Password protocol) protocol to
	// verify passwords.
	//
	//     * ALLOW_CUSTOM_AUTH: Enable Lambda trigger based
	// authentication.
	//
	//     * ALLOW_USER_PASSWORD_AUTH: Enable user password-based
	// authentication. In this flow, Cognito receives the password in the request
	// instead of using the SRP protocol to verify passwords.
	//
	//     *
	// ALLOW_USER_SRP_AUTH: Enable SRP based authentication.
	//
	//     *
	// ALLOW_REFRESH_TOKEN_AUTH: Enable authflow to refresh tokens.
	ExplicitAuthFlows []types.ExplicitAuthFlowsType

	// Boolean to specify whether you want to generate a secret for the user pool
	// client being created.
	GenerateSecret *bool

	// A list of allowed logout URLs for the identity providers.
	LogoutURLs []*string

	// Use this setting to choose which errors and responses are returned by Cognito
	// APIs during authentication, account confirmation, and password recovery when the
	// user does not exist in the user pool. When set to ENABLED and the user does not
	// exist, authentication returns an error indicating either the username or
	// password was incorrect, and account confirmation and password recovery return a
	// response indicating a code was sent to a simulated destination. When set to
	// LEGACY, those APIs will return a UserNotFoundException exception if the user
	// does not exist in the user pool. Valid values include:
	//
	//     * ENABLED - This
	// prevents user existence-related errors.
	//
	//     * LEGACY - This represents the old
	// behavior of Cognito where user existence related errors are not prevented.
	//
	// This
	// setting affects the behavior of following APIs:
	//
	//     * AdminInitiateAuth ()
	//
	//
	// * AdminRespondToAuthChallenge ()
	//
	//     * InitiateAuth ()
	//
	//     *
	// RespondToAuthChallenge ()
	//
	//     * ForgotPassword ()
	//
	//     * ConfirmForgotPassword
	// ()
	//
	//     * ConfirmSignUp ()
	//
	//     * ResendConfirmationCode ()
	//
	// After February 15th
	// 2020, the value of PreventUserExistenceErrors will default to ENABLED for newly
	// created user pool clients if no value is provided.
	PreventUserExistenceErrors types.PreventUserExistenceErrorTypes

	// The read attributes.
	ReadAttributes []*string

	// The time limit, in days, after which the refresh token is no longer valid and
	// cannot be used.
	RefreshTokenValidity *int32

	// A list of provider names for the identity providers that are supported on this
	// client. The following are supported: COGNITO, Facebook, Google and
	// LoginWithAmazon.
	SupportedIdentityProviders []*string

	// The user pool attributes that the app client can write to. If your app client
	// allows users to sign in through an identity provider, this array must include
	// all attributes that are mapped to identity provider attributes. Amazon Cognito
	// updates mapped attributes when users sign in to your application through an
	// identity provider. If your app client lacks write access to a mapped attribute,
	// Amazon Cognito throws an error when it attempts to update the attribute. For
	// more information, see Specifying Identity Provider Attribute Mappings for Your
	// User Pool
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html).
	WriteAttributes []*string
}

// Represents the response from the server to create a user pool client.
type CreateUserPoolClientOutput struct {

	// The user pool client that was just created.
	UserPoolClient *types.UserPoolClientType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateUserPoolClientMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserPoolClient{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserPoolClient{}, middleware.After)
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
	addOpCreateUserPoolClientValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserPoolClient(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateUserPoolClient(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "CreateUserPoolClient",
	}
}
