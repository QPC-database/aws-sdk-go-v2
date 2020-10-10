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

// Returns the detailed parameter list for a particular DB cluster parameter group.
// For more information on Amazon Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) DescribeDBClusterParameters(ctx context.Context, params *DescribeDBClusterParametersInput, optFns ...func(*Options)) (*DescribeDBClusterParametersOutput, error) {
	if params == nil {
		params = &DescribeDBClusterParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterParameters", params, optFns, addOperationDescribeDBClusterParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBClusterParametersInput struct {

	// The name of a specific DB cluster parameter group to return parameter details
	// for. Constraints:
	//
	//     * If supplied, must match the name of an existing
	// DBClusterParameterGroup.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// This parameter isn't currently supported.
	Filters []*types.Filter

	// An optional pagination token provided by a previous DescribeDBClusterParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// A value that indicates to return only parameters for a specific source.
	// Parameter sources can be engine, service, or customer.
	Source *string
}

// Provides details about a DB cluster parameter group including the parameters in
// the DB cluster parameter group.
type DescribeDBClusterParametersOutput struct {

	// An optional pagination token provided by a previous DescribeDBClusterParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// Provides a list of parameters for the DB cluster parameter group.
	Parameters []*types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBClusterParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterParameters{}, middleware.After)
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
	addOpDescribeDBClusterParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterParameters(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDBClusterParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterParameters",
	}
}
