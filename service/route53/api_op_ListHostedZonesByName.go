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

// Retrieves a list of your hosted zones in lexicographic order. The response
// includes a HostedZones child element for each hosted zone created by the current
// AWS account. ListHostedZonesByName sorts hosted zones by name with the labels
// reversed. For example:  <p> <code>com.example.www.</code> </p> <p>Note the
// trailing dot, which can change the sort order in some circumstances.</p> <p>If
// the domain name includes escape characters or Punycode,
// <code>ListHostedZonesByName</code> alphabetizes the domain name using the
// escaped or Punycoded value, which is the format that Amazon Route 53 saves in
// its database. For example, to create a hosted zone for exämple.com, you specify
// ex\344mple.com for the domain name. <code>ListHostedZonesByName</code>
// alphabetizes it as:</p> <p> <code>com.ex\344mple.</code> </p> <p>The labels are
// reversed and alphabetized using the escaped value. For more information about
// valid domain name formats, including internationalized domain names, see <a
// href="https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html">DNS
// Domain Name Format</a> in the <i>Amazon Route 53 Developer Guide</i>.</p>
// <p>Route 53 returns up to 100 items in each response. If you have a lot of
// hosted zones, use the <code>MaxItems</code> parameter to list them in groups of
// up to 100. The response includes values that help navigate from one group of
// <code>MaxItems</code> hosted zones to the next:</p> <ul> <li> <p>The
// <code>DNSName</code> and <code>HostedZoneId</code> elements in the response
// contain the values, if any, specified for the <code>dnsname</code> and
// <code>hostedzoneid</code> parameters in the request that produced the current
// response.</p> </li> <li> <p>The <code>MaxItems</code> element in the response
// contains the value, if any, that you specified for the <code>maxitems</code>
// parameter in the request that produced the current response.</p> </li> <li>
// <p>If the value of <code>IsTruncated</code> in the response is true, there are
// more hosted zones associated with the current AWS account. </p> <p>If
// <code>IsTruncated</code> is false, this response includes the last hosted zone
// that is associated with the current account. The <code>NextDNSName</code>
// element and <code>NextHostedZoneId</code> elements are omitted from the
// response.</p> </li> <li> <p>The <code>NextDNSName</code> and
// <code>NextHostedZoneId</code> elements in the response contain the domain name
// and the hosted zone ID of the next hosted zone that is associated with the
// current AWS account. If you want to list more hosted zones, make another call to
// <code>ListHostedZonesByName</code>, and specify the value of
// <code>NextDNSName</code> and <code>NextHostedZoneId</code> in the
// <code>dnsname</code> and <code>hostedzoneid</code> parameters, respectively.</p>
// </li> </ul>
func (c *Client) ListHostedZonesByName(ctx context.Context, params *ListHostedZonesByNameInput, optFns ...func(*Options)) (*ListHostedZonesByNameOutput, error) {
	if params == nil {
		params = &ListHostedZonesByNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHostedZonesByName", params, optFns, addOperationListHostedZonesByNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHostedZonesByNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves a list of the public and private hosted zones that are associated with
// the current AWS account in ASCII order by domain name.
type ListHostedZonesByNameInput struct {

	// (Optional) For your first request to ListHostedZonesByName, include the dnsname
	// parameter only if you want to specify the name of the first hosted zone in the
	// response. If you don't include the dnsname parameter, Amazon Route 53 returns
	// all of the hosted zones that were created by the current AWS account, in ASCII
	// order. For subsequent requests, include both dnsname and hostedzoneid
	// parameters. For dnsname, specify the value of NextDNSName from the previous
	// response.
	DNSName *string

	// (Optional) For your first request to ListHostedZonesByName, do not include the
	// hostedzoneid parameter. If you have more hosted zones than the value of
	// maxitems, ListHostedZonesByName returns only the first maxitems hosted zones. To
	// get the next group of maxitems hosted zones, submit another request to
	// ListHostedZonesByName and include both dnsname and hostedzoneid parameters. For
	// the value of hostedzoneid, specify the value of the NextHostedZoneId element
	// from the previous response.
	HostedZoneId *string

	// The maximum number of hosted zones to be included in the response body for this
	// request. If you have more than maxitems hosted zones, then the value of the
	// IsTruncated element in the response is true, and the values of NextDNSName and
	// NextHostedZoneId specify the first hosted zone in the next group of maxitems
	// hosted zones.
	MaxItems *string
}

// A complex type that contains the response information for the request.
type ListHostedZonesByNameOutput struct {

	// A complex type that contains general information about the hosted zone.
	//
	// This member is required.
	HostedZones []*types.HostedZone

	// A flag that indicates whether there are more hosted zones to be listed. If the
	// response was truncated, you can get the next group of maxitems hosted zones by
	// calling ListHostedZonesByName again and specifying the values of NextDNSName and
	// NextHostedZoneId elements in the dnsname and hostedzoneid parameters.
	//
	// This member is required.
	IsTruncated *bool

	// The value that you specified for the maxitems parameter in the call to
	// ListHostedZonesByName that produced the current response.
	//
	// This member is required.
	MaxItems *string

	// For the second and subsequent calls to ListHostedZonesByName, DNSName is the
	// value that you specified for the dnsname parameter in the request that produced
	// the current response.
	DNSName *string

	// The ID that Amazon Route 53 assigned to the hosted zone when you created it.
	HostedZoneId *string

	// If IsTruncated is true, the value of NextDNSName is the name of the first hosted
	// zone in the next group of maxitems hosted zones. Call ListHostedZonesByName
	// again and specify the value of NextDNSName and NextHostedZoneId in the dnsname
	// and hostedzoneid parameters, respectively. This element is present only if
	// IsTruncated is true.
	NextDNSName *string

	// If IsTruncated is true, the value of NextHostedZoneId identifies the first
	// hosted zone in the next group of maxitems hosted zones. Call
	// ListHostedZonesByName again and specify the value of NextDNSName and
	// NextHostedZoneId in the dnsname and hostedzoneid parameters, respectively. This
	// element is present only if IsTruncated is true.
	NextHostedZoneId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListHostedZonesByNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListHostedZonesByName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListHostedZonesByName{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListHostedZonesByName(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListHostedZonesByName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListHostedZonesByName",
	}
}
