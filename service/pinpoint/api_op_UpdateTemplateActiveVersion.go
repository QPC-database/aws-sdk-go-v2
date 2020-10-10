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

// Changes the status of a specific version of a message template to active.
func (c *Client) UpdateTemplateActiveVersion(ctx context.Context, params *UpdateTemplateActiveVersionInput, optFns ...func(*Options)) (*UpdateTemplateActiveVersionOutput, error) {
	if params == nil {
		params = &UpdateTemplateActiveVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTemplateActiveVersion", params, optFns, addOperationUpdateTemplateActiveVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTemplateActiveVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTemplateActiveVersionInput struct {

	// Specifies which version of a message template to use as the active version of
	// the template.
	//
	// This member is required.
	TemplateActiveVersionRequest *types.TemplateActiveVersionRequest

	// The name of the message template. A template name must start with an
	// alphanumeric character and can contain a maximum of 128 characters. The
	// characters can be alphanumeric characters, underscores (_), or hyphens (-).
	// Template names are case sensitive.
	//
	// This member is required.
	TemplateName *string

	// The type of channel that the message template is designed for. Valid values are:
	// EMAIL, PUSH, SMS, and VOICE.
	//
	// This member is required.
	TemplateType *string
}

type UpdateTemplateActiveVersionOutput struct {

	// Provides information about an API request or response.
	//
	// This member is required.
	MessageBody *types.MessageBody

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateTemplateActiveVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTemplateActiveVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTemplateActiveVersion{}, middleware.After)
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
	addOpUpdateTemplateActiveVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTemplateActiveVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateTemplateActiveVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateTemplateActiveVersion",
	}
}
