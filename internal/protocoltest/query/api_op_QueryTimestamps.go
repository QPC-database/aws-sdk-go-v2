// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This test serializes timestamps.
//
//     * Timestamps are serialized as RFC 3339
// date-time values by default.
//
//     * A timestampFormat trait on a member changes
// the format.
//
//     * A timestampFormat trait on the shape targeted by the member
// changes the format.
func (c *Client) QueryTimestamps(ctx context.Context, params *QueryTimestampsInput, optFns ...func(*Options)) (*QueryTimestampsOutput, error) {
	if params == nil {
		params = &QueryTimestampsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QueryTimestamps", params, optFns, addOperationQueryTimestampsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryTimestampsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryTimestampsInput struct {
	EpochMember *time.Time

	EpochTarget *time.Time

	NormalFormat *time.Time
}

type QueryTimestampsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationQueryTimestampsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpQueryTimestamps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpQueryTimestamps{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opQueryTimestamps(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opQueryTimestamps(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "QueryTimestamps",
	}
}
