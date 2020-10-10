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

// Returns one or more cluster subnet group objects, which contain metadata about
// your cluster subnet groups. By default, this operation returns information about
// all cluster subnet groups that are defined in you AWS account. If you specify
// both tag keys and tag values in the same request, Amazon Redshift returns all
// subnet groups that match any combination of the specified keys and values. For
// example, if you have owner and environment for tag keys, and admin and test for
// tag values, all subnet groups that have any combination of those values are
// returned. If both tag keys and values are omitted from the request, subnet
// groups are returned regardless of whether they have tag keys or values
// associated with them.
func (c *Client) DescribeClusterSubnetGroups(ctx context.Context, params *DescribeClusterSubnetGroupsInput, optFns ...func(*Options)) (*DescribeClusterSubnetGroupsOutput, error) {
	if params == nil {
		params = &DescribeClusterSubnetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusterSubnetGroups", params, optFns, addOperationDescribeClusterSubnetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterSubnetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeClusterSubnetGroupsInput struct {

	// The name of the cluster subnet group for which information is requested.
	ClusterSubnetGroupName *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterSubnetGroups () request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records by
	// providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// A tag key or keys for which you want to return all matching cluster subnet
	// groups that are associated with the specified key or keys. For example, suppose
	// that you have subnet groups that are tagged with keys called owner and
	// environment. If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the subnet groups that have either or both of
	// these tag keys associated with them.
	TagKeys []*string

	// A tag value or values for which you want to return all matching cluster subnet
	// groups that are associated with the specified tag value or values. For example,
	// suppose that you have subnet groups that are tagged with values called admin and
	// test. If you specify both of these tag values in the request, Amazon Redshift
	// returns a response with the subnet groups that have either or both of these tag
	// values associated with them.
	TagValues []*string
}

// Contains the output from the DescribeClusterSubnetGroups () action.
type DescribeClusterSubnetGroupsOutput struct {

	// A list of ClusterSubnetGroup () instances.
	ClusterSubnetGroups []*types.ClusterSubnetGroup

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeClusterSubnetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterSubnetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterSubnetGroups{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterSubnetGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeClusterSubnetGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusterSubnetGroups",
	}
}
