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

// Starts deletion of channel. The associated outputs are also deleted.
func (c *Client) DeleteChannel(ctx context.Context, params *DeleteChannelInput, optFns ...func(*Options)) (*DeleteChannelOutput, error) {
	if params == nil {
		params = &DeleteChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteChannel", params, optFns, addOperationDeleteChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DeleteChannelRequest
type DeleteChannelInput struct {

	// Unique ID of the channel.
	//
	// This member is required.
	ChannelId *string
}

// Placeholder documentation for DeleteChannelResponse
type DeleteChannelOutput struct {

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

func addOperationDeleteChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteChannel{}, middleware.After)
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
	addOpDeleteChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteChannel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DeleteChannel",
	}
}
