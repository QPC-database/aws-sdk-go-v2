// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the GCM channel for an application or updates the status and settings of
// the GCM channel for an application.
func (c *Client) UpdateGcmChannel(ctx context.Context, params *UpdateGcmChannelInput, optFns ...func(*Options)) (*UpdateGcmChannelOutput, error) {
	if params == nil {
		params = &UpdateGcmChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGcmChannel", params, optFns, addOperationUpdateGcmChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGcmChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGcmChannelInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// Specifies the status and settings of the GCM channel for an application. This
	// channel enables Amazon Pinpoint to send push notifications through the Firebase
	// Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM), service.
	//
	// This member is required.
	GCMChannelRequest *types.GCMChannelRequest
}

type UpdateGcmChannelOutput struct {

	// Provides information about the status and settings of the GCM channel for an
	// application. The GCM channel enables Amazon Pinpoint to send push notifications
	// through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging
	// (GCM), service.
	//
	// This member is required.
	GCMChannelResponse *types.GCMChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateGcmChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGcmChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGcmChannel{}, middleware.After)
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
	addOpUpdateGcmChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGcmChannel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateGcmChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateGcmChannel",
	}
}
