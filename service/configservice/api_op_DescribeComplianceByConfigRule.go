// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Indicates whether the specified AWS Config rules are compliant. If a rule is
// noncompliant, this action returns the number of AWS resources that do not comply
// with the rule. A rule is compliant if all of the evaluated resources comply with
// it. It is noncompliant if any of these resources do not comply. If AWS Config
// has no current evaluation results for the rule, it returns INSUFFICIENT_DATA.
// This result might indicate one of the following conditions:
//
//     * AWS Config
// has never invoked an evaluation for the rule. To check whether it has, use the
// DescribeConfigRuleEvaluationStatus action to get the
// LastSuccessfulInvocationTime and LastFailedInvocationTime.
//
//     * The rule's AWS
// Lambda function is failing to send evaluation results to AWS Config. Verify that
// the role you assigned to your configuration recorder includes the
// config:PutEvaluations permission. If the rule is a custom rule, verify that the
// AWS Lambda execution role includes the config:PutEvaluations permission.
//
//     *
// The rule's AWS Lambda function has returned NOT_APPLICABLE for all evaluation
// results. This can occur if the resources were deleted or removed from the rule's
// scope.
func (c *Client) DescribeComplianceByConfigRule(ctx context.Context, params *DescribeComplianceByConfigRuleInput, optFns ...func(*Options)) (*DescribeComplianceByConfigRuleOutput, error) {
	if params == nil {
		params = &DescribeComplianceByConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeComplianceByConfigRule", params, optFns, addOperationDescribeComplianceByConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeComplianceByConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeComplianceByConfigRuleInput struct {

	// Filters the results by compliance. The allowed values are COMPLIANT and
	// NON_COMPLIANT.
	ComplianceTypes []types.ComplianceType

	// Specify one or more AWS Config rule names to filter the results by rule.
	ConfigRuleNames []*string

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

//
type DescribeComplianceByConfigRuleOutput struct {

	// Indicates whether each of the specified AWS Config rules is compliant.
	ComplianceByConfigRules []*types.ComplianceByConfigRule

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeComplianceByConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeComplianceByConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeComplianceByConfigRule{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeComplianceByConfigRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeComplianceByConfigRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeComplianceByConfigRule",
	}
}
