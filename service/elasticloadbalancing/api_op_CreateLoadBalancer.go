// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Classic Load Balancer.  <p>You can add listeners, security groups,
// subnets, and tags when you create your load balancer, or you can add them later
// using <a>CreateLoadBalancerListeners</a>,
// <a>ApplySecurityGroupsToLoadBalancer</a>, <a>AttachLoadBalancerToSubnets</a>,
// and <a>AddTags</a>.</p> <p>To describe your current load balancers, see
// <a>DescribeLoadBalancers</a>. When you are finished with a load balancer, you
// can delete it using <a>DeleteLoadBalancer</a>.</p> <p>You can create up to 20
// load balancers per region per account. You can request an increase for the
// number of load balancers for your account. For more information, see <a
// href="https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-limits.html">Limits
// for Your Classic Load Balancer</a> in the <i>Classic Load Balancers
// Guide</i>.</p>
func (c *Client) CreateLoadBalancer(ctx context.Context, params *CreateLoadBalancerInput, optFns ...func(*Options)) (*CreateLoadBalancerOutput, error) {
	if params == nil {
		params = &CreateLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoadBalancer", params, optFns, addOperationCreateLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateLoadBalancer.
type CreateLoadBalancerInput struct {

	// The listeners. For more information, see Listeners for Your Classic Load
	// Balancer
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html)
	// in the Classic Load Balancers Guide.
	//
	// This member is required.
	Listeners []*types.Listener

	// The name of the load balancer. This name must be unique within your set of load
	// balancers for the region, must have a maximum of 32 characters, must contain
	// only alphanumeric characters or hyphens, and cannot begin or end with a hyphen.
	//
	// This member is required.
	LoadBalancerName *string

	// One or more Availability Zones from the same region as the load balancer. You
	// must specify at least one Availability Zone. You can add more Availability Zones
	// after you create the load balancer using EnableAvailabilityZonesForLoadBalancer
	// ().
	AvailabilityZones []*string

	// The type of a load balancer. Valid only for load balancers in a VPC. By default,
	// Elastic Load Balancing creates an Internet-facing load balancer with a DNS name
	// that resolves to public IP addresses. For more information about Internet-facing
	// and Internal load balancers, see Load Balancer Scheme
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/how-elastic-load-balancing-works.html#load-balancer-scheme)
	// in the Elastic Load Balancing User Guide. Specify internal to create a load
	// balancer with a DNS name that resolves to private IP addresses.
	Scheme *string

	// The IDs of the security groups to assign to the load balancer.
	SecurityGroups []*string

	// The IDs of the subnets in your VPC to attach to the load balancer. Specify one
	// subnet per Availability Zone specified in AvailabilityZones.
	Subnets []*string

	// A list of tags to assign to the load balancer. For more information about
	// tagging your load balancer, see Tag Your Classic Load Balancer
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/add-remove-tags.html)
	// in the Classic Load Balancers Guide.
	Tags []*types.Tag
}

// Contains the output for CreateLoadBalancer.
type CreateLoadBalancerOutput struct {

	// The DNS name of the load balancer.
	DNSName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLoadBalancer{}, middleware.After)
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
	addOpCreateLoadBalancerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoadBalancer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateLoadBalancer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateLoadBalancer",
	}
}
