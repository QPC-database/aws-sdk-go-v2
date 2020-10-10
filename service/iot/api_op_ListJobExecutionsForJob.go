// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the job executions for a job.
func (c *Client) ListJobExecutionsForJob(ctx context.Context, params *ListJobExecutionsForJobInput, optFns ...func(*Options)) (*ListJobExecutionsForJobOutput, error) {
	if params == nil {
		params = &ListJobExecutionsForJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobExecutionsForJob", params, optFns, addOperationListJobExecutionsForJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobExecutionsForJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobExecutionsForJobInput struct {

	// The unique identifier you assigned to this job when it was created.
	//
	// This member is required.
	JobId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string

	// The status of the job.
	Status types.JobExecutionStatus
}

type ListJobExecutionsForJobOutput struct {

	// A list of job execution summaries.
	ExecutionSummaries []*types.JobExecutionSummaryForJob

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJobExecutionsForJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobExecutionsForJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobExecutionsForJob{}, middleware.After)
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
	addOpListJobExecutionsForJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJobExecutionsForJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListJobExecutionsForJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListJobExecutionsForJob",
	}
}
