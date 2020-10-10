// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get information about a dedicated IP address, including the name of the
// dedicated IP pool that it's associated with, as well information about the
// automatic warm-up process for the address.
func (c *Client) GetDedicatedIp(ctx context.Context, params *GetDedicatedIpInput, optFns ...func(*Options)) (*GetDedicatedIpOutput, error) {
	if params == nil {
		params = &GetDedicatedIpInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDedicatedIp", params, optFns, addOperationGetDedicatedIpMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDedicatedIpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain more information about a dedicated IP address.
type GetDedicatedIpInput struct {

	// The IP address that you want to obtain more information about. The value you
	// specify has to be a dedicated IP address that's assocaited with your Amazon
	// Pinpoint account.
	//
	// This member is required.
	Ip *string
}

// Information about a dedicated IP address.
type GetDedicatedIpOutput struct {

	// An object that contains information about a dedicated IP address.
	DedicatedIp *types.DedicatedIp

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDedicatedIpMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDedicatedIp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDedicatedIp{}, middleware.After)
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
	addOpGetDedicatedIpValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDedicatedIp(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDedicatedIp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetDedicatedIp",
	}
}
