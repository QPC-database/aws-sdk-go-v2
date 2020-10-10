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

// Returns metadata information for a specific bot. You must provide the bot name
// and the bot version or alias. This operation requires permissions for the
// lex:GetBot action.
func (c *Client) GetBot(ctx context.Context, params *GetBotInput, optFns ...func(*Options)) (*GetBotOutput, error) {
	if params == nil {
		params = &GetBotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBot", params, optFns, addOperationGetBotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotInput struct {

	// The name of the bot. The name is case sensitive.
	//
	// This member is required.
	Name *string

	// The version or alias of the bot.
	//
	// This member is required.
	VersionOrAlias *string
}

type GetBotOutput struct {

	// The message that Amazon Lex returns when the user elects to end the conversation
	// without completing it. For more information, see PutBot ().
	AbortStatement *types.Statement

	// Checksum of the bot used to identify a specific revision of the bot's $LATEST
	// version.
	Checksum *string

	// For each Amazon Lex bot created with the Amazon Lex Model Building Service, you
	// must specify whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to the Children's Online Privacy Protection Act (COPPA)
	// by specifying true or false in the childDirected field. By specifying true in
	// the childDirected field, you confirm that your use of Amazon Lex is related to a
	// website, program, or other application that is directed or targeted, in whole or
	// in part, to children under age 13 and subject to COPPA. By specifying false in
	// the childDirected field, you confirm that your use of Amazon Lex is not related
	// to a website, program, or other application that is directed or targeted, in
	// whole or in part, to children under age 13 and subject to COPPA. You may not
	// specify a default value for the childDirected field that does not accurately
	// reflect whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to COPPA. If your use of Amazon Lex relates to a
	// website, program, or other application that is directed in whole or in part, to
	// children under age 13, you must obtain any required verifiable parental consent
	// under COPPA. For information regarding the use of Amazon Lex in connection with
	// websites, programs, or other applications that are directed or targeted, in
	// whole or in part, to children under age 13, see the Amazon Lex FAQ.
	// (https://aws.amazon.com/lex/faqs#data-security)
	ChildDirected *bool

	// The message Amazon Lex uses when it doesn't understand the user's request. For
	// more information, see PutBot ().
	ClarificationPrompt *types.Prompt

	// The date that the bot was created.
	CreatedDate *time.Time

	// A description of the bot.
	Description *string

	// Indicates whether user utterances should be sent to Amazon Comprehend for
	// sentiment analysis.
	DetectSentiment *bool

	// If status is FAILED, Amazon Lex explains why it failed to build the bot.
	FailureReason *string

	// The maximum time in seconds that Amazon Lex retains the data gathered in a
	// conversation. For more information, see PutBot ().
	IdleSessionTTLInSeconds *int32

	// An array of intent objects. For more information, see PutBot ().
	Intents []*types.Intent

	// The date that the bot was updated. When you create a resource, the creation date
	// and last updated date are the same.
	LastUpdatedDate *time.Time

	// The target locale for the bot.
	Locale types.Locale

	// The name of the bot.
	Name *string

	// The status of the bot. When the status is BUILDING Amazon Lex is building the
	// bot for testing and use. If the status of the bot is READY_BASIC_TESTING, you
	// can test the bot using the exact utterances specified in the bot's intents. When
	// the bot is ready for full testing or to run, the status is READY. If there was a
	// problem with building the bot, the status is FAILED and the failureReason field
	// explains why the bot did not build. If the bot was saved but not built, the
	// status is NOT_BUILT.
	Status types.Status

	// The version of the bot. For a new bot, the version is always $LATEST.
	Version *string

	// The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the
	// user. For more information, see PutBot ().
	VoiceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBot{}, middleware.After)
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
	addOpGetBotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBot",
	}
}
