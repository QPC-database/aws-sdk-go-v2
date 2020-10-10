// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a user's current mailbox quota for a specified organization and user.
func (c *Client) UpdateMailboxQuota(ctx context.Context, params *UpdateMailboxQuotaInput, optFns ...func(*Options)) (*UpdateMailboxQuotaOutput, error) {
	if params == nil {
		params = &UpdateMailboxQuotaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMailboxQuota", params, optFns, addOperationUpdateMailboxQuotaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMailboxQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMailboxQuotaInput struct {

	// The updated mailbox quota, in MB, for the specified user.
	//
	// This member is required.
	MailboxQuota *int32

	// The identifier for the organization that contains the user for whom to update
	// the mailbox quota.
	//
	// This member is required.
	OrganizationId *string

	// The identifer for the user for whom to update the mailbox quota.
	//
	// This member is required.
	UserId *string
}

type UpdateMailboxQuotaOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateMailboxQuotaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMailboxQuota{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMailboxQuota{}, middleware.After)
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
	addOpUpdateMailboxQuotaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMailboxQuota(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateMailboxQuota(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "UpdateMailboxQuota",
	}
}
