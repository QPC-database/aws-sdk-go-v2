// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the invitations for resource sharing that you've received.
func (c *Client) GetResourceShareInvitations(ctx context.Context, params *GetResourceShareInvitationsInput, optFns ...func(*Options)) (*GetResourceShareInvitationsOutput, error) {
	if params == nil {
		params = &GetResourceShareInvitationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceShareInvitations", params, optFns, addOperationGetResourceShareInvitationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceShareInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceShareInvitationsInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []*string

	// The Amazon Resource Names (ARN) of the invitations.
	ResourceShareInvitationArns []*string
}

type GetResourceShareInvitationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the invitations.
	ResourceShareInvitations []*types.ResourceShareInvitation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetResourceShareInvitationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResourceShareInvitations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourceShareInvitations{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceShareInvitations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetResourceShareInvitations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "GetResourceShareInvitations",
	}
}
