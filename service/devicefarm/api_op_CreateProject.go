// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a project.
func (c *Client) CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) {
	if params == nil {
		params = &CreateProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProject", params, optFns, addOperationCreateProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the create project operation.
type CreateProjectInput struct {

	// The project's name.
	//
	// This member is required.
	Name *string

	// Sets the execution timeout value (in minutes) for a project. All test runs in
	// this project use the specified execution timeout value unless overridden when
	// scheduling a run.
	DefaultJobTimeoutMinutes *int32
}

// Represents the result of a create project request.
type CreateProjectOutput struct {

	// The newly created project.
	Project *types.Project

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProject{}, middleware.After)
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
	addOpCreateProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProject(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateProject",
	}
}
