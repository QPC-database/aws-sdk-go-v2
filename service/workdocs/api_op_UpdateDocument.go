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

// Updates the specified attributes of a document. The user must have access to
// both the document and its parent folder, if applicable.
func (c *Client) UpdateDocument(ctx context.Context, params *UpdateDocumentInput, optFns ...func(*Options)) (*UpdateDocumentOutput, error) {
	if params == nil {
		params = &UpdateDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDocument", params, optFns, addOperationUpdateDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDocumentInput struct {

	// The ID of the document.
	//
	// This member is required.
	DocumentId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The name of the document.
	Name *string

	// The ID of the parent folder.
	ParentFolderId *string

	// The resource state of the document. Only ACTIVE and RECYCLED are supported.
	ResourceState types.ResourceStateType
}

type UpdateDocumentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDocument{}, middleware.After)
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
	addOpUpdateDocumentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDocument(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDocument(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "UpdateDocument",
	}
}
