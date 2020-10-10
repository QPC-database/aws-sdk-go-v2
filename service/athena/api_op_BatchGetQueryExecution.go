// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details of a single query execution or a list of up to 50 query
// executions, which you provide as an array of query execution ID strings.
// Requires you to have access to the workgroup in which the queries ran. To get a
// list of query execution IDs, use ListQueryExecutionsInput$WorkGroup (). Query
// executions differ from named (saved) queries. Use BatchGetNamedQueryInput () to
// get details about named queries.
func (c *Client) BatchGetQueryExecution(ctx context.Context, params *BatchGetQueryExecutionInput, optFns ...func(*Options)) (*BatchGetQueryExecutionOutput, error) {
	if params == nil {
		params = &BatchGetQueryExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetQueryExecution", params, optFns, addOperationBatchGetQueryExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetQueryExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetQueryExecutionInput struct {

	// An array of query execution IDs.
	//
	// This member is required.
	QueryExecutionIds []*string
}

type BatchGetQueryExecutionOutput struct {

	// Information about a query execution.
	QueryExecutions []*types.QueryExecution

	// Information about the query executions that failed to run.
	UnprocessedQueryExecutionIds []*types.UnprocessedQueryExecutionId

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchGetQueryExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetQueryExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetQueryExecution{}, middleware.After)
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
	addOpBatchGetQueryExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetQueryExecution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchGetQueryExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "BatchGetQueryExecution",
	}
}
