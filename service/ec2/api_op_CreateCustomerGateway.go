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

// Provides information to AWS about your VPN customer gateway device. The customer
// gateway is the appliance at your end of the VPN connection. (The device on the
// AWS side of the VPN connection is the virtual private gateway.) You must provide
// the Internet-routable IP address of the customer gateway's external interface.
// The IP address must be static and can be behind a device performing network
// address translation (NAT). For devices that use Border Gateway Protocol (BGP),
// you can also provide the device's BGP Autonomous System Number (ASN). You can
// use an existing ASN assigned to your network. If you don't have an ASN already,
// you can use a private ASN (in the 64512 - 65534 range). Amazon EC2 supports all
// 2-byte ASN numbers in the range of 1 - 65534, with the exception of 7224, which
// is reserved in the us-east-1 Region, and 9059, which is reserved in the
// eu-west-1 Region. For more information, see AWS Site-to-Site VPN
// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the AWS
// Site-to-Site VPN User Guide. To create more than one customer gateway with the
// same VPN type, IP address, and BGP ASN, specify a unique device name for each
// customer gateway. Identical requests return information about the existing
// customer gateway and do not create new customer gateways.
func (c *Client) CreateCustomerGateway(ctx context.Context, params *CreateCustomerGatewayInput, optFns ...func(*Options)) (*CreateCustomerGatewayOutput, error) {
	if params == nil {
		params = &CreateCustomerGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomerGateway", params, optFns, addOperationCreateCustomerGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomerGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateCustomerGateway.
type CreateCustomerGatewayInput struct {

	// For devices that support BGP, the customer gateway's BGP ASN. Default: 65000
	//
	// This member is required.
	BgpAsn *int32

	// The type of VPN connection that this customer gateway supports (ipsec.1).
	//
	// This member is required.
	Type types.GatewayType

	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn *string

	// A name for the customer gateway device. Length Constraints: Up to 255
	// characters.
	DeviceName *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The Internet-routable IP address for the customer gateway's outside interface.
	// The address must be static.
	PublicIp *string

	// The tags to apply to the customer gateway.
	TagSpecifications []*types.TagSpecification
}

// Contains the output of CreateCustomerGateway.
type CreateCustomerGatewayOutput struct {

	// Information about the customer gateway.
	CustomerGateway *types.CustomerGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCustomerGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateCustomerGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateCustomerGateway{}, middleware.After)
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
	addOpCreateCustomerGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomerGateway(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateCustomerGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateCustomerGateway",
	}
}
