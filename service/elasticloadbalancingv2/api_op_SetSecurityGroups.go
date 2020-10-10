// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates the specified security groups with the specified Application Load
// Balancer. The specified security groups override the previously associated
// security groups. You can't specify a security group for a Network Load Balancer.
func (c *Client) SetSecurityGroups(ctx context.Context, params *SetSecurityGroupsInput, optFns ...func(*Options)) (*SetSecurityGroupsOutput, error) {
	if params == nil {
		params = &SetSecurityGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetSecurityGroups", params, optFns, addOperationSetSecurityGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetSecurityGroupsInput struct {

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// This member is required.
	LoadBalancerArn *string

	// The IDs of the security groups.
	//
	// This member is required.
	SecurityGroups []*string
}

type SetSecurityGroupsOutput struct {

	// The IDs of the security groups associated with the load balancer.
	SecurityGroupIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetSecurityGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetSecurityGroups{}, middleware.After)
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
	addOpSetSecurityGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetSecurityGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetSecurityGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "SetSecurityGroups",
	}
}
