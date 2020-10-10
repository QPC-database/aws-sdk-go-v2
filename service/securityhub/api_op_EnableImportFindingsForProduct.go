// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the integration of a partner product with Security Hub. Integrated
// products send findings to Security Hub. When you enable a product integration, a
// permissions policy that grants permission for the product to send findings to
// Security Hub is applied.
func (c *Client) EnableImportFindingsForProduct(ctx context.Context, params *EnableImportFindingsForProductInput, optFns ...func(*Options)) (*EnableImportFindingsForProductOutput, error) {
	if params == nil {
		params = &EnableImportFindingsForProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableImportFindingsForProduct", params, optFns, addOperationEnableImportFindingsForProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableImportFindingsForProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableImportFindingsForProductInput struct {

	// The ARN of the product to enable the integration for.
	//
	// This member is required.
	ProductArn *string
}

type EnableImportFindingsForProductOutput struct {

	// The ARN of your subscription to the product to enable integrations for.
	ProductSubscriptionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableImportFindingsForProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEnableImportFindingsForProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableImportFindingsForProduct{}, middleware.After)
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
	addOpEnableImportFindingsForProductValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableImportFindingsForProduct(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEnableImportFindingsForProduct(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "EnableImportFindingsForProduct",
	}
}
