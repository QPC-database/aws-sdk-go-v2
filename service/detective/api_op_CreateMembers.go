// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends a request to invite the specified AWS accounts to be member accounts in
// the behavior graph. This operation can only be called by the master account for
// a behavior graph. CreateMembers verifies the accounts and then sends invitations
// to the verified accounts. The request provides the behavior graph ARN and the
// list of accounts to invite. The response separates the requested accounts into
// two lists:
//
//     * The accounts that CreateMembers was able to start the
// verification for. This list includes member accounts that are being verified,
// that have passed verification and are being sent an invitation, and that have
// failed verification.
//
//     * The accounts that CreateMembers was unable to
// process. This list includes accounts that were already invited to be member
// accounts in the behavior graph.
func (c *Client) CreateMembers(ctx context.Context, params *CreateMembersInput, optFns ...func(*Options)) (*CreateMembersOutput, error) {
	if params == nil {
		params = &CreateMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMembers", params, optFns, addOperationCreateMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMembersInput struct {

	// The list of AWS accounts to invite to become member accounts in the behavior
	// graph. For each invited account, the account list contains the account
	// identifier and the AWS account root user email address.
	//
	// This member is required.
	Accounts []*types.Account

	// The ARN of the behavior graph to invite the member accounts to contribute their
	// data to.
	//
	// This member is required.
	GraphArn *string

	// Customized message text to include in the invitation email message to the
	// invited member accounts.
	Message *string
}

type CreateMembersOutput struct {

	// The set of member account invitation requests that Detective was able to
	// process. This includes accounts that are being verified, that failed
	// verification, and that passed verification and are being sent an invitation.
	Members []*types.MemberDetail

	// The list of accounts for which Detective was unable to process the invitation
	// request. For each account, the list provides the reason why the request could
	// not be processed. The list includes accounts that are already member accounts in
	// the behavior graph.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMembers{}, middleware.After)
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
	addOpCreateMembersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMembers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "CreateMembers",
	}
}
