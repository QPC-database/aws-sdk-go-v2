// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a message template for messages that are sent through a push
// notification channel.
func (c *Client) CreatePushTemplate(ctx context.Context, params *CreatePushTemplateInput, optFns ...func(*Options)) (*CreatePushTemplateOutput, error) {
	if params == nil {
		params = &CreatePushTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePushTemplate", params, optFns, addOperationCreatePushTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePushTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePushTemplateInput struct {

	// Specifies the content and settings for a message template that can be used in
	// messages that are sent through a push notification channel.
	//
	// This member is required.
	PushNotificationTemplateRequest *types.PushNotificationTemplateRequest

	// The name of the message template. A template name must start with an
	// alphanumeric character and can contain a maximum of 128 characters. The
	// characters can be alphanumeric characters, underscores (_), or hyphens (-).
	// Template names are case sensitive.
	//
	// This member is required.
	TemplateName *string
}

type CreatePushTemplateOutput struct {

	// Provides information about a request to create a message template.
	//
	// This member is required.
	CreateTemplateMessageBody *types.CreateTemplateMessageBody

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePushTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePushTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePushTemplate{}, middleware.After)
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
	addOpCreatePushTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePushTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreatePushTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "CreatePushTemplate",
	}
}
