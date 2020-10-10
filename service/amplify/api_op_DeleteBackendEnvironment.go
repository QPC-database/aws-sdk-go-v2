// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a backend environment for an Amplify app.
func (c *Client) DeleteBackendEnvironment(ctx context.Context, params *DeleteBackendEnvironmentInput, optFns ...func(*Options)) (*DeleteBackendEnvironmentOutput, error) {
	if params == nil {
		params = &DeleteBackendEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBackendEnvironment", params, optFns, addOperationDeleteBackendEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBackendEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the delete backend environment request.
type DeleteBackendEnvironmentInput struct {

	// The unique ID of an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of a backend environment of an Amplify app.
	//
	// This member is required.
	EnvironmentName *string
}

// The result structure of the delete backend environment result.
type DeleteBackendEnvironmentOutput struct {

	// Describes the backend environment for an Amplify app.
	//
	// This member is required.
	BackendEnvironment *types.BackendEnvironment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBackendEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBackendEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBackendEnvironment{}, middleware.After)
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
	addOpDeleteBackendEnvironmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBackendEnvironment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteBackendEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "DeleteBackendEnvironment",
	}
}
