// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes one or more rules from the specified IP access control group.
func (c *Client) RevokeIpRules(ctx context.Context, params *RevokeIpRulesInput, optFns ...func(*Options)) (*RevokeIpRulesOutput, error) {
	if params == nil {
		params = &RevokeIpRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeIpRules", params, optFns, addOperationRevokeIpRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeIpRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeIpRulesInput struct {

	// The identifier of the group.
	//
	// This member is required.
	GroupId *string

	// The rules to remove from the group.
	//
	// This member is required.
	UserRules []*string
}

type RevokeIpRulesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRevokeIpRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRevokeIpRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRevokeIpRules{}, middleware.After)
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
	addOpRevokeIpRulesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeIpRules(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRevokeIpRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "RevokeIpRules",
	}
}
