// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the display name of a trial.
func (c *Client) UpdateTrial(ctx context.Context, params *UpdateTrialInput, optFns ...func(*Options)) (*UpdateTrialOutput, error) {
	if params == nil {
		params = &UpdateTrialInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTrial", params, optFns, addOperationUpdateTrialMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTrialInput struct {

	// The name of the trial to update.
	//
	// This member is required.
	TrialName *string

	// The name of the trial as displayed. The name doesn't need to be unique. If
	// DisplayName isn't specified, TrialName is displayed.
	DisplayName *string
}

type UpdateTrialOutput struct {

	// The Amazon Resource Name (ARN) of the trial.
	TrialArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateTrialMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTrial{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTrial{}, middleware.After)
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
	addOpUpdateTrialValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTrial(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateTrial(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateTrial",
	}
}
