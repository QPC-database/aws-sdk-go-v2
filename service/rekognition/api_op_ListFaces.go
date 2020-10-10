// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns metadata for faces in the specified collection. This metadata includes
// information such as the bounding box coordinates, the confidence (that the
// bounding box contains a face), and face ID. For an example, see Listing Faces in
// a Collection in the Amazon Rekognition Developer Guide.  <p>This operation
// requires permissions to perform the <code>rekognition:ListFaces</code>
// action.</p>
func (c *Client) ListFaces(ctx context.Context, params *ListFacesInput, optFns ...func(*Options)) (*ListFacesOutput, error) {
	if params == nil {
		params = &ListFacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFaces", params, optFns, addOperationListFacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFacesInput struct {

	// ID of the collection from which to list the faces.
	//
	// This member is required.
	CollectionId *string

	// Maximum number of faces to return.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Rekognition returns a pagination token in the response. You
	// can use this pagination token to retrieve the next set of faces.
	NextToken *string
}

type ListFacesOutput struct {

	// Version number of the face detection model associated with the input collection
	// (CollectionId).
	FaceModelVersion *string

	// An array of Face objects.
	Faces []*types.Face

	// If the response is truncated, Amazon Rekognition returns this token that you can
	// use in the subsequent request to retrieve the next set of faces.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFaces{}, middleware.After)
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
	addOpListFacesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFaces(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListFaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "ListFaces",
	}
}
