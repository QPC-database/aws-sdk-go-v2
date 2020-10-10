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

// Describes Amazon RDS instances. Required Permissions: To use this action, an IAM
// user must have a Show, Deploy, or Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information about
// user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
// This call accepts only one resource-identifying parameter.
func (c *Client) DescribeRdsDbInstances(ctx context.Context, params *DescribeRdsDbInstancesInput, optFns ...func(*Options)) (*DescribeRdsDbInstancesOutput, error) {
	if params == nil {
		params = &DescribeRdsDbInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRdsDbInstances", params, optFns, addOperationDescribeRdsDbInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRdsDbInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRdsDbInstancesInput struct {

	// The ID of the stack with which the instances are registered. The operation
	// returns descriptions of all registered Amazon RDS instances.
	//
	// This member is required.
	StackId *string

	// An array containing the ARNs of the instances to be described.
	RdsDbInstanceArns []*string
}

// Contains the response to a DescribeRdsDbInstances request.
type DescribeRdsDbInstancesOutput struct {

	// An a array of RdsDbInstance objects that describe the instances.
	RdsDbInstances []*types.RdsDbInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRdsDbInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRdsDbInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRdsDbInstances{}, middleware.After)
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
	addOpDescribeRdsDbInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRdsDbInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeRdsDbInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeRdsDbInstances",
	}
}
