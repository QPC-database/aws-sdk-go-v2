// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a patch baseline.
func (c *Client) DeletePatchBaseline(ctx context.Context, params *DeletePatchBaselineInput, optFns ...func(*Options)) (*DeletePatchBaselineOutput, error) {
	if params == nil {
		params = &DeletePatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePatchBaseline", params, optFns, addOperationDeletePatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePatchBaselineInput struct {

	// The ID of the patch baseline to delete.
	//
	// This member is required.
	BaselineId *string
}

type DeletePatchBaselineOutput struct {

	// The ID of the deleted patch baseline.
	BaselineId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeletePatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePatchBaseline{}, middleware.After)
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
	addOpDeletePatchBaselineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePatchBaseline(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeletePatchBaseline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeletePatchBaseline",
	}
}
