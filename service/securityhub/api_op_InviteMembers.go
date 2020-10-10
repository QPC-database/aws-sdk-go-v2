// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Invites other AWS accounts to become member accounts for the Security Hub master
// account that the invitation is sent from. Before you can use this action to
// invite a member, you must first use the CreateMembers () action to create the
// member account in Security Hub. When the account owner accepts the invitation to
// become a member account and enables Security Hub, the master account can view
// the findings generated from the member account.
func (c *Client) InviteMembers(ctx context.Context, params *InviteMembersInput, optFns ...func(*Options)) (*InviteMembersOutput, error) {
	if params == nil {
		params = &InviteMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InviteMembers", params, optFns, addOperationInviteMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InviteMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InviteMembersInput struct {

	// The list of account IDs of the AWS accounts to invite to Security Hub as
	// members.
	AccountIds []*string
}

type InviteMembersOutput struct {

	// The list of AWS accounts that could not be processed. For each account, the list
	// includes the account ID and the email address.
	UnprocessedAccounts []*types.Result

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInviteMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInviteMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInviteMembers{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opInviteMembers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opInviteMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "InviteMembers",
	}
}
