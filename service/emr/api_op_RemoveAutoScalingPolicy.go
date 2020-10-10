// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes an automatic scaling policy from a specified instance group within an
// EMR cluster.
func (c *Client) RemoveAutoScalingPolicy(ctx context.Context, params *RemoveAutoScalingPolicyInput, optFns ...func(*Options)) (*RemoveAutoScalingPolicyOutput, error) {
	if params == nil {
		params = &RemoveAutoScalingPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveAutoScalingPolicy", params, optFns, addOperationRemoveAutoScalingPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveAutoScalingPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveAutoScalingPolicyInput struct {

	// Specifies the ID of a cluster. The instance group to which the automatic scaling
	// policy is applied is within this cluster.
	//
	// This member is required.
	ClusterId *string

	// Specifies the ID of the instance group to which the scaling policy is applied.
	//
	// This member is required.
	InstanceGroupId *string
}

type RemoveAutoScalingPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveAutoScalingPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveAutoScalingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveAutoScalingPolicy{}, middleware.After)
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
	addOpRemoveAutoScalingPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveAutoScalingPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemoveAutoScalingPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "RemoveAutoScalingPolicy",
	}
}
