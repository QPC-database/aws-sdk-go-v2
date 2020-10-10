// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Runs commands on one or more managed instances.
func (c *Client) SendCommand(ctx context.Context, params *SendCommandInput, optFns ...func(*Options)) (*SendCommandOutput, error) {
	if params == nil {
		params = &SendCommandInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendCommand", params, optFns, addOperationSendCommandMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendCommandOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendCommandInput struct {

	// Required. The name of the Systems Manager document to run. This can be a public
	// document or a custom document.
	//
	// This member is required.
	DocumentName *string

	// Enables Systems Manager to send Run Command output to Amazon CloudWatch Logs.
	CloudWatchOutputConfig *types.CloudWatchOutputConfig

	// User-specified information about the command, such as a brief description of
	// what the command should do.
	Comment *string

	// The Sha256 or Sha1 hash created by the system when the document was created.
	// Sha1 hashes have been deprecated.
	DocumentHash *string

	// Sha256 or Sha1. Sha1 hashes have been deprecated.
	DocumentHashType types.DocumentHashType

	// The SSM document version to use in the request. You can specify $DEFAULT,
	// $LATEST, or a specific version number. If you run commands by using the AWS CLI,
	// then you must escape the first two options by using a backslash. If you specify
	// a version number, then you don't need to use the backslash. For example:
	// --document-version "\$DEFAULT" --document-version "\$LATEST" --document-version
	// "3"
	DocumentVersion *string

	// The IDs of the instances where the command should run. Specifying instance IDs
	// is most useful when you are targeting a limited number of instances, though you
	// can specify up to 50 IDs. To target a larger number of instances, or if you
	// prefer not to list individual instance IDs, we recommend using the Targets
	// option instead. Using Targets, which accepts tag key-value pairs to identify the
	// instances to send commands to, you can a send command to tens, hundreds, or
	// thousands of instances at once. For more information about how to use targets,
	// see Using targets and rate controls to send commands to a fleet
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/send-commands-multiple.html)
	// in the AWS Systems Manager User Guide.
	InstanceIds []*string

	// (Optional) The maximum number of instances that are allowed to run the command
	// at the same time. You can specify a number such as 10 or a percentage such as
	// 10%. The default value is 50. For more information about how to use
	// MaxConcurrency, see Using concurrency controls
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/send-commands-multiple.html#send-commands-velocity)
	// in the AWS Systems Manager User Guide.
	MaxConcurrency *string

	// The maximum number of errors allowed without the command failing. When the
	// command fails one more time beyond the value of MaxErrors, the systems stops
	// sending the command to additional targets. You can specify a number like 10 or a
	// percentage like 10%. The default value is 0. For more information about how to
	// use MaxErrors, see Using error controls
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/send-commands-multiple.html#send-commands-maxerrors)
	// in the AWS Systems Manager User Guide.
	MaxErrors *string

	// Configurations for sending notifications.
	NotificationConfig *types.NotificationConfig

	// The name of the S3 bucket where command execution responses should be stored.
	OutputS3BucketName *string

	// The directory structure within the S3 bucket where the responses should be
	// stored.
	OutputS3KeyPrefix *string

	// (Deprecated) You can no longer specify this parameter. The system ignores it.
	// Instead, Systems Manager automatically determines the Region of the S3 bucket.
	OutputS3Region *string

	// The required and optional parameters specified in the document being run.
	Parameters map[string][]*string

	// The ARN of the IAM service role to use to publish Amazon Simple Notification
	// Service (Amazon SNS) notifications for Run Command commands.
	ServiceRoleArn *string

	// An array of search criteria that targets instances using a Key,Value combination
	// that you specify. Specifying targets is most useful when you want to send a
	// command to a large number of instances at once. Using Targets, which accepts tag
	// key-value pairs to identify instances, you can send a command to tens, hundreds,
	// or thousands of instances at once. To send a command to a smaller number of
	// instances, you can use the InstanceIds option instead. For more information
	// about how to use targets, see Sending commands to a fleet
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/send-commands-multiple.html)
	// in the AWS Systems Manager User Guide.
	Targets []*types.Target

	// If this time is reached and the command has not already started running, it will
	// not run.
	TimeoutSeconds *int32
}

type SendCommandOutput struct {

	// The request as it was received by Systems Manager. Also provides the command ID
	// which can be used future references to this request.
	Command *types.Command

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSendCommandMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendCommand{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendCommand{}, middleware.After)
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
	addOpSendCommandValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendCommand(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSendCommand(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "SendCommand",
	}
}
