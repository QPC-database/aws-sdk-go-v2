// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a user account for the specified Amazon Connect instance.
func (c *Client) CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) {
	if params == nil {
		params = &CreateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUser", params, optFns, addOperationCreateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The phone settings for the user.
	//
	// This member is required.
	PhoneConfig *types.UserPhoneConfig

	// The identifier of the routing profile for the user.
	//
	// This member is required.
	RoutingProfileId *string

	// The identifier of the security profile for the user.
	//
	// This member is required.
	SecurityProfileIds []*string

	// The user name for the account. For instances not using SAML for identity
	// management, the user name can include up to 20 characters. If you are using SAML
	// for identity management, the user name can include up to 64 characters from
	// [a-zA-Z0-9_-.\@]+.
	//
	// This member is required.
	Username *string

	// The identifier of the user account in the directory used for identity
	// management. If Amazon Connect cannot access the directory, you can specify this
	// identifier to authenticate users. If you include the identifier, we assume that
	// Amazon Connect cannot access the directory. Otherwise, the identity information
	// is used to authenticate users from your directory. This parameter is required if
	// you are using an existing directory for identity management in Amazon Connect
	// when Amazon Connect cannot access your directory to authenticate users. If you
	// are using SAML for identity management and include this parameter, an error is
	// returned.
	DirectoryUserId *string

	// The identifier of the hierarchy group for the user.
	HierarchyGroupId *string

	// The information about the identity of the user.
	IdentityInfo *types.UserIdentityInfo

	// The password for the user account. A password is required if you are using
	// Amazon Connect for identity management. Otherwise, it is an error to include a
	// password.
	Password *string

	// One or more tags.
	Tags map[string]*string
}

type CreateUserOutput struct {

	// The Amazon Resource Name (ARN) of the user account.
	UserArn *string

	// The identifier of the user account.
	UserId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateUser{}, middleware.After)
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
	addOpCreateUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUser(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "CreateUser",
	}
}
