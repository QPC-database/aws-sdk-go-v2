// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches a key policy to the specified customer master key (CMK). You cannot
// perform this operation on a CMK in a different AWS account. For more information
// about key policies, see Key Policies
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) PutKeyPolicy(ctx context.Context, params *PutKeyPolicyInput, optFns ...func(*Options)) (*PutKeyPolicyOutput, error) {
	if params == nil {
		params = &PutKeyPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutKeyPolicy", params, optFns, addOperationPutKeyPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutKeyPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutKeyPolicyInput struct {

	// A unique identifier for the customer master key (CMK). Specify the key ID or the
	// Amazon Resource Name (ARN) of the CMK. For example:
	//
	//     * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys () or DescribeKey ().
	//
	// This member is required.
	KeyId *string

	// The key policy to attach to the CMK. The key policy must meet the following
	// criteria:
	//
	//     * If you don't set BypassPolicyLockoutSafetyCheck to true, the
	// key policy must allow the principal that is making the PutKeyPolicy request to
	// make a subsequent PutKeyPolicy request on the CMK. This reduces the risk that
	// the CMK becomes unmanageable. For more information, refer to the scenario in the
	// Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	// section of the AWS Key Management Service Developer Guide.
	//
	//     * Each statement
	// in the key policy must contain one or more principals. The principals in the key
	// policy must exist and be visible to AWS KMS. When you create a new AWS principal
	// (for example, an IAM user or role), you might need to enforce a delay before
	// including the new principal in a key policy because the new principal might not
	// be immediately visible to AWS KMS. For more information, see Changes that I make
	// are not always immediately visible
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency)
	// in the AWS Identity and Access Management User Guide.
	//
	// The key policy cannot
	// exceed 32 kilobytes (32768 bytes). For more information, see Resource Quotas
	// (https://docs.aws.amazon.com/kms/latest/developerguide/resource-limits.html) in
	// the AWS Key Management Service Developer Guide.
	//
	// This member is required.
	Policy *string

	// The name of the key policy. The only valid value is default.
	//
	// This member is required.
	PolicyName *string

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the CMK becomes unmanageable.
	// Do not set this value to true indiscriminately. For more information, refer to
	// the scenario in the Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	// section in the AWS Key Management Service Developer Guide. Use this parameter
	// only when you intend to prevent the principal that is making the request from
	// making a subsequent PutKeyPolicy request on the CMK. The default value is false.
	BypassPolicyLockoutSafetyCheck *bool
}

type PutKeyPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutKeyPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutKeyPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutKeyPolicy{}, middleware.After)
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
	addOpPutKeyPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutKeyPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutKeyPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "PutKeyPolicy",
	}
}
