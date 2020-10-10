// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows a user to enter a confirmation code to reset a forgotten password.
func (c *Client) ConfirmForgotPassword(ctx context.Context, params *ConfirmForgotPasswordInput, optFns ...func(*Options)) (*ConfirmForgotPasswordOutput, error) {
	if params == nil {
		params = &ConfirmForgotPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfirmForgotPassword", params, optFns, addOperationConfirmForgotPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfirmForgotPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request representing the confirmation for a password reset.
type ConfirmForgotPasswordInput struct {

	// The app client ID of the app associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The confirmation code sent by a user's request to retrieve a forgotten password.
	// For more information, see
	//
	// This member is required.
	ConfirmationCode *string

	// The password sent by a user's request to retrieve a forgotten password.
	//
	// This member is required.
	Password *string

	// The user name of the user for whom you want to enter a code to retrieve a
	// forgotten password.
	//
	// This member is required.
	Username *string

	// The Amazon Pinpoint analytics metadata for collecting metrics for
	// ConfirmForgotPassword calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the
	// ConfirmForgotPassword API action, Amazon Cognito invokes the function that is
	// assigned to the post confirmation trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. This
	// payload contains a clientMetadata attribute, which provides the data that you
	// assigned to the ClientMetadata parameter in your ConfirmForgotPassword request.
	// In your function code in AWS Lambda, you can process the clientMetadata value to
	// enhance your workflow for your specific needs. For more information, see
	// Customizing User Pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	//     * Amazon Cognito
	// does not store the ClientMetadata value. This data is available only to AWS
	// Lambda triggers that are assigned to a user pool to support custom workflows. If
	// your user pool configuration does not include triggers, the ClientMetadata
	// parameter serves no purpose.
	//
	//     * Amazon Cognito does not validate the
	// ClientMetadata value.
	//
	//     * Amazon Cognito does not encrypt the the
	// ClientMetadata value, so don't use it to provide sensitive information.
	ClientMetadata map[string]*string

	// A keyed-hash message authentication code (HMAC) calculated using the secret key
	// of a user pool client and username plus the client ID in the message.
	SecretHash *string

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	UserContextData *types.UserContextDataType
}

// The response from the server that results from a user's request to retrieve a
// forgotten password.
type ConfirmForgotPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationConfirmForgotPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpConfirmForgotPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpConfirmForgotPassword{}, middleware.After)
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
	addOpConfirmForgotPasswordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opConfirmForgotPassword(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opConfirmForgotPassword(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConfirmForgotPassword",
	}
}
