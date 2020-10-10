// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets a stack policy for a specified stack.
func (c *Client) SetStackPolicy(ctx context.Context, params *SetStackPolicyInput, optFns ...func(*Options)) (*SetStackPolicyOutput, error) {
	if params == nil {
		params = &SetStackPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetStackPolicy", params, optFns, addOperationSetStackPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetStackPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the SetStackPolicy () action.
type SetStackPolicyInput struct {

	// The name or unique stack ID that you want to associate a policy with.
	//
	// This member is required.
	StackName *string

	// Structure containing the stack policy body. For more information, go to  Prevent
	// Updates to Stack Resources
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html)
	// in the AWS CloudFormation User Guide. You can specify either the StackPolicyBody
	// or the StackPolicyURL parameter, but not both.
	StackPolicyBody *string

	// Location of a file containing the stack policy. The URL must point to a policy
	// (maximum size: 16 KB) located in an S3 bucket in the same Region as the stack.
	// You can specify either the StackPolicyBody or the StackPolicyURL parameter, but
	// not both.
	StackPolicyURL *string
}

type SetStackPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetStackPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetStackPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetStackPolicy{}, middleware.After)
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
	addOpSetStackPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetStackPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetStackPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "SetStackPolicy",
	}
}
