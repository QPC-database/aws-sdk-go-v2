// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels a job in an AWS Batch job queue. Jobs that are in the SUBMITTED,
// PENDING, or RUNNABLE state are cancelled. Jobs that have progressed to STARTING
// or RUNNING are not cancelled (but the API operation still succeeds, even if no
// job is cancelled); these jobs must be terminated with the TerminateJob ()
// operation.
func (c *Client) CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) {
	if params == nil {
		params = &CancelJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelJob", params, optFns, addOperationCancelJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelJobInput struct {

	// The AWS Batch job ID of the job to cancel.
	//
	// This member is required.
	JobId *string

	// A message to attach to the job that explains the reason for canceling it. This
	// message is returned by future DescribeJobs () operations on the job. This
	// message is also recorded in the AWS Batch activity logs.
	//
	// This member is required.
	Reason *string
}

type CancelJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelJob{}, middleware.After)
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
	addOpCancelJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCancelJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "CancelJob",
	}
}
