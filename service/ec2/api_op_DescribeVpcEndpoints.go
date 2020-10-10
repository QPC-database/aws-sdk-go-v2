// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your VPC endpoints.
func (c *Client) DescribeVpcEndpoints(ctx context.Context, params *DescribeVpcEndpointsInput, optFns ...func(*Options)) (*DescribeVpcEndpointsOutput, error) {
	if params == nil {
		params = &DescribeVpcEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcEndpoints", params, optFns, addOperationDescribeVpcEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeVpcEndpoints.
type DescribeVpcEndpointsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * service-name - The name of the service.
	//
	//     *
	// vpc-id - The ID of the VPC in which the endpoint resides.
	//
	//     * vpc-endpoint-id
	// - The ID of the endpoint.
	//
	//     * vpc-endpoint-state - The state of the endpoint
	// (pendingAcceptance | pending | available | deleting | deleted | rejected |
	// failed).
	//
	//     * tag: - The key/value combination of a tag assigned to the
	// resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	//     * tag-key - The key of a tag assigned to the resource. Use this
	// filter to find all resources assigned a tag with a specific key, regardless of
	// the tag value.
	Filters []*types.Filter

	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// Constraint: If the value is greater than 1,000, we return only 1,000 items.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// prior call.)
	NextToken *string

	// One or more endpoint IDs.
	VpcEndpointIds []*string
}

// Contains the output of DescribeVpcEndpoints.
type DescribeVpcEndpointsOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the endpoints.
	VpcEndpoints []*types.VpcEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpcEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpoints{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpoints(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeVpcEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcEndpoints",
	}
}
