// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Performs all the write operations in a batch. Either all the operations succeed
// or none.
func (c *Client) BatchWrite(ctx context.Context, params *BatchWriteInput, optFns ...func(*Options)) (*BatchWriteOutput, error) {
	if params == nil {
		params = &BatchWriteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchWrite", params, optFns, addOperationBatchWriteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchWriteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchWriteInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory (). For
	// more information, see arns ().
	//
	// This member is required.
	DirectoryArn *string

	// A list of operations that are part of the batch.
	//
	// This member is required.
	Operations []*types.BatchWriteOperation
}

type BatchWriteOutput struct {

	// A list of all the responses for each batch write.
	Responses []*types.BatchWriteOperationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchWriteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchWrite{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchWrite{}, middleware.After)
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
	addOpBatchWriteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchWrite(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchWrite(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "BatchWrite",
	}
}
