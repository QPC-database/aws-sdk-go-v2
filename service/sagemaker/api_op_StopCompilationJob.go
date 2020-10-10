// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a model compilation job. To stop a job, Amazon SageMaker sends the
// algorithm the SIGTERM signal. This gracefully shuts the job down. If the job
// hasn't stopped, it sends the SIGKILL signal. When it receives a
// StopCompilationJob request, Amazon SageMaker changes the
// CompilationJobSummary$CompilationJobStatus () of the job to Stopping. After
// Amazon SageMaker stops the job, it sets the
// CompilationJobSummary$CompilationJobStatus () to Stopped.
func (c *Client) StopCompilationJob(ctx context.Context, params *StopCompilationJobInput, optFns ...func(*Options)) (*StopCompilationJobOutput, error) {
	if params == nil {
		params = &StopCompilationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopCompilationJob", params, optFns, addOperationStopCompilationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopCompilationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopCompilationJobInput struct {

	// The name of the model compilation job to stop.
	//
	// This member is required.
	CompilationJobName *string
}

type StopCompilationJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopCompilationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopCompilationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopCompilationJob{}, middleware.After)
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
	addOpStopCompilationJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopCompilationJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopCompilationJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StopCompilationJob",
	}
}
