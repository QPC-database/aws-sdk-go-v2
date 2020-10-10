// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an image pipeline.
func (c *Client) DeleteImagePipeline(ctx context.Context, params *DeleteImagePipelineInput, optFns ...func(*Options)) (*DeleteImagePipelineOutput, error) {
	if params == nil {
		params = &DeleteImagePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteImagePipeline", params, optFns, addOperationDeleteImagePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteImagePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteImagePipelineInput struct {

	// The Amazon Resource Name (ARN) of the image pipeline to delete.
	//
	// This member is required.
	ImagePipelineArn *string
}

type DeleteImagePipelineOutput struct {

	// The Amazon Resource Name (ARN) of the image pipeline that was deleted.
	ImagePipelineArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteImagePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteImagePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteImagePipeline{}, middleware.After)
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
	addOpDeleteImagePipelineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteImagePipeline(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteImagePipeline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "DeleteImagePipeline",
	}
}
