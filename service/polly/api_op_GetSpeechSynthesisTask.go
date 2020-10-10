// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a specific SpeechSynthesisTask object based on its TaskID. This object
// contains information about the given speech synthesis task, including the status
// of the task, and a link to the S3 bucket containing the output of the task.
func (c *Client) GetSpeechSynthesisTask(ctx context.Context, params *GetSpeechSynthesisTaskInput, optFns ...func(*Options)) (*GetSpeechSynthesisTaskOutput, error) {
	if params == nil {
		params = &GetSpeechSynthesisTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSpeechSynthesisTask", params, optFns, addOperationGetSpeechSynthesisTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSpeechSynthesisTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSpeechSynthesisTaskInput struct {

	// The Amazon Polly generated identifier for a speech synthesis task.
	//
	// This member is required.
	TaskId *string
}

type GetSpeechSynthesisTaskOutput struct {

	// SynthesisTask object that provides information from the requested task,
	// including output format, creation time, task status, and so on.
	SynthesisTask *types.SynthesisTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSpeechSynthesisTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSpeechSynthesisTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSpeechSynthesisTask{}, middleware.After)
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
	addOpGetSpeechSynthesisTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSpeechSynthesisTask(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSpeechSynthesisTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "GetSpeechSynthesisTask",
	}
}
