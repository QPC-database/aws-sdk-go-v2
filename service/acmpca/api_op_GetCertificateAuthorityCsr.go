// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the certificate signing request (CSR) for your private certificate
// authority (CA). The CSR is created when you call the CreateCertificateAuthority
// () action. Sign the CSR with your ACM Private CA-hosted or on-premises root or
// subordinate CA. Then import the signed certificate back into ACM Private CA by
// calling the ImportCertificateAuthorityCertificate () action. The CSR is returned
// as a base64 PEM-encoded string.
func (c *Client) GetCertificateAuthorityCsr(ctx context.Context, params *GetCertificateAuthorityCsrInput, optFns ...func(*Options)) (*GetCertificateAuthorityCsrOutput, error) {
	if params == nil {
		params = &GetCertificateAuthorityCsrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCertificateAuthorityCsr", params, optFns, addOperationGetCertificateAuthorityCsrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCertificateAuthorityCsrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCertificateAuthorityCsrInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called the
	// CreateCertificateAuthority () action. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string
}

type GetCertificateAuthorityCsrOutput struct {

	// The base64 PEM-encoded certificate signing request (CSR) for your private CA
	// certificate.
	Csr *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCertificateAuthorityCsrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCertificateAuthorityCsr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCertificateAuthorityCsr{}, middleware.After)
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
	addOpGetCertificateAuthorityCsrValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCertificateAuthorityCsr(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCertificateAuthorityCsr(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "GetCertificateAuthorityCsr",
	}
}
