// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the users in the specified group. Calling this action requires developer
// credentials.
func (c *Client) ListUsersInGroup(ctx context.Context, params *ListUsersInGroupInput, optFns ...func(*Options)) (*ListUsersInGroupOutput, error) {
	if params == nil {
		params = &ListUsersInGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsersInGroup", params, optFns, addOperationListUsersInGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsersInGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUsersInGroupInput struct {

	// The name of the group.
	//
	// This member is required.
	GroupName *string

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The limit of the request to list users.
	Limit *int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

type ListUsersInGroupOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The users returned in the request to list users.
	Users []*types.UserType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUsersInGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUsersInGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUsersInGroup{}, middleware.After)
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
	addOpListUsersInGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListUsersInGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListUsersInGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ListUsersInGroup",
	}
}
