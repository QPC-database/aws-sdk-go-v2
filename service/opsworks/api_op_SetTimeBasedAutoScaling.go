// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specify the time-based auto scaling configuration for a specified instance. For
// more information, see Managing Load with Time-based and Load-based Instances
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-autoscaling.html).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) SetTimeBasedAutoScaling(ctx context.Context, params *SetTimeBasedAutoScalingInput, optFns ...func(*Options)) (*SetTimeBasedAutoScalingOutput, error) {
	if params == nil {
		params = &SetTimeBasedAutoScalingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTimeBasedAutoScaling", params, optFns, addOperationSetTimeBasedAutoScalingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTimeBasedAutoScalingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetTimeBasedAutoScalingInput struct {

	// The instance ID.
	//
	// This member is required.
	InstanceId *string

	// An AutoScalingSchedule with the instance schedule.
	AutoScalingSchedule *types.WeeklyAutoScalingSchedule
}

type SetTimeBasedAutoScalingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetTimeBasedAutoScalingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetTimeBasedAutoScaling{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetTimeBasedAutoScaling{}, middleware.After)
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
	addOpSetTimeBasedAutoScalingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetTimeBasedAutoScaling(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetTimeBasedAutoScaling(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "SetTimeBasedAutoScaling",
	}
}
