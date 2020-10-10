// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables detailed monitoring for a running instance. Otherwise, basic monitoring
// is enabled. For more information, see Monitoring your instances and volumes
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch.html) in
// the Amazon Elastic Compute Cloud User Guide. To disable detailed monitoring, see
// .
func (c *Client) MonitorInstances(ctx context.Context, params *MonitorInstancesInput, optFns ...func(*Options)) (*MonitorInstancesOutput, error) {
	if params == nil {
		params = &MonitorInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MonitorInstances", params, optFns, addOperationMonitorInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MonitorInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MonitorInstancesInput struct {

	// The IDs of the instances.
	//
	// This member is required.
	InstanceIds []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type MonitorInstancesOutput struct {

	// The monitoring information.
	InstanceMonitorings []*types.InstanceMonitoring

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationMonitorInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpMonitorInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpMonitorInstances{}, middleware.After)
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
	addOpMonitorInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opMonitorInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opMonitorInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "MonitorInstances",
	}
}
