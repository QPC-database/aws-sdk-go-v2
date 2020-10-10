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

// Creates a network interface in the specified subnet. For more information about
// network interfaces, see Elastic Network Interfaces
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateNetworkInterface(ctx context.Context, params *CreateNetworkInterfaceInput, optFns ...func(*Options)) (*CreateNetworkInterfaceOutput, error) {
	if params == nil {
		params = &CreateNetworkInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNetworkInterface", params, optFns, addOperationCreateNetworkInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNetworkInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateNetworkInterface.
type CreateNetworkInterfaceInput struct {

	// The ID of the subnet to associate with the network interface.
	//
	// This member is required.
	SubnetId *string

	// A description for the network interface.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The IDs of one or more security groups.
	Groups []*string

	// Indicates the type of network interface. To create an Elastic Fabric Adapter
	// (EFA), specify efa. For more information, see  Elastic Fabric Adapter
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/efa.html) in the Amazon
	// Elastic Compute Cloud User Guide.
	InterfaceType types.NetworkInterfaceCreationType

	// The number of IPv6 addresses to assign to a network interface. Amazon EC2
	// automatically selects the IPv6 addresses from the subnet range. You can't use
	// this option if specifying specific IPv6 addresses. If your subnet has the
	// AssignIpv6AddressOnCreation attribute set to true, you can specify 0 to override
	// this setting.
	Ipv6AddressCount *int32

	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your
	// subnet. You can't use this option if you're specifying a number of IPv6
	// addresses.
	Ipv6Addresses []*types.InstanceIpv6Address

	// The primary private IPv4 address of the network interface. If you don't specify
	// an IPv4 address, Amazon EC2 selects one for you from the subnet's IPv4 CIDR
	// range. If you specify an IP address, you cannot indicate any IP addresses
	// specified in privateIpAddresses as primary (only one IP address can be
	// designated as primary).
	PrivateIpAddress *string

	// One or more private IPv4 addresses.
	PrivateIpAddresses []*types.PrivateIpAddressSpecification

	// The number of secondary private IPv4 addresses to assign to a network interface.
	// When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these
	// IP addresses within the subnet's IPv4 CIDR range. You can't specify this option
	// and specify more than one private IP address using privateIpAddresses. The
	// number of IP addresses you can assign to a network interface varies by instance
	// type. For more information, see IP Addresses Per ENI Per Instance Type
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI)
	// in the Amazon Virtual Private Cloud User Guide.
	SecondaryPrivateIpAddressCount *int32

	// The tags to apply to the new network interface.
	TagSpecifications []*types.TagSpecification
}

// Contains the output of CreateNetworkInterface.
type CreateNetworkInterfaceOutput struct {

	// Information about the network interface.
	NetworkInterface *types.NetworkInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateNetworkInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateNetworkInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateNetworkInterface{}, middleware.After)
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
	addOpCreateNetworkInterfaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetworkInterface(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateNetworkInterface(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateNetworkInterface",
	}
}
