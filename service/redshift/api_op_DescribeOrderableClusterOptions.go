// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of orderable cluster options. Before you create a new cluster you
// can use this operation to find what options are available, such as the EC2
// Availability Zones (AZ) in the specific AWS Region that you can specify, and the
// node types you can request. The node types differ by available storage, memory,
// CPU and price. With the cost involved you might want to obtain a list of cluster
// options in the specific region and specify values when creating a cluster. For
// more information about managing clusters, go to Amazon Redshift Clusters
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in
// the Amazon Redshift Cluster Management Guide.
func (c *Client) DescribeOrderableClusterOptions(ctx context.Context, params *DescribeOrderableClusterOptionsInput, optFns ...func(*Options)) (*DescribeOrderableClusterOptionsOutput, error) {
	if params == nil {
		params = &DescribeOrderableClusterOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrderableClusterOptions", params, optFns, addOperationDescribeOrderableClusterOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrderableClusterOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeOrderableClusterOptionsInput struct {

	// The version filter value. Specify this parameter to show only the available
	// offerings matching the specified version. Default: All versions. Constraints:
	// Must be one of the version returned from DescribeClusterVersions ().
	ClusterVersion *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeOrderableClusterOptions ()
	// request exceed the value specified in MaxRecords, AWS returns a value in the
	// Marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The node type filter value. Specify this parameter to show only the available
	// offerings matching the specified node type.
	NodeType *string
}

// Contains the output from the DescribeOrderableClusterOptions () action.
type DescribeOrderableClusterOptionsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// An OrderableClusterOption structure containing information about orderable
	// options for the cluster.
	OrderableClusterOptions []*types.OrderableClusterOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeOrderableClusterOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeOrderableClusterOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeOrderableClusterOptions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrderableClusterOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeOrderableClusterOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeOrderableClusterOptions",
	}
}
