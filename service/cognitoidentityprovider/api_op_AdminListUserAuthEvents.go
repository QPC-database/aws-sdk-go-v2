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

// Lists a history of user activity and any risks detected as part of Amazon
// Cognito advanced security.
func (c *Client) AdminListUserAuthEvents(ctx context.Context, params *AdminListUserAuthEventsInput, optFns ...func(*Options)) (*AdminListUserAuthEventsOutput, error) {
	if params == nil {
		params = &AdminListUserAuthEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminListUserAuthEvents", params, optFns, addOperationAdminListUserAuthEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminListUserAuthEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminListUserAuthEventsInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user pool username or an alias.
	//
	// This member is required.
	Username *string

	// The maximum number of authentication events to return.
	MaxResults *int32

	// A pagination token.
	NextToken *string
}

type AdminListUserAuthEventsOutput struct {

	// The response object. It includes the EventID, EventType, CreationDate,
	// EventRisk, and EventResponse.
	AuthEvents []*types.AuthEventType

	// A pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAdminListUserAuthEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminListUserAuthEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminListUserAuthEvents{}, middleware.After)
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
	addOpAdminListUserAuthEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminListUserAuthEvents(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAdminListUserAuthEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminListUserAuthEvents",
	}
}
