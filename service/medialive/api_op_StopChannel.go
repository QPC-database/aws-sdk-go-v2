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

// Stops a running channel
func (c *Client) StopChannel(ctx context.Context, params *StopChannelInput, optFns ...func(*Options)) (*StopChannelOutput, error) {
	if params == nil {
		params = &StopChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopChannel", params, optFns, addOperationStopChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for StopChannelRequest
type StopChannelInput struct {

	// A request to stop a running channel
	//
	// This member is required.
	ChannelId *string
}

// Placeholder documentation for StopChannelResponse
type StopChannelOutput struct {

	// The unique arn of the channel.
	Arn *string

	// The class for this channel. STANDARD for a channel with two pipelines or
	// SINGLE_PIPELINE for a channel with one pipeline.
	ChannelClass types.ChannelClass

	// A list of destinations of the channel. For UDP outputs, there is one destination
	// per output. For other types (HLS, for example), there is one destination per
	// packager.
	Destinations []*types.OutputDestination

	// The endpoints where outgoing connections initiate from
	EgressEndpoints []*types.ChannelEgressEndpoint

	// Encoder Settings
	EncoderSettings *types.EncoderSettings

	// The unique id of the channel.
	Id *string

	// List of input attachments for channel.
	InputAttachments []*types.InputAttachment

	// Placeholder documentation for InputSpecification
	InputSpecification *types.InputSpecification

	// The log level being written to CloudWatch Logs.
	LogLevel types.LogLevel

	// The name of the channel. (user-mutable)
	Name *string

	// Runtime details for the pipelines of a running channel.
	PipelineDetails []*types.PipelineDetail

	// The number of currently healthy pipelines.
	PipelinesRunningCount *int32

	// The Amazon Resource Name (ARN) of the role assumed when running the Channel.
	RoleArn *string

	// Placeholder documentation for ChannelState
	State types.ChannelState

	// A collection of key-value pairs.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopChannel{}, middleware.After)
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
	addOpStopChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopChannel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "StopChannel",
	}
}
