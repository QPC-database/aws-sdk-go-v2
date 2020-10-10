// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the value that Amazon Route 53 returns in response to a DNS request for a
// specified record name and type. You can optionally specify the IP address of a
// DNS resolver, an EDNS0 client subnet IP address, and a subnet mask.
func (c *Client) TestDNSAnswer(ctx context.Context, params *TestDNSAnswerInput, optFns ...func(*Options)) (*TestDNSAnswerOutput, error) {
	if params == nil {
		params = &TestDNSAnswerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestDNSAnswer", params, optFns, addOperationTestDNSAnswerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestDNSAnswerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets the value that Amazon Route 53 returns in response to a DNS request for a
// specified record name and type. You can optionally specify the IP address of a
// DNS resolver, an EDNS0 client subnet IP address, and a subnet mask.
type TestDNSAnswerInput struct {

	// The ID of the hosted zone that you want Amazon Route 53 to simulate a query for.
	//
	// This member is required.
	HostedZoneId *string

	// The name of the resource record set that you want Amazon Route 53 to simulate a
	// query for.
	//
	// This member is required.
	RecordName *string

	// The type of the resource record set.
	//
	// This member is required.
	RecordType types.RRType

	// If the resolver that you specified for resolverip supports EDNS0, specify the
	// IPv4 or IPv6 address of a client in the applicable location, for example,
	// 192.0.2.44 or 2001:db8:85a3::8a2e:370:7334.
	EDNS0ClientSubnetIP *string

	// If you specify an IP address for edns0clientsubnetip, you can optionally specify
	// the number of bits of the IP address that you want the checking tool to include
	// in the DNS query. For example, if you specify 192.0.2.44 for edns0clientsubnetip
	// and 24 for edns0clientsubnetmask, the checking tool will simulate a request from
	// 192.0.2.0/24. The default value is 24 bits for IPv4 addresses and 64 bits for
	// IPv6 addresses. The range of valid values depends on whether edns0clientsubnetip
	// is an IPv4 or an IPv6 address:
	//
	//     * IPv4: Specify a value between 0 and 32
	//
	//
	// * IPv6: Specify a value between 0 and 128
	EDNS0ClientSubnetMask *string

	// If you want to simulate a request from a specific DNS resolver, specify the IP
	// address for that resolver. If you omit this value, TestDnsAnswer uses the IP
	// address of a DNS resolver in the AWS US East (N. Virginia) Region (us-east-1).
	ResolverIP *string
}

// A complex type that contains the response to a TestDNSAnswer request.
type TestDNSAnswerOutput struct {

	// The Amazon Route 53 name server used to respond to the request.
	//
	// This member is required.
	Nameserver *string

	// The protocol that Amazon Route 53 used to respond to the request, either UDP or
	// TCP.
	//
	// This member is required.
	Protocol *string

	// A list that contains values that Amazon Route 53 returned for this resource
	// record set.
	//
	// This member is required.
	RecordData []*string

	// The name of the resource record set that you submitted a request for.
	//
	// This member is required.
	RecordName *string

	// The type of the resource record set that you submitted a request for.
	//
	// This member is required.
	RecordType types.RRType

	// A code that indicates whether the request is valid or not. The most common
	// response code is NOERROR, meaning that the request is valid. If the response is
	// not valid, Amazon Route 53 returns a response code that describes the error. For
	// a list of possible response codes, see DNS RCODES
	// (http://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-6)
	// on the IANA website.
	//
	// This member is required.
	ResponseCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTestDNSAnswerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpTestDNSAnswer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpTestDNSAnswer{}, middleware.After)
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
	addOpTestDNSAnswerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestDNSAnswer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTestDNSAnswer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "TestDNSAnswer",
	}
}
