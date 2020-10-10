// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a network profile by the network profile ARN.
func (c *Client) UpdateNetworkProfile(ctx context.Context, params *UpdateNetworkProfileInput, optFns ...func(*Options)) (*UpdateNetworkProfileOutput, error) {
	if params == nil {
		params = &UpdateNetworkProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNetworkProfile", params, optFns, addOperationUpdateNetworkProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNetworkProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNetworkProfileInput struct {

	// The ARN of the network profile associated with a device.
	//
	// This member is required.
	NetworkProfileArn *string

	// The ARN of the Private Certificate Authority (PCA) created in AWS Certificate
	// Manager (ACM). This is used to issue certificates to the devices.
	CertificateAuthorityArn *string

	// The current password of the Wi-Fi network.
	CurrentPassword *string

	// Detailed information about a device's network profile.
	Description *string

	// The name of the network profile associated with a device.
	NetworkProfileName *string

	// The next, or subsequent, password of the Wi-Fi network. This password is
	// asynchronously transmitted to the device and is used when the password of the
	// network changes to NextPassword.
	NextPassword *string

	// The root certificate(s) of your authentication server that will be installed on
	// your devices and used to trust your authentication server during EAP
	// negotiation.
	TrustAnchors []*string
}

type UpdateNetworkProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateNetworkProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNetworkProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNetworkProfile{}, middleware.After)
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
	addOpUpdateNetworkProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNetworkProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateNetworkProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "UpdateNetworkProfile",
	}
}
