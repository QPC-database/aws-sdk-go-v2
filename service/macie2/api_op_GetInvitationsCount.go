// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the count of Amazon Macie membership invitations that were received by
// an account.
func (c *Client) GetInvitationsCount(ctx context.Context, params *GetInvitationsCountInput, optFns ...func(*Options)) (*GetInvitationsCountOutput, error) {
	if params == nil {
		params = &GetInvitationsCountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInvitationsCount", params, optFns, addOperationGetInvitationsCountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInvitationsCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInvitationsCountInput struct {
}

type GetInvitationsCountOutput struct {

	// The total number of invitations that were received by the account, not including
	// the currently accepted invitation.
	InvitationsCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInvitationsCountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInvitationsCount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInvitationsCount{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInvitationsCount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetInvitationsCount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetInvitationsCount",
	}
}
