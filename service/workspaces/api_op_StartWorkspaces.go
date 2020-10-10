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

// Starts the specified WorkSpaces. You cannot start a WorkSpace unless it has a
// running mode of AutoStop and a state of STOPPED.
func (c *Client) StartWorkspaces(ctx context.Context, params *StartWorkspacesInput, optFns ...func(*Options)) (*StartWorkspacesOutput, error) {
	if params == nil {
		params = &StartWorkspacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartWorkspaces", params, optFns, addOperationStartWorkspacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartWorkspacesInput struct {

	// The WorkSpaces to start. You can specify up to 25 WorkSpaces.
	//
	// This member is required.
	StartWorkspaceRequests []*types.StartRequest
}

type StartWorkspacesOutput struct {

	// Information about the WorkSpaces that could not be started.
	FailedRequests []*types.FailedWorkspaceChangeRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartWorkspacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartWorkspaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartWorkspaces{}, middleware.After)
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
	addOpStartWorkspacesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartWorkspaces(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartWorkspaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "StartWorkspaces",
	}
}
