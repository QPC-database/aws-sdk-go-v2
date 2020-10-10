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

// Retrieves information about the status and settings of the GCM channel for an
// application.
func (c *Client) GetGcmChannel(ctx context.Context, params *GetGcmChannelInput, optFns ...func(*Options)) (*GetGcmChannelOutput, error) {
	if params == nil {
		params = &GetGcmChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGcmChannel", params, optFns, addOperationGetGcmChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGcmChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGcmChannelInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string
}

type GetGcmChannelOutput struct {

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

func addOperationGetGcmChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGcmChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGcmChannel{}, middleware.After)
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
	addOpGetGcmChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetGcmChannel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetGcmChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetGcmChannel",
	}
}
