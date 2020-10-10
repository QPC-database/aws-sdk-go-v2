// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Defines a webhook and returns a unique webhook URL generated by CodePipeline.
// This URL can be supplied to third party source hosting providers to call every
// time there's a code change. When CodePipeline receives a POST request on this
// URL, the pipeline defined in the webhook is started as long as the POST request
// satisfied the authentication and filtering requirements supplied when defining
// the webhook. RegisterWebhookWithThirdParty and DeregisterWebhookWithThirdParty
// APIs can be used to automatically configure supported third parties to call the
// generated webhook URL.
func (c *Client) PutWebhook(ctx context.Context, params *PutWebhookInput, optFns ...func(*Options)) (*PutWebhookOutput, error) {
	if params == nil {
		params = &PutWebhookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutWebhook", params, optFns, addOperationPutWebhookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutWebhookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutWebhookInput struct {

	// The detail provided in an input file to create the webhook, such as the webhook
	// name, the pipeline name, and the action name. Give the webhook a unique name
	// that helps you identify it. You might name the webhook after the pipeline and
	// action it targets so that you can easily recognize what it's used for later.
	//
	// This member is required.
	Webhook *types.WebhookDefinition

	// The tags for the webhook.
	Tags []*types.Tag
}

type PutWebhookOutput struct {

	// The detail returned from creating the webhook, such as the webhook name, webhook
	// URL, and webhook ARN.
	Webhook *types.ListWebhookItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutWebhookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutWebhook{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutWebhook{}, middleware.After)
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
	addOpPutWebhookValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutWebhook(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutWebhook(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "PutWebhook",
	}
}
