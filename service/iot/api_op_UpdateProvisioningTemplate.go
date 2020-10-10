// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a fleet provisioning template.
func (c *Client) UpdateProvisioningTemplate(ctx context.Context, params *UpdateProvisioningTemplateInput, optFns ...func(*Options)) (*UpdateProvisioningTemplateOutput, error) {
	if params == nil {
		params = &UpdateProvisioningTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProvisioningTemplate", params, optFns, addOperationUpdateProvisioningTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProvisioningTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProvisioningTemplateInput struct {

	// The name of the fleet provisioning template.
	//
	// This member is required.
	TemplateName *string

	// The ID of the default provisioning template version.
	DefaultVersionId *int32

	// The description of the fleet provisioning template.
	Description *string

	// True to enable the fleet provisioning template, otherwise false.
	Enabled *bool

	// Updates the pre-provisioning hook template.
	PreProvisioningHook *types.ProvisioningHook

	// The ARN of the role associated with the provisioning template. This IoT role
	// grants permission to provision a device.
	ProvisioningRoleArn *string

	// Removes pre-provisioning hook template.
	RemovePreProvisioningHook *bool
}

type UpdateProvisioningTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateProvisioningTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProvisioningTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProvisioningTemplate{}, middleware.After)
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
	addOpUpdateProvisioningTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProvisioningTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateProvisioningTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateProvisioningTemplate",
	}
}
