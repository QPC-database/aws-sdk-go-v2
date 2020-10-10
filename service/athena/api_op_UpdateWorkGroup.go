// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the workgroup with the specified name. The workgroup's name cannot be
// changed.
func (c *Client) UpdateWorkGroup(ctx context.Context, params *UpdateWorkGroupInput, optFns ...func(*Options)) (*UpdateWorkGroupOutput, error) {
	if params == nil {
		params = &UpdateWorkGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkGroup", params, optFns, addOperationUpdateWorkGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkGroupInput struct {

	// The specified workgroup that will be updated.
	//
	// This member is required.
	WorkGroup *string

	// The workgroup configuration that will be updated for the given workgroup.
	ConfigurationUpdates *types.WorkGroupConfigurationUpdates

	// The workgroup description.
	Description *string

	// The workgroup state that will be updated for the given workgroup.
	State types.WorkGroupState
}

type UpdateWorkGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateWorkGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWorkGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWorkGroup{}, middleware.After)
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
	addOpUpdateWorkGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateWorkGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "UpdateWorkGroup",
	}
}
