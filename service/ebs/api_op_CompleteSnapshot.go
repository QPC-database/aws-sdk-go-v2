// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Seals and completes the snapshot after all of the required blocks of data have
// been written to it. Completing the snapshot changes the status to completed. You
// cannot write new blocks to a snapshot after it has been completed.
func (c *Client) CompleteSnapshot(ctx context.Context, params *CompleteSnapshotInput, optFns ...func(*Options)) (*CompleteSnapshotOutput, error) {
	if params == nil {
		params = &CompleteSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteSnapshot", params, optFns, addOperationCompleteSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompleteSnapshotInput struct {

	// The number of blocks that were written to the snapshot.
	//
	// This member is required.
	ChangedBlocksCount *int32

	// The ID of the snapshot.
	//
	// This member is required.
	SnapshotId *string

	// An aggregated Base-64 SHA256 checksum based on the checksums of each written
	// block. To generate the aggregated checksum using the linear aggregation method,
	// arrange the checksums for each written block in ascending order of their block
	// index, concatenate them to form a single string, and then generate the checksum
	// on the entire string using the SHA256 algorithm.
	Checksum *string

	// The aggregation method used to generate the checksum. Currently, the only
	// supported aggregation method is LINEAR.
	ChecksumAggregationMethod types.ChecksumAggregationMethod

	// The algorithm used to generate the checksum. Currently, the only supported
	// algorithm is SHA256.
	ChecksumAlgorithm types.ChecksumAlgorithm
}

type CompleteSnapshotOutput struct {

	// The status of the snapshot.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCompleteSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCompleteSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCompleteSnapshot{}, middleware.After)
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
	addOpCompleteSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteSnapshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCompleteSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "CompleteSnapshot",
	}
}
