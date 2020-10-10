// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves version metadata for the specified document.
func (c *Client) GetDocumentVersion(ctx context.Context, params *GetDocumentVersionInput, optFns ...func(*Options)) (*GetDocumentVersionOutput, error) {
	if params == nil {
		params = &GetDocumentVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDocumentVersion", params, optFns, addOperationGetDocumentVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDocumentVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDocumentVersionInput struct {

	// The ID of the document.
	//
	// This member is required.
	DocumentId *string

	// The version ID of the document.
	//
	// This member is required.
	VersionId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// A comma-separated list of values. Specify "SOURCE" to include a URL for the
	// source document.
	Fields *string

	// Set this to TRUE to include custom metadata in the response.
	IncludeCustomMetadata *bool
}

type GetDocumentVersionOutput struct {

	// The custom metadata on the document version.
	CustomMetadata map[string]*string

	// The version metadata.
	Metadata *types.DocumentVersionMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDocumentVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDocumentVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDocumentVersion{}, middleware.After)
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
	addOpGetDocumentVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDocumentVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "GetDocumentVersion",
	}
}
