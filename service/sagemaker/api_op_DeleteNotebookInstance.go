// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an Amazon SageMaker notebook instance. Before you can delete a notebook
// instance, you must call the StopNotebookInstance API. When you delete a notebook
// instance, you lose all of your data. Amazon SageMaker removes the ML compute
// instance, and deletes the ML storage volume and the network interface associated
// with the notebook instance.
func (c *Client) DeleteNotebookInstance(ctx context.Context, params *DeleteNotebookInstanceInput, optFns ...func(*Options)) (*DeleteNotebookInstanceOutput, error) {
	if params == nil {
		params = &DeleteNotebookInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteNotebookInstance", params, optFns, addOperationDeleteNotebookInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteNotebookInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteNotebookInstanceInput struct {

	// The name of the Amazon SageMaker notebook instance to delete.
	//
	// This member is required.
	NotebookInstanceName *string
}

type DeleteNotebookInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteNotebookInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteNotebookInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteNotebookInstance{}, middleware.After)
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
	addOpDeleteNotebookInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteNotebookInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteNotebookInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DeleteNotebookInstance",
	}
}
