// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all global tables that have a replica in the specified Region. This
// operation only applies to Version 2017.11.29
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
// of global tables.
func (c *Client) ListGlobalTables(ctx context.Context, params *ListGlobalTablesInput, optFns ...func(*Options)) (*ListGlobalTablesOutput, error) {
	if params == nil {
		params = &ListGlobalTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGlobalTables", params, optFns, addOperationListGlobalTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGlobalTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGlobalTablesInput struct {

	// The first global table name that this operation will evaluate.
	ExclusiveStartGlobalTableName *string

	// The maximum number of table names to return, if the parameter is not specified
	// DynamoDB defaults to 100. If the number of global tables DynamoDB finds reaches
	// this limit, it stops the operation and returns the table names collected up to
	// that point, with a table name in the LastEvaluatedGlobalTableName to apply in a
	// subsequent operation to the ExclusiveStartGlobalTableName parameter.
	Limit *int32

	// Lists the global tables in a specific Region.
	RegionName *string
}

type ListGlobalTablesOutput struct {

	// List of global table names.
	GlobalTables []*types.GlobalTable

	// Last evaluated global table name.
	LastEvaluatedGlobalTableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListGlobalTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListGlobalTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListGlobalTables{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListGlobalTables(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)
	return nil
}

func newServiceMetadataMiddleware_opListGlobalTables(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ListGlobalTables",
	}
}
