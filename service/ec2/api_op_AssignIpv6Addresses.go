// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns one or more IPv6 addresses to the specified network interface. You can
// specify one or more specific IPv6 addresses, or you can specify the number of
// IPv6 addresses to be automatically assigned from within the subnet's IPv6 CIDR
// block range. You can assign as many IPv6 addresses to a network interface as you
// can assign private IPv4 addresses, and the limit varies per instance type. For
// information, see IP Addresses Per Network Interface Per Instance Type
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI)
// in the Amazon Elastic Compute Cloud User Guide. You must specify either the IPv6
// addresses or the IPv6 address count in the request.
func (c *Client) AssignIpv6Addresses(ctx context.Context, params *AssignIpv6AddressesInput, optFns ...func(*Options)) (*AssignIpv6AddressesOutput, error) {
	if params == nil {
		params = &AssignIpv6AddressesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssignIpv6Addresses", params, optFns, addOperationAssignIpv6AddressesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssignIpv6AddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssignIpv6AddressesInput struct {

	// The ID of the network interface.
	//
	// This member is required.
	NetworkInterfaceId *string

	// The number of IPv6 addresses to assign to the network interface. Amazon EC2
	// automatically selects the IPv6 addresses from the subnet range. You can't use
	// this option if specifying specific IPv6 addresses.
	Ipv6AddressCount *int32

	// One or more specific IPv6 addresses to be assigned to the network interface. You
	// can't use this option if you're specifying a number of IPv6 addresses.
	Ipv6Addresses []*string
}

type AssignIpv6AddressesOutput struct {

	// The IPv6 addresses assigned to the network interface.
	AssignedIpv6Addresses []*string

	// The ID of the network interface.
	NetworkInterfaceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssignIpv6AddressesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssignIpv6Addresses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssignIpv6Addresses{}, middleware.After)
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
	addOpAssignIpv6AddressesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssignIpv6Addresses(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssignIpv6Addresses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssignIpv6Addresses",
	}
}
