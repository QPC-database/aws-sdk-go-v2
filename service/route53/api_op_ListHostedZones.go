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

// Retrieves a list of the public and private hosted zones that are associated with
// the current AWS account. The response includes a HostedZones child element for
// each hosted zone. Amazon Route 53 returns a maximum of 100 items in each
// response. If you have a lot of hosted zones, you can use the maxitems parameter
// to list them in groups of up to 100.
func (c *Client) ListHostedZones(ctx context.Context, params *ListHostedZonesInput, optFns ...func(*Options)) (*ListHostedZonesOutput, error) {
	if params == nil {
		params = &ListHostedZonesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHostedZones", params, optFns, addOperationListHostedZonesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHostedZonesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to retrieve a list of the public and private hosted zones that are
// associated with the current AWS account.
type ListHostedZonesInput struct {

	// If you're using reusable delegation sets and you want to list all of the hosted
	// zones that are associated with a reusable delegation set, specify the ID of that
	// reusable delegation set.
	DelegationSetId *string

	// If the value of IsTruncated in the previous response was true, you have more
	// hosted zones. To get more hosted zones, submit another ListHostedZones request.
	// For the value of marker, specify the value of NextMarker from the previous
	// response, which is the ID of the first hosted zone that Amazon Route 53 will
	// return if you submit another request. If the value of IsTruncated in the
	// previous response was false, there are no more hosted zones to get.
	Marker *string

	// (Optional) The maximum number of hosted zones that you want Amazon Route 53 to
	// return. If you have more than maxitems hosted zones, the value of IsTruncated in
	// the response is true, and the value of NextMarker is the hosted zone ID of the
	// first hosted zone that Route 53 will return if you submit another request.
	MaxItems *string
}

type ListHostedZonesOutput struct {

	// A complex type that contains general information about the hosted zone.
	//
	// This member is required.
	HostedZones []*types.HostedZone

	// A flag indicating whether there are more hosted zones to be listed. If the
	// response was truncated, you can get more hosted zones by submitting another
	// ListHostedZones request and specifying the value of NextMarker in the marker
	// parameter.
	//
	// This member is required.
	IsTruncated *bool

	// For the second and subsequent calls to ListHostedZones, Marker is the value that
	// you specified for the marker parameter in the request that produced the current
	// response.
	//
	// This member is required.
	Marker *string

	// The value that you specified for the maxitems parameter in the call to
	// ListHostedZones that produced the current response.
	//
	// This member is required.
	MaxItems *string

	// If IsTruncated is true, the value of NextMarker identifies the first hosted zone
	// in the next group of hosted zones. Submit another ListHostedZones request, and
	// specify the value of NextMarker from the response in the marker parameter. This
	// element is present only if IsTruncated is true.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListHostedZonesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListHostedZones{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListHostedZones{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListHostedZones(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addSanitizeURLMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListHostedZones(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListHostedZones",
	}
}
