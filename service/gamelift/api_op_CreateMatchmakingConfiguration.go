// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Defines a new matchmaking configuration for use with FlexMatch. A matchmaking
// configuration sets out guidelines for matching players and getting the matches
// into games. You can set up multiple matchmaking configurations to handle the
// scenarios needed for your game. Each matchmaking ticket (StartMatchmaking () or
// StartMatchBackfill ()) specifies a configuration for the match and provides
// player attributes to support the configuration being used. To create a
// matchmaking configuration, at a minimum you must specify the following:
// configuration name; a rule set that governs how to evaluate players and find
// acceptable matches; a game session queue to use when placing a new game session
// for the match; and the maximum time allowed for a matchmaking attempt. There are
// two ways to track the progress of matchmaking tickets: (1) polling ticket status
// with DescribeMatchmaking (); or (2) receiving notifications with Amazon Simple
// Notification Service (SNS). To use notifications, you first need to set up an
// SNS topic to receive the notifications, and provide the topic ARN in the
// matchmaking configuration. Since notifications promise only "best effort"
// delivery, we recommend calling DescribeMatchmaking if no notifications are
// received within 30 seconds. Learn more  Design a FlexMatch Matchmaker
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-configuration.html)
// Setting up Notifications for Matchmaking
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-notification.html)
// Related operations
//
//     * CreateMatchmakingConfiguration ()
//
//     *
// DescribeMatchmakingConfigurations ()
//
//     * UpdateMatchmakingConfiguration ()
//
//
// * DeleteMatchmakingConfiguration ()
//
//     * CreateMatchmakingRuleSet ()
//
//     *
// DescribeMatchmakingRuleSets ()
//
//     * ValidateMatchmakingRuleSet ()
//
//     *
// DeleteMatchmakingRuleSet ()
func (c *Client) CreateMatchmakingConfiguration(ctx context.Context, params *CreateMatchmakingConfigurationInput, optFns ...func(*Options)) (*CreateMatchmakingConfigurationOutput, error) {
	if params == nil {
		params = &CreateMatchmakingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMatchmakingConfiguration", params, optFns, addOperationCreateMatchmakingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMatchmakingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type CreateMatchmakingConfigurationInput struct {

	// A flag that determines whether a match that was created with this configuration
	// must be accepted by the matched players. To require acceptance, set to TRUE.
	//
	// This member is required.
	AcceptanceRequired *bool

	// Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html))
	// that is assigned to a GameLift game session queue resource and uniquely
	// identifies it. ARNs are unique across all Regions. These queues are used when
	// placing game sessions for matches that are created with this matchmaking
	// configuration. Queues can be located in any Region.
	//
	// This member is required.
	GameSessionQueueArns []*string

	// A unique identifier for a matchmaking configuration. This name is used to
	// identify the configuration associated with a matchmaking request or ticket.
	//
	// This member is required.
	Name *string

	// The maximum duration, in seconds, that a matchmaking ticket can remain in
	// process before timing out. Requests that fail due to timing out can be
	// resubmitted as needed.
	//
	// This member is required.
	RequestTimeoutSeconds *int32

	// A unique identifier for a matchmaking rule set to use with this configuration.
	// You can use either the rule set name or ARN value. A matchmaking configuration
	// can only use rule sets that are defined in the same Region.
	//
	// This member is required.
	RuleSetName *string

	// The length of time (in seconds) to wait for players to accept a proposed match.
	// If any player rejects the match or fails to accept before the timeout, the
	// ticket continues to look for an acceptable match.
	AcceptanceTimeoutSeconds *int32

	// The number of player slots in a match to keep open for future players. For
	// example, assume that the configuration's rule set specifies a match for a single
	// 12-person team. If the additional player count is set to 2, only 10 players are
	// initially selected for the match.
	AdditionalPlayerCount *int32

	// The method used to backfill game sessions that are created with this matchmaking
	// configuration. Specify MANUAL when your game manages backfill requests manually
	// or does not use the match backfill feature. Specify AUTOMATIC to have GameLift
	// create a StartMatchBackfill () request whenever a game session has one or more
	// open slots. Learn more about manual and automatic backfill in  Backfill Existing
	// Games with FlexMatch
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-backfill.html).
	BackfillMode types.BackfillMode

	// Information to be added to all events related to this matchmaking configuration.
	CustomEventData *string

	// A human-readable description of the matchmaking configuration.
	Description *string

	// A set of custom properties for a game session, formatted as key-value pairs.
	// These properties are passed to a game server process in the GameSession ()
	// object with a request to start a new game session (see Start a Game Session
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	// This information is added to the new GameSession () object that is created for a
	// successful match.
	GameProperties []*types.GameProperty

	// A set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process in the GameSession () object with a
	// request to start a new game session (see Start a Game Session
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	// This information is added to the new GameSession () object that is created for a
	// successful match.
	GameSessionData *string

	// An SNS topic ARN that is set up to receive matchmaking notifications.
	NotificationTarget *string

	// A list of labels to assign to the new matchmaking configuration resource. Tags
	// are developer-defined key-value pairs. Tagging AWS resources are useful for
	// resource management, access management and cost allocation. For more
	// information, see  Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
	// General Reference. Once the resource is created, you can use TagResource (),
	// UntagResource (), and ListTagsForResource () to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference for
	// actual tagging limits.
	Tags []*types.Tag
}

// Represents the returned data in response to a request action.
type CreateMatchmakingConfigurationOutput struct {

	// Object that describes the newly created matchmaking configuration.
	Configuration *types.MatchmakingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMatchmakingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMatchmakingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMatchmakingConfiguration{}, middleware.After)
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
	addOpCreateMatchmakingConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMatchmakingConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateMatchmakingConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateMatchmakingConfiguration",
	}
}
