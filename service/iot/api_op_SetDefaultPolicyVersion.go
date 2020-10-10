// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the specified version of the specified policy as the policy's default
// (operative) version. This action affects all certificates to which the policy is
// attached. To list the principals the policy is attached to, use the
// ListPrincipalPolicy API.
func (c *Client) SetDefaultPolicyVersion(ctx context.Context, params *SetDefaultPolicyVersionInput, optFns ...func(*Options)) (*SetDefaultPolicyVersionOutput, error) {
	if params == nil {
		params = &SetDefaultPolicyVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetDefaultPolicyVersion", params, optFns, addOperationSetDefaultPolicyVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetDefaultPolicyVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the SetDefaultPolicyVersion operation.
type SetDefaultPolicyVersionInput struct {

	// The policy name.
	//
	// This member is required.
	PolicyName *string

	// The policy version ID.
	//
	// This member is required.
	PolicyVersionId *string
}

type SetDefaultPolicyVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetDefaultPolicyVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSetDefaultPolicyVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSetDefaultPolicyVersion{}, middleware.After)
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
	addOpSetDefaultPolicyVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetDefaultPolicyVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetDefaultPolicyVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "SetDefaultPolicyVersion",
	}
}
