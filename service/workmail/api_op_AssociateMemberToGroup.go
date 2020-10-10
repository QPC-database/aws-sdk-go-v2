// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a member (user or group) to the group's set.
func (c *Client) AssociateMemberToGroup(ctx context.Context, params *AssociateMemberToGroupInput, optFns ...func(*Options)) (*AssociateMemberToGroupOutput, error) {
	if params == nil {
		params = &AssociateMemberToGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateMemberToGroup", params, optFns, addOperationAssociateMemberToGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateMemberToGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateMemberToGroupInput struct {

	// The group to which the member (user or group) is associated.
	//
	// This member is required.
	GroupId *string

	// The member (user or group) to associate to the group.
	//
	// This member is required.
	MemberId *string

	// The organization under which the group exists.
	//
	// This member is required.
	OrganizationId *string
}

type AssociateMemberToGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateMemberToGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateMemberToGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateMemberToGroup{}, middleware.After)
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
	addOpAssociateMemberToGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateMemberToGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssociateMemberToGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "AssociateMemberToGroup",
	}
}
