// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing room profile by room profile ARN.
func (c *Client) UpdateProfile(ctx context.Context, params *UpdateProfileInput, optFns ...func(*Options)) (*UpdateProfileOutput, error) {
	if params == nil {
		params = &UpdateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProfile", params, optFns, addOperationUpdateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProfileInput struct {

	// The updated address for the room profile.
	Address *string

	// The updated distance unit for the room profile.
	DistanceUnit types.DistanceUnit

	// Sets the profile as default if selected. If this is missing, no update is done
	// to the default status.
	IsDefault *bool

	// The updated locale for the room profile. (This is currently only available to a
	// limited preview audience.)
	Locale *string

	// The updated maximum volume limit for the room profile.
	MaxVolumeLimit *int32

	// The updated meeting room settings of a room profile.
	MeetingRoomConfiguration *types.UpdateMeetingRoomConfiguration

	// Whether the PSTN setting of the room profile is enabled.
	PSTNEnabled *bool

	// The ARN of the room profile to update. Required.
	ProfileArn *string

	// The updated name for the room profile.
	ProfileName *string

	// Whether the setup mode of the profile is enabled.
	SetupModeDisabled *bool

	// The updated temperature unit for the room profile.
	TemperatureUnit types.TemperatureUnit

	// The updated timezone for the room profile.
	Timezone *string

	// The updated wake word for the room profile.
	WakeWord types.WakeWord
}

type UpdateProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProfile{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "UpdateProfile",
	}
}
