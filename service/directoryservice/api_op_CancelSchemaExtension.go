// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels an in-progress schema extension to a Microsoft AD directory. Once a
// schema extension has started replicating to all domain controllers, the task can
// no longer be canceled. A schema extension can be canceled during any of the
// following states; Initializing, CreatingSnapshot, and UpdatingSchema.
func (c *Client) CancelSchemaExtension(ctx context.Context, params *CancelSchemaExtensionInput, optFns ...func(*Options)) (*CancelSchemaExtensionOutput, error) {
	if params == nil {
		params = &CancelSchemaExtensionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelSchemaExtension", params, optFns, addOperationCancelSchemaExtensionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelSchemaExtensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelSchemaExtensionInput struct {

	// The identifier of the directory whose schema extension will be canceled.
	//
	// This member is required.
	DirectoryId *string

	// The identifier of the schema extension that will be canceled.
	//
	// This member is required.
	SchemaExtensionId *string
}

type CancelSchemaExtensionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelSchemaExtensionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelSchemaExtension{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelSchemaExtension{}, middleware.After)
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
	addOpCancelSchemaExtensionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelSchemaExtension(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCancelSchemaExtension(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "CancelSchemaExtension",
	}
}
