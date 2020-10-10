// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a paginated list of TypedLink facet names for a particular schema. For
// more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) ListTypedLinkFacetNames(ctx context.Context, params *ListTypedLinkFacetNamesInput, optFns ...func(*Options)) (*ListTypedLinkFacetNamesOutput, error) {
	if params == nil {
		params = &ListTypedLinkFacetNamesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTypedLinkFacetNames", params, optFns, addOperationListTypedLinkFacetNamesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTypedLinkFacetNamesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypedLinkFacetNamesInput struct {

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns ().
	//
	// This member is required.
	SchemaArn *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListTypedLinkFacetNamesOutput struct {

	// The names of typed link facets that exist within the schema.
	FacetNames []*string

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTypedLinkFacetNamesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTypedLinkFacetNames{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTypedLinkFacetNames{}, middleware.After)
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
	addOpListTypedLinkFacetNamesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTypedLinkFacetNames(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTypedLinkFacetNames(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListTypedLinkFacetNames",
	}
}
