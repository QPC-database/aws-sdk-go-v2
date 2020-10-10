// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about an intent. In addition to the intent name, you must
// specify the intent version. This operation requires permissions to perform the
// lex:GetIntent action.
func (c *Client) GetIntent(ctx context.Context, params *GetIntentInput, optFns ...func(*Options)) (*GetIntentOutput, error) {
	if params == nil {
		params = &GetIntentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIntent", params, optFns, addOperationGetIntentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIntentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIntentInput struct {

	// The name of the intent. The name is case sensitive.
	//
	// This member is required.
	Name *string

	// The version of the intent.
	//
	// This member is required.
	Version *string
}

type GetIntentOutput struct {

	// Checksum of the intent.
	Checksum *string

	// After the Lambda function specified in the fulfillmentActivity element fulfills
	// the intent, Amazon Lex conveys this statement to the user.
	ConclusionStatement *types.Statement

	// If defined in the bot, Amazon Lex uses prompt to confirm the intent before
	// fulfilling the user's request. For more information, see PutIntent ().
	ConfirmationPrompt *types.Prompt

	// The date that the intent was created.
	CreatedDate *time.Time

	// A description of the intent.
	Description *string

	// If defined in the bot, Amazon Amazon Lex invokes this Lambda function for each
	// user input. For more information, see PutIntent ().
	DialogCodeHook *types.CodeHook

	// If defined in the bot, Amazon Lex uses this prompt to solicit additional user
	// activity after the intent is fulfilled. For more information, see PutIntent ().
	FollowUpPrompt *types.FollowUpPrompt

	// Describes how the intent is fulfilled. For more information, see PutIntent ().
	FulfillmentActivity *types.FulfillmentActivity

	// Configuration information, if any, to connect to an Amazon Kendra index with the
	// AMAZON.KendraSearchIntent intent.
	KendraConfiguration *types.KendraConfiguration

	// The date that the intent was updated. When you create a resource, the creation
	// date and the last updated date are the same.
	LastUpdatedDate *time.Time

	// The name of the intent.
	Name *string

	// A unique identifier for a built-in intent.
	ParentIntentSignature *string

	// If the user answers "no" to the question defined in confirmationPrompt, Amazon
	// Lex responds with this statement to acknowledge that the intent was canceled.
	RejectionStatement *types.Statement

	// An array of sample utterances configured for the intent.
	SampleUtterances []*string

	// An array of intent slots configured for the intent.
	Slots []*types.Slot

	// The version of the intent.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetIntentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIntent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntent{}, middleware.After)
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
	addOpGetIntentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntent(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetIntent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetIntent",
	}
}
