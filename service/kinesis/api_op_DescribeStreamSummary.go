// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides a summarized description of the specified Kinesis data stream without
// the shard list. The information returned includes the stream name, Amazon
// Resource Name (ARN), status, record retention period, approximate creation time,
// monitoring, encryption details, and open shard count.
func (c *Client) DescribeStreamSummary(ctx context.Context, params *DescribeStreamSummaryInput, optFns ...func(*Options)) (*DescribeStreamSummaryOutput, error) {
	if params == nil {
		params = &DescribeStreamSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStreamSummary", params, optFns, addOperationDescribeStreamSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStreamSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStreamSummaryInput struct {

	// The name of the stream to describe.
	//
	// This member is required.
	StreamName *string
}

type DescribeStreamSummaryOutput struct {

	// A StreamDescriptionSummary () containing information about the stream.
	//
	// This member is required.
	StreamDescriptionSummary *types.StreamDescriptionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeStreamSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStreamSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStreamSummary{}, middleware.After)
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
	addOpDescribeStreamSummaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStreamSummary(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeStreamSummary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DescribeStreamSummary",
	}
}
