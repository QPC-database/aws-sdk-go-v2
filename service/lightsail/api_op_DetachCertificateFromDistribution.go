// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches an SSL/TLS certificate from your Amazon Lightsail content delivery
// network (CDN) distribution.  <p>After the certificate is detached, your
// distribution stops accepting traffic for all of the domains that are associated
// with the certificate.</p>
func (c *Client) DetachCertificateFromDistribution(ctx context.Context, params *DetachCertificateFromDistributionInput, optFns ...func(*Options)) (*DetachCertificateFromDistributionOutput, error) {
	if params == nil {
		params = &DetachCertificateFromDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachCertificateFromDistribution", params, optFns, addOperationDetachCertificateFromDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachCertificateFromDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachCertificateFromDistributionInput struct {

	// The name of the distribution from which to detach the certificate.  <p>Use the
	// <code>GetDistributions</code> action to get a list of distribution names that
	// you can specify.</p>
	//
	// This member is required.
	DistributionName *string
}

type DetachCertificateFromDistributionOutput struct {

	// An object that describes the result of the action, such as the status of the
	// request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachCertificateFromDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetachCertificateFromDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachCertificateFromDistribution{}, middleware.After)
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
	addOpDetachCertificateFromDistributionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachCertificateFromDistribution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetachCertificateFromDistribution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DetachCertificateFromDistribution",
	}
}
