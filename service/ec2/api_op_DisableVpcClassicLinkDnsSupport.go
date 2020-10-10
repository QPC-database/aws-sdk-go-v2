// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables ClassicLink DNS support for a VPC. If disabled, DNS hostnames resolve
// to public IP addresses when addressed between a linked EC2-Classic instance and
// instances in the VPC to which it's linked. For more information, see ClassicLink
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
// the Amazon Elastic Compute Cloud User Guide. You must specify a VPC ID in the
// request.
func (c *Client) DisableVpcClassicLinkDnsSupport(ctx context.Context, params *DisableVpcClassicLinkDnsSupportInput, optFns ...func(*Options)) (*DisableVpcClassicLinkDnsSupportOutput, error) {
	if params == nil {
		params = &DisableVpcClassicLinkDnsSupportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableVpcClassicLinkDnsSupport", params, optFns, addOperationDisableVpcClassicLinkDnsSupportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableVpcClassicLinkDnsSupportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableVpcClassicLinkDnsSupportInput struct {

	// The ID of the VPC.
	VpcId *string
}

type DisableVpcClassicLinkDnsSupportOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableVpcClassicLinkDnsSupportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisableVpcClassicLinkDnsSupport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisableVpcClassicLinkDnsSupport{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableVpcClassicLinkDnsSupport(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisableVpcClassicLinkDnsSupport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisableVpcClassicLinkDnsSupport",
	}
}
