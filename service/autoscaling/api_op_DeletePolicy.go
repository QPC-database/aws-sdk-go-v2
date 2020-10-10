// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified scaling policy. Deleting either a step scaling policy or a
// simple scaling policy deletes the underlying alarm action, but does not delete
// the alarm, even if it no longer has an associated action. For more information,
// see Deleting a Scaling Policy
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/deleting-scaling-policy.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) {
	if params == nil {
		params = &DeletePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePolicy", params, optFns, addOperationDeletePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePolicyInput struct {

	// The name or Amazon Resource Name (ARN) of the policy.
	//
	// This member is required.
	PolicyName *string

	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
}

type DeletePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeletePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeletePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeletePolicy{}, middleware.After)
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
	addOpDeletePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeletePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DeletePolicy",
	}
}
