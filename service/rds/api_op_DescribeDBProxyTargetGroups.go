// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about DB proxy target groups, represented by
// DBProxyTargetGroup data structures.
func (c *Client) DescribeDBProxyTargetGroups(ctx context.Context, params *DescribeDBProxyTargetGroupsInput, optFns ...func(*Options)) (*DescribeDBProxyTargetGroupsOutput, error) {
	if params == nil {
		params = &DescribeDBProxyTargetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBProxyTargetGroups", params, optFns, addOperationDescribeDBProxyTargetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBProxyTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBProxyTargetGroupsInput struct {

	// The identifier of the DBProxy associated with the target group.
	//
	// This member is required.
	DBProxyName *string

	// This parameter is not currently supported.
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

	// The identifier of the DBProxyTargetGroup to describe.
	TargetGroupName *string
}

type DescribeDBProxyTargetGroupsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// An arbitrary number of DBProxyTargetGroup objects, containing details of the
	// corresponding target groups.
	TargetGroups []*types.DBProxyTargetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBProxyTargetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBProxyTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBProxyTargetGroups{}, middleware.After)
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
	addOpDescribeDBProxyTargetGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBProxyTargetGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDBProxyTargetGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBProxyTargetGroups",
	}
}
