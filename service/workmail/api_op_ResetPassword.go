// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows the administrator to reset the password for a user.
func (c *Client) ResetPassword(ctx context.Context, params *ResetPasswordInput, optFns ...func(*Options)) (*ResetPasswordOutput, error) {
	if params == nil {
		params = &ResetPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetPassword", params, optFns, addOperationResetPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetPasswordInput struct {

	// The identifier of the organization that contains the user for which the password
	// is reset.
	//
	// This member is required.
	OrganizationId *string

	// The new password for the user.
	//
	// This member is required.
	Password *string

	// The identifier of the user for whom the password is reset.
	//
	// This member is required.
	UserId *string
}

type ResetPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResetPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResetPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetPassword{}, middleware.After)
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
	addOpResetPasswordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetPassword(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opResetPassword(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "ResetPassword",
	}
}
