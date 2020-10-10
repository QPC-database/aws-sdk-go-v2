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

// Retrieves the content and settings of a message template for messages that are
// sent through the email channel.
func (c *Client) GetEmailTemplate(ctx context.Context, params *GetEmailTemplateInput, optFns ...func(*Options)) (*GetEmailTemplateOutput, error) {
	if params == nil {
		params = &GetEmailTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEmailTemplate", params, optFns, addOperationGetEmailTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEmailTemplateInput struct {

	// The name of the message template. A template name must start with an
	// alphanumeric character and can contain a maximum of 128 characters. The
	// characters can be alphanumeric characters, underscores (_), or hyphens (-).
	// Template names are case sensitive.
	//
	// This member is required.
	TemplateName *string

	// The unique identifier for the version of the message template to update,
	// retrieve information about, or delete. To retrieve identifiers and other
	// information for all the versions of a template, use the Template Versions
	// resource. If specified, this value must match the identifier for an existing
	// template version. If specified for an update operation, this value must match
	// the identifier for the latest existing version of the template. This restriction
	// helps ensure that race conditions don't occur. If you don't specify a value for
	// this parameter, Amazon Pinpoint does the following:
	//
	//     * For a get operation,
	// retrieves information about the active version of the template.
	//
	//     * For an
	// update operation, saves the updates to (overwrites) the latest existing version
	// of the template, if the create-new-version parameter isn't used or is set to
	// false.
	//
	//     * For a delete operation, deletes the template, including all
	// versions of the template.
	Version *string
}

type GetEmailTemplateOutput struct {

	// Provides information about the content and settings for a message template that
	// can be used in messages that are sent through the email channel.
	//
	// This member is required.
	EmailTemplateResponse *types.EmailTemplateResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEmailTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEmailTemplate{}, middleware.After)
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
	addOpGetEmailTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEmailTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetEmailTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetEmailTemplate",
	}
}
