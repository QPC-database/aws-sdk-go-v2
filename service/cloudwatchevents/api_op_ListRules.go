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

// Lists your Amazon EventBridge rules. You can either list all the rules or you
// can provide a prefix to match to the rule names.  <p>ListRules does not list the
// targets of a rule. To see the targets associated with a rule, use
// <a>ListTargetsByRule</a>.</p>
func (c *Client) ListRules(ctx context.Context, params *ListRulesInput, optFns ...func(*Options)) (*ListRulesOutput, error) {
	if params == nil {
		params = &ListRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRules", params, optFns, addOperationListRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRulesInput struct {

	// Limits the results to show only the rules associated with the specified event
	// bus.
	EventBusName *string

	// The maximum number of results to return.
	Limit *int32

	// The prefix matching the rule name.
	NamePrefix *string

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string
}

type ListRulesOutput struct {

	// Indicates whether there are additional results to retrieve. If there are no more
	// results, the value is null.
	NextToken *string

	// The rules that match the specified criteria.
	Rules []*types.Rule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRules{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRules(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListRules",
	}
}
