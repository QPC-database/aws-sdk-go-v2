// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the tags for a project.
func (c *Client) ListTagsForProject(ctx context.Context, params *ListTagsForProjectInput, optFns ...func(*Options)) (*ListTagsForProjectOutput, error) {
	if params == nil {
		params = &ListTagsForProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForProject", params, optFns, addOperationListTagsForProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsForProjectInput struct {

	// The ID of the project to get tags for.
	//
	// This member is required.
	Id *string

	// Reserved for future use.
	MaxResults *int32

	// Reserved for future use.
	NextToken *string
}

type ListTagsForProjectOutput struct {

	// Reserved for future use.
	NextToken *string

	// The tags for the project.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsForProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTagsForProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTagsForProject{}, middleware.After)
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
	addOpListTagsForProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForProject(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTagsForProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "ListTagsForProject",
	}
}
