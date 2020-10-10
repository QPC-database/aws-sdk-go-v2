// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Provides information regarding the user.
func (c *Client) DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) {
	if params == nil {
		params = &DescribeUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUser", params, optFns, addOperationDescribeUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserInput struct {

	// The identifier for the organization under which the user exists.
	//
	// This member is required.
	OrganizationId *string

	// The identifier for the user to be described.
	//
	// This member is required.
	UserId *string
}

type DescribeUserOutput struct {

	// The date and time at which the user was disabled for Amazon WorkMail usage, in
	// UNIX epoch time format.
	DisabledDate *time.Time

	// The display name of the user.
	DisplayName *string

	// The email of the user.
	Email *string

	// The date and time at which the user was enabled for Amazon WorkMail usage, in
	// UNIX epoch time format.
	EnabledDate *time.Time

	// The name for the user.
	Name *string

	// The state of a user: enabled (registered to Amazon WorkMail) or disabled
	// (deregistered or never registered to WorkMail).
	State types.EntityState

	// The identifier for the described user.
	UserId *string

	// In certain cases, other entities are modeled as users. If interoperability is
	// enabled, resources are imported into Amazon WorkMail as users. Because different
	// WorkMail organizations rely on different directory types, administrators can
	// distinguish between an unregistered user (account is disabled and has a user
	// role) and the directory administrators. The values are USER, RESOURCE, and
	// SYSTEM_USER.
	UserRole types.UserRole

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUser{}, middleware.After)
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
	addOpDescribeUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUser(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DescribeUser",
	}
}
