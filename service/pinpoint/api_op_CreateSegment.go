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

// Creates a new segment for an application or updates the configuration,
// dimension, and other settings for an existing segment that's associated with an
// application.
func (c *Client) CreateSegment(ctx context.Context, params *CreateSegmentInput, optFns ...func(*Options)) (*CreateSegmentOutput, error) {
	if params == nil {
		params = &CreateSegmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSegment", params, optFns, addOperationCreateSegmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSegmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSegmentInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// Specifies the configuration, dimension, and other settings for a segment. A
	// WriteSegmentRequest object can include a Dimensions object or a SegmentGroups
	// object, but not both.
	//
	// This member is required.
	WriteSegmentRequest *types.WriteSegmentRequest
}

type CreateSegmentOutput struct {

	// Provides information about the configuration, dimension, and other settings for
	// a segment.
	//
	// This member is required.
	SegmentResponse *types.SegmentResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSegmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSegment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSegment{}, middleware.After)
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
	addOpCreateSegmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSegment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateSegment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "CreateSegment",
	}
}
