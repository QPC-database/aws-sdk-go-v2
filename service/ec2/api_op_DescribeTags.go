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

// Describes the specified tags for your EC2 resources. For more information about
// tags, see Tagging Your Resources
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) {
	if params == nil {
		params = &DescribeTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTags", params, optFns, addOperationDescribeTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTagsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters.
	//
	//     * key - The tag key.
	//
	//     * resource-id - The ID of the
	// resource.
	//
	//     * resource-type - The resource type (customer-gateway |
	// dedicated-host | dhcp-options | elastic-ip | fleet | fpga-image |
	// host-reservation | image | instance | internet-gateway | key-pair |
	// launch-template | natgateway | network-acl | network-interface | placement-group
	// | reserved-instances | route-table | security-group | snapshot |
	// spot-instances-request | subnet | volume | vpc | vpc-endpoint |
	// vpc-endpoint-service | vpc-peering-connection | vpn-connection | vpn-gateway).
	//
	//
	// * tag: - The key/value combination of the tag. For example, specify "tag:Owner"
	// for the filter name and "TeamA" for the filter value to find resources with the
	// tag "Owner=TeamA".
	//
	//     * value - The tag value.
	Filters []*types.Filter

	// The maximum number of results to return in a single call. This value can be
	// between 5 and 1000. To retrieve the remaining results, make another call with
	// the returned NextToken value.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string
}

type DescribeTagsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The tags.
	Tags []*types.TagDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTags{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTags(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTags",
	}
}
