// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the specified Contributor Insights rules. When rules are disabled, they
// do not analyze log groups and do not incur costs.
func (c *Client) DisableInsightRules(ctx context.Context, params *DisableInsightRulesInput, optFns ...func(*Options)) (*DisableInsightRulesOutput, error) {
	if params == nil {
		params = &DisableInsightRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableInsightRules", params, optFns, addOperationDisableInsightRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableInsightRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableInsightRulesInput struct {

	// An array of the rule names to disable. If you need to find out the names of your
	// rules, use DescribeInsightRules
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeInsightRules.html).
	//
	// This member is required.
	RuleNames []*string
}

type DisableInsightRulesOutput struct {

	// An array listing the rules that could not be disabled. You cannot disable
	// built-in rules.
	Failures []*types.PartialFailure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableInsightRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisableInsightRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableInsightRules{}, middleware.After)
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
	addOpDisableInsightRulesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableInsightRules(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisableInsightRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DisableInsightRules",
	}
}
