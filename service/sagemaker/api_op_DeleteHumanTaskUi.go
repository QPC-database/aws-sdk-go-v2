// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this operation to delete a human task user interface (worker task template).
// To see a list of human task user interfaces (work task templates) in your
// account, use . When you delete a worker task template, it no longer appears when
// you call ListHumanTaskUis.
func (c *Client) DeleteHumanTaskUi(ctx context.Context, params *DeleteHumanTaskUiInput, optFns ...func(*Options)) (*DeleteHumanTaskUiOutput, error) {
	if params == nil {
		params = &DeleteHumanTaskUiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteHumanTaskUi", params, optFns, addOperationDeleteHumanTaskUiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteHumanTaskUiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteHumanTaskUiInput struct {

	// The name of the human task user interface (work task template) you want to
	// delete.
	//
	// This member is required.
	HumanTaskUiName *string
}

type DeleteHumanTaskUiOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteHumanTaskUiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteHumanTaskUi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteHumanTaskUi{}, middleware.After)
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
	addOpDeleteHumanTaskUiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHumanTaskUi(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteHumanTaskUi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DeleteHumanTaskUi",
	}
}
