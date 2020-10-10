// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a client with AWS SSO. This allows clients to initiate device
// authorization. The output should be persisted for reuse through many
// authentication requests.
func (c *Client) RegisterClient(ctx context.Context, params *RegisterClientInput, optFns ...func(*Options)) (*RegisterClientOutput, error) {
	if params == nil {
		params = &RegisterClientInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterClient", params, optFns, addOperationRegisterClientMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterClientInput struct {

	// The friendly name of the client.
	//
	// This member is required.
	ClientName *string

	// The type of client. The service supports only public as a client type. Anything
	// other than public will be rejected by the service.
	//
	// This member is required.
	ClientType *string

	// The list of scopes that are defined by the client. Upon authorization, this list
	// is used to restrict permissions when granting an access token.
	Scopes []*string
}

type RegisterClientOutput struct {

	// The endpoint where the client can request authorization.
	AuthorizationEndpoint *string

	// The unique identifier string for each client. This client uses this identifier
	// to get authenticated by the service in subsequent calls.
	ClientId *string

	// Indicates the time at which the clientId and clientSecret were issued.
	ClientIdIssuedAt *int64

	// A secret string generated for the client. The client will use this string to get
	// authenticated by the service in subsequent calls.
	ClientSecret *string

	// Indicates the time at which the clientId and clientSecret will become invalid.
	ClientSecretExpiresAt *int64

	// The endpoint where the client can get an access token.
	TokenEndpoint *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterClientMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterClient{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterClient{}, middleware.After)
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
	addOpRegisterClientValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterClient(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterClient(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterClient",
	}
}
