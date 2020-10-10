// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes one or more partitions in a batch operation.
func (c *Client) BatchDeletePartition(ctx context.Context, params *BatchDeletePartitionInput, optFns ...func(*Options)) (*BatchDeletePartitionOutput, error) {
	if params == nil {
		params = &BatchDeletePartitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeletePartition", params, optFns, addOperationBatchDeletePartitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeletePartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeletePartitionInput struct {

	// The name of the catalog database in which the table in question resides.
	//
	// This member is required.
	DatabaseName *string

	// A list of PartitionInput structures that define the partitions to be deleted.
	//
	// This member is required.
	PartitionsToDelete []*types.PartitionValueList

	// The name of the table that contains the partitions to be deleted.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the partition to be deleted resides. If none is
	// provided, the AWS account ID is used by default.
	CatalogId *string
}

type BatchDeletePartitionOutput struct {

	// The errors encountered when trying to delete the requested partitions.
	Errors []*types.PartitionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDeletePartitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeletePartition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeletePartition{}, middleware.After)
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
	addOpBatchDeletePartitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeletePartition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchDeletePartition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchDeletePartition",
	}
}
