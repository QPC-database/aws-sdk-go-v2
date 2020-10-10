// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the status of the connections that have been made between the
// replication instance and an endpoint. Connections are created when you test an
// endpoint.
func (c *Client) DescribeConnections(ctx context.Context, params *DescribeConnectionsInput, optFns ...func(*Options)) (*DescribeConnectionsOutput, error) {
	if params == nil {
		params = &DescribeConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnections", params, optFns, addOperationDescribeConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeConnectionsInput struct {

	// The filters applied to the connection. Valid filter names: endpoint-arn |
	// replication-instance-arn
	Filters []*types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

//
type DescribeConnectionsOutput struct {

	// A description of the connections.
	Connections []*types.Connection

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConnections{}, middleware.After)
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
	addOpDescribeConnectionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnections(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeConnections",
	}
}
