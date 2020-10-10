// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a CA certificate with AWS IoT. This CA certificate can then be used to
// sign device certificates, which can be then registered with AWS IoT. You can
// register up to 10 CA certificates per AWS account that have the same subject
// field. This enables you to have up to 10 certificate authorities sign your
// device certificates. If you have more than one CA certificate registered, make
// sure you pass the CA certificate when you register your device certificates with
// the RegisterCertificate API.
func (c *Client) RegisterCACertificate(ctx context.Context, params *RegisterCACertificateInput, optFns ...func(*Options)) (*RegisterCACertificateOutput, error) {
	if params == nil {
		params = &RegisterCACertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterCACertificate", params, optFns, addOperationRegisterCACertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterCACertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the RegisterCACertificate operation.
type RegisterCACertificateInput struct {

	// The CA certificate.
	//
	// This member is required.
	CaCertificate *string

	// The private key verification certificate.
	//
	// This member is required.
	VerificationCertificate *string

	// Allows this CA certificate to be used for auto registration of device
	// certificates.
	AllowAutoRegistration *bool

	// Information about the registration configuration.
	RegistrationConfig *types.RegistrationConfig

	// A boolean value that specifies if the CA certificate is set to active.
	SetAsActive *bool

	// Metadata which can be used to manage the CA certificate. For URI Request
	// parameters use format: ...key1=value1&key2=value2... For the CLI command-line
	// parameter use format: &&tags "key1=value1&key2=value2..." For the cli-input-json
	// file use format: "tags": "key1=value1&key2=value2..."
	Tags []*types.Tag
}

// The output from the RegisterCACertificateResponse operation.
type RegisterCACertificateOutput struct {

	// The CA certificate ARN.
	CertificateArn *string

	// The CA certificate identifier.
	CertificateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterCACertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterCACertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterCACertificate{}, middleware.After)
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
	addOpRegisterCACertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCACertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterCACertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RegisterCACertificate",
	}
}
