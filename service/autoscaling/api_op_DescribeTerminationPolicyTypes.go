// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the termination policies supported by Amazon EC2 Auto Scaling. For
// more information, see Controlling Which Auto Scaling Instances Terminate During
// Scale In
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DescribeTerminationPolicyTypes(ctx context.Context, params *DescribeTerminationPolicyTypesInput, optFns ...func(*Options)) (*DescribeTerminationPolicyTypesOutput, error) {
	if params == nil {
		params = &DescribeTerminationPolicyTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTerminationPolicyTypes", params, optFns, addOperationDescribeTerminationPolicyTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTerminationPolicyTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTerminationPolicyTypesInput struct {
}

type DescribeTerminationPolicyTypesOutput struct {

	// The termination policies supported by Amazon EC2 Auto Scaling: OldestInstance,
	// OldestLaunchConfiguration, NewestInstance, ClosestToNextInstanceHour, Default,
	// OldestLaunchTemplate, and AllocationStrategy.
	TerminationPolicyTypes []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTerminationPolicyTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTerminationPolicyTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTerminationPolicyTypes{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTerminationPolicyTypes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTerminationPolicyTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeTerminationPolicyTypes",
	}
}
