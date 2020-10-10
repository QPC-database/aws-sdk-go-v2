// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a specified approval rule template.
func (c *Client) GetApprovalRuleTemplate(ctx context.Context, params *GetApprovalRuleTemplateInput, optFns ...func(*Options)) (*GetApprovalRuleTemplateOutput, error) {
	if params == nil {
		params = &GetApprovalRuleTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApprovalRuleTemplate", params, optFns, addOperationGetApprovalRuleTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApprovalRuleTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApprovalRuleTemplateInput struct {

	// The name of the approval rule template for which you want to get information.
	//
	// This member is required.
	ApprovalRuleTemplateName *string
}

type GetApprovalRuleTemplateOutput struct {

	// The content and structure of the approval rule template.
	//
	// This member is required.
	ApprovalRuleTemplate *types.ApprovalRuleTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetApprovalRuleTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetApprovalRuleTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetApprovalRuleTemplate{}, middleware.After)
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
	addOpGetApprovalRuleTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetApprovalRuleTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetApprovalRuleTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetApprovalRuleTemplate",
	}
}
