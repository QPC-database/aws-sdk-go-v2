// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates device authorization by requesting a pair of verification codes from
// the authorization service.
func (c *Client) StartDeviceAuthorization(ctx context.Context, params *StartDeviceAuthorizationInput, optFns ...func(*Options)) (*StartDeviceAuthorizationOutput, error) {
	if params == nil {
		params = &StartDeviceAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDeviceAuthorization", params, optFns, addOperationStartDeviceAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDeviceAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDeviceAuthorizationInput struct {

	// The unique identifier string for the client that is registered with AWS SSO.
	// This value should come from the persisted result of the RegisterClient () API
	// operation.
	//
	// This member is required.
	ClientId *string

	// A secret string that is generated for the client. This value should come from
	// the persisted result of the RegisterClient () API operation.
	//
	// This member is required.
	ClientSecret *string

	// The URL for the AWS SSO user portal. For more information, see Using the User
	// Portal
	// (https://docs.aws.amazon.com/singlesignon/latest/userguide/using-the-portal.html)
	// in the AWS Single Sign-On User Guide.
	//
	// This member is required.
	StartUrl *string
}

type StartDeviceAuthorizationOutput struct {

	// The short-lived code that is used by the device when polling for a session
	// token.
	DeviceCode *string

	// Indicates the number of seconds in which the verification code will become
	// invalid.
	ExpiresIn *int32

	// Indicates the number of seconds the client must wait between attempts when
	// polling for a session.
	Interval *int32

	// A one-time user verification code. This is needed to authorize an in-use device.
	UserCode *string

	// The URI of the verification page that takes the userCode to authorize the
	// device.
	VerificationUri *string

	// An alternate URL that the client can use to automatically launch a browser. This
	// process skips the manual step in which the user visits the verification page and
	// enters their code.
	VerificationUriComplete *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartDeviceAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartDeviceAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartDeviceAuthorization{}, middleware.After)
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
	addOpStartDeviceAuthorizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeviceAuthorization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartDeviceAuthorization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartDeviceAuthorization",
	}
}
