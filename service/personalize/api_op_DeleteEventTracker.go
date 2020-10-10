// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the event tracker. Does not delete the event-interactions dataset from
// the associated dataset group. For more information on event trackers, see
// CreateEventTracker ().
func (c *Client) DeleteEventTracker(ctx context.Context, params *DeleteEventTrackerInput, optFns ...func(*Options)) (*DeleteEventTrackerOutput, error) {
	if params == nil {
		params = &DeleteEventTrackerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEventTracker", params, optFns, addOperationDeleteEventTrackerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEventTrackerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEventTrackerInput struct {

	// The Amazon Resource Name (ARN) of the event tracker to delete.
	//
	// This member is required.
	EventTrackerArn *string
}

type DeleteEventTrackerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteEventTrackerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEventTracker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEventTracker{}, middleware.After)
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
	addOpDeleteEventTrackerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEventTracker(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteEventTracker(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DeleteEventTracker",
	}
}
