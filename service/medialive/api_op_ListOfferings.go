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

// List offerings available for purchase.
func (c *Client) ListOfferings(ctx context.Context, params *ListOfferingsInput, optFns ...func(*Options)) (*ListOfferingsOutput, error) {
	if params == nil {
		params = &ListOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOfferings", params, optFns, addOperationListOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListOfferingsRequest
type ListOfferingsInput struct {

	// Filter by channel class, 'STANDARD' or 'SINGLE_PIPELINE'
	ChannelClass *string

	// Filter to offerings that match the configuration of an existing channel, e.g.
	// '2345678' (a channel ID)
	ChannelConfiguration *string

	// Filter by codec, 'AVC', 'HEVC', 'MPEG2', or 'AUDIO'
	Codec *string

	// Filter by offering duration, e.g. '12'
	Duration *string

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// Filter by bitrate, 'MAX_10_MBPS', 'MAX_20_MBPS', or 'MAX_50_MBPS'
	MaximumBitrate *string

	// Filter by framerate, 'MAX_30_FPS' or 'MAX_60_FPS'
	MaximumFramerate *string

	// Placeholder documentation for __string
	NextToken *string

	// Filter by resolution, 'SD', 'HD', 'FHD', or 'UHD'
	Resolution *string

	// Filter by resource type, 'INPUT', 'OUTPUT', 'MULTIPLEX', or 'CHANNEL'
	ResourceType *string

	// Filter by special feature, 'ADVANCED_AUDIO' or 'AUDIO_NORMALIZATION'
	SpecialFeature *string

	// Filter by video quality, 'STANDARD', 'ENHANCED', or 'PREMIUM'
	VideoQuality *string
}

// Placeholder documentation for ListOfferingsResponse
type ListOfferingsOutput struct {

	// Token to retrieve the next page of results
	NextToken *string

	// List of offerings
	Offerings []*types.Offering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOfferings{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOfferings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListOfferings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "ListOfferings",
	}
}
