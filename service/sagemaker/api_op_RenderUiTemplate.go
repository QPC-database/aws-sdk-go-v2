// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Renders the UI template so that you can preview the worker's experience.
func (c *Client) RenderUiTemplate(ctx context.Context, params *RenderUiTemplateInput, optFns ...func(*Options)) (*RenderUiTemplateOutput, error) {
	if params == nil {
		params = &RenderUiTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RenderUiTemplate", params, optFns, addOperationRenderUiTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RenderUiTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RenderUiTemplateInput struct {

	// The Amazon Resource Name (ARN) that has access to the S3 objects that are used
	// by the template.
	//
	// This member is required.
	RoleArn *string

	// A RenderableTask object containing a representative task to render.
	//
	// This member is required.
	Task *types.RenderableTask

	// The HumanTaskUiArn of the worker UI that you want to render. Do not provide a
	// HumanTaskUiArn if you use the UiTemplate parameter. See a list of available
	// Human Ui Amazon Resource Names (ARNs) in UiConfig ().
	HumanTaskUiArn *string

	// A Template object containing the worker UI template to render.
	UiTemplate *types.UiTemplate
}

type RenderUiTemplateOutput struct {

	// A list of one or more RenderingError objects if any were encountered while
	// rendering the template. If there were no errors, the list is empty.
	//
	// This member is required.
	Errors []*types.RenderingError

	// A Liquid template that renders the HTML for the worker UI.
	//
	// This member is required.
	RenderedContent *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRenderUiTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRenderUiTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRenderUiTemplate{}, middleware.After)
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
	addOpRenderUiTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRenderUiTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRenderUiTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "RenderUiTemplate",
	}
}
