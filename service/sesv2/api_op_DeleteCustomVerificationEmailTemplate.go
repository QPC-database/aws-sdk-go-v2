// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an existing custom verification email template. For more information
// about custom verification email templates, see Using Custom Verification Email
// Templates
// (https://docs.aws.amazon.com/es/latest/DeveloperGuide/send-email-verify-address-custom.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) DeleteCustomVerificationEmailTemplate(ctx context.Context, params *DeleteCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*DeleteCustomVerificationEmailTemplateOutput, error) {
	if params == nil {
		params = &DeleteCustomVerificationEmailTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomVerificationEmailTemplate", params, optFns, addOperationDeleteCustomVerificationEmailTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomVerificationEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete an existing custom verification email template.
type DeleteCustomVerificationEmailTemplateInput struct {

	// The name of the custom verification email template that you want to delete.
	//
	// This member is required.
	TemplateName *string
}

// If the action is successful, the service sends back an HTTP 200 response with an
// empty HTTP body.
type DeleteCustomVerificationEmailTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCustomVerificationEmailTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteCustomVerificationEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteCustomVerificationEmailTemplate{}, middleware.After)
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
	addOpDeleteCustomVerificationEmailTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomVerificationEmailTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCustomVerificationEmailTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteCustomVerificationEmailTemplate",
	}
}
