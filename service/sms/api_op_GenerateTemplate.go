// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates an Amazon CloudFormation template based on the current launch
// configuration and writes it to an Amazon S3 object in the customer’s Amazon S3
// bucket.
func (c *Client) GenerateTemplate(ctx context.Context, params *GenerateTemplateInput, optFns ...func(*Options)) (*GenerateTemplateOutput, error) {
	if params == nil {
		params = &GenerateTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateTemplate", params, optFns, addOperationGenerateTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateTemplateInput struct {

	// ID of the application associated with the Amazon CloudFormation template.
	AppId *string

	// Format for generating the Amazon CloudFormation template.
	TemplateFormat types.OutputFormat
}

type GenerateTemplateOutput struct {

	// Location of the Amazon S3 object.
	S3Location *types.S3Location

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGenerateTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateTemplate{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGenerateTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "GenerateTemplate",
	}
}
