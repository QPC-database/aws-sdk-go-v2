// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers the IAM role that grants Amazon Inspector access to AWS Services
// needed to perform security assessments.
func (c *Client) RegisterCrossAccountAccessRole(ctx context.Context, params *RegisterCrossAccountAccessRoleInput, optFns ...func(*Options)) (*RegisterCrossAccountAccessRoleOutput, error) {
	if params == nil {
		params = &RegisterCrossAccountAccessRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterCrossAccountAccessRole", params, optFns, addOperationRegisterCrossAccountAccessRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterCrossAccountAccessRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterCrossAccountAccessRoleInput struct {

	// The ARN of the IAM role that grants Amazon Inspector access to AWS Services
	// needed to perform security assessments.
	//
	// This member is required.
	RoleArn *string
}

type RegisterCrossAccountAccessRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterCrossAccountAccessRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterCrossAccountAccessRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterCrossAccountAccessRole{}, middleware.After)
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
	addOpRegisterCrossAccountAccessRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCrossAccountAccessRole(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterCrossAccountAccessRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "RegisterCrossAccountAccessRole",
	}
}
