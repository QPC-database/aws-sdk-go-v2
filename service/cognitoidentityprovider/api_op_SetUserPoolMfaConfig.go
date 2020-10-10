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

// Set the user pool multi-factor authentication (MFA) configuration.
func (c *Client) SetUserPoolMfaConfig(ctx context.Context, params *SetUserPoolMfaConfigInput, optFns ...func(*Options)) (*SetUserPoolMfaConfigOutput, error) {
	if params == nil {
		params = &SetUserPoolMfaConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetUserPoolMfaConfig", params, optFns, addOperationSetUserPoolMfaConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetUserPoolMfaConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetUserPoolMfaConfigInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The MFA configuration. Valid values include:
	//
	//     * OFF MFA will not be used for
	// any users.
	//
	//     * ON MFA is required for all users to sign in.
	//
	//     * OPTIONAL
	// MFA will be required only for individual users who have an MFA factor enabled.
	MfaConfiguration types.UserPoolMfaType

	// The SMS text message MFA configuration.
	SmsMfaConfiguration *types.SmsMfaConfigType

	// The software token MFA configuration.
	SoftwareTokenMfaConfiguration *types.SoftwareTokenMfaConfigType
}

type SetUserPoolMfaConfigOutput struct {

	// The MFA configuration. Valid values include:
	//
	//     * OFF MFA will not be used for
	// any users.
	//
	//     * ON MFA is required for all users to sign in.
	//
	//     * OPTIONAL
	// MFA will be required only for individual users who have an MFA factor enabled.
	MfaConfiguration types.UserPoolMfaType

	// The SMS text message MFA configuration.
	SmsMfaConfiguration *types.SmsMfaConfigType

	// The software token MFA configuration.
	SoftwareTokenMfaConfiguration *types.SoftwareTokenMfaConfigType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetUserPoolMfaConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetUserPoolMfaConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetUserPoolMfaConfig{}, middleware.After)
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
	addOpSetUserPoolMfaConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetUserPoolMfaConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetUserPoolMfaConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "SetUserPoolMfaConfig",
	}
}
