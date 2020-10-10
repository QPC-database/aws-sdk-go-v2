// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the room details for the specified Amazon Chime Enterprise account.
// Optionally, filter the results by a member ID (user ID or bot ID) to see a list
// of rooms that the member belongs to.
func (c *Client) ListRooms(ctx context.Context, params *ListRoomsInput, optFns ...func(*Options)) (*ListRoomsOutput, error) {
	if params == nil {
		params = &ListRoomsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRooms", params, optFns, addOperationListRoomsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoomsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoomsInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The member ID (user ID or bot ID).
	MemberId *string

	// The token to use to retrieve the next page of results.
	NextToken *string
}

type ListRoomsOutput struct {

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The room details.
	Rooms []*types.Room

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRoomsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRooms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRooms{}, middleware.After)
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
	addOpListRoomsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRooms(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListRooms(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListRooms",
	}
}
