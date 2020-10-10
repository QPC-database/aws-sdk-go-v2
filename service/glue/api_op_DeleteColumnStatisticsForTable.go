// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves table statistics of columns.
func (c *Client) DeleteColumnStatisticsForTable(ctx context.Context, params *DeleteColumnStatisticsForTableInput, optFns ...func(*Options)) (*DeleteColumnStatisticsForTableOutput, error) {
	if params == nil {
		params = &DeleteColumnStatisticsForTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteColumnStatisticsForTable", params, optFns, addOperationDeleteColumnStatisticsForTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteColumnStatisticsForTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteColumnStatisticsForTableInput struct {

	// The name of the column.
	//
	// This member is required.
	ColumnName *string

	// The name of the catalog database where the partitions reside.
	//
	// This member is required.
	DatabaseName *string

	// The name of the partitions' table.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the partitions in question reside. If none is
	// supplied, the AWS account ID is used by default.
	CatalogId *string
}

type DeleteColumnStatisticsForTableOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteColumnStatisticsForTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteColumnStatisticsForTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteColumnStatisticsForTable{}, middleware.After)
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
	addOpDeleteColumnStatisticsForTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteColumnStatisticsForTable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteColumnStatisticsForTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteColumnStatisticsForTable",
	}
}
