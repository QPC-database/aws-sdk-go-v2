// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of certificate ARNs and domain names. You can request that only
// certificates that match a specific status be listed. You can also filter by
// specific attributes of the certificate. Default filtering returns only RSA_2048
// certificates. For more information, see Filters ().
func (c *Client) ListCertificates(ctx context.Context, params *ListCertificatesInput, optFns ...func(*Options)) (*ListCertificatesOutput, error) {
	if params == nil {
		params = &ListCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCertificates", params, optFns, addOperationListCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCertificatesInput struct {

	// Filter the certificate list by status value.
	CertificateStatuses []types.CertificateStatus

	// Filter the certificate list. For more information, see the Filters () structure.
	Includes *types.Filters

	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response. If additional items exist beyond the number you
	// specify, the NextToken element is sent in the response. Use this NextToken value
	// in a subsequent request to retrieve additional items.
	MaxItems *int32

	// Use this parameter only when paginating results and only in a subsequent request
	// after you receive a response with truncated results. Set it to the value of
	// NextToken from the response you just received.
	NextToken *string
}

type ListCertificatesOutput struct {

	// A list of ACM certificates.
	CertificateSummaryList []*types.CertificateSummary

	// When the list is truncated, this value is present and contains the value to use
	// for the NextToken parameter in a subsequent pagination request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCertificates{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCertificates(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCertificates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm",
		OperationName: "ListCertificates",
	}
}
