// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Defines the settings you will use for the human review workflow user interface.
// Reviewers will see a three-panel interface with an instruction area, the item to
// review, and an input area.
func (c *Client) CreateHumanTaskUi(ctx context.Context, params *CreateHumanTaskUiInput, optFns ...func(*Options)) (*CreateHumanTaskUiOutput, error) {
	if params == nil {
		params = &CreateHumanTaskUiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHumanTaskUi", params, optFns, addOperationCreateHumanTaskUiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHumanTaskUiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHumanTaskUiInput struct {

	// The name of the user interface you are creating.
	//
	// This member is required.
	HumanTaskUiName *string

	// The Liquid template for the worker user interface.
	//
	// This member is required.
	UiTemplate *types.UiTemplate

	// An array of key-value pairs that contain metadata to help you categorize and
	// organize a human review workflow user interface. Each tag consists of a key and
	// a value, both of which you define.
	Tags []*types.Tag
}

type CreateHumanTaskUiOutput struct {

	// The Amazon Resource Name (ARN) of the human review workflow user interface you
	// create.
	//
	// This member is required.
	HumanTaskUiArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHumanTaskUiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHumanTaskUi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHumanTaskUi{}, middleware.After)
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
	addOpCreateHumanTaskUiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHumanTaskUi(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateHumanTaskUi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateHumanTaskUi",
	}
}
