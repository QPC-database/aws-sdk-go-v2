// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// For an existing AWS CodeBuild build project that has its source code stored in a
// GitHub or Bitbucket repository, stops AWS CodeBuild from rebuilding the source
// code every time a code change is pushed to the repository.
func (c *Client) DeleteWebhook(ctx context.Context, params *DeleteWebhookInput, optFns ...func(*Options)) (*DeleteWebhookOutput, error) {
	if params == nil {
		params = &DeleteWebhookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteWebhook", params, optFns, addOperationDeleteWebhookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteWebhookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWebhookInput struct {

	// The name of the AWS CodeBuild project.
	//
	// This member is required.
	ProjectName *string
}

type DeleteWebhookOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteWebhookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteWebhook{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteWebhook{}, middleware.After)
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
	addOpDeleteWebhookValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWebhook(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteWebhook(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "DeleteWebhook",
	}
}
