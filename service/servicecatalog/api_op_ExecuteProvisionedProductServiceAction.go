// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Executes a self-service action against a provisioned product.
func (c *Client) ExecuteProvisionedProductServiceAction(ctx context.Context, params *ExecuteProvisionedProductServiceActionInput, optFns ...func(*Options)) (*ExecuteProvisionedProductServiceActionOutput, error) {
	if params == nil {
		params = &ExecuteProvisionedProductServiceActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteProvisionedProductServiceAction", params, optFns, addOperationExecuteProvisionedProductServiceActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteProvisionedProductServiceActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteProvisionedProductServiceActionInput struct {

	// An idempotency token that uniquely identifies the execute request.
	//
	// This member is required.
	ExecuteToken *string

	// The identifier of the provisioned product.
	//
	// This member is required.
	ProvisionedProductId *string

	// The self-service action identifier. For example, act-fs7abcd89wxyz.
	//
	// This member is required.
	ServiceActionId *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string

	// A map of all self-service action parameters and their values. If a provided
	// parameter is of a special type, such as TARGET, the provided value will override
	// the default value generated by AWS Service Catalog. If the parameters field is
	// not provided, no additional parameters are passed and default values will be
	// used for any special parameters such as TARGET.
	Parameters map[string][]*string
}

type ExecuteProvisionedProductServiceActionOutput struct {

	// An object containing detailed information about the result of provisioning the
	// product.
	RecordDetail *types.RecordDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExecuteProvisionedProductServiceActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpExecuteProvisionedProductServiceAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpExecuteProvisionedProductServiceAction{}, middleware.After)
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
	addIdempotencyToken_opExecuteProvisionedProductServiceActionMiddleware(stack, options)
	addOpExecuteProvisionedProductServiceActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteProvisionedProductServiceAction(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpExecuteProvisionedProductServiceAction struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpExecuteProvisionedProductServiceAction) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpExecuteProvisionedProductServiceAction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ExecuteProvisionedProductServiceActionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ExecuteProvisionedProductServiceActionInput ")
	}

	if input.ExecuteToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ExecuteToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opExecuteProvisionedProductServiceActionMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpExecuteProvisionedProductServiceAction{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opExecuteProvisionedProductServiceAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ExecuteProvisionedProductServiceAction",
	}
}
