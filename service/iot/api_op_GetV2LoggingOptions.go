// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the fine grained logging options.
func (c *Client) GetV2LoggingOptions(ctx context.Context, params *GetV2LoggingOptionsInput, optFns ...func(*Options)) (*GetV2LoggingOptionsOutput, error) {
	if params == nil {
		params = &GetV2LoggingOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetV2LoggingOptions", params, optFns, addOperationGetV2LoggingOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetV2LoggingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetV2LoggingOptionsInput struct {
}

type GetV2LoggingOptionsOutput struct {

	// The default log level.
	DefaultLogLevel types.LogLevel

	// Disables all logs.
	DisableAllLogs *bool

	// The IAM role ARN AWS IoT uses to write to your CloudWatch logs.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetV2LoggingOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetV2LoggingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetV2LoggingOptions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetV2LoggingOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetV2LoggingOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetV2LoggingOptions",
	}
}
