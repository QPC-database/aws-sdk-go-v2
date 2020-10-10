// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified rule. DescribeRule does not list the targets of a rule.
// To see the targets associated with a rule, use ListTargetsByRule ().
func (c *Client) DescribeRule(ctx context.Context, params *DescribeRuleInput, optFns ...func(*Options)) (*DescribeRuleOutput, error) {
	if params == nil {
		params = &DescribeRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRule", params, optFns, addOperationDescribeRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRuleInput struct {

	// The name of the rule.
	//
	// This member is required.
	Name *string

	// The event bus associated with the rule. If you omit this, the default event bus
	// is used.
	EventBusName *string
}

type DescribeRuleOutput struct {

	// The Amazon Resource Name (ARN) of the rule.
	Arn *string

	// The description of the rule.
	Description *string

	// The event bus associated with the rule.
	EventBusName *string

	// The event pattern. For more information, see Events and Event Patterns
	// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)
	// in the Amazon EventBridge User Guide.
	EventPattern *string

	// If this is a managed rule, created by an AWS service on your behalf, this field
	// displays the principal name of the AWS service that created the rule.
	ManagedBy *string

	// The name of the rule.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role associated with the rule.
	RoleArn *string

	// The scheduling expression. For example, "cron(0 20 * * ? *)", "rate(5 minutes)".
	ScheduleExpression *string

	// Specifies whether the rule is enabled or disabled.
	State types.RuleState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRule{}, middleware.After)
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
	addOpDescribeRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DescribeRule",
	}
}
