// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns detailed metadata about a task that is being executed.
func (c *Client) DescribeTaskExecution(ctx context.Context, params *DescribeTaskExecutionInput, optFns ...func(*Options)) (*DescribeTaskExecutionOutput, error) {
	if params == nil {
		params = &DescribeTaskExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTaskExecution", params, optFns, addOperationDescribeTaskExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTaskExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTaskExecutionRequest
type DescribeTaskExecutionInput struct {

	// The Amazon Resource Name (ARN) of the task that is being executed.
	//
	// This member is required.
	TaskExecutionArn *string
}

// DescribeTaskExecutionResponse
type DescribeTaskExecutionOutput struct {

	// The physical number of bytes transferred over the network.
	BytesTransferred *int64

	// The number of logical bytes written to the destination AWS storage resource.
	BytesWritten *int64

	// The estimated physical number of bytes that is to be transferred over the
	// network.
	EstimatedBytesToTransfer *int64

	// The expected number of files that is to be transferred over the network. This
	// value is calculated during the PREPARING phase, before the TRANSFERRING phase.
	// This value is the expected number of files to be transferred. It's calculated
	// based on comparing the content of the source and destination locations and
	// finding the delta that needs to be transferred.
	EstimatedFilesToTransfer *int64

	// A list of filter rules that determines which files to exclude from a task. The
	// list should contain a single filter string that consists of the patterns to
	// exclude. The patterns are delimited by "|" (that is, a pipe), for example:
	// "/folder1|/folder2"
	Excludes []*types.FilterRule

	// The actual number of files that was transferred over the network. This value is
	// calculated and updated on an ongoing basis during the TRANSFERRING phase. It's
	// updated periodically when each file is read from the source and sent over the
	// network. If failures occur during a transfer, this value can be less than
	// EstimatedFilesToTransfer. This value can also be greater than
	// EstimatedFilesTransferred in some cases. This element is implementation-specific
	// for some location types, so don't use it as an indicator for a correct file
	// number or to monitor your task execution.
	FilesTransferred *int64

	// A list of filter rules that determines which files to include when running a
	// task. The list should contain a single filter string that consists of the
	// patterns to include. The patterns are delimited by "|" (that is, a pipe), for
	// example: "/folder1|/folder2"
	Includes []*types.FilterRule

	// Represents the options that are available to control the behavior of a
	// StartTaskExecution () operation. Behavior includes preserving metadata such as
	// user ID (UID), group ID (GID), and file permissions, and also overwriting files
	// in the destination, data integrity verification, and so on. A task has a set of
	// default options associated with it. If you don't specify an option in
	// StartTaskExecution (), the default value is used. You can override the defaults
	// options on each task execution by specifying an overriding Options value to
	// StartTaskExecution ().
	Options *types.Options

	// The result of the task execution.
	Result *types.TaskExecutionResultDetail

	// The time that the task execution was started.
	StartTime *time.Time

	// The status of the task execution.  <p>For detailed information about task
	// execution statuses, see Understanding Task Statuses in the <i>AWS DataSync User
	// Guide.</i> </p>
	Status types.TaskExecutionStatus

	// The Amazon Resource Name (ARN) of the task execution that was described.
	// TaskExecutionArn is hierarchical and includes TaskArn for the task that was
	// executed. For example, a TaskExecution value with the ARN
	// arn:aws:datasync:us-east-1:111222333444:task/task-0208075f79cedf4a2/execution/exec-08ef1e88ec491019b
	// executed the task with the ARN
	// arn:aws:datasync:us-east-1:111222333444:task/task-0208075f79cedf4a2.
	TaskExecutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTaskExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTaskExecution{}, middleware.After)
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
	addOpDescribeTaskExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTaskExecution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTaskExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeTaskExecution",
	}
}
