// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specifies a domain to be associated to Amazon WorkLink.
func (c *Client) AssociateDomain(ctx context.Context, params *AssociateDomainInput, optFns ...func(*Options)) (*AssociateDomainOutput, error) {
	if params == nil {
		params = &AssociateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateDomain", params, optFns, addOperationAssociateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDomainInput struct {

	// The ARN of an issued ACM certificate that is valid for the domain being
	// associated.
	//
	// This member is required.
	AcmCertificateArn *string

	// The fully qualified domain name (FQDN).
	//
	// This member is required.
	DomainName *string

	// The Amazon Resource Name (ARN) of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The name to display.
	DisplayName *string
}

type AssociateDomainOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateDomain{}, middleware.After)
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
	addOpAssociateDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDomain(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssociateDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "AssociateDomain",
	}
}
