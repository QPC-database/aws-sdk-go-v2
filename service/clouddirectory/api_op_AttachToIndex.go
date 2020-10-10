// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches the specified object to the specified index.
func (c *Client) AttachToIndex(ctx context.Context, params *AttachToIndexInput, optFns ...func(*Options)) (*AttachToIndexOutput, error) {
	if params == nil {
		params = &AttachToIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachToIndex", params, optFns, addOperationAttachToIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachToIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachToIndexInput struct {

	// The Amazon Resource Name (ARN) of the directory where the object and index
	// exist.
	//
	// This member is required.
	DirectoryArn *string

	// A reference to the index that you are attaching the object to.
	//
	// This member is required.
	IndexReference *types.ObjectReference

	// A reference to the object that you are attaching to the index.
	//
	// This member is required.
	TargetReference *types.ObjectReference
}

type AttachToIndexOutput struct {

	// The ObjectIdentifier of the object that was attached to the index.
	AttachedObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAttachToIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAttachToIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAttachToIndex{}, middleware.After)
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
	addOpAttachToIndexValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachToIndex(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAttachToIndex(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "AttachToIndex",
	}
}
