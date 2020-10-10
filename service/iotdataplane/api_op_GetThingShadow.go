// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the shadow for the specified thing. For more information, see
// GetThingShadow
// (http://docs.aws.amazon.com/iot/latest/developerguide/API_GetThingShadow.html)
// in the AWS IoT Developer Guide.
func (c *Client) GetThingShadow(ctx context.Context, params *GetThingShadowInput, optFns ...func(*Options)) (*GetThingShadowOutput, error) {
	if params == nil {
		params = &GetThingShadowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetThingShadow", params, optFns, addOperationGetThingShadowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetThingShadowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the GetThingShadow operation.
type GetThingShadowInput struct {

	// The name of the thing.
	//
	// This member is required.
	ThingName *string

	// The name of the shadow.
	ShadowName *string
}

// The output from the GetThingShadow operation.
type GetThingShadowOutput struct {

	// The state information, in JSON format.
	Payload []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetThingShadowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetThingShadow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetThingShadow{}, middleware.After)
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
	addOpGetThingShadowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetThingShadow(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetThingShadow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdata",
		OperationName: "GetThingShadow",
	}
}
