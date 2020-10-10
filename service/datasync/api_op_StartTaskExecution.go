// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a specific invocation of a task. A TaskExecution value represents an
// individual run of a task. Each task can have at most one TaskExecution at a
// time. TaskExecution has the following transition phases: INITIALIZING |
// PREPARING | TRANSFERRING | VERIFYING | SUCCESS/FAILURE.  <p>For detailed
// information, see the Task Execution section in the Components and Terminology
// topic in the <i>AWS DataSync User Guide</i>.</p>
func (c *Client) StartTaskExecution(ctx context.Context, params *StartTaskExecutionInput, optFns ...func(*Options)) (*StartTaskExecutionOutput, error) {
	if params == nil {
		params = &StartTaskExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTaskExecution", params, optFns, addOperationStartTaskExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTaskExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// StartTaskExecutionRequest
type StartTaskExecutionInput struct {

	// The Amazon Resource Name (ARN) of the task to start.
	//
	// This member is required.
	TaskArn *string

	// A list of filter rules that determines which files to include when running a
	// task. The pattern should contain a single filter string that consists of the
	// patterns to include. The patterns are delimited by "|" (that is, a pipe). For
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
	OverrideOptions *types.Options
}

// StartTaskExecutionResponse
type StartTaskExecutionOutput struct {

	// The Amazon Resource Name (ARN) of the specific task execution that was started.
	TaskExecutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartTaskExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTaskExecution{}, middleware.After)
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
	addOpStartTaskExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartTaskExecution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartTaskExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "StartTaskExecution",
	}
}
