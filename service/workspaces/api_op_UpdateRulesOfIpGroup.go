// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces the current rules of the specified IP access control group with the
// specified rules.
func (c *Client) UpdateRulesOfIpGroup(ctx context.Context, params *UpdateRulesOfIpGroupInput, optFns ...func(*Options)) (*UpdateRulesOfIpGroupOutput, error) {
	if params == nil {
		params = &UpdateRulesOfIpGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRulesOfIpGroup", params, optFns, addOperationUpdateRulesOfIpGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRulesOfIpGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRulesOfIpGroupInput struct {

	// The identifier of the group.
	//
	// This member is required.
	GroupId *string

	// One or more rules.
	//
	// This member is required.
	UserRules []*types.IpRuleItem
}

type UpdateRulesOfIpGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRulesOfIpGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRulesOfIpGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRulesOfIpGroup{}, middleware.After)
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
	addOpUpdateRulesOfIpGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRulesOfIpGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateRulesOfIpGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "UpdateRulesOfIpGroup",
	}
}
