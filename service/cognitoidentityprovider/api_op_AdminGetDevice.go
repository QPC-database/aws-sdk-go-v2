// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the device, as an administrator. Calling this action requires developer
// credentials.
func (c *Client) AdminGetDevice(ctx context.Context, params *AdminGetDeviceInput, optFns ...func(*Options)) (*AdminGetDeviceOutput, error) {
	if params == nil {
		params = &AdminGetDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminGetDevice", params, optFns, addOperationAdminGetDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminGetDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get the device, as an administrator.
type AdminGetDeviceInput struct {

	// The device key.
	//
	// This member is required.
	DeviceKey *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user name.
	//
	// This member is required.
	Username *string
}

// Gets the device response, as an administrator.
type AdminGetDeviceOutput struct {

	// The device.
	//
	// This member is required.
	Device *types.DeviceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAdminGetDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminGetDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminGetDevice{}, middleware.After)
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
	addOpAdminGetDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminGetDevice(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAdminGetDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminGetDevice",
	}
}
