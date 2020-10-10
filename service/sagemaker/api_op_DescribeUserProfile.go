// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a user profile. For more information, see CreateUserProfile.
func (c *Client) DescribeUserProfile(ctx context.Context, params *DescribeUserProfileInput, optFns ...func(*Options)) (*DescribeUserProfileOutput, error) {
	if params == nil {
		params = &DescribeUserProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUserProfile", params, optFns, addOperationDescribeUserProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserProfileInput struct {

	// The domain ID.
	//
	// This member is required.
	DomainId *string

	// The user profile name.
	//
	// This member is required.
	UserProfileName *string
}

type DescribeUserProfileOutput struct {

	// The creation time.
	CreationTime *time.Time

	// The ID of the domain that contains the profile.
	DomainId *string

	// The failure reason.
	FailureReason *string

	// The ID of the user's profile in the Amazon Elastic File System (EFS) volume.
	HomeEfsFileSystemUid *string

	// The last modified time.
	LastModifiedTime *time.Time

	// The SSO user identifier.
	SingleSignOnUserIdentifier *string

	// The SSO user value.
	SingleSignOnUserValue *string

	// The status.
	Status types.UserProfileStatus

	// The user profile Amazon Resource Name (ARN).
	UserProfileArn *string

	// The user profile name.
	UserProfileName *string

	// A collection of settings.
	UserSettings *types.UserSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeUserProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUserProfile{}, middleware.After)
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
	addOpDescribeUserProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUserProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeUserProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeUserProfile",
	}
}
