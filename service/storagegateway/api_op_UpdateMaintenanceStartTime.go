// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a gateway's weekly maintenance start time information, including day and
// time of the week. The maintenance time is the time in your gateway's time zone.
func (c *Client) UpdateMaintenanceStartTime(ctx context.Context, params *UpdateMaintenanceStartTimeInput, optFns ...func(*Options)) (*UpdateMaintenanceStartTimeOutput, error) {
	if params == nil {
		params = &UpdateMaintenanceStartTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMaintenanceStartTime", params, optFns, addOperationUpdateMaintenanceStartTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMaintenanceStartTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the following fields:  <ul> <li> <p>
// <a>UpdateMaintenanceStartTimeInput$DayOfMonth</a> </p> </li> <li> <p>
// <a>UpdateMaintenanceStartTimeInput$DayOfWeek</a> </p> </li> <li> <p>
// <a>UpdateMaintenanceStartTimeInput$HourOfDay</a> </p> </li> <li> <p>
// <a>UpdateMaintenanceStartTimeInput$MinuteOfHour</a> </p> </li> </ul>
type UpdateMaintenanceStartTimeInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string

	// The hour component of the maintenance start time represented as hh, where hh is
	// the hour (00 to 23). The hour of the day is in the time zone of the gateway.
	//
	// This member is required.
	HourOfDay *int32

	// The minute component of the maintenance start time represented as mm, where mm
	// is the minute (00 to 59). The minute of the hour is in the time zone of the
	// gateway.
	//
	// This member is required.
	MinuteOfHour *int32

	// The day of the month component of the maintenance start time represented as an
	// ordinal number from 1 to 28, where 1 represents the first day of the month and
	// 28 represents the last day of the month.
	DayOfMonth *int32

	// The day of the week component of the maintenance start time week represented as
	// an ordinal number from 0 to 6, where 0 represents Sunday and 6 Saturday.
	DayOfWeek *int32
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway whose
// maintenance start time is updated.
type UpdateMaintenanceStartTimeOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateMaintenanceStartTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMaintenanceStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMaintenanceStartTime{}, middleware.After)
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
	addOpUpdateMaintenanceStartTimeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMaintenanceStartTime(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateMaintenanceStartTime(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateMaintenanceStartTime",
	}
}
