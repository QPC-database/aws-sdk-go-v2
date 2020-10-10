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

// Modify the default properties used to create WorkSpaces.
func (c *Client) ModifyWorkspaceCreationProperties(ctx context.Context, params *ModifyWorkspaceCreationPropertiesInput, optFns ...func(*Options)) (*ModifyWorkspaceCreationPropertiesOutput, error) {
	if params == nil {
		params = &ModifyWorkspaceCreationPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyWorkspaceCreationProperties", params, optFns, addOperationModifyWorkspaceCreationPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyWorkspaceCreationPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyWorkspaceCreationPropertiesInput struct {

	// The identifier of the directory.
	//
	// This member is required.
	ResourceId *string

	// The default properties for creating WorkSpaces.
	//
	// This member is required.
	WorkspaceCreationProperties *types.WorkspaceCreationProperties
}

type ModifyWorkspaceCreationPropertiesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyWorkspaceCreationPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyWorkspaceCreationProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyWorkspaceCreationProperties{}, middleware.After)
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
	addOpModifyWorkspaceCreationPropertiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyWorkspaceCreationProperties(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyWorkspaceCreationProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ModifyWorkspaceCreationProperties",
	}
}
