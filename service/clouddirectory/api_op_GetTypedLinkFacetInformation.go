// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the identity attribute order for a specific TypedLinkFacet (). For more
// information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) GetTypedLinkFacetInformation(ctx context.Context, params *GetTypedLinkFacetInformationInput, optFns ...func(*Options)) (*GetTypedLinkFacetInformationOutput, error) {
	if params == nil {
		params = &GetTypedLinkFacetInformationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTypedLinkFacetInformation", params, optFns, addOperationGetTypedLinkFacetInformationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTypedLinkFacetInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTypedLinkFacetInformationInput struct {

	// The unique name of the typed link facet.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns ().
	//
	// This member is required.
	SchemaArn *string
}

type GetTypedLinkFacetInformationOutput struct {

	// The order of identity attributes for the facet, from most significant to least
	// significant. The ability to filter typed links considers the order that the
	// attributes are defined on the typed link facet. When providing ranges to typed
	// link selection, any inexact ranges must be specified at the end. Any attributes
	// that do not have a range specified are presumed to match the entire range.
	// Filters are interpreted in the order of the attributes on the typed link facet,
	// not the order in which they are supplied to any API calls. For more information
	// about identity attributes, see Typed Links
	// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
	IdentityAttributeOrder []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTypedLinkFacetInformationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTypedLinkFacetInformation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTypedLinkFacetInformation{}, middleware.After)
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
	addOpGetTypedLinkFacetInformationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTypedLinkFacetInformation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetTypedLinkFacetInformation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "GetTypedLinkFacetInformation",
	}
}
