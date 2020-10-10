// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a receipt rule set by cloning an existing one. All receipt rules and
// configurations are copied to the new receipt rule set and are completely
// independent of the source rule set. For information about setting up rule sets,
// see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rule-set.html).
// You can execute this operation no more than once per second.
func (c *Client) CloneReceiptRuleSet(ctx context.Context, params *CloneReceiptRuleSetInput, optFns ...func(*Options)) (*CloneReceiptRuleSetOutput, error) {
	if params == nil {
		params = &CloneReceiptRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CloneReceiptRuleSet", params, optFns, addOperationCloneReceiptRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CloneReceiptRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create a receipt rule set by cloning an existing one.
// You use receipt rule sets to receive email with Amazon SES. For more
// information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type CloneReceiptRuleSetInput struct {

	// The name of the rule set to clone.
	//
	// This member is required.
	OriginalRuleSetName *string

	// The name of the rule set to create. The name must:
	//
	//     * This value can only
	// contain ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes
	// (-).
	//
	//     * Start and end with a letter or number.
	//
	//     * Contain less than 64
	// characters.
	//
	// This member is required.
	RuleSetName *string
}

// An empty element returned on a successful request.
type CloneReceiptRuleSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCloneReceiptRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCloneReceiptRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCloneReceiptRuleSet{}, middleware.After)
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
	addOpCloneReceiptRuleSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCloneReceiptRuleSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCloneReceiptRuleSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CloneReceiptRuleSet",
	}
}
