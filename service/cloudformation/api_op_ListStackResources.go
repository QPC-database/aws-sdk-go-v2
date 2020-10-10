// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns descriptions of all resources of the specified stack. For deleted
// stacks, ListStackResources returns resource information for up to 90 days after
// the stack has been deleted.
func (c *Client) ListStackResources(ctx context.Context, params *ListStackResourcesInput, optFns ...func(*Options)) (*ListStackResourcesOutput, error) {
	if params == nil {
		params = &ListStackResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStackResources", params, optFns, addOperationListStackResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStackResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListStackResource () action.
type ListStackResourcesInput struct {

	// The name or the unique stack ID that is associated with the stack, which are not
	// always interchangeable:
	//
	//     * Running stacks: You can specify either the
	// stack's name or its unique stack ID.
	//
	//     * Deleted stacks: You must specify the
	// unique stack ID.
	//
	// Default: There is no default value.
	//
	// This member is required.
	StackName *string

	// A string that identifies the next page of stack resources that you want to
	// retrieve.
	NextToken *string
}

// The output for a ListStackResources () action.
type ListStackResourcesOutput struct {

	// If the output exceeds 1 MB, a string that identifies the next page of stack
	// resources. If no additional page exists, this value is null.
	NextToken *string

	// A list of StackResourceSummary structures.
	StackResourceSummaries []*types.StackResourceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListStackResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListStackResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListStackResources{}, middleware.After)
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
	addOpListStackResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListStackResources(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListStackResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListStackResources",
	}
}
