// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a user from a group so that the user is no longer a member of the group.
func (c *Client) DeleteGroupMembership(ctx context.Context, params *DeleteGroupMembershipInput, optFns ...func(*Options)) (*DeleteGroupMembershipOutput, error) {
	if params == nil {
		params = &DeleteGroupMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGroupMembership", params, optFns, addOperationDeleteGroupMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGroupMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGroupMembershipInput struct {

	// The ID for the AWS account that the group is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The name of the group that you want to delete the user from.
	//
	// This member is required.
	GroupName *string

	// The name of the user that you want to delete from the group membership.
	//
	// This member is required.
	MemberName *string

	// The namespace. Currently, you should set this to default.
	//
	// This member is required.
	Namespace *string
}

type DeleteGroupMembershipOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteGroupMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteGroupMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteGroupMembership{}, middleware.After)
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
	addOpDeleteGroupMembershipValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGroupMembership(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteGroupMembership(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteGroupMembership",
	}
}
