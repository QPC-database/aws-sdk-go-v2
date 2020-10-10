// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the policies for the specified Auto Scaling group.
func (c *Client) DescribePolicies(ctx context.Context, params *DescribePoliciesInput, optFns ...func(*Options)) (*DescribePoliciesOutput, error) {
	if params == nil {
		params = &DescribePoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePolicies", params, optFns, addOperationDescribePoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePoliciesInput struct {

	// The name of the Auto Scaling group.
	AutoScalingGroupName *string

	// The maximum number of items to be returned with each call. The default value is
	// 50 and the maximum value is 100.
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The names of one or more policies. If you omit this parameter, all policies are
	// described. If a group name is provided, the results are limited to that group.
	// This list is limited to 50 items. If you specify an unknown policy name, it is
	// ignored with no error.
	PolicyNames []*string

	// One or more policy types. The valid values are SimpleScaling, StepScaling, and
	// TargetTrackingScaling.
	PolicyTypes []*string
}

type DescribePoliciesOutput struct {

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// The scaling policies.
	ScalingPolicies []*types.ScalingPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribePolicies{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePolicies(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribePolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribePolicies",
	}
}
