// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers with AWS CodeDeploy a revision for the specified application.
func (c *Client) RegisterApplicationRevision(ctx context.Context, params *RegisterApplicationRevisionInput, optFns ...func(*Options)) (*RegisterApplicationRevisionOutput, error) {
	if params == nil {
		params = &RegisterApplicationRevisionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterApplicationRevision", params, optFns, addOperationRegisterApplicationRevisionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterApplicationRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a RegisterApplicationRevision operation.
type RegisterApplicationRevisionInput struct {

	// The name of an AWS CodeDeploy application associated with the IAM user or AWS
	// account.
	//
	// This member is required.
	ApplicationName *string

	// Information about the application revision to register, including type and
	// location.
	//
	// This member is required.
	Revision *types.RevisionLocation

	// A comment about the revision.
	Description *string
}

type RegisterApplicationRevisionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterApplicationRevisionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterApplicationRevision{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterApplicationRevision{}, middleware.After)
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
	addOpRegisterApplicationRevisionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterApplicationRevision(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterApplicationRevision(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "RegisterApplicationRevision",
	}
}
