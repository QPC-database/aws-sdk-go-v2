// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Reboots a broker. Note: This API is asynchronous.
func (c *Client) RebootBroker(ctx context.Context, params *RebootBrokerInput, optFns ...func(*Options)) (*RebootBrokerOutput, error) {
	if params == nil {
		params = &RebootBrokerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebootBroker", params, optFns, addOperationRebootBrokerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebootBrokerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebootBrokerInput struct {

	// The unique ID that Amazon MQ generates for the broker.
	//
	// This member is required.
	BrokerId *string
}

type RebootBrokerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRebootBrokerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRebootBroker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRebootBroker{}, middleware.After)
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
	addOpRebootBrokerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRebootBroker(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRebootBroker(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "RebootBroker",
	}
}
