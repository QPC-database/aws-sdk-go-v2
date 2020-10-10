// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a user from a project. Removing a user from a project also removes the
// IAM policies from that user that allowed access to the project and its
// resources. Disassociating a team member does not remove that user's profile from
// AWS CodeStar. It does not remove the user from IAM.
func (c *Client) DisassociateTeamMember(ctx context.Context, params *DisassociateTeamMemberInput, optFns ...func(*Options)) (*DisassociateTeamMemberOutput, error) {
	if params == nil {
		params = &DisassociateTeamMemberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateTeamMember", params, optFns, addOperationDisassociateTeamMemberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateTeamMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateTeamMemberInput struct {

	// The ID of the AWS CodeStar project from which you want to remove a team member.
	//
	// This member is required.
	ProjectId *string

	// The Amazon Resource Name (ARN) of the IAM user or group whom you want to remove
	// from the project.
	//
	// This member is required.
	UserArn *string
}

type DisassociateTeamMemberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateTeamMemberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateTeamMember{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateTeamMember{}, middleware.After)
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
	addOpDisassociateTeamMemberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateTeamMember(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateTeamMember(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "DisassociateTeamMember",
	}
}
