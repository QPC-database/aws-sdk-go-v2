// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used by activity workers to report to the service that the ActivityTask ()
// represented by the specified taskToken is still making progress. The worker can
// also specify details of the progress, for example percent complete, using the
// details parameter. This action can also be used by the worker as a mechanism to
// check if cancellation is being requested for the activity task. If a
// cancellation is being attempted for the specified task, then the boolean
// cancelRequested flag returned by the service is set to true. This action resets
// the taskHeartbeatTimeout clock. The taskHeartbeatTimeout is specified in
// RegisterActivityType (). This action doesn't in itself create an event in the
// workflow execution history. However, if the task times out, the workflow
// execution history contains a ActivityTaskTimedOut event that contains the
// information from the last heartbeat generated by the activity worker. The
// taskStartToCloseTimeout of an activity type is the maximum duration of an
// activity task, regardless of the number of RecordActivityTaskHeartbeat ()
// requests received. The taskStartToCloseTimeout is also specified in
// RegisterActivityType (). This operation is only useful for long-lived activities
// to report liveliness of the task and to determine if a cancellation is being
// attempted. If the cancelRequested flag returns true, a cancellation is being
// attempted. If the worker can cancel the activity, it should respond with
// RespondActivityTaskCanceled (). Otherwise, it should ignore the cancellation
// request. Access Control You can use IAM policies to control this action's access
// to Amazon SWF resources as follows:
//
//     * Use a Resource element with the
// domain name to limit the action to only specified domains.
//
//     * Use an Action
// element to allow or deny permission to call this action.
//
//     * You cannot use
// an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have
// sufficient permissions to invoke the action, or the parameter values fall
// outside the specified constraints, the action fails. The associated event
// attribute's cause parameter is set to OPERATION_NOT_PERMITTED. For details and
// example IAM policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) RecordActivityTaskHeartbeat(ctx context.Context, params *RecordActivityTaskHeartbeatInput, optFns ...func(*Options)) (*RecordActivityTaskHeartbeatOutput, error) {
	if params == nil {
		params = &RecordActivityTaskHeartbeatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RecordActivityTaskHeartbeat", params, optFns, addOperationRecordActivityTaskHeartbeatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RecordActivityTaskHeartbeatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecordActivityTaskHeartbeatInput struct {

	// The taskToken of the ActivityTask (). taskToken is generated by the service and
	// should be treated as an opaque value. If the task is passed to another process,
	// its taskToken must also be passed. This enables it to provide its progress and
	// respond with results.
	//
	// This member is required.
	TaskToken *string

	// If specified, contains details about the progress of the task.
	Details *string
}

// Status information about an activity task.
type RecordActivityTaskHeartbeatOutput struct {

	// Set to true if cancellation of the task is requested.
	//
	// This member is required.
	CancelRequested *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRecordActivityTaskHeartbeatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRecordActivityTaskHeartbeat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRecordActivityTaskHeartbeat{}, middleware.After)
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
	addOpRecordActivityTaskHeartbeatValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRecordActivityTaskHeartbeat(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRecordActivityTaskHeartbeat(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "RecordActivityTaskHeartbeat",
	}
}
