// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the name and/or the path of the specified IAM group. You should
// understand the implications of changing a group's path or name. For more
// information, see Renaming Users and Groups
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_WorkingWithGroupsAndUsers.html)
// in the IAM User Guide. The person making the request (the principal), must have
// permission to change the role group with the old name and the new name. For
// example, to change the group named Managers to MGRs, the principal must have a
// policy that allows them to update both groups. If the principal has permission
// to update the Managers group, but not the MGRs group, then the update fails. For
// more information about permissions, see Access Management
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access.html).
func (c *Client) UpdateGroup(ctx context.Context, params *UpdateGroupInput, optFns ...func(*Options)) (*UpdateGroupOutput, error) {
	if params == nil {
		params = &UpdateGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGroup", params, optFns, addOperationUpdateGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGroupInput struct {

	// Name of the IAM group to update. If you're changing the name of the group, this
	// is the original name. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// This member is required.
	GroupName *string

	// New name for the IAM group. Only include this if changing the group's name. IAM
	// user, group, role, and policy names must be unique within the account. Names are
	// not distinguished by case. For example, you cannot create resources named both
	// "MyResource" and "myresource".
	NewGroupName *string

	// New path for the IAM group. Only include this if changing the group's path. This
	// parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a
	// string of characters consisting of either a forward slash (/) by itself or a
	// string that must begin and end with forward slashes. In addition, it can contain
	// any ASCII character from the ! (\u0021) through the DEL character (\u007F),
	// including most punctuation characters, digits, and upper and lowercased letters.
	NewPath *string
}

type UpdateGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateGroup{}, middleware.After)
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
	addOpUpdateGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateGroup",
	}
}
