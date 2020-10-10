// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the priorities of the specified rules. You can reorder the rules as long as
// there are no priority conflicts in the new order. Any existing rules that you do
// not specify retain their current priority.
func (c *Client) SetRulePriorities(ctx context.Context, params *SetRulePrioritiesInput, optFns ...func(*Options)) (*SetRulePrioritiesOutput, error) {
	if params == nil {
		params = &SetRulePrioritiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetRulePriorities", params, optFns, addOperationSetRulePrioritiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetRulePrioritiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetRulePrioritiesInput struct {

	// The rule priorities.
	//
	// This member is required.
	RulePriorities []*types.RulePriorityPair
}

type SetRulePrioritiesOutput struct {

	// Information about the rules.
	Rules []*types.Rule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetRulePrioritiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetRulePriorities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetRulePriorities{}, middleware.After)
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
	addOpSetRulePrioritiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetRulePriorities(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetRulePriorities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "SetRulePriorities",
	}
}
