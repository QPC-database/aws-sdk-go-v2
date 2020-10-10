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

// Creates an intent or replaces an existing intent. To define the interaction
// between the user and your bot, you use one or more intents. For a pizza ordering
// bot, for example, you would create an OrderPizza intent. To create an intent or
// replace an existing intent, you must provide the following:
//
//     * Intent name.
// For example, OrderPizza.
//
//     * Sample utterances. For example, "Can I order a
// pizza, please." and "I want to order a pizza."
//
//     * Information to be
// gathered. You specify slot types for the information that your bot will request
// from the user. You can specify standard slot types, such as a date or a time, or
// custom slot types such as the size and crust of a pizza.
//
//     * How the intent
// will be fulfilled. You can provide a Lambda function or configure the intent to
// return the intent information to the client application. If you use a Lambda
// function, when all of the intent information is available, Amazon Lex invokes
// your Lambda function. If you configure your intent to return the intent
// information to the client application.
//
// You can specify other optional
// information in the request, such as:  <ul> <li> <p>A confirmation prompt to ask
// the user to confirm an intent. For example, "Shall I order your pizza?"</p>
// </li> <li> <p>A conclusion statement to send to the user after the intent has
// been fulfilled. For example, "I placed your pizza order."</p> </li> <li> <p>A
// follow-up prompt that asks the user for additional activity. For example, asking
// "Do you want to order a drink with your pizza?"</p> </li> </ul> <p>If you
// specify an existing intent name to update the intent, Amazon Lex replaces the
// values in the <code>$LATEST</code> version of the intent with the values in the
// request. Amazon Lex removes fields that you don't provide in the request. If you
// don't specify the required fields, Amazon Lex throws an exception. When you
// update the <code>$LATEST</code> version of an intent, the <code>status</code>
// field of any bot that uses the <code>$LATEST</code> version of the intent is set
// to <code>NOT_BUILT</code>.</p> <p>For more information, see
// <a>how-it-works</a>.</p> <p>This operation requires permissions for the
// <code>lex:PutIntent</code> action.</p>
func (c *Client) PutIntent(ctx context.Context, params *PutIntentInput, optFns ...func(*Options)) (*PutIntentOutput, error) {
	if params == nil {
		params = &PutIntentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutIntent", params, optFns, addOperationPutIntentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutIntentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutIntentInput struct {

	// The name of the intent. The name is not case sensitive. The name can't match a
	// built-in intent name, or a built-in intent name with "AMAZON." removed. For
	// example, because there is a built-in intent called AMAZON.HelpIntent, you can't
	// create a custom intent called HelpIntent. For a list of built-in intents, see
	// Standard Built-in Intents
	// (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	//
	// This member is required.
	Name *string

	// Identifies a specific revision of the $LATEST version. When you create a new
	// intent, leave the checksum field blank. If you specify a checksum you get a
	// BadRequestException exception. When you want to update a intent, set the
	// checksum field to the checksum of the most recent revision of the $LATEST
	// version. If you don't specify the  checksum field, or if the checksum does not
	// match the $LATEST version, you get a PreconditionFailedException exception.
	Checksum *string

	// The statement that you want Amazon Lex to convey to the user after the intent is
	// successfully fulfilled by the Lambda function. This element is relevant only if
	// you provide a Lambda function in the fulfillmentActivity. If you return the
	// intent to the client application, you can't specify this element. The
	// followUpPrompt and conclusionStatement are mutually exclusive. You can specify
	// only one.
	ConclusionStatement *types.Statement

	// Prompts the user to confirm the intent. This question should have a yes or no
	// answer. Amazon Lex uses this prompt to ensure that the user acknowledges that
	// the intent is ready for fulfillment. For example, with the OrderPizza intent,
	// you might want to confirm that the order is correct before placing it. For other
	// intents, such as intents that simply respond to user questions, you might not
	// need to ask the user for confirmation before providing the information. You you
	// must provide both the rejectionStatement and the confirmationPrompt, or neither.
	ConfirmationPrompt *types.Prompt

	// When set to true a new numbered version of the intent is created. This is the
	// same as calling the CreateIntentVersion operation. If you do not specify
	// createVersion, the default is false.
	CreateVersion *bool

	// A description of the intent.
	Description *string

	// Specifies a Lambda function to invoke for each user input. You can invoke this
	// Lambda function to personalize user interaction. For example, suppose your bot
	// determines that the user is John. Your Lambda function might retrieve John's
	// information from a backend database and prepopulate some of the values. For
	// example, if you find that John is gluten intolerant, you might set the
	// corresponding intent slot, GlutenIntolerant, to true. You might find John's
	// phone number and set the corresponding session attribute.
	DialogCodeHook *types.CodeHook

	// Amazon Lex uses this prompt to solicit additional activity after fulfilling an
	// intent. For example, after the OrderPizza intent is fulfilled, you might prompt
	// the user to order a drink. The action that Amazon Lex takes depends on the
	// user's response, as follows:
	//
	//     * If the user says "Yes" it responds with the
	// clarification prompt that is configured for the bot.
	//
	//     * if the user says
	// "Yes" and continues with an utterance that triggers an intent it starts a
	// conversation for the intent.
	//
	//     * If the user says "No" it responds with the
	// rejection statement configured for the the follow-up prompt.
	//
	//     * If it
	// doesn't recognize the utterance it repeats the follow-up prompt again.
	//
	//
	// <p>The <code>followUpPrompt</code> field and the
	// <code>conclusionStatement</code> field are mutually exclusive. You can specify
	// only one. </p>
	FollowUpPrompt *types.FollowUpPrompt

	// Required. Describes how the intent is fulfilled. For example, after a user
	// provides all of the information for a pizza order, fulfillmentActivity defines
	// how the bot places an order with a local pizza store. You might configure Amazon
	// Lex to return all of the intent information to the client application, or direct
	// it to invoke a Lambda function that can process the intent (for example, place
	// an order with a pizzeria).
	FulfillmentActivity *types.FulfillmentActivity

	// Configuration information required to use the AMAZON.KendraSearchIntent intent
	// to connect to an Amazon Kendra index. For more information, see
	// AMAZON.KendraSearchIntent
	// (http://docs.aws.amazon.com/lex/latest/dg/built-in-intent-kendra-search.html).
	KendraConfiguration *types.KendraConfiguration

	// A unique identifier for the built-in intent to base this intent on. To find the
	// signature for an intent, see Standard Built-in Intents
	// (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	ParentIntentSignature *string

	// When the user answers "no" to the question defined in confirmationPrompt, Amazon
	// Lex responds with this statement to acknowledge that the intent was canceled.
	// You must provide both the rejectionStatement and the confirmationPrompt, or
	// neither.
	RejectionStatement *types.Statement

	// An array of utterances (strings) that a user might say to signal the intent. For
	// example, "I want {PizzaSize} pizza", "Order {Quantity} {PizzaSize} pizzas".
	// <p>In each utterance, a slot name is enclosed in curly braces. </p>
	SampleUtterances []*string

	// An array of intent slots. At runtime, Amazon Lex elicits required slot values
	// from the user using prompts defined in the slots. For more information, see
	// how-it-works ().
	Slots []*types.Slot
}

type PutIntentOutput struct {

	// Checksum of the $LATESTversion of the intent created or updated.
	Checksum *string

	// After the Lambda function specified in thefulfillmentActivityintent fulfills the
	// intent, Amazon Lex conveys this statement to the user.
	ConclusionStatement *types.Statement

	// If defined in the intent, Amazon Lex prompts the user to confirm the intent
	// before fulfilling it.
	ConfirmationPrompt *types.Prompt

	// True if a new version of the intent was created. If the createVersion field was
	// not specified in the request, the createVersion field is set to false in the
	// response.
	CreateVersion *bool

	// The date that the intent was created.
	CreatedDate *time.Time

	// A description of the intent.
	Description *string

	// If defined in the intent, Amazon Lex invokes this Lambda function for each user
	// input.
	DialogCodeHook *types.CodeHook

	// If defined in the intent, Amazon Lex uses this prompt to solicit additional user
	// activity after the intent is fulfilled.
	FollowUpPrompt *types.FollowUpPrompt

	// If defined in the intent, Amazon Lex invokes this Lambda function to fulfill the
	// intent after the user provides all of the information required by the intent.
	FulfillmentActivity *types.FulfillmentActivity

	// Configuration information, if any, required to connect to an Amazon Kendra index
	// and use the AMAZON.KendraSearchIntent intent.
	KendraConfiguration *types.KendraConfiguration

	// The date that the intent was updated. When you create a resource, the creation
	// date and last update dates are the same.
	LastUpdatedDate *time.Time

	// The name of the intent.
	Name *string

	// A unique identifier for the built-in intent that this intent is based on.
	ParentIntentSignature *string

	// If the user answers "no" to the question defined in confirmationPrompt Amazon
	// Lex responds with this statement to acknowledge that the intent was canceled.
	RejectionStatement *types.Statement

	// An array of sample utterances that are configured for the intent.
	SampleUtterances []*string

	// An array of intent slots that are configured for the intent.
	Slots []*types.Slot

	// The version of the intent. For a new intent, the version is always $LATEST.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutIntentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutIntent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutIntent{}, middleware.After)
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
	addOpPutIntentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutIntent(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutIntent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "PutIntent",
	}
}
