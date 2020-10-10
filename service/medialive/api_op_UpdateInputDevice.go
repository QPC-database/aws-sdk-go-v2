// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the parameters for the input device.
func (c *Client) UpdateInputDevice(ctx context.Context, params *UpdateInputDeviceInput, optFns ...func(*Options)) (*UpdateInputDeviceOutput, error) {
	if params == nil {
		params = &UpdateInputDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInputDevice", params, optFns, addOperationUpdateInputDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInputDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update an input device.
type UpdateInputDeviceInput struct {

	// The unique ID of the input device. For example, hd-123456789abcdef.
	//
	// This member is required.
	InputDeviceId *string

	// The settings that you want to apply to the input device.
	HdDeviceSettings *types.InputDeviceConfigurableSettings

	// The name that you assigned to this input device (not the unique ID).
	Name *string
}

// Placeholder documentation for UpdateInputDeviceResponse
type UpdateInputDeviceOutput struct {

	// The unique ARN of the input device.
	Arn *string

	// The state of the connection between the input device and AWS.
	ConnectionState types.InputDeviceConnectionState

	// The status of the action to synchronize the device configuration. If you change
	// the configuration of the input device (for example, the maximum bitrate),
	// MediaLive sends the new data to the device. The device might not update itself
	// immediately. SYNCED means the device has updated its configuration. SYNCING
	// means that it has not updated its configuration.
	DeviceSettingsSyncState types.DeviceSettingsSyncState

	// Settings that describe an input device that is type HD.
	HdDeviceSettings *types.InputDeviceHdSettings

	// The unique ID of the input device.
	Id *string

	// The network MAC address of the input device.
	MacAddress *string

	// A name that you specify for the input device.
	Name *string

	// The network settings for the input device.
	NetworkSettings *types.InputDeviceNetworkSettings

	// The unique serial number of the input device.
	SerialNumber *string

	// The type of the input device.
	Type types.InputDeviceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateInputDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateInputDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateInputDevice{}, middleware.After)
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
	addOpUpdateInputDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInputDevice(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateInputDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateInputDevice",
	}
}
