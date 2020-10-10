// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update an Input Security Group's Whilelists.
func (c *Client) UpdateInputSecurityGroup(ctx context.Context, params *UpdateInputSecurityGroupInput, optFns ...func(*Options)) (*UpdateInputSecurityGroupOutput, error) {
	if params == nil {
		params = &UpdateInputSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInputSecurityGroup", params, optFns, addOperationUpdateInputSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInputSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to update some combination of the Input Security Group name and the
// IPv4 CIDRs the Input Security Group should allow.
type UpdateInputSecurityGroupInput struct {

	// The id of the Input Security Group to update.
	//
	// This member is required.
	InputSecurityGroupId *string

	// A collection of key-value pairs.
	Tags map[string]*string

	// List of IPv4 CIDR addresses to whitelist
	WhitelistRules []*types.InputWhitelistRuleCidr
}

// Placeholder documentation for UpdateInputSecurityGroupResponse
type UpdateInputSecurityGroupOutput struct {

	// An Input Security Group
	SecurityGroup *types.InputSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateInputSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateInputSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateInputSecurityGroup{}, middleware.After)
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
	addOpUpdateInputSecurityGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInputSecurityGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateInputSecurityGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateInputSecurityGroup",
	}
}
