// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the identity provider configuration of the specified fleet.
func (c *Client) DescribeIdentityProviderConfiguration(ctx context.Context, params *DescribeIdentityProviderConfigurationInput, optFns ...func(*Options)) (*DescribeIdentityProviderConfigurationOutput, error) {
	if params == nil {
		params = &DescribeIdentityProviderConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIdentityProviderConfiguration", params, optFns, addOperationDescribeIdentityProviderConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIdentityProviderConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIdentityProviderConfigurationInput struct {

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string
}

type DescribeIdentityProviderConfigurationOutput struct {

	// The SAML metadata document provided by the user’s identity provider.
	IdentityProviderSamlMetadata *string

	// The type of identity provider.
	IdentityProviderType types.IdentityProviderType

	// The SAML metadata document uploaded to the user’s identity provider.
	ServiceProviderSamlMetadata *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeIdentityProviderConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeIdentityProviderConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeIdentityProviderConfiguration{}, middleware.After)
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
	addOpDescribeIdentityProviderConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIdentityProviderConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeIdentityProviderConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DescribeIdentityProviderConfiguration",
	}
}
