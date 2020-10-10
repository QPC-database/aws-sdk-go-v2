// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about a fleet provisioning template.
func (c *Client) DescribeProvisioningTemplate(ctx context.Context, params *DescribeProvisioningTemplateInput, optFns ...func(*Options)) (*DescribeProvisioningTemplateOutput, error) {
	if params == nil {
		params = &DescribeProvisioningTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProvisioningTemplate", params, optFns, addOperationDescribeProvisioningTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProvisioningTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProvisioningTemplateInput struct {

	// The name of the fleet provisioning template.
	//
	// This member is required.
	TemplateName *string
}

type DescribeProvisioningTemplateOutput struct {

	// The date when the fleet provisioning template was created.
	CreationDate *time.Time

	// The default fleet template version ID.
	DefaultVersionId *int32

	// The description of the fleet provisioning template.
	Description *string

	// True if the fleet provisioning template is enabled, otherwise false.
	Enabled *bool

	// The date when the fleet provisioning template was last modified.
	LastModifiedDate *time.Time

	// Gets information about a pre-provisioned hook.
	PreProvisioningHook *types.ProvisioningHook

	// The ARN of the role associated with the provisioning template. This IoT role
	// grants permission to provision a device.
	ProvisioningRoleArn *string

	// The ARN of the fleet provisioning template.
	TemplateArn *string

	// The JSON formatted contents of the fleet provisioning template.
	TemplateBody *string

	// The name of the fleet provisioning template.
	TemplateName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeProvisioningTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProvisioningTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProvisioningTemplate{}, middleware.After)
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
	addOpDescribeProvisioningTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProvisioningTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeProvisioningTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeProvisioningTemplate",
	}
}
