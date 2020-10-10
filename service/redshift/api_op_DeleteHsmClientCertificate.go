// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified HSM client certificate.
func (c *Client) DeleteHsmClientCertificate(ctx context.Context, params *DeleteHsmClientCertificateInput, optFns ...func(*Options)) (*DeleteHsmClientCertificateOutput, error) {
	if params == nil {
		params = &DeleteHsmClientCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteHsmClientCertificate", params, optFns, addOperationDeleteHsmClientCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteHsmClientCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteHsmClientCertificateInput struct {

	// The identifier of the HSM client certificate to be deleted.
	//
	// This member is required.
	HsmClientCertificateIdentifier *string
}

type DeleteHsmClientCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteHsmClientCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteHsmClientCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteHsmClientCertificate{}, middleware.After)
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
	addOpDeleteHsmClientCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHsmClientCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteHsmClientCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DeleteHsmClientCertificate",
	}
}
